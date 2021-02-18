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
	        "code": "0x608060405234801561001057600080fd5b50600436106100a95760003560e01c8063bd18dae311610071578063bd18dae3146101ab578063e54caf92146101db578063ef2fa85f1461020b578063ef4c169e14610229578063efe7827214610247578063ff695ef014610277576100a9565b80631129753f146100ae57806313bb431c146100de5780632a2434a2146101115780636879c67b14610147578063a57d0e2514610178575b600080fd5b6100c860048036038101906100c39190611ff1565b6102a7565b6040516100d591906129e1565b60405180910390f35b6100f860048036038101906100f39190612144565b610381565b6040516101089493929190612c65565b60405180910390f35b61012b600480360381019061012691906120a3565b610701565b60405161013e9796959493929190612d11565b60405180910390f35b610161600480360381019061015c919061201a565b6108e8565b60405161016f9291906129fc565b60405180910390f35b610192600480360381019061018d919061224a565b610a7d565b6040516101a29493929190612cb1565b60405180910390f35b6101c560048036038101906101c091906121d3565b611504565b6040516101d291906129e1565b60405180910390f35b6101f560048036038101906101f09190612108565b6117fe565b60405161020291906129e1565b60405180910390f35b61021361191c565b60405161022091906129c6565b60405180910390f35b610231611945565b60405161023e91906129e1565b60405180910390f35b610261600480360381019061025c9190611ff1565b611bb6565b60405161026e9190612cf6565b60405180910390f35b610291600480360381019061028c91906120cc565b611c13565b60405161029e91906129e1565b60405180910390f35b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610338576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161032f90612aa5565b60405180910390fd5b816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060019050919050565b600080600060603273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146103f6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103ed90612b45565b60405180910390fd5b60011515600160008b63ffffffff1663ffffffff16815260200190815260200160002060000160009054906101000a900460ff1615151461046c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161046390612c45565b60405180910390fd5b60008560405160200161047f9190612979565b604051602081830303815290604052805190602001209050600015156003600083815260200190815260200160002060000160009054906101000a900460ff16151514610501576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104f890612a85565b60405180910390fd5b600061050e898c8c611d4b565b90503373ffffffffffffffffffffffffffffffffffffffff164173ffffffffffffffffffffffffffffffffffffffff161480610589575073010000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff164173ffffffffffffffffffffffffffffffffffffffff16145b6105c8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105bf90612c05565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff164173ffffffffffffffffffffffffffffffffffffffff16148015610643575073010000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff164173ffffffffffffffffffffffffffffffffffffffff1614155b156106b1576040518060600160405280600115158152602001898152602001828152506003600084815260200190815260200160002060008201518160000160006101000a81548160ff02191690831515021790555060208201518160010155604082015181600201559050505b8a600160008d63ffffffff1663ffffffff16815260200190815260200160002060000160159054906101000a900467ffffffffffffffff1689899550955095509550505095509550955095915050565b600080600080600080600060011515600160008a63ffffffff1663ffffffff16815260200190815260200160002060000160009054906101000a900460ff16151514610782576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161077990612c45565b60405180910390fd5b600160008963ffffffff1663ffffffff16815260200190815260200160002060000160019054906101000a900467ffffffffffffffff16600160008a63ffffffff1663ffffffff168152602001908152602001600020600001600d9054906101000a900467ffffffffffffffff16600160008b63ffffffff1663ffffffff16815260200190815260200160002060000160099054906101000a900461ffff16600160008c63ffffffff1663ffffffff168152602001908152602001600020600001600b9054906101000a900461ffff16600160008d63ffffffff1663ffffffff16815260200190815260200160002060000160159054906101000a900467ffffffffffffffff16600160008e63ffffffff1663ffffffff16815260200190815260200160002060010154600160008f63ffffffff1663ffffffff168152602001908152602001600020600301549650965096509650965096509650919395979092949650565b60008060011515600360008a815260200190815260200160002060000160009054906101000a900460ff16151514610955576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161094c90612c25565b60405180910390fd5b600088886040516020016109699190612cf6565b604051602081830303815290604052805190602001208888886040516020016109929190612cf6565b60405160208183030381529060405280519060200120886040516020016109b99190612cf6565b604051602081830303815290604052805190602001206040516020016109e496959493929190612909565b60405160208183030381529060405280519060200120905080600360008b81526020019081526020016000206001015414610a54576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a4b90612a25565b60405180910390fd5b6001600360008b8152602001908152602001600020600201549250925050965096945050505050565b6000806000803273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610af1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ae890612b45565b60405180910390fd5b60011515600160008a63ffffffff1663ffffffff16815260200190815260200160002060000160009054906101000a900460ff16151514610b67576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b5e90612c45565b60405180910390fd5b600160008963ffffffff1663ffffffff16815260200190815260200160002060000160099054906101000a900461ffff1661ffff16600160008a63ffffffff1663ffffffff16815260200190815260200160002060000160159054906101000a900467ffffffffffffffff160167ffffffffffffffff168767ffffffffffffffff1614610c29576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c2090612b25565b60405180910390fd5b600160008963ffffffff1663ffffffff16815260200190815260200160002060000160099054906101000a900461ffff1661ffff166001870102600160008a63ffffffff1663ffffffff16815260200190815260200160002060000160019054906101000a900467ffffffffffffffff160167ffffffffffffffff168767ffffffffffffffff1614610cf0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ce790612ae5565b60405180910390fd5b600160008963ffffffff1663ffffffff168152602001908152602001600020600101544211610d54576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d4b90612b65565b60405180910390fd5b600160008963ffffffff1663ffffffff16815260200190815260200160002060020154600160008a63ffffffff1663ffffffff168152602001908152602001600020600301546002021015610e3757600160008963ffffffff1663ffffffff16815260200190815260200160002060030154600202600160008a63ffffffff1663ffffffff1681526020019081526020016000206001015442036003021015610e32576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e2990612a45565b60405180910390fd5b610ec4565b600160008963ffffffff1663ffffffff16815260200190815260200160002060030154600f600160008b63ffffffff1663ffffffff168152602001908152602001600020600101544203011015610ec3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610eba90612a45565b60405180910390fd5b5b600088604051602001610ed79190612990565b6040516020818303038152906040528051906020012087604051602001610efe91906129ab565b60405160208183030381529060405280519060200120604051602001610f259291906128dd565b604051602081830303815290604052805190602001209050600015156002600083815260200190815260200160002060000160009054906101000a900460ff16151514610fa7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f9e90612ba5565b60405180910390fd5b60008767ffffffffffffffff1611156110a257600089604051602001610fcd9190612990565b6040516020818303038152906040528051906020012060018903604051602001610ff791906129ab565b6040516020818303038152906040528051906020012060405160200161101e9291906128dd565b604051602081830303815290604052805190602001209050600115156002600083815260200190815260200160002060000160009054906101000a900460ff161515146110a0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161109790612b05565b60405180910390fd5b505b3373ffffffffffffffffffffffffffffffffffffffff164173ffffffffffffffffffffffffffffffffffffffff16148061111b575073010000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff164173ffffffffffffffffffffffffffffffffffffffff16145b61115a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161115190612c05565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff164173ffffffffffffffffffffffffffffffffffffffff161480156111d5575073010000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff164173ffffffffffffffffffffffffffffffffffffffff1614155b156114b9576001600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900467ffffffffffffffff1601600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506040518060600160405280600115158152602001878152602001428152506002600083815260200190815260200160002060008201518160000160006101000a81548160ff021916908315150217905550602082015181600101556040820151816002015590505060018701600160008b63ffffffff1663ffffffff168152602001908152602001600020600001600d6101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555087600160008b63ffffffff1663ffffffff16815260200190815260200160002060000160156101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060006002600160008c63ffffffff1663ffffffff168152602001908152602001600020600101544203600160008d63ffffffff1663ffffffff1681526020019081526020016000206003015401816113e757fe5b049050600160008b63ffffffff1663ffffffff1681526020019081526020016000206002015460020281111561146857600160008b63ffffffff1663ffffffff16815260200190815260200160002060020154600202600160008c63ffffffff1663ffffffff16815260200190815260200160002060030181905550611490565b80600160008c63ffffffff1663ffffffff168152602001908152602001600020600301819055505b42600160008c63ffffffff1663ffffffff16815260200190815260200160002060010181905550505b8860018903600160008c63ffffffff1663ffffffff168152602001908152602001600020600001600b9054906101000a900461ffff1688945094509450945050945094509450949050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611595576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161158c90612aa5565b60405180910390fd5b60001515600160008863ffffffff1663ffffffff16815260200190815260200160002060000160009054906101000a900460ff1615151461160b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161160290612ac5565b60405180910390fd5b60008461ffff1611611652576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161164990612b85565b60405180910390fd5b6040518061012001604052806001151581526020018667ffffffffffffffff1681526020018561ffff1681526020018461ffff168152602001600067ffffffffffffffff1681526020018667ffffffffffffffff1681526020014281526020018381526020016000815250600160008863ffffffff1663ffffffff16815260200190815260200160002060008201518160000160006101000a81548160ff02191690831515021790555060208201518160000160016101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060408201518160000160096101000a81548161ffff021916908361ffff160217905550606082015181600001600b6101000a81548161ffff021916908361ffff160217905550608082015181600001600d6101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060a08201518160000160156101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060c0820151816001015560e0820151816002015561010082015181600301559050506001905095945050505050565b600060011515600160008563ffffffff1663ffffffff16815260200190815260200160002060000160009054906101000a900460ff16151514611876576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161186d90612c45565b60405180910390fd5b6000836040516020016118899190612990565b60405160208183030381529060405280519060200120836040516020016118b091906129ab565b604051602081830303815290604052805190602001206040516020016118d79291906128dd565b6040516020818303038152906040528051906020012090506002600082815260200190815260200160002060000160009054906101000a900460ff1691505092915050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000801515600060149054906101000a900460ff1615151461199c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161199390612bc5565b60405180910390fd5b7310000000000000000000000000000000000000006000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040518061012001604052806001151581526020016303a38d8a67ffffffffffffffff168152602001601e61ffff168152602001600061ffff168152602001600067ffffffffffffffff1681526020016303a38d8a67ffffffffffffffff168152602001428152602001607881526020016000815250600160008063ffffffff16815260200190815260200160002060008201518160000160006101000a81548160ff02191690831515021790555060208201518160000160016101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060408201518160000160096101000a81548161ffff021916908361ffff160217905550606082015181600001600b6101000a81548161ffff021916908361ffff160217905550608082015181600001600d6101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060a08201518160000160156101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060c0820151816001015560e0820151816002015561010082015181600301559050506001600060146101000a81548160ff0219169083151502179055506001905090565b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900467ffffffffffffffff169050919050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611ca4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611c9b90612aa5565b60405180910390fd5b60011515600160008563ffffffff1663ffffffff16815260200190815260200160002060000160009054906101000a900460ff16151514611d1a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611d1190612c45565b60405180910390fd5b81600160008563ffffffff1663ffffffff168152602001908152602001600020600201819055506001905092915050565b600060011515600160008563ffffffff1663ffffffff16815260200190815260200160002060000160009054906101000a900460ff16151514611dc3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611dba90612c45565b60405180910390fd5b600083604051602001611dd69190612990565b6040516020818303038152906040528051906020012083604051602001611dfd91906129ab565b60405160208183030381529060405280519060200120604051602001611e249291906128dd565b604051602081830303815290604052805190602001209050600115156002600083815260200190815260200160002060000160009054906101000a900460ff16151514611ea6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611e9d90612a65565b60405180910390fd5b84600260008381526020019081526020016000206001015414611efe576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611ef590612be5565b60405180910390fd5b60026000828152602001908152602001600020600201549150509392505050565b600081359050611f2e81612f29565b92915050565b600081359050611f4381612f40565b92915050565b600082601f830112611f5a57600080fd5b8135611f6d611f6882612db1565b612d80565b91508082526020830160208301858383011115611f8957600080fd5b611f94838284612e8c565b50505092915050565b600081359050611fac81612f57565b92915050565b600081359050611fc181612f6e565b92915050565b600081359050611fd681612f85565b92915050565b600081359050611feb81612f9c565b92915050565b60006020828403121561200357600080fd5b600061201184828501611f1f565b91505092915050565b60008060008060008060c0878903121561203357600080fd5b600061204189828a01611f34565b965050602061205289828a01611fdc565b955050604061206389828a01611f34565b945050606061207489828a01611f34565b935050608061208589828a01611fdc565b92505060a061209689828a01611fdc565b9150509295509295509295565b6000602082840312156120b557600080fd5b60006120c384828501611fc7565b91505092915050565b600080604083850312156120df57600080fd5b60006120ed85828601611fc7565b92505060206120fe85828601611fb2565b9150509250929050565b6000806040838503121561211b57600080fd5b600061212985828601611fc7565b925050602061213a85828601611fdc565b9150509250929050565b600080600080600060a0868803121561215c57600080fd5b600061216a88828901611fc7565b955050602061217b88828901611fdc565b945050604061218c88828901611f34565b935050606061219d88828901611f34565b925050608086013567ffffffffffffffff8111156121ba57600080fd5b6121c688828901611f49565b9150509295509295909350565b600080600080600060a086880312156121eb57600080fd5b60006121f988828901611fc7565b955050602061220a88828901611fdc565b945050604061221b88828901611f9d565b935050606061222c88828901611f9d565b925050608061223d88828901611fb2565b9150509295509295909350565b6000806000806080858703121561226057600080fd5b600061226e87828801611fc7565b945050602061227f87828801611fdc565b935050604061229087828801611fdc565b92505060606122a187828801611f34565b91505092959194509250565b6122b681612e08565b82525050565b6122c581612e1a565b82525050565b6122d481612e26565b82525050565b6122eb6122e682612e26565b612ece565b82525050565b60006122fc82612de1565b6123068185612dec565b9350612316818560208601612e9b565b61231f81612efe565b840191505092915050565b600061233582612de1565b61233f8185612dfd565b935061234f818560208601612e9b565b80840191505092915050565b6000612368601383612dec565b91507f696e76616c6964207061796d656e7448617368000000000000000000000000006000830152602082019050919050565b60006123a8602c83612dec565b91507f6e6f7420656e6f7567682074696d6520656c61707365642073696e636520707260008301527f696f722066696e616c69747900000000000000000000000000000000000000006020830152604082019050919050565b600061240e603283612dec565b91507f66696e616c69736564436c61696d506572696f64735b6c6f636174696f6e486160008301527f73685d20646f6573206e6f7420657869737400000000000000000000000000006020830152604082019050919050565b6000612474601383612dec565b91507f7478496420616c72656164792070726f76656e000000000000000000000000006000830152602082019050919050565b60006124b4602083612dec565b91507f6d73672e73656e64657220213d20676f7665726e616e6365436f6e74726163746000830152602082019050919050565b60006124f4601683612dec565b91507f636861696e496420616c726561647920657869737473000000000000000000006000830152602082019050919050565b6000612534601883612dec565b91507f696e76616c696420636c61696d506572696f64496e64657800000000000000006000830152602082019050919050565b6000612574602783612dec565b91507f70726576696f757320636c61696d20706572696f64206e6f742079657420666960008301527f6e616c69736564000000000000000000000000000000000000000000000000006020830152604082019050919050565b60006125da600e83612dec565b91507f696e76616c6964206c65646765720000000000000000000000000000000000006000830152602082019050919050565b600061261a601783612dec565b91507f6d73672e73656e64657220213d2074782e6f726967696e0000000000000000006000830152602082019050919050565b600061265a603583612dec565b91507f626c6f636b2e74696d657374616d70203c3d20636861696e735b636861696e4960008301527f645d2e66696e616c6973656454696d657374616d7000000000000000000000006020830152604082019050919050565b60006126c0601683612dec565b91507f636c61696d506572696f644c656e677468203d3d2030000000000000000000006000830152602082019050919050565b6000612700601e83612dec565b91507f6c6f636174696f6e4861736820616c72656164792066696e616c6973656400006000830152602082019050919050565b6000612740601483612dec565b91507f696e697469616c6973656420213d2066616c73650000000000000000000000006000830152602082019050919050565b6000612780601783612dec565b91507f496e76616c696420636c61696d506572696f64486173680000000000000000006000830152602082019050919050565b60006127c0601c83612dec565b91507f496e76616c696420626c6f636b2e636f696e626173652076616c7565000000006000830152602082019050919050565b6000612800601383612dec565b91507f7478496420646f6573206e6f74206578697374000000000000000000000000006000830152602082019050919050565b6000612840601683612dec565b91507f636861696e496420646f6573206e6f74206578697374000000000000000000006000830152602082019050919050565b61287c81612e30565b82525050565b61288b81612e5e565b82525050565b61289a81612e68565b82525050565b6128b16128ac82612e68565b612ed8565b82525050565b6128c081612e78565b82525050565b6128d76128d282612e78565b612eea565b82525050565b60006128e982856122da565b6020820191506128f982846122da565b6020820191508190509392505050565b600061291582896122da565b60208201915061292582886122da565b60208201915061293582876122da565b60208201915061294582866122da565b60208201915061295582856122da565b60208201915061296582846122da565b602082019150819050979650505050505050565b6000612985828461232a565b915081905092915050565b600061299c82846128a0565b60048201915081905092915050565b60006129b782846128c6565b60088201915081905092915050565b60006020820190506129db60008301846122ad565b92915050565b60006020820190506129f660008301846122bc565b92915050565b6000604082019050612a1160008301856122bc565b612a1e6020830184612882565b9392505050565b60006020820190508181036000830152612a3e8161235b565b9050919050565b60006020820190508181036000830152612a5e8161239b565b9050919050565b60006020820190508181036000830152612a7e81612401565b9050919050565b60006020820190508181036000830152612a9e81612467565b9050919050565b60006020820190508181036000830152612abe816124a7565b9050919050565b60006020820190508181036000830152612ade816124e7565b9050919050565b60006020820190508181036000830152612afe81612527565b9050919050565b60006020820190508181036000830152612b1e81612567565b9050919050565b60006020820190508181036000830152612b3e816125cd565b9050919050565b60006020820190508181036000830152612b5e8161260d565b9050919050565b60006020820190508181036000830152612b7e8161264d565b9050919050565b60006020820190508181036000830152612b9e816126b3565b9050919050565b60006020820190508181036000830152612bbe816126f3565b9050919050565b60006020820190508181036000830152612bde81612733565b9050919050565b60006020820190508181036000830152612bfe81612773565b9050919050565b60006020820190508181036000830152612c1e816127b3565b9050919050565b60006020820190508181036000830152612c3e816127f3565b9050919050565b60006020820190508181036000830152612c5e81612833565b9050919050565b6000608082019050612c7a6000830187612891565b612c8760208301866128b7565b612c9460408301856122cb565b8181036060830152612ca681846122f1565b905095945050505050565b6000608082019050612cc66000830187612891565b612cd360208301866128b7565b612ce06040830185612873565b612ced60608301846122cb565b95945050505050565b6000602082019050612d0b60008301846128b7565b92915050565b600060e082019050612d26600083018a6128b7565b612d3360208301896128b7565b612d406040830188612873565b612d4d6060830187612873565b612d5a60808301866128b7565b612d6760a0830185612882565b612d7460c0830184612882565b98975050505050505050565b6000604051905081810181811067ffffffffffffffff82111715612da757612da6612efc565b5b8060405250919050565b600067ffffffffffffffff821115612dcc57612dcb612efc565b5b601f19601f8301169050602081019050919050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b6000612e1382612e3e565b9050919050565b60008115159050919050565b6000819050919050565b600061ffff82169050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600063ffffffff82169050919050565b600067ffffffffffffffff82169050919050565b82818337600083830152505050565b60005b83811015612eb9578082015181840152602081019050612e9e565b83811115612ec8576000848401525b50505050565b6000819050919050565b6000612ee382612f1c565b9050919050565b6000612ef582612f0f565b9050919050565bfe5b6000601f19601f8301169050919050565b60008160c01b9050919050565b60008160e01b9050919050565b612f3281612e08565b8114612f3d57600080fd5b50565b612f4981612e26565b8114612f5457600080fd5b50565b612f6081612e30565b8114612f6b57600080fd5b50565b612f7781612e5e565b8114612f8257600080fd5b50565b612f8e81612e68565b8114612f9957600080fd5b50565b612fa581612e78565b8114612fb057600080fd5b5056fea26469706673582212208eccc66cb15dcc6fbc62fc414802e16e79fdaff14c1aea0d6d50a90dc33bcafa64736f6c63430007030033"
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
	      }
	    },
	    "number": "0x0",
	    "gasUsed": "0x0",
	    "parentHash": "0x0000000000000000000000000000000000000000000000000000000000000000"
	  }`
)