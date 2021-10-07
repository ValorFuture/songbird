// (c) 2021, Flare Networks Limited. All rights reserved.
// Please see the file LICENSE for licensing terms.

// SPDX-License-Identifier: MIT
pragma solidity 0.7.6;

contract StateConnector {

//====================================================================
// Data Structures
//====================================================================

    struct HashExists {
        bool        exists;
        bool        proven;
        uint256     revealTime;
        uint64      indexValue;
        bytes32     hashValue;
    }

    uint256 private constant TWO_PHASE_COMMIT_LOWER_BOUND = 30;
    uint256 private constant TWO_PHASE_COMMIT_UPPER_BOUND = 1 days;
    address private constant GENESIS_COINBASE = address(0x0100000000000000000000000000000000000000);

    // Finalised payment hashes
    mapping(bytes32 => HashExists) private payments;

//====================================================================
// Events
//====================================================================

    event PaymentProven(bool prove, uint64 chainId, uint64 ledger, string txId, bytes32 paymentHash);

//====================================================================
// Constructor
//====================================================================

    constructor() {
    }

//====================================================================
// Functions
//====================================================================  

    function setPaymentFinality(
        bool prove,
        uint64 chainId,
        uint64 ledger,
        bytes32 paymentHash,
        string memory txId
    ) external returns (
        uint256 _instructions,
        bytes32 _paymentHash,
        string memory _txId
    ) {
        require(ledger > 0, "ledger == 0");
        require(paymentHash > 0x0, "paymentHash == 0x0");
        require(bytes(txId).length > 0, "txId is empty string");
        require(block.coinbase == msg.sender || block.coinbase == GENESIS_COINBASE, "invalid block.coinbase value");

        bytes32 txIdHash = keccak256(abi.encodePacked(txId));
        bytes32 finalisedPaymentLoc = keccak256(abi.encodePacked("finalisedPayment", chainId, paymentHash, txIdHash));
        require(!payments[finalisedPaymentLoc].proven, "payment already proven");

        bool initialCommit;
        bytes32 proposedLoc = keccak256(abi.encodePacked(prove, chainId, ledger, paymentHash, txIdHash));
        if (payments[proposedLoc].exists) {
            require(block.timestamp >= payments[proposedLoc].revealTime, 
                "block.timestamp < payments[proposedLoc].revealTime");
            require(payments[proposedLoc].hashValue == paymentHash, 
                "invalid paymentHash");
            require(payments[proposedLoc].revealTime + TWO_PHASE_COMMIT_UPPER_BOUND > block.timestamp,
                "reveal is too late");
        } else if (block.coinbase != msg.sender && block.coinbase == GENESIS_COINBASE) {
            initialCommit = true;
        }

        if (block.coinbase == msg.sender && block.coinbase != GENESIS_COINBASE) {
            if (!payments[proposedLoc].exists) {
                payments[proposedLoc] = HashExists(
                    true,
                    false,
                    block.timestamp + TWO_PHASE_COMMIT_LOWER_BOUND,
                    ledger,
                    paymentHash
                );
            } else {
                payments[finalisedPaymentLoc] = HashExists(
                    true,
                    prove,
                    block.timestamp,
                    ledger,
                    paymentHash
                );
                emit PaymentProven(prove, chainId, ledger, txId, paymentHash);
            }
        }

        return ((initialCommit?1:0)*2**196 + (prove?1:0)*2**128 + chainId*2**64 + ledger, paymentHash, txId);
    }

    function getPaymentFinality(
        uint64 chainId,
        bytes32 txIdHash,
        bytes32 destinationHash,
        uint64 amount,
        bytes32 currencyHash
    ) external view returns (
        bool finality,
        uint64 ledger
    ) {
        bytes32 paymentHash = keccak256(abi.encodePacked(
            txIdHash,
            destinationHash,
            keccak256(abi.encode(amount)),
            currencyHash));
        bytes32 finalisedPaymentLoc = keccak256(abi.encodePacked("finalisedPayment", chainId, paymentHash, txIdHash));
        require(payments[finalisedPaymentLoc].exists, "payment does not exist");
        require(payments[finalisedPaymentLoc].hashValue == paymentHash, "invalid paymentHash");

        return (
            payments[finalisedPaymentLoc].proven,
            payments[finalisedPaymentLoc].indexValue
        );
    }

}
