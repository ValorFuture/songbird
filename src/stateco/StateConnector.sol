// (c) 2021, Flare Networks Limited. All rights reserved.
// Please see the file LICENSE for licensing terms.

// SPDX-License-Identifier: MIT
pragma solidity 0.8.9;

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

    event PaymentSet(uint64 chainId, uint64 ledger, bytes32 txId, uint16 utxo, bytes32 paymentHash);

//====================================================================
// Constructor
//====================================================================

    constructor() {
    }

//====================================================================
// Functions
//====================================================================  

    function setPaymentFinality(
        uint64 chainId,
        uint64 ledger,
        bytes32 txId,
        uint16 utxo,
        bytes32 paymentHash
    ) external returns (
        uint256 _instructions,
        bytes32 _txId,
        bytes32 _paymentHash
    ) {
        require(ledger > 0, "ledger == 0");
        require(txId > 0x0, "txId == 0x0");
        require(paymentHash > 0x0, "paymentHash == 0x0");
        require(block.coinbase == msg.sender || block.coinbase == GENESIS_COINBASE, "invalid block.coinbase value");

        bytes32 finalisedPaymentLocation = keccak256(abi.encodePacked(
            keccak256(abi.encode("FlareStateConnector_FINALISED")),
            keccak256(abi.encode(chainId)),
            keccak256(abi.encode(ledger)),
            txId,
            keccak256(abi.encode(utxo)),
            paymentHash
        ));
        require(!payments[finalisedPaymentLocation].proven, "payment already proven");

        bool initialCommit;
        bytes32 proposedPaymentLocation = keccak256(abi.encodePacked(
            keccak256(abi.encode("FlareStateConnector_PROPOSED")),
            keccak256(abi.encode(msg.sender)),
            keccak256(abi.encode(chainId)),
            keccak256(abi.encode(ledger)),
            txId,
            keccak256(abi.encode(utxo)),
            paymentHash
        ));
        if (payments[proposedPaymentLocation].exists) {
            require(block.timestamp >= payments[proposedPaymentLocation].revealTime, 
                "block.timestamp < payments[proposedPaymentLocation].revealTime");
            require(payments[proposedPaymentLocation].revealTime + TWO_PHASE_COMMIT_UPPER_BOUND > block.timestamp,
                "reveal is too late");
            require(payments[proposedPaymentLocation].indexValue == ledger, 
                "invalid ledger");
            require(payments[proposedPaymentLocation].hashValue == paymentHash, 
                "invalid paymentHash");
        } else if (block.coinbase != msg.sender && block.coinbase == GENESIS_COINBASE) {
            initialCommit = true;
        }

        if (block.coinbase == msg.sender && block.coinbase != GENESIS_COINBASE) {
            if (!payments[proposedPaymentLocation].exists) {
                payments[proposedPaymentLocation] = HashExists(
                    true,
                    false,
                    block.timestamp + TWO_PHASE_COMMIT_LOWER_BOUND,
                    ledger,
                    paymentHash
                );
            } else {
                payments[finalisedPaymentLocation] = HashExists(
                    true,
                    true,
                    block.timestamp,
                    ledger,
                    paymentHash
                );
                emit PaymentSet(chainId, ledger, txId, utxo, paymentHash);
            }
        }

        return (
            uint256(initialCommit?1:0)<<192 | uint256(chainId)<<128 | uint256(ledger)<<64 | uint256(utxo),
            txId,
            paymentHash
        );
    }

    function getPaymentFinality(
        uint64 chainId,
        uint64 ledger,
        bytes32 txIdHash,
        uint16 utxo,
        bytes32 originHash,
        bytes32 destinationHash,
        bytes32 currencyHash,
        uint64 amount
    ) external view returns (
        bool _proven
    ) {
        bytes32 paymentHash = keccak256(abi.encodePacked(
            keccak256(abi.encode("FlareStateConnector_PAYMENTHASH")),
            keccak256(abi.encode(chainId)),
            keccak256(abi.encode(ledger)),
            txIdHash,
            keccak256(abi.encode(utxo)),
            originHash,
            destinationHash,
            currencyHash,
            keccak256(abi.encode(amount))
        ));
        bytes32 finalisedPaymentLoc = keccak256(abi.encodePacked(
            keccak256(abi.encode("FlareStateConnector_FINALISED")),
            keccak256(abi.encode(chainId)),
            keccak256(abi.encode(ledger)),
            txIdHash,
            keccak256(abi.encode(utxo)),
            paymentHash
        ));
        require(payments[finalisedPaymentLoc].exists, "payment does not exist");
        require(payments[finalisedPaymentLoc].proven, "payment is not yet proven");
        require(payments[finalisedPaymentLoc].indexValue == ledger, "invalid ledger value");
        require(payments[finalisedPaymentLoc].hashValue == paymentHash, "invalid paymentHash");

        return (payments[finalisedPaymentLoc].proven);
    }

}
