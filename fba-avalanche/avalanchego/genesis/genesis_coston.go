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
	        "code": "0x608060405234801561001057600080fd5b50600436106100b45760003560e01c8063d02a133511610071578063d02a1335146101b9578063dab99b50146101ef578063ef2fa85f1461021f578063ef4c169e1461023d578063efe782721461025b578063f6f56af71461028b576100b4565b806307003bb4146100b9578063093370c0146100d75780631129753f146101075780635465dfc414610137578063b172b2221461016b578063bd18dae314610189575b600080fd5b6100c16102bb565b6040516100ce9190612520565b60405180910390f35b6100f160048036038101906100ec9190611dcc565b6102ce565b6040516100fe9190612520565b60405180910390f35b610121600480360381019061011c9190611d3e565b6103ec565b60405161012e9190612520565b60405180910390f35b610151600480360381019061014c9190611f20565b6104c6565b60405161016295949392919061273b565b60405180910390f35b610173610e24565b6040516101809190612505565b60405180910390f35b6101a3600480360381019061019e9190611ea9565b610e48565b6040516101b09190612520565b60405180910390f35b6101d360048036038101906101ce9190611da3565b6110fb565b6040516101e697969594939291906127a9565b60405180910390f35b61020960048036038101906102049190611d67565b6112e2565b6040516102169190612520565b60405180910390f35b6102276113b0565b6040516102349190612505565b60405180910390f35b6102456113d9565b6040516102529190612520565b60405180910390f35b61027560048036038101906102709190611d15565b61164a565b604051610282919061278e565b60405180910390f35b6102a560048036038101906102a09190611e08565b6116a7565b6040516102b29190612520565b60405180910390f35b600060149054906101000a900460ff1681565b600060011515600160008563ffffffff1663ffffffff16815260200190815260200160002060000160009054906101000a900460ff16151514610346576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161033d9061271b565b60405180910390fd5b60008360405160200161035991906124cf565b604051602081830303815290604052805190602001208360405160200161038091906124ea565b604051602081830303815290604052805190602001206040516020016103a79291906124a3565b6040516020818303038152906040528051906020012090506002600082815260200190815260200160002060000160009054906101000a900460ff1691505092915050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461047d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610474906125bb565b60405180910390fd5b816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060019050919050565b60008060008060003273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461053c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105339061267b565b60405180910390fd5b60011515600160008b63ffffffff1663ffffffff16815260200190815260200160002060000160009054906101000a900460ff161515146105b2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105a99061271b565b60405180910390fd5b600160008a63ffffffff1663ffffffff16815260200190815260200160002060000160099054906101000a900461ffff1661ffff16600160008b63ffffffff1663ffffffff16815260200190815260200160002060000160159054906101000a900467ffffffffffffffff160167ffffffffffffffff168867ffffffffffffffff1614610674576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161066b9061263b565b60405180910390fd5b600160008a63ffffffff1663ffffffff16815260200190815260200160002060000160099054906101000a900461ffff1661ffff166001880102600160008b63ffffffff1663ffffffff16815260200190815260200160002060000160019054906101000a900467ffffffffffffffff160167ffffffffffffffff168867ffffffffffffffff161461073b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610732906125fb565b60405180910390fd5b600160008a63ffffffff1663ffffffff16815260200190815260200160002060030154600202600160008b63ffffffff1663ffffffff16815260200190815260200160002060010154420360030210156107ca576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107c19061255b565b60405180910390fd5b6000896040516020016107dd91906124cf565b604051602081830303815290604052805190602001208860405160200161080491906124ea565b6040516020818303038152906040528051906020012060405160200161082b9291906124a3565b604051602081830303815290604052805190602001209050600015156002600083815260200190815260200160002060000160009054906101000a900460ff161515146108ad576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108a49061269b565b60405180910390fd5b60008867ffffffffffffffff1611156109a85760008a6040516020016108d391906124cf565b6040516020818303038152906040528051906020012060018a036040516020016108fd91906124ea565b604051602081830303815290604052805190602001206040516020016109249291906124a3565b604051602081830303815290604052805190602001209050600115156002600083815260200190815260200160002060000160009054906101000a900460ff161515146109a6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161099d9061261b565b60405180910390fd5b505b3373ffffffffffffffffffffffffffffffffffffffff164173ffffffffffffffffffffffffffffffffffffffff161480610a21575073010000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff164173ffffffffffffffffffffffffffffffffffffffff16145b610a60576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a57906126db565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff164173ffffffffffffffffffffffffffffffffffffffff16148015610adb575073010000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff164173ffffffffffffffffffffffffffffffffffffffff1614155b15610da9576001600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900467ffffffffffffffff1601600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506040518060400160405280600115158152602001888152506002600083815260200190815260200160002060008201518160000160006101000a81548160ff0219169083151502179055506020820151816001015590505060018801600160008c63ffffffff1663ffffffff168152602001908152602001600020600001600d6101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555088600160008c63ffffffff1663ffffffff16815260200190815260200160002060000160156101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060006002600160008d63ffffffff1663ffffffff168152602001908152602001600020600101544203600160008e63ffffffff1663ffffffff168152602001908152602001600020600301540181610cdd57fe5b049050600160008c63ffffffff1663ffffffff16815260200190815260200160002060020154811115610d5857600160008c63ffffffff1663ffffffff16815260200190815260200160002060020154600160008d63ffffffff1663ffffffff16815260200190815260200160002060030181905550610d80565b80600160008d63ffffffff1663ffffffff168152602001908152602001600020600301819055505b42600160008d63ffffffff1663ffffffff16815260200190815260200160002060010181905550505b8989600160008d63ffffffff1663ffffffff16815260200190815260200160002060000160099054906101000a900461ffff16600160008e63ffffffff1663ffffffff168152602001908152602001600020600001600b9054906101000a900461ffff168a9550955095509550955050945094509450945094565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610ed9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ed0906125bb565b60405180910390fd5b60001515600160008863ffffffff1663ffffffff16815260200190815260200160002060000160009054906101000a900460ff16151514610f4f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f46906125db565b60405180910390fd5b6040518061012001604052806001151581526020018667ffffffffffffffff1681526020018561ffff1681526020018461ffff168152602001600067ffffffffffffffff1681526020018667ffffffffffffffff1681526020014281526020018381526020016000815250600160008863ffffffff1663ffffffff16815260200190815260200160002060008201518160000160006101000a81548160ff02191690831515021790555060208201518160000160016101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060408201518160000160096101000a81548161ffff021916908361ffff160217905550606082015181600001600b6101000a81548161ffff021916908361ffff160217905550608082015181600001600d6101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060a08201518160000160156101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060c0820151816001015560e0820151816002015561010082015181600301559050506001905095945050505050565b600080600080600080600060011515600160008a63ffffffff1663ffffffff16815260200190815260200160002060000160009054906101000a900460ff1615151461117c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111739061271b565b60405180910390fd5b600160008963ffffffff1663ffffffff16815260200190815260200160002060000160019054906101000a900467ffffffffffffffff16600160008a63ffffffff1663ffffffff168152602001908152602001600020600001600d9054906101000a900467ffffffffffffffff16600160008b63ffffffff1663ffffffff16815260200190815260200160002060000160099054906101000a900461ffff16600160008c63ffffffff1663ffffffff168152602001908152602001600020600001600b9054906101000a900461ffff16600160008d63ffffffff1663ffffffff16815260200190815260200160002060000160159054906101000a900467ffffffffffffffff16600160008e63ffffffff1663ffffffff16815260200190815260200160002060010154600160008f63ffffffff1663ffffffff168152602001908152602001600020600301549650965096509650965096509650919395979092949650565b6000600115156003600085815260200190815260200160002060000160009054906101000a900460ff1615151461134e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611345906126fb565b60405180910390fd5b816003600085815260200190815260200160002060010154146113a6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161139d9061253b565b60405180910390fd5b6001905092915050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000801515600060149054906101000a900460ff16151514611430576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611427906126bb565b60405180910390fd5b7310000000000000000000000000000000000000006000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040518061012001604052806001151581526020016303a38d8a67ffffffffffffffff168152602001601e61ffff168152602001600061ffff168152602001600067ffffffffffffffff1681526020016303a38d8a67ffffffffffffffff168152602001428152602001607881526020016000815250600160008063ffffffff16815260200190815260200160002060008201518160000160006101000a81548160ff02191690831515021790555060208201518160000160016101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060408201518160000160096101000a81548161ffff021916908361ffff160217905550606082015181600001600b6101000a81548161ffff021916908361ffff160217905550608082015181600001600d6101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060a08201518160000160156101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060c0820151816001015560e0820151816002015561010082015181600301559050506001600060146101000a81548160ff0219169083151502179055506001905090565b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900467ffffffffffffffff169050919050565b60003273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611717576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161170e9061267b565b60405180910390fd5b60011515600160008963ffffffff1663ffffffff16815260200190815260200160002060000160009054906101000a900460ff1615151461178d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016117849061271b565b60405180910390fd5b6001151561179c8689896119d7565b1515146117de576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016117d59061265b565b60405180910390fd5b600115156117ed868685611b55565b15151461182f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016118269061259b565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff164173ffffffffffffffffffffffffffffffffffffffff1614806118a8575073010000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff164173ffffffffffffffffffffffffffffffffffffffff16145b6118e7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016118de906126db565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff164173ffffffffffffffffffffffffffffffffffffffff16148015611962575073010000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff164173ffffffffffffffffffffffffffffffffffffffff1614155b156119c8576040518060400160405280600115158152602001848152506003600086815260200190815260200160002060008201518160000160006101000a81548160ff02191690831515021790555060208201518160010155905050600190506119cd565b600090505b9695505050505050565b600060011515600160008563ffffffff1663ffffffff16815260200190815260200160002060000160009054906101000a900460ff16151514611a4f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a469061271b565b60405180910390fd5b600083604051602001611a6291906124cf565b6040516020818303038152906040528051906020012083604051602001611a8991906124ea565b60405160208183030381529060405280519060200120604051602001611ab09291906124a3565b604051602081830303815290604052805190602001209050600115156002600083815260200190815260200160002060000160009054906101000a900460ff16151514611b32576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611b299061257b565b60405180910390fd5b846002600083815260200190815260200160002060010154149150509392505050565b60008083905060005b8351811015611bf1576000848281518110611b7557fe5b6020026020010151905080831015611bb7578281604051602001611b9a9291906124a3565b604051602081830303815290604052805190602001209250611be3565b8083604051602001611bca9291906124a3565b6040516020818303038152906040528051906020012092505b508080600101915050611b5e565b508481149150509392505050565b600081359050611c0e81612966565b92915050565b600081359050611c238161297d565b92915050565b600082601f830112611c3a57600080fd5b8135611c4d611c4882612849565b612818565b91508181835260208401935060208101905083856020840282011115611c7257600080fd5b60005b83811015611ca25781611c888882611cac565b845260208401935060208301925050600181019050611c75565b5050505092915050565b600081359050611cbb81612994565b92915050565b600081359050611cd0816129ab565b92915050565b600081359050611ce5816129c2565b92915050565b600081359050611cfa816129d9565b92915050565b600081359050611d0f816129f0565b92915050565b600060208284031215611d2757600080fd5b6000611d3584828501611bff565b91505092915050565b600060208284031215611d5057600080fd5b6000611d5e84828501611c14565b91505092915050565b60008060408385031215611d7a57600080fd5b6000611d8885828601611cac565b9250506020611d9985828601611cac565b9150509250929050565b600060208284031215611db557600080fd5b6000611dc384828501611ceb565b91505092915050565b60008060408385031215611ddf57600080fd5b6000611ded85828601611ceb565b9250506020611dfe85828601611d00565b9150509250929050565b60008060008060008060c08789031215611e2157600080fd5b6000611e2f89828a01611ceb565b9650506020611e4089828a01611d00565b9550506040611e5189828a01611cac565b9450506060611e6289828a01611cac565b9350506080611e7389828a01611cac565b92505060a087013567ffffffffffffffff811115611e9057600080fd5b611e9c89828a01611c29565b9150509295509295509295565b600080600080600060a08688031215611ec157600080fd5b6000611ecf88828901611ceb565b9550506020611ee088828901611d00565b9450506040611ef188828901611cc1565b9350506060611f0288828901611cc1565b9250506080611f1388828901611cd6565b9150509295509295909350565b60008060008060808587031215611f3657600080fd5b6000611f4487828801611ceb565b9450506020611f5587828801611d00565b9350506040611f6687828801611d00565b9250506060611f7787828801611cac565b91505092959194509250565b611f8c81612886565b82525050565b611f9b816128aa565b82525050565b611faa816128b6565b82525050565b611fc1611fbc826128b6565b61291c565b82525050565b6000611fd4601383612875565b91507f696e76616c6964207061796d656e7448617368000000000000000000000000006000830152602082019050919050565b6000612014602c83612875565b91507f6e6f7420656e6f7567682074696d6520656c61707365642073696e636520707260008301527f696f722066696e616c69747900000000000000000000000000000000000000006020830152604082019050919050565b600061207a603283612875565b91507f66696e616c69736564436c61696d506572696f64735b6c6f636174696f6e486160008301527f73685d20646f6573206e6f7420657869737400000000000000000000000000006020830152604082019050919050565b60006120e0601483612875565b91507f5061796d656e74206e6f742076657269666965640000000000000000000000006000830152602082019050919050565b6000612120602083612875565b91507f6d73672e73656e64657220213d20676f7665726e616e6365436f6e74726163746000830152602082019050919050565b6000612160601683612875565b91507f636861696e496420616c726561647920657869737473000000000000000000006000830152602082019050919050565b60006121a0601883612875565b91507f696e76616c696420636c61696d506572696f64496e64657800000000000000006000830152602082019050919050565b60006121e0602783612875565b91507f70726576696f757320636c61696d20706572696f64206e6f742079657420666960008301527f6e616c69736564000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000612246600e83612875565b91507f696e76616c6964206c65646765720000000000000000000000000000000000006000830152602082019050919050565b6000612286601a83612875565b91507f436c61696d20706572696f64206e6f742066696e616c697365640000000000006000830152602082019050919050565b60006122c6601783612875565b91507f6d73672e73656e64657220213d2074782e6f726967696e0000000000000000006000830152602082019050919050565b6000612306601e83612875565b91507f6c6f636174696f6e4861736820616c72656164792066696e616c6973656400006000830152602082019050919050565b6000612346601483612875565b91507f696e697469616c6973656420213d2066616c73650000000000000000000000006000830152602082019050919050565b6000612386601c83612875565b91507f496e76616c696420626c6f636b2e636f696e626173652076616c7565000000006000830152602082019050919050565b60006123c6601383612875565b91507f7478496420646f6573206e6f74206578697374000000000000000000000000006000830152602082019050919050565b6000612406601683612875565b91507f636861696e496420646f6573206e6f74206578697374000000000000000000006000830152602082019050919050565b612442816128c0565b82525050565b612451816128ee565b82525050565b612460816128f8565b82525050565b612477612472826128f8565b612926565b82525050565b61248681612908565b82525050565b61249d61249882612908565b612938565b82525050565b60006124af8285611fb0565b6020820191506124bf8284611fb0565b6020820191508190509392505050565b60006124db8284612466565b60048201915081905092915050565b60006124f6828461248c565b60088201915081905092915050565b600060208201905061251a6000830184611f83565b92915050565b60006020820190506125356000830184611f92565b92915050565b6000602082019050818103600083015261255481611fc7565b9050919050565b6000602082019050818103600083015261257481612007565b9050919050565b600060208201905081810360008301526125948161206d565b9050919050565b600060208201905081810360008301526125b4816120d3565b9050919050565b600060208201905081810360008301526125d481612113565b9050919050565b600060208201905081810360008301526125f481612153565b9050919050565b6000602082019050818103600083015261261481612193565b9050919050565b60006020820190508181036000830152612634816121d3565b9050919050565b6000602082019050818103600083015261265481612239565b9050919050565b6000602082019050818103600083015261267481612279565b9050919050565b60006020820190508181036000830152612694816122b9565b9050919050565b600060208201905081810360008301526126b4816122f9565b9050919050565b600060208201905081810360008301526126d481612339565b9050919050565b600060208201905081810360008301526126f481612379565b9050919050565b60006020820190508181036000830152612714816123b9565b9050919050565b60006020820190508181036000830152612734816123f9565b9050919050565b600060a0820190506127506000830188612457565b61275d602083018761247d565b61276a6040830186612439565b6127776060830185612439565b6127846080830184611fa1565b9695505050505050565b60006020820190506127a3600083018461247d565b92915050565b600060e0820190506127be600083018a61247d565b6127cb602083018961247d565b6127d86040830188612439565b6127e56060830187612439565b6127f2608083018661247d565b6127ff60a0830185612448565b61280c60c0830184612448565b98975050505050505050565b6000604051905081810181811067ffffffffffffffff8211171561283f5761283e61294a565b5b8060405250919050565b600067ffffffffffffffff8211156128645761286361294a565b5b602082029050602081019050919050565b600082825260208201905092915050565b6000612891826128ce565b9050919050565b60006128a3826128ce565b9050919050565b60008115159050919050565b6000819050919050565b600061ffff82169050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600063ffffffff82169050919050565b600067ffffffffffffffff82169050919050565b6000819050919050565b600061293182612959565b9050919050565b60006129438261294c565b9050919050565bfe5b60008160c01b9050919050565b60008160e01b9050919050565b61296f81612886565b811461297a57600080fd5b50565b61298681612898565b811461299157600080fd5b50565b61299d816128b6565b81146129a857600080fd5b50565b6129b4816128c0565b81146129bf57600080fd5b50565b6129cb816128ee565b81146129d657600080fd5b50565b6129e2816128f8565b81146129ed57600080fd5b50565b6129f981612908565b8114612a0457600080fd5b5056fea2646970667358221220c8e01ef1e58d5a4c5adfc551d62f6a1d3aa3ec7da59d9e85f842331524e2d84464736f6c63430007030033"
	      },
	      "9b200C53C8635a3d5972e1Dc5a7661713143987C": {
	        "balance": "0x314dc6448d9338c15B0a00000000"
	      },
	      "3D3f2a9B0E231ffa0dBBE6Ef12f33Dc66E72Bf17": {
	        "balance": "0x314dc6448d9338c15B0a00000000"
	      },
	      "8AA75C4751ca8B419f345EA62bD91D4032ed40Ab": {
	        "balance": "0x314dc6448d9338c15B0a00000000"
	      },
	      "5e60D64E668F388e0a4f70930B67126fCcC4B543": {
	        "balance": "0x314dc6448d9338c15B0a00000000"
	      }
	    },
	    "number": "0x0",
	    "gasUsed": "0x0",
	    "parentHash": "0x0000000000000000000000000000000000000000000000000000000000000000"
	  }`
)