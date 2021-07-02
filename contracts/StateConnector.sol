// SPDX-License-Identifier: MIT
pragma solidity 0.7.6;

contract StateConnector {

//====================================================================
// Data Structures
//====================================================================

    address private governanceContract;
    bool private initialised;
    uint32 private numChains;
    uint256 private initialiseTime;
    uint64 private rewardPeriodTimespan;

    struct Chain {
        bool        exists;
        uint64      genesisLedger;
        uint64      ledgerHistorySize; // Range of ledgers below finalisedLedgerIndex that can searched when proving a payment
        uint16      claimPeriodLength; // Number of ledgers in a claim period
        uint16      numConfirmations; // Number of confirmations required to consider this claim period finalised
        uint64      finalisedClaimPeriodIndex;
        uint64      finalisedLedgerIndex;
        uint256     finalisedTimestamp;
        uint256     timeDiffExpected;
        uint256     timeDiffAvg;
    }

    struct HashExists {
        bool        exists;
        bytes32     commitHash;
        uint256     commitTime;
        uint256     permittedRevealTime;
        bytes32     revealHash;
        uint64      index;
        uint64      indexSearchRegion;
        bool        proven;
    }

    // Chain ID mapping to Chain struct
    mapping(uint32 => Chain) private chains;
    // msg.sender => Location hash => claim period
    mapping(address => mapping(bytes32 => HashExists)) private proposedClaimPeriods;
    // Location hash => claim period
    mapping(bytes32 => HashExists) private finalisedClaimPeriods;
    // Finalised payment hashes
    mapping(uint32 => mapping(bytes32 => HashExists)) private finalisedPayments;
    // Mapping of how many claim periods an address has successfully mined
    mapping(address => mapping(uint256 => uint64)) private claimPeriodsMined;
    // Accounts that the governance contract voted to block from submitting proofs
    mapping(address => bool) private governanceBlockedAccounts;

//====================================================================
// Constructor for pre-compiled code
//====================================================================

    constructor() {
    }

    modifier onlyGovernance() {
        require(msg.sender == governanceContract, 'msg.sender != governanceContract');
        _;
    }

    modifier chainExists(uint32 chainId) {
        require(chains[chainId].exists, 'chainId does not exist');
        _;
    }

    modifier senderNotGovernanceBlocked() {
        require(!governanceBlockedAccounts[msg.sender], 'this account is governance blocked');
        _;
    }

    function initialiseChains() public returns (bool success) {
        require(!initialised, 'initialised != false');
        governanceContract = 0xfffEc6C83c8BF5c3F4AE0cCF8c45CE20E4560BD7;
        chains[0] = Chain(true, 689300, 0, 1, 4, 0, 689300, block.timestamp, 900, 0); //BTC
        chains[1] = Chain(true, 2069000, 0, 1, 12, 0, 2069000, block.timestamp, 150, 0); //LTC
        chains[2] = Chain(true, 3768500, 0, 2, 40, 0, 3768500, block.timestamp, 120, 0); //DOGE
        chains[3] = Chain(true, 64583400, 0, 30, 0, 0, 64583400, block.timestamp, 120, 0); //XRP
        chains[4] = Chain(true, 35863000, 0, 20, 0, 0, 35863000, block.timestamp, 120, 0); //XLM 
        numChains = 5;
        rewardPeriodTimespan = 604800;
        initialiseTime = block.timestamp;
        initialised = true;
        return true;
    }

//====================================================================
// Functions
//====================================================================  

    function getGovernanceContract() external view returns (address _governanceContract) {
        return governanceContract;
    }

    function setGovernanceContract(address _governanceContract) external onlyGovernance {
        require(_governanceContract != address(0x0), '_governanceContract == 0x0');
        governanceContract = _governanceContract;
    }

    function addChain(uint64 genesisLedger, uint64 ledgerHistorySize, uint16 claimPeriodLength, uint16 numConfirmations, uint256 timeDiffExpected) external onlyGovernance {
        require(!chains[numChains].exists, 'chainId already exists'); // Can happen if numChains is overflowed
        require(claimPeriodLength > 0, 'claimPeriodLength == 0');
        chains[numChains] = Chain(true, genesisLedger, ledgerHistorySize, claimPeriodLength, numConfirmations, 0, genesisLedger, block.timestamp, timeDiffExpected, 0);
        numChains = numChains+1;
    }

    // Solution if an underlying chain loses liveness is to disable that chain temporarily
    function disableChain(uint32 chainId) external onlyGovernance chainExists(chainId) {
        chains[chainId].exists = false;
    }

    function enableChain(uint32 chainId) external onlyGovernance {
        require(chainId < numChains, 'chainId >= numChains');
        require(chains[chainId].exists == false, 'chains[chainId].exists == true');
        chains[chainId].exists = true;
    }

    function updateChainTiming(uint32 chainId, uint64 ledgerHistorySize, uint256 timeDiffExpected) external onlyGovernance chainExists(chainId) {
        chains[chainId].ledgerHistorySize = ledgerHistorySize;
        chains[chainId].timeDiffExpected = timeDiffExpected;
    }

    function getClaimPeriodsMined(address miner, uint256 rewardSchedule) external view returns (uint64 numMined) {
        return claimPeriodsMined[miner][rewardSchedule];
    }

    function getRewardPeriod() private view returns (uint256 rewardSchedule) {
        require(block.timestamp > initialiseTime, "block.timestamp <= initialiseTime");
        return (block.timestamp - initialiseTime)/rewardPeriodTimespan;
    }

    function blockAddress(address blockedAddress) external onlyGovernance {
        require(blockedAddress != governanceContract, 'blockedAddress == governanceContract');
        governanceBlockedAccounts[blockedAddress] = true;
    }

    function unblockAddress(address blockedAddress) external onlyGovernance {
        require(governanceBlockedAccounts[blockedAddress], '!governanceBlockedAccounts[blockedAddress]');
        governanceBlockedAccounts[blockedAddress] = false;
    }

    function getLatestIndex(uint32 chainId) external view chainExists(chainId) returns (uint64 genesisLedger, uint64 finalisedClaimPeriodIndex, uint16 claimPeriodLength, uint64 finalisedLedgerIndex, uint256 finalisedTimestamp, uint256 timeDiffAvg) {
        finalisedTimestamp = chains[chainId].finalisedTimestamp;
        timeDiffAvg = chains[chainId].timeDiffAvg;
        if (proposedClaimPeriods[msg.sender][keccak256(abi.encodePacked(chainId,chains[chainId].finalisedClaimPeriodIndex))].exists) {
            finalisedTimestamp = 0;
            timeDiffAvg = proposedClaimPeriods[msg.sender][keccak256(abi.encodePacked(chainId,chains[chainId].finalisedClaimPeriodIndex))].permittedRevealTime;
        }
        return (chains[chainId].genesisLedger, chains[chainId].finalisedClaimPeriodIndex, chains[chainId].claimPeriodLength, chains[chainId].finalisedLedgerIndex, finalisedTimestamp, timeDiffAvg);
    }

    function getClaimPeriodIndexFinality(uint32 chainId, uint64 claimPeriodIndex) external view chainExists(chainId) returns (bool finality) {
        bytes32 locationHash =  keccak256(abi.encodePacked(chainId,claimPeriodIndex));
        return (finalisedClaimPeriods[locationHash].exists);
    }

    function proveClaimPeriodFinality(uint32 chainId, uint64 ledger, uint64 claimPeriodIndex, bytes32 claimPeriodHash) external chainExists(chainId) senderNotGovernanceBlocked returns (uint32 _chainId, uint64 _ledger, uint16 _numConfirmations, bytes32 _claimPeriodHash) {
        require(claimPeriodHash > 0x0, 'claimPeriodHash == 0x0');
        require(block.coinbase == msg.sender || block.coinbase == address(0x0100000000000000000000000000000000000000), 'invalid block.coinbase value');
        require(ledger == chains[chainId].finalisedLedgerIndex + chains[chainId].claimPeriodLength, 'invalid ledger');
        require(claimPeriodIndex == chains[chainId].finalisedClaimPeriodIndex, 'invalid claimPeriodIndex');
        require(block.timestamp > chains[chainId].finalisedTimestamp, 'block.timestamp <= chains[chainId].finalisedTimestamp');
        if (2*chains[chainId].timeDiffAvg < chains[chainId].timeDiffExpected) {
        	require(3*(block.timestamp-chains[chainId].finalisedTimestamp) >= 2*chains[chainId].timeDiffAvg, 'not enough time elapsed since prior finality');
        } else {
        	require(block.timestamp-chains[chainId].finalisedTimestamp+15 >= chains[chainId].timeDiffAvg, 'not enough time elapsed since prior finality');
        }
        bytes32 locationHash = keccak256(abi.encodePacked(chainId,claimPeriodIndex));
        require(!finalisedClaimPeriods[locationHash].proven, 'locationHash already finalised');
        if (claimPeriodIndex > 0) {
            bytes32 prevLocationHash =  keccak256(abi.encodePacked(chainId,claimPeriodIndex-1));
            require(finalisedClaimPeriods[prevLocationHash].proven, 'previous claim period not yet finalised');
        }
        if (proposedClaimPeriods[msg.sender][locationHash].exists) {
            require(block.timestamp >= proposedClaimPeriods[msg.sender][locationHash].permittedRevealTime, 'block.timestamp < proposedClaimPeriods[msg.sender][locationHash].permittedRevealTime');
            require(proposedClaimPeriods[msg.sender][locationHash].commitHash == keccak256(abi.encodePacked(msg.sender,claimPeriodHash)), 'proposedClaimPeriods[msg.sender][locationHash].commitHash != keccak256(abi.encodePacked(msg.sender,claimPeriodHash))');
        } else if (block.coinbase != msg.sender && block.coinbase == address(0x0100000000000000000000000000000000000000)) {
            claimPeriodHash = 0x0;
        }
        if (block.coinbase == msg.sender && block.coinbase != address(0x0100000000000000000000000000000000000000)) {
            if (!proposedClaimPeriods[msg.sender][locationHash].exists) {
                proposedClaimPeriods[msg.sender][locationHash] = HashExists(true, claimPeriodHash, block.timestamp, block.timestamp+chains[chainId].timeDiffAvg/2, 0x0, ledger, 0, false);
            } else {
                finalisedClaimPeriods[locationHash] = HashExists(true, 0x0, proposedClaimPeriods[msg.sender][locationHash].commitTime, block.timestamp, claimPeriodHash, ledger, 0, true);
                // Node checked claimPeriodHash, and it was valid
                uint256 currentRewardPeriod = getRewardPeriod();
                claimPeriodsMined[msg.sender][currentRewardPeriod] = claimPeriodsMined[msg.sender][currentRewardPeriod]+1;
                chains[chainId].finalisedClaimPeriodIndex = claimPeriodIndex+1;
                chains[chainId].finalisedLedgerIndex = ledger;
                uint256 timeDiffAvgUpdate = (chains[chainId].timeDiffAvg + (proposedClaimPeriods[msg.sender][locationHash].commitTime-chains[chainId].finalisedTimestamp))/2;
                if (timeDiffAvgUpdate > 2*chains[chainId].timeDiffExpected) {
                    chains[chainId].timeDiffAvg = 2*chains[chainId].timeDiffExpected;
                } else {
                    chains[chainId].timeDiffAvg = timeDiffAvgUpdate;
                }
                chains[chainId].finalisedTimestamp = proposedClaimPeriods[msg.sender][locationHash].commitTime;
            }
        }
        return (chainId, ledger-1, chains[chainId].numConfirmations, claimPeriodHash);
    }

    function getClaimPeriodFinality(bytes32 claimPeriodHash, uint32 chainId, uint64 claimPeriodIndex) private view chainExists(chainId) returns (bool finality) {
        bytes32 locationHash =  keccak256(abi.encodePacked(chainId,claimPeriodIndex));
        require(finalisedClaimPeriods[locationHash].exists, 'finalisedClaimPeriods[locationHash] does not exist');
        require(finalisedClaimPeriods[locationHash].revealHash == claimPeriodHash, 'invalid claimPeriodHash');
        return true;
    }

    // If ledger == payment's ledger -> return true
    function provePaymentFinality(uint32 chainId, bytes32 paymentHash, uint64 ledger, string memory txId) external chainExists(chainId) returns (uint32 _chainId, uint64 _ledger, uint64 finalisedLedgerIndex, bytes32 _paymentHash, string memory _txId) {
        bytes32 txIdHash = keccak256(abi.encodePacked(txId));
        require(!finalisedPayments[chainId][txIdHash].proven, 'txId already proven');
        require(ledger < chains[chainId].finalisedLedgerIndex, 'ledger >= chains[chainId].finalisedLedgerIndex');
        uint64 indexSearchRegion = chains[chainId].genesisLedger;
        if (chains[chainId].ledgerHistorySize > 0) {
            require(chains[chainId].finalisedLedgerIndex - chains[chainId].genesisLedger > chains[chainId].ledgerHistorySize, 'chains[chainId].finalisedLedgerIndex - chains[chainId].genesisLedger <= chains[chainId].ledgerHistorySize');
            require(chains[chainId].finalisedLedgerIndex - ledger < chains[chainId].ledgerHistorySize, 'chains[chainId].finalisedLedgerIndex - ledger >= chains[chainId].ledgerHistorySize');
            indexSearchRegion = chains[chainId].finalisedLedgerIndex - chains[chainId].ledgerHistorySize;
        }
        require(block.coinbase == msg.sender || block.coinbase == address(0x0100000000000000000000000000000000000000), 'invalid block.coinbase value');
        if (block.coinbase == msg.sender && block.coinbase != address(0x0100000000000000000000000000000000000000)) {
        	finalisedPayments[chainId][txIdHash] = HashExists(true, 0x0, 0, block.timestamp, paymentHash, ledger, indexSearchRegion, true);
        }
        return (chainId, ledger, chains[chainId].finalisedLedgerIndex, paymentHash, txId);
    }

    // If ledger < payment's ledger or payment does not exist within data-available region -> return true
    function disprovePaymentFinality(uint32 chainId, bytes32 paymentHash, uint64 ledger, string memory txId) external chainExists(chainId) returns (uint32 _chainId, uint64 _ledger, uint64 finalisedLedgerIndex, bytes32 _paymentHash, string memory _txId) {
        bytes32 txIdHash = keccak256(abi.encodePacked(txId));
        require(!finalisedPayments[chainId][txIdHash].proven, 'txId already proven');
        require(finalisedPayments[chainId][txIdHash].index < ledger, 'finalisedPayments[chainId][txIdHash].index >= ledger');
        require(ledger < chains[chainId].finalisedLedgerIndex, 'ledger >= chains[chainId].finalisedLedgerIndex');
        uint64 indexSearchRegion = chains[chainId].genesisLedger;
        if (chains[chainId].ledgerHistorySize > 0) {
            require(chains[chainId].finalisedLedgerIndex - chains[chainId].genesisLedger > chains[chainId].ledgerHistorySize, 'chains[chainId].finalisedLedgerIndex - chains[chainId].genesisLedger <= chains[chainId].ledgerHistorySize');
            require(chains[chainId].finalisedLedgerIndex - ledger < chains[chainId].ledgerHistorySize, 'chains[chainId].finalisedLedgerIndex - ledger >= chains[chainId].ledgerHistorySize');
            indexSearchRegion = chains[chainId].finalisedLedgerIndex - chains[chainId].ledgerHistorySize;
        }
        require(block.coinbase == msg.sender || block.coinbase == address(0x0100000000000000000000000000000000000000), 'invalid block.coinbase value');
        if (block.coinbase == msg.sender && block.coinbase != address(0x0100000000000000000000000000000000000000)) {
        	finalisedPayments[chainId][txIdHash] = HashExists(true, 0x0, 0, block.timestamp, paymentHash, ledger, indexSearchRegion, false);
        }
        return (chainId, ledger, chains[chainId].finalisedLedgerIndex, paymentHash, txId);
    }

    function getPaymentFinality(uint32 chainId, bytes32 txId, bytes32 sourceHash, bytes32 destinationHash, uint64 destinationTag, uint64 amount, bytes32 currencyHash) external view chainExists(chainId) returns (uint64 ledger, uint64 indexSearchRegion, bool finality) {
        require(finalisedPayments[chainId][txId].exists, 'txId does not exist');
        bytes32 paymentHash = keccak256(abi.encodePacked(
        							txId,
        							sourceHash,
        							destinationHash,
        							keccak256(abi.encode(destinationTag)),
        							keccak256(abi.encode(amount)),
                                    currencyHash));
    	require(finalisedPayments[chainId][txId].revealHash == paymentHash, 'invalid paymentHash');
    	return (finalisedPayments[chainId][txId].index, finalisedPayments[chainId][txId].indexSearchRegion, finalisedPayments[chainId][txId].proven);
    }
}