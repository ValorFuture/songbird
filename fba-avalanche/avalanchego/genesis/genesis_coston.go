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
	        "code": "0x608060405234801561001057600080fd5b50600436106100cf5760003560e01c8063bd18dae31161008c578063ef2fa85f11610066578063ef2fa85f1461026b578063ef4c169e14610289578063efe78272146102a7578063f6f56af7146102d7576100cf565b8063bd18dae3146101d4578063d02a133514610204578063dab99b501461023a576100cf565b806307003bb4146100d4578063093370c0146100f25780631129753f146101225780632e73b5c8146101525780635465dfc414610182578063b172b222146101b6575b600080fd5b6100dc610309565b6040516100e991906127dd565b60405180910390f35b61010c60048036038101906101079190612049565b61031c565b60405161011991906127dd565b60405180910390f35b61013c60048036038101906101379190611f7f565b61043a565b60405161014991906127dd565b60405180910390f35b61016c6004803603810190610167919061200d565b610514565b60405161017991906127dd565b60405180910390f35b61019c6004803603810190610197919061219d565b61064c565b6040516101ad959493929190612a78565b60405180910390f35b6101be610fba565b6040516101cb91906127c2565b60405180910390f35b6101ee60048036038101906101e99190612126565b610fde565b6040516101fb91906127dd565b60405180910390f35b61021e60048036038101906102199190611fe4565b611291565b6040516102319796959493929190612ae6565b60405180910390f35b610254600480360381019061024f9190611fa8565b611478565b6040516102629291906127f8565b60405180910390f35b610273611561565b60405161028091906127c2565b60405180910390f35b61029161158a565b60405161029e91906127dd565b60405180910390f35b6102c160048036038101906102bc9190611f56565b6117fb565b6040516102ce9190612acb565b60405180910390f35b6102f160048036038101906102ec9190612085565b611858565b60405161030093929190612a41565b60405180910390f35b600060149054906101000a900460ff1681565b600060011515600160008563ffffffff1663ffffffff16815260200190815260200160002060000160009054906101000a900460ff16151514610394576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161038b90612a21565b60405180910390fd5b6000836040516020016103a7919061278c565b60405160208183030381529060405280519060200120836040516020016103ce91906127a7565b604051602081830303815290604052805190602001206040516020016103f5929190612760565b6040516020818303038152906040528051906020012090506002600082815260200190815260200160002060000160009054906101000a900460ff1691505092915050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146104cb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104c2906128c1565b60405180910390fd5b816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060019050919050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146105a5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161059c906128c1565b60405180910390fd5b60011515600160008563ffffffff1663ffffffff16815260200190815260200160002060000160009054906101000a900460ff1615151461061b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161061290612a21565b60405180910390fd5b81600160008563ffffffff1663ffffffff168152602001908152602001600020600201819055506001905092915050565b60008060008060003273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146106c2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106b990612981565b60405180910390fd5b60011515600160008b63ffffffff1663ffffffff16815260200190815260200160002060000160009054906101000a900460ff16151514610738576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161072f90612a21565b60405180910390fd5b600160008a63ffffffff1663ffffffff16815260200190815260200160002060000160099054906101000a900461ffff1661ffff16600160008b63ffffffff1663ffffffff16815260200190815260200160002060000160159054906101000a900467ffffffffffffffff160167ffffffffffffffff168867ffffffffffffffff16146107fa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107f190612961565b60405180910390fd5b600160008a63ffffffff1663ffffffff16815260200190815260200160002060000160099054906101000a900461ffff1661ffff166001880102600160008b63ffffffff1663ffffffff16815260200190815260200160002060000160019054906101000a900467ffffffffffffffff160167ffffffffffffffff168867ffffffffffffffff16146108c1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108b890612901565b60405180910390fd5b600160008a63ffffffff1663ffffffff16815260200190815260200160002060030154600202600160008b63ffffffff1663ffffffff1681526020019081526020016000206001015442036003021015610950576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161094790612861565b60405180910390fd5b600089604051602001610963919061278c565b604051602081830303815290604052805190602001208860405160200161098a91906127a7565b604051602081830303815290604052805190602001206040516020016109b1929190612760565b604051602081830303815290604052805190602001209050600015156002600083815260200190815260200160002060000160009054906101000a900460ff16151514610a33576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a2a906129a1565b60405180910390fd5b60008867ffffffffffffffff161115610b2e5760008a604051602001610a59919061278c565b6040516020818303038152906040528051906020012060018a03604051602001610a8391906127a7565b60405160208183030381529060405280519060200120604051602001610aaa929190612760565b604051602081830303815290604052805190602001209050600115156002600083815260200190815260200160002060000160009054906101000a900460ff16151514610b2c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b2390612921565b60405180910390fd5b505b3373ffffffffffffffffffffffffffffffffffffffff164173ffffffffffffffffffffffffffffffffffffffff161480610ba7575073010000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff164173ffffffffffffffffffffffffffffffffffffffff16145b610be6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bdd906129e1565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff164173ffffffffffffffffffffffffffffffffffffffff16148015610c61575073010000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff164173ffffffffffffffffffffffffffffffffffffffff1614155b15610f3f576001600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900467ffffffffffffffff1601600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506040518060600160405280600115158152602001888152602001428152506002600083815260200190815260200160002060008201518160000160006101000a81548160ff021916908315150217905550602082015181600101556040820151816002015590505060018801600160008c63ffffffff1663ffffffff168152602001908152602001600020600001600d6101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555088600160008c63ffffffff1663ffffffff16815260200190815260200160002060000160156101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060006002600160008d63ffffffff1663ffffffff168152602001908152602001600020600101544203600160008e63ffffffff1663ffffffff168152602001908152602001600020600301540181610e7357fe5b049050600160008c63ffffffff1663ffffffff16815260200190815260200160002060020154811115610eee57600160008c63ffffffff1663ffffffff16815260200190815260200160002060020154600160008d63ffffffff1663ffffffff16815260200190815260200160002060030181905550610f16565b80600160008d63ffffffff1663ffffffff168152602001908152602001600020600301819055505b42600160008d63ffffffff1663ffffffff16815260200190815260200160002060010181905550505b8989600160008d63ffffffff1663ffffffff16815260200190815260200160002060000160099054906101000a900461ffff16600160008e63ffffffff1663ffffffff168152602001908152602001600020600001600b9054906101000a900461ffff168a9550955095509550955050945094509450945094565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461106f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611066906128c1565b60405180910390fd5b60001515600160008863ffffffff1663ffffffff16815260200190815260200160002060000160009054906101000a900460ff161515146110e5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110dc906128e1565b60405180910390fd5b6040518061012001604052806001151581526020018667ffffffffffffffff1681526020018561ffff1681526020018461ffff168152602001600067ffffffffffffffff1681526020018667ffffffffffffffff1681526020014281526020018381526020016000815250600160008863ffffffff1663ffffffff16815260200190815260200160002060008201518160000160006101000a81548160ff02191690831515021790555060208201518160000160016101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060408201518160000160096101000a81548161ffff021916908361ffff160217905550606082015181600001600b6101000a81548161ffff021916908361ffff160217905550608082015181600001600d6101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060a08201518160000160156101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060c0820151816001015560e0820151816002015561010082015181600301559050506001905095945050505050565b600080600080600080600060011515600160008a63ffffffff1663ffffffff16815260200190815260200160002060000160009054906101000a900460ff16151514611312576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161130990612a21565b60405180910390fd5b600160008963ffffffff1663ffffffff16815260200190815260200160002060000160019054906101000a900467ffffffffffffffff16600160008a63ffffffff1663ffffffff168152602001908152602001600020600001600d9054906101000a900467ffffffffffffffff16600160008b63ffffffff1663ffffffff16815260200190815260200160002060000160099054906101000a900461ffff16600160008c63ffffffff1663ffffffff168152602001908152602001600020600001600b9054906101000a900461ffff16600160008d63ffffffff1663ffffffff16815260200190815260200160002060000160159054906101000a900467ffffffffffffffff16600160008e63ffffffff1663ffffffff16815260200190815260200160002060010154600160008f63ffffffff1663ffffffff168152602001908152602001600020600301549650965096509650965096509650919395979092949650565b600080600115156003600086815260200190815260200160002060000160009054906101000a900460ff161515146114e5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114dc90612a01565b60405180910390fd5b8260036000868152602001908152602001600020600101541461153d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161153490612841565b60405180910390fd5b60016003600086815260200190815260200160002060020154915091509250929050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000801515600060149054906101000a900460ff161515146115e1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016115d8906129c1565b60405180910390fd5b7310000000000000000000000000000000000000006000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040518061012001604052806001151581526020016303a38d8a67ffffffffffffffff168152602001601e61ffff168152602001600061ffff168152602001600067ffffffffffffffff1681526020016303a38d8a67ffffffffffffffff168152602001428152602001607881526020016000815250600160008063ffffffff16815260200190815260200160002060008201518160000160006101000a81548160ff02191690831515021790555060208201518160000160016101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060408201518160000160096101000a81548161ffff021916908361ffff160217905550606082015181600001600b6101000a81548161ffff021916908361ffff160217905550608082015181600001600d6101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060a08201518160000160156101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060c0820151816001015560e0820151816002015561010082015181600301559050506001600060146101000a81548160ff0219169083151502179055506001905090565b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900467ffffffffffffffff169050919050565b60008060003273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146118cb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016118c290612981565b60405180910390fd5b60011515600160008b63ffffffff1663ffffffff16815260200190815260200160002060000160009054906101000a900460ff16151514611941576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161193890612a21565b60405180910390fd5b600015156003600088815260200190815260200160002060000160009054906101000a900460ff161515146119ab576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016119a2906128a1565b60405180910390fd5b60006119b8888b8b611bc2565b9050600115156119c9898988611d96565b151514611a0b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a0290612941565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff164173ffffffffffffffffffffffffffffffffffffffff161480611a84575073010000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff164173ffffffffffffffffffffffffffffffffffffffff16145b611ac3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611aba906129e1565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff164173ffffffffffffffffffffffffffffffffffffffff16148015611b3e575073010000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff164173ffffffffffffffffffffffffffffffffffffffff1614155b15611bac576040518060600160405280600115158152602001878152602001828152506003600089815260200190815260200160002060008201518160000160006101000a81548160ff02191690831515021790555060208201518160010155604082015181600201559050505b8987879350935093505096509650969350505050565b600060011515600160008563ffffffff1663ffffffff16815260200190815260200160002060000160009054906101000a900460ff16151514611c3a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611c3190612a21565b60405180910390fd5b600083604051602001611c4d919061278c565b6040516020818303038152906040528051906020012083604051602001611c7491906127a7565b60405160208183030381529060405280519060200120604051602001611c9b929190612760565b604051602081830303815290604052805190602001209050600115156002600083815260200190815260200160002060000160009054906101000a900460ff16151514611d1d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611d1490612881565b60405180910390fd5b84600260008381526020019081526020016000206001015414611d75576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611d6c90612821565b60405180910390fd5b60026000828152602001908152602001600020600201549150509392505050565b60008083905060005b8351811015611e32576000848281518110611db657fe5b6020026020010151905080831015611df8578281604051602001611ddb929190612760565b604051602081830303815290604052805190602001209250611e24565b8083604051602001611e0b929190612760565b6040516020818303038152906040528051906020012092505b508080600101915050611d9f565b508481149150509392505050565b600081359050611e4f81612ca3565b92915050565b600081359050611e6481612cba565b92915050565b600082601f830112611e7b57600080fd5b8135611e8e611e8982612b86565b612b55565b91508181835260208401935060208101905083856020840282011115611eb357600080fd5b60005b83811015611ee35781611ec98882611eed565b845260208401935060208301925050600181019050611eb6565b5050505092915050565b600081359050611efc81612cd1565b92915050565b600081359050611f1181612ce8565b92915050565b600081359050611f2681612cff565b92915050565b600081359050611f3b81612d16565b92915050565b600081359050611f5081612d2d565b92915050565b600060208284031215611f6857600080fd5b6000611f7684828501611e40565b91505092915050565b600060208284031215611f9157600080fd5b6000611f9f84828501611e55565b91505092915050565b60008060408385031215611fbb57600080fd5b6000611fc985828601611eed565b9250506020611fda85828601611eed565b9150509250929050565b600060208284031215611ff657600080fd5b600061200484828501611f2c565b91505092915050565b6000806040838503121561202057600080fd5b600061202e85828601611f2c565b925050602061203f85828601611f17565b9150509250929050565b6000806040838503121561205c57600080fd5b600061206a85828601611f2c565b925050602061207b85828601611f41565b9150509250929050565b60008060008060008060c0878903121561209e57600080fd5b60006120ac89828a01611f2c565b96505060206120bd89828a01611f41565b95505060406120ce89828a01611eed565b94505060606120df89828a01611eed565b93505060806120f089828a01611eed565b92505060a087013567ffffffffffffffff81111561210d57600080fd5b61211989828a01611e6a565b9150509295509295509295565b600080600080600060a0868803121561213e57600080fd5b600061214c88828901611f2c565b955050602061215d88828901611f41565b945050604061216e88828901611f02565b935050606061217f88828901611f02565b925050608061219088828901611f17565b9150509295509295909350565b600080600080608085870312156121b357600080fd5b60006121c187828801611f2c565b94505060206121d287828801611f41565b93505060406121e387828801611f41565b92505060606121f487828801611eed565b91505092959194509250565b61220981612bc3565b82525050565b61221881612be7565b82525050565b61222781612bf3565b82525050565b61223e61223982612bf3565b612c59565b82525050565b6000612251600c83612bb2565b91507f496e76616c696420726f6f7400000000000000000000000000000000000000006000830152602082019050919050565b6000612291601383612bb2565b91507f696e76616c6964207061796d656e7448617368000000000000000000000000006000830152602082019050919050565b60006122d1602c83612bb2565b91507f6e6f7420656e6f7567682074696d6520656c61707365642073696e636520707260008301527f696f722066696e616c69747900000000000000000000000000000000000000006020830152604082019050919050565b6000612337603283612bb2565b91507f66696e616c69736564436c61696d506572696f64735b6c6f636174696f6e486160008301527f73685d20646f6573206e6f7420657869737400000000000000000000000000006020830152604082019050919050565b600061239d601383612bb2565b91507f7478496420616c72656164792070726f76656e000000000000000000000000006000830152602082019050919050565b60006123dd602083612bb2565b91507f6d73672e73656e64657220213d20676f7665726e616e6365436f6e74726163746000830152602082019050919050565b600061241d601683612bb2565b91507f636861696e496420616c726561647920657869737473000000000000000000006000830152602082019050919050565b600061245d601883612bb2565b91507f696e76616c696420636c61696d506572696f64496e64657800000000000000006000830152602082019050919050565b600061249d602783612bb2565b91507f70726576696f757320636c61696d20706572696f64206e6f742079657420666960008301527f6e616c69736564000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000612503601483612bb2565b91507f496e76616c6964206d65726b6c652070726f6f660000000000000000000000006000830152602082019050919050565b6000612543600e83612bb2565b91507f696e76616c6964206c65646765720000000000000000000000000000000000006000830152602082019050919050565b6000612583601783612bb2565b91507f6d73672e73656e64657220213d2074782e6f726967696e0000000000000000006000830152602082019050919050565b60006125c3601e83612bb2565b91507f6c6f636174696f6e4861736820616c72656164792066696e616c6973656400006000830152602082019050919050565b6000612603601483612bb2565b91507f696e697469616c6973656420213d2066616c73650000000000000000000000006000830152602082019050919050565b6000612643601c83612bb2565b91507f496e76616c696420626c6f636b2e636f696e626173652076616c7565000000006000830152602082019050919050565b6000612683601383612bb2565b91507f7478496420646f6573206e6f74206578697374000000000000000000000000006000830152602082019050919050565b60006126c3601683612bb2565b91507f636861696e496420646f6573206e6f74206578697374000000000000000000006000830152602082019050919050565b6126ff81612bfd565b82525050565b61270e81612c2b565b82525050565b61271d81612c35565b82525050565b61273461272f82612c35565b612c63565b82525050565b61274381612c45565b82525050565b61275a61275582612c45565b612c75565b82525050565b600061276c828561222d565b60208201915061277c828461222d565b6020820191508190509392505050565b60006127988284612723565b60048201915081905092915050565b60006127b38284612749565b60088201915081905092915050565b60006020820190506127d76000830184612200565b92915050565b60006020820190506127f2600083018461220f565b92915050565b600060408201905061280d600083018561220f565b61281a6020830184612705565b9392505050565b6000602082019050818103600083015261283a81612244565b9050919050565b6000602082019050818103600083015261285a81612284565b9050919050565b6000602082019050818103600083015261287a816122c4565b9050919050565b6000602082019050818103600083015261289a8161232a565b9050919050565b600060208201905081810360008301526128ba81612390565b9050919050565b600060208201905081810360008301526128da816123d0565b9050919050565b600060208201905081810360008301526128fa81612410565b9050919050565b6000602082019050818103600083015261291a81612450565b9050919050565b6000602082019050818103600083015261293a81612490565b9050919050565b6000602082019050818103600083015261295a816124f6565b9050919050565b6000602082019050818103600083015261297a81612536565b9050919050565b6000602082019050818103600083015261299a81612576565b9050919050565b600060208201905081810360008301526129ba816125b6565b9050919050565b600060208201905081810360008301526129da816125f6565b9050919050565b600060208201905081810360008301526129fa81612636565b9050919050565b60006020820190508181036000830152612a1a81612676565b9050919050565b60006020820190508181036000830152612a3a816126b6565b9050919050565b6000606082019050612a566000830186612714565b612a63602083018561221e565b612a70604083018461221e565b949350505050565b600060a082019050612a8d6000830188612714565b612a9a602083018761273a565b612aa760408301866126f6565b612ab460608301856126f6565b612ac1608083018461221e565b9695505050505050565b6000602082019050612ae0600083018461273a565b92915050565b600060e082019050612afb600083018a61273a565b612b08602083018961273a565b612b1560408301886126f6565b612b2260608301876126f6565b612b2f608083018661273a565b612b3c60a0830185612705565b612b4960c0830184612705565b98975050505050505050565b6000604051905081810181811067ffffffffffffffff82111715612b7c57612b7b612c87565b5b8060405250919050565b600067ffffffffffffffff821115612ba157612ba0612c87565b5b602082029050602081019050919050565b600082825260208201905092915050565b6000612bce82612c0b565b9050919050565b6000612be082612c0b565b9050919050565b60008115159050919050565b6000819050919050565b600061ffff82169050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600063ffffffff82169050919050565b600067ffffffffffffffff82169050919050565b6000819050919050565b6000612c6e82612c96565b9050919050565b6000612c8082612c89565b9050919050565bfe5b60008160c01b9050919050565b60008160e01b9050919050565b612cac81612bc3565b8114612cb757600080fd5b50565b612cc381612bd5565b8114612cce57600080fd5b50565b612cda81612bf3565b8114612ce557600080fd5b50565b612cf181612bfd565b8114612cfc57600080fd5b50565b612d0881612c2b565b8114612d1357600080fd5b50565b612d1f81612c35565b8114612d2a57600080fd5b50565b612d3681612c45565b8114612d4157600080fd5b5056fea26469706673582212204c84245e6785e6860218721ef7f68f77b2136441902b005500330b65401b2c4364736f6c63430007030033"
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