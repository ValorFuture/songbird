package genesis

import (
	"time"
	"github.com/ava-labs/avalanchego/utils/units"
)

var (
	costonGenesisConfigJSON = `{
		"networkID": 16,
		"allocations": [],
		"startTime": 0,
		"initialStakeDuration": 0,
		"initialStakeDurationOffset": 0,
		"initialStakedFunds": [],
		"initialStakers": [],
		"message": "coston"
	}`
	
	// CostonParams are the params used for local networks
	CostonParams = Params{
		TxFee:              units.MilliAvax,
		CreationTxFee:      10 * units.MilliAvax,
		UptimeRequirement:  .6, // 60%
		MinValidatorStake:  1 * units.Avax,
		MaxValidatorStake:  3 * units.MegaAvax,
		MinDelegatorStake:  1 * units.Avax,
		MinDelegationFee:   20000, // 2%
		MinStakeDuration:   24 * time.Hour,
		MaxStakeDuration:   365 * 24 * time.Hour,
		StakeMintingPeriod: 365 * 24 * time.Hour,
	}
)

const (
	CostonGenesis = `{
	    "config": {
	      "chainId": 16,
	      "homesteadBlock": 0,
	      "daoForkBlock": 0,
	      "daoForkSupport": true,
	      "eip150Block": 0,
	      "eip150Hash": "0x2086799aeebeae135c246c65021c82b4e15a2c451340993aacfd2751886514f0",
	      "eip155Block": 0,
	      "eip158Block": 0,
	      "byzantiumBlock": 0,
	      "constantinopleBlock": 0,
	      "petersburgBlock": 0
	    },
	    "nonce": "0x0",
	    "timestamp": "0x0",
	    "extraData": "0x00",
	    "gasLimit": "0x5f5e100",
	    "difficulty": "0x0",
	    "mixHash": "0x0000000000000000000000000000000000000000000000000000000000000000",
	    "coinbase": "0x0100000000000000000000000000000000000000",
	    "alloc": {
	      "1000000000000000000000000000000000000001": {
	        "balance": "0x0",
	        "code": "0x608060405234801561001057600080fd5b50600436106100a95760003560e01c8063bd18dae311610071578063bd18dae3146101ab578063e54caf92146101db578063ef2fa85f1461020b578063ef4c169e14610229578063efe7827214610247578063ff695ef014610277576100a9565b80631129753f146100ae57806313bb431c146100de5780632a2434a2146101115780636879c67b14610147578063a57d0e2514610178575b600080fd5b6100c860048036038101906100c39190611f15565b6102a7565b6040516100d591906128c5565b60405180910390f35b6100f860048036038101906100f39190612068565b610381565b6040516101089493929190612b29565b60405180910390f35b61012b60048036038101906101269190611fc7565b610693565b60405161013e9796959493929190612bd5565b60405180910390f35b610161600480360381019061015c9190611f3e565b61087a565b60405161016f9291906128e0565b60405180910390f35b610192600480360381019061018d919061216e565b610a0f565b6040516101a29493929190612b75565b60405180910390f35b6101c560048036038101906101c091906120f7565b611428565b6040516101d291906128c5565b60405180910390f35b6101f560048036038101906101f0919061202c565b611722565b60405161020291906128c5565b60405180910390f35b610213611840565b60405161022091906128aa565b60405180910390f35b610231611869565b60405161023e91906128c5565b60405180910390f35b610261600480360381019061025c9190611f15565b611ada565b60405161026e9190612bba565b60405180910390f35b610291600480360381019061028c9190611ff0565b611b37565b60405161029e91906128c5565b60405180910390f35b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610338576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161032f90612989565b60405180910390fd5b816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060019050919050565b6000806000606060011515600160008b63ffffffff1663ffffffff16815260200190815260200160002060000160009054906101000a900460ff161515146103fe576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103f590612b09565b60405180910390fd5b600085604051602001610411919061285d565b604051602081830303815290604052805190602001209050600015156003600083815260200190815260200160002060000160009054906101000a900460ff16151514610493576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161048a90612969565b60405180910390fd5b60006104a0898c8c611c6f565b90503373ffffffffffffffffffffffffffffffffffffffff164173ffffffffffffffffffffffffffffffffffffffff16148061051b575073010000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff164173ffffffffffffffffffffffffffffffffffffffff16145b61055a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161055190612ac9565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff164173ffffffffffffffffffffffffffffffffffffffff161480156105d5575073010000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff164173ffffffffffffffffffffffffffffffffffffffff1614155b15610643576040518060600160405280600115158152602001898152602001828152506003600084815260200190815260200160002060008201518160000160006101000a81548160ff02191690831515021790555060208201518160010155604082015181600201559050505b8a600160008d63ffffffff1663ffffffff16815260200190815260200160002060000160159054906101000a900467ffffffffffffffff1689899550955095509550505095509550955095915050565b600080600080600080600060011515600160008a63ffffffff1663ffffffff16815260200190815260200160002060000160009054906101000a900460ff16151514610714576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161070b90612b09565b60405180910390fd5b600160008963ffffffff1663ffffffff16815260200190815260200160002060000160019054906101000a900467ffffffffffffffff16600160008a63ffffffff1663ffffffff168152602001908152602001600020600001600d9054906101000a900467ffffffffffffffff16600160008b63ffffffff1663ffffffff16815260200190815260200160002060000160099054906101000a900461ffff16600160008c63ffffffff1663ffffffff168152602001908152602001600020600001600b9054906101000a900461ffff16600160008d63ffffffff1663ffffffff16815260200190815260200160002060000160159054906101000a900467ffffffffffffffff16600160008e63ffffffff1663ffffffff16815260200190815260200160002060010154600160008f63ffffffff1663ffffffff168152602001908152602001600020600301549650965096509650965096509650919395979092949650565b60008060011515600360008a815260200190815260200160002060000160009054906101000a900460ff161515146108e7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108de90612ae9565b60405180910390fd5b600088886040516020016108fb9190612bba565b604051602081830303815290604052805190602001208888886040516020016109249190612bba565b604051602081830303815290604052805190602001208860405160200161094b9190612bba565b60405160208183030381529060405280519060200120604051602001610976969594939291906127ed565b60405160208183030381529060405280519060200120905080600360008b815260200190815260200160002060010154146109e6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109dd90612909565b60405180910390fd5b6001600360008b8152602001908152602001600020600201549250925050965096945050505050565b60008060008060011515600160008a63ffffffff1663ffffffff16815260200190815260200160002060000160009054906101000a900460ff16151514610a8b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a8290612b09565b60405180910390fd5b600160008963ffffffff1663ffffffff16815260200190815260200160002060000160099054906101000a900461ffff1661ffff16600160008a63ffffffff1663ffffffff16815260200190815260200160002060000160159054906101000a900467ffffffffffffffff160167ffffffffffffffff168767ffffffffffffffff1614610b4d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b4490612a09565b60405180910390fd5b600160008963ffffffff1663ffffffff16815260200190815260200160002060000160099054906101000a900461ffff1661ffff166001870102600160008a63ffffffff1663ffffffff16815260200190815260200160002060000160019054906101000a900467ffffffffffffffff160167ffffffffffffffff168767ffffffffffffffff1614610c14576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c0b906129c9565b60405180910390fd5b600160008963ffffffff1663ffffffff168152602001908152602001600020600101544211610c78576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c6f90612a29565b60405180910390fd5b600160008963ffffffff1663ffffffff16815260200190815260200160002060020154600160008a63ffffffff1663ffffffff168152602001908152602001600020600301546002021015610d5b57600160008963ffffffff1663ffffffff16815260200190815260200160002060030154600202600160008a63ffffffff1663ffffffff1681526020019081526020016000206001015442036003021015610d56576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d4d90612929565b60405180910390fd5b610de8565b600160008963ffffffff1663ffffffff16815260200190815260200160002060030154600f600160008b63ffffffff1663ffffffff168152602001908152602001600020600101544203011015610de7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610dde90612929565b60405180910390fd5b5b600088604051602001610dfb9190612874565b6040516020818303038152906040528051906020012087604051602001610e22919061288f565b60405160208183030381529060405280519060200120604051602001610e499291906127c1565b604051602081830303815290604052805190602001209050600015156002600083815260200190815260200160002060000160009054906101000a900460ff16151514610ecb576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ec290612a69565b60405180910390fd5b60008767ffffffffffffffff161115610fc657600089604051602001610ef19190612874565b6040516020818303038152906040528051906020012060018903604051602001610f1b919061288f565b60405160208183030381529060405280519060200120604051602001610f429291906127c1565b604051602081830303815290604052805190602001209050600115156002600083815260200190815260200160002060000160009054906101000a900460ff16151514610fc4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fbb906129e9565b60405180910390fd5b505b3373ffffffffffffffffffffffffffffffffffffffff164173ffffffffffffffffffffffffffffffffffffffff16148061103f575073010000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff164173ffffffffffffffffffffffffffffffffffffffff16145b61107e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161107590612ac9565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff164173ffffffffffffffffffffffffffffffffffffffff161480156110f9575073010000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff164173ffffffffffffffffffffffffffffffffffffffff1614155b156113dd576001600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900467ffffffffffffffff1601600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506040518060600160405280600115158152602001878152602001428152506002600083815260200190815260200160002060008201518160000160006101000a81548160ff021916908315150217905550602082015181600101556040820151816002015590505060018701600160008b63ffffffff1663ffffffff168152602001908152602001600020600001600d6101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555087600160008b63ffffffff1663ffffffff16815260200190815260200160002060000160156101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060006002600160008c63ffffffff1663ffffffff168152602001908152602001600020600101544203600160008d63ffffffff1663ffffffff16815260200190815260200160002060030154018161130b57fe5b049050600160008b63ffffffff1663ffffffff1681526020019081526020016000206002015460020281111561138c57600160008b63ffffffff1663ffffffff16815260200190815260200160002060020154600202600160008c63ffffffff1663ffffffff168152602001908152602001600020600301819055506113b4565b80600160008c63ffffffff1663ffffffff168152602001908152602001600020600301819055505b42600160008c63ffffffff1663ffffffff16815260200190815260200160002060010181905550505b8860018903600160008c63ffffffff1663ffffffff168152602001908152602001600020600001600b9054906101000a900461ffff1688945094509450945050945094509450949050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146114b9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114b090612989565b60405180910390fd5b60001515600160008863ffffffff1663ffffffff16815260200190815260200160002060000160009054906101000a900460ff1615151461152f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611526906129a9565b60405180910390fd5b60008461ffff1611611576576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161156d90612a49565b60405180910390fd5b6040518061012001604052806001151581526020018667ffffffffffffffff1681526020018561ffff1681526020018461ffff168152602001600067ffffffffffffffff1681526020018667ffffffffffffffff1681526020014281526020018381526020016000815250600160008863ffffffff1663ffffffff16815260200190815260200160002060008201518160000160006101000a81548160ff02191690831515021790555060208201518160000160016101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060408201518160000160096101000a81548161ffff021916908361ffff160217905550606082015181600001600b6101000a81548161ffff021916908361ffff160217905550608082015181600001600d6101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060a08201518160000160156101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060c0820151816001015560e0820151816002015561010082015181600301559050506001905095945050505050565b600060011515600160008563ffffffff1663ffffffff16815260200190815260200160002060000160009054906101000a900460ff1615151461179a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161179190612b09565b60405180910390fd5b6000836040516020016117ad9190612874565b60405160208183030381529060405280519060200120836040516020016117d4919061288f565b604051602081830303815290604052805190602001206040516020016117fb9291906127c1565b6040516020818303038152906040528051906020012090506002600082815260200190815260200160002060000160009054906101000a900460ff1691505092915050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000801515600060149054906101000a900460ff161515146118c0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016118b790612a89565b60405180910390fd5b7310000000000000000000000000000000000000006000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040518061012001604052806001151581526020016303a38d8a67ffffffffffffffff168152602001601e61ffff168152602001600061ffff168152602001600067ffffffffffffffff1681526020016303a38d8a67ffffffffffffffff168152602001428152602001607881526020016000815250600160008063ffffffff16815260200190815260200160002060008201518160000160006101000a81548160ff02191690831515021790555060208201518160000160016101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060408201518160000160096101000a81548161ffff021916908361ffff160217905550606082015181600001600b6101000a81548161ffff021916908361ffff160217905550608082015181600001600d6101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060a08201518160000160156101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060c0820151816001015560e0820151816002015561010082015181600301559050506001600060146101000a81548160ff0219169083151502179055506001905090565b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900467ffffffffffffffff169050919050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611bc8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611bbf90612989565b60405180910390fd5b60011515600160008563ffffffff1663ffffffff16815260200190815260200160002060000160009054906101000a900460ff16151514611c3e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611c3590612b09565b60405180910390fd5b81600160008563ffffffff1663ffffffff168152602001908152602001600020600201819055506001905092915050565b600060011515600160008563ffffffff1663ffffffff16815260200190815260200160002060000160009054906101000a900460ff16151514611ce7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611cde90612b09565b60405180910390fd5b600083604051602001611cfa9190612874565b6040516020818303038152906040528051906020012083604051602001611d21919061288f565b60405160208183030381529060405280519060200120604051602001611d489291906127c1565b604051602081830303815290604052805190602001209050600115156002600083815260200190815260200160002060000160009054906101000a900460ff16151514611dca576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611dc190612949565b60405180910390fd5b84600260008381526020019081526020016000206001015414611e22576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611e1990612aa9565b60405180910390fd5b60026000828152602001908152602001600020600201549150509392505050565b600081359050611e5281612ded565b92915050565b600081359050611e6781612e04565b92915050565b600082601f830112611e7e57600080fd5b8135611e91611e8c82612c75565b612c44565b91508082526020830160208301858383011115611ead57600080fd5b611eb8838284612d50565b50505092915050565b600081359050611ed081612e1b565b92915050565b600081359050611ee581612e32565b92915050565b600081359050611efa81612e49565b92915050565b600081359050611f0f81612e60565b92915050565b600060208284031215611f2757600080fd5b6000611f3584828501611e43565b91505092915050565b60008060008060008060c08789031215611f5757600080fd5b6000611f6589828a01611e58565b9650506020611f7689828a01611f00565b9550506040611f8789828a01611e58565b9450506060611f9889828a01611e58565b9350506080611fa989828a01611f00565b92505060a0611fba89828a01611f00565b9150509295509295509295565b600060208284031215611fd957600080fd5b6000611fe784828501611eeb565b91505092915050565b6000806040838503121561200357600080fd5b600061201185828601611eeb565b925050602061202285828601611ed6565b9150509250929050565b6000806040838503121561203f57600080fd5b600061204d85828601611eeb565b925050602061205e85828601611f00565b9150509250929050565b600080600080600060a0868803121561208057600080fd5b600061208e88828901611eeb565b955050602061209f88828901611f00565b94505060406120b088828901611e58565b93505060606120c188828901611e58565b925050608086013567ffffffffffffffff8111156120de57600080fd5b6120ea88828901611e6d565b9150509295509295909350565b600080600080600060a0868803121561210f57600080fd5b600061211d88828901611eeb565b955050602061212e88828901611f00565b945050604061213f88828901611ec1565b935050606061215088828901611ec1565b925050608061216188828901611ed6565b9150509295509295909350565b6000806000806080858703121561218457600080fd5b600061219287828801611eeb565b94505060206121a387828801611f00565b93505060406121b487828801611f00565b92505060606121c587828801611e58565b91505092959194509250565b6121da81612ccc565b82525050565b6121e981612cde565b82525050565b6121f881612cea565b82525050565b61220f61220a82612cea565b612d92565b82525050565b600061222082612ca5565b61222a8185612cb0565b935061223a818560208601612d5f565b61224381612dc2565b840191505092915050565b600061225982612ca5565b6122638185612cc1565b9350612273818560208601612d5f565b80840191505092915050565b600061228c601383612cb0565b91507f696e76616c6964207061796d656e7448617368000000000000000000000000006000830152602082019050919050565b60006122cc602c83612cb0565b91507f6e6f7420656e6f7567682074696d6520656c61707365642073696e636520707260008301527f696f722066696e616c69747900000000000000000000000000000000000000006020830152604082019050919050565b6000612332603283612cb0565b91507f66696e616c69736564436c61696d506572696f64735b6c6f636174696f6e486160008301527f73685d20646f6573206e6f7420657869737400000000000000000000000000006020830152604082019050919050565b6000612398601383612cb0565b91507f7478496420616c72656164792070726f76656e000000000000000000000000006000830152602082019050919050565b60006123d8602083612cb0565b91507f6d73672e73656e64657220213d20676f7665726e616e6365436f6e74726163746000830152602082019050919050565b6000612418601683612cb0565b91507f636861696e496420616c726561647920657869737473000000000000000000006000830152602082019050919050565b6000612458601883612cb0565b91507f696e76616c696420636c61696d506572696f64496e64657800000000000000006000830152602082019050919050565b6000612498602783612cb0565b91507f70726576696f757320636c61696d20706572696f64206e6f742079657420666960008301527f6e616c69736564000000000000000000000000000000000000000000000000006020830152604082019050919050565b60006124fe600e83612cb0565b91507f696e76616c6964206c65646765720000000000000000000000000000000000006000830152602082019050919050565b600061253e603583612cb0565b91507f626c6f636b2e74696d657374616d70203c3d20636861696e735b636861696e4960008301527f645d2e66696e616c6973656454696d657374616d7000000000000000000000006020830152604082019050919050565b60006125a4601683612cb0565b91507f636c61696d506572696f644c656e677468203d3d2030000000000000000000006000830152602082019050919050565b60006125e4601e83612cb0565b91507f6c6f636174696f6e4861736820616c72656164792066696e616c6973656400006000830152602082019050919050565b6000612624601483612cb0565b91507f696e697469616c6973656420213d2066616c73650000000000000000000000006000830152602082019050919050565b6000612664601783612cb0565b91507f496e76616c696420636c61696d506572696f64486173680000000000000000006000830152602082019050919050565b60006126a4601c83612cb0565b91507f496e76616c696420626c6f636b2e636f696e626173652076616c7565000000006000830152602082019050919050565b60006126e4601383612cb0565b91507f7478496420646f6573206e6f74206578697374000000000000000000000000006000830152602082019050919050565b6000612724601683612cb0565b91507f636861696e496420646f6573206e6f74206578697374000000000000000000006000830152602082019050919050565b61276081612cf4565b82525050565b61276f81612d22565b82525050565b61277e81612d2c565b82525050565b61279561279082612d2c565b612d9c565b82525050565b6127a481612d3c565b82525050565b6127bb6127b682612d3c565b612dae565b82525050565b60006127cd82856121fe565b6020820191506127dd82846121fe565b6020820191508190509392505050565b60006127f982896121fe565b60208201915061280982886121fe565b60208201915061281982876121fe565b60208201915061282982866121fe565b60208201915061283982856121fe565b60208201915061284982846121fe565b602082019150819050979650505050505050565b6000612869828461224e565b915081905092915050565b60006128808284612784565b60048201915081905092915050565b600061289b82846127aa565b60088201915081905092915050565b60006020820190506128bf60008301846121d1565b92915050565b60006020820190506128da60008301846121e0565b92915050565b60006040820190506128f560008301856121e0565b6129026020830184612766565b9392505050565b600060208201905081810360008301526129228161227f565b9050919050565b60006020820190508181036000830152612942816122bf565b9050919050565b6000602082019050818103600083015261296281612325565b9050919050565b600060208201905081810360008301526129828161238b565b9050919050565b600060208201905081810360008301526129a2816123cb565b9050919050565b600060208201905081810360008301526129c28161240b565b9050919050565b600060208201905081810360008301526129e28161244b565b9050919050565b60006020820190508181036000830152612a028161248b565b9050919050565b60006020820190508181036000830152612a22816124f1565b9050919050565b60006020820190508181036000830152612a4281612531565b9050919050565b60006020820190508181036000830152612a6281612597565b9050919050565b60006020820190508181036000830152612a82816125d7565b9050919050565b60006020820190508181036000830152612aa281612617565b9050919050565b60006020820190508181036000830152612ac281612657565b9050919050565b60006020820190508181036000830152612ae281612697565b9050919050565b60006020820190508181036000830152612b02816126d7565b9050919050565b60006020820190508181036000830152612b2281612717565b9050919050565b6000608082019050612b3e6000830187612775565b612b4b602083018661279b565b612b5860408301856121ef565b8181036060830152612b6a8184612215565b905095945050505050565b6000608082019050612b8a6000830187612775565b612b97602083018661279b565b612ba46040830185612757565b612bb160608301846121ef565b95945050505050565b6000602082019050612bcf600083018461279b565b92915050565b600060e082019050612bea600083018a61279b565b612bf7602083018961279b565b612c046040830188612757565b612c116060830187612757565b612c1e608083018661279b565b612c2b60a0830185612766565b612c3860c0830184612766565b98975050505050505050565b6000604051905081810181811067ffffffffffffffff82111715612c6b57612c6a612dc0565b5b8060405250919050565b600067ffffffffffffffff821115612c9057612c8f612dc0565b5b601f19601f8301169050602081019050919050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b6000612cd782612d02565b9050919050565b60008115159050919050565b6000819050919050565b600061ffff82169050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600063ffffffff82169050919050565b600067ffffffffffffffff82169050919050565b82818337600083830152505050565b60005b83811015612d7d578082015181840152602081019050612d62565b83811115612d8c576000848401525b50505050565b6000819050919050565b6000612da782612de0565b9050919050565b6000612db982612dd3565b9050919050565bfe5b6000601f19601f8301169050919050565b60008160c01b9050919050565b60008160e01b9050919050565b612df681612ccc565b8114612e0157600080fd5b50565b612e0d81612cea565b8114612e1857600080fd5b50565b612e2481612cf4565b8114612e2f57600080fd5b50565b612e3b81612d22565b8114612e4657600080fd5b50565b612e5281612d2c565b8114612e5d57600080fd5b50565b612e6981612d3c565b8114612e7457600080fd5b5056fea26469706673582212204e325d3656c065ae6015a20c40becc84b1d5405bbd84f0cfa85d94df01ad440664736f6c63430007030033"
	      },
	      "1000000000000000000000000000000000000002": {
	        "balance": "0x0",
	        "code": "0x608060405234801561001057600080fd5b50600436106100365760003560e01c80637fec8d381461003b5780638be2fb8614610059575b600080fd5b610043610077565b6040516100509190610154565b60405180910390f35b6100616100ca565b60405161006e919061018f565b60405180910390f35b6000805443116100bc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016100b39061016f565b60405180910390fd5b436000819055506001905090565b60005481565b6100d9816101bb565b82525050565b60006100ec6025836101aa565b91507f626c6f636b2e6e756d626572203c3d2073797374656d4c61737454726967676560008301527f72656441740000000000000000000000000000000000000000000000000000006020830152604082019050919050565b61014e816101c7565b82525050565b600060208201905061016960008301846100d0565b92915050565b60006020820190508181036000830152610188816100df565b9050919050565b60006020820190506101a46000830184610145565b92915050565b600082825260208201905092915050565b60008115159050919050565b600081905091905056fea264697066735822122027763a81724f350e677ad02735ad09a2664074ff366eee639087afaba88b66f464736f6c63430007030033"
	      },
	      "9b200C53C8635a3d5972e1Dc5a7661713143987C": {
	        "balance": "0x314dc6448d9338c15B0a00000000"
	      },
	      "B35ABD91a7Da229a1a844229A652C3b34b59C59A": {
	        "balance": "0x314dc6448d9338c15B0a00000000"
	      },
	      "c783df8a850f42e7F7e57013759C285caa701eB6": {
	        "balance": "0x314dc6448d9338c15B0a00000000"
	      },
	      "eAD9C93b79Ae7C1591b1FB5323BD777E86e150d4": {
	        "balance": "0x314dc6448d9338c15B0a00000000"
	      },
	      "E5904695748fe4A84b40b3fc79De2277660BD1D3": {
	        "balance": "0x314dc6448d9338c15B0a00000000"
	      },
	      "92561F28Ec438Ee9831D00D1D59fbDC981b762b2": {
	        "balance": "0x314dc6448d9338c15B0a00000000"
	      },
	      "2fFd013AaA7B5a7DA93336C2251075202b33FB2B": {
	        "balance": "0x314dc6448d9338c15B0a00000000"
	      },
	      "9FC9C2DfBA3b6cF204C37a5F690619772b926e39": {
	        "balance": "0x314dc6448d9338c15B0a00000000"
	      },
	      "FbC51a9582D031f2ceaaD3959256596C5D3a5468": {
	        "balance": "0x314dc6448d9338c15B0a00000000"
	      },
	      "84Fae3d3Cba24A97817b2a18c2421d462dbBCe9f": {
	        "balance": "0x314dc6448d9338c15B0a00000000"
	      },
	      "fa3bdc8709226da0da13a4d904c8b66f16c3c8ba": {
	        "balance": "0x314dc6448d9338c15B0a00000000"
	      },
	      "6c365935ca8710200c7595f0a72eb6023a7706cd": {
	        "balance": "0x314dc6448d9338c15B0a00000000"
	      }
	    },
	    "number": "0x0",
	    "gasUsed": "0x0",
	    "parentHash": "0x0000000000000000000000000000000000000000000000000000000000000000"
	  }`
)