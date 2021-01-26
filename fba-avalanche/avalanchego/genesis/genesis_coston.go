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
	        "code": "0x608060405234801561001057600080fd5b50600436106100a95760003560e01c8063760f6a5a11610071578063760f6a5a1461018d578063b172b222146101c1578063ef2fa85f146101df578063ef4c169e146101fd578063efe782721461021b578063f08ca9871461024b576100a9565b806307003bb4146100ae5780631129753f146100cc5780632cd2f4c6146100fc578063407f908c1461012d578063550d098e1461015d575b600080fd5b6100b6610280565b6040516100c39190611f63565b60405180910390f35b6100e660048036038101906100e191906117b3565b610293565b6040516100f39190611f63565b60405180910390f35b610116600480360381019061011191906118a4565b61036d565b604051610124929190611f7e565b60405180910390f35b61014760048036038101906101429190611982565b6105df565b6040516101549190611f63565b60405180910390f35b61017760048036038101906101729190611805565b6107a1565b6040516101849190611f63565b60405180910390f35b6101a760048036038101906101a29190611841565b6108b3565b6040516101b8959493929190611fa7565b60405180910390f35b6101c96110aa565b6040516101d69190611f48565b60405180910390f35b6101e76110ce565b6040516101f49190611f48565b60405180910390f35b6102056110f7565b6040516102129190611f63565b60405180910390f35b6102356004803603810190610230919061178a565b611289565b604051610242919061226d565b60405180910390f35b610265600480360381019061026091906117dc565b6112d2565b60405161027796959493929190612288565b60405180910390f35b600060149054906101000a900460ff1681565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610324576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161031b906120cd565b60405180910390fd5b816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060019050919050565b60008060011515600160008d815260200190815260200160002060000160009054906101000a900460ff161515146103da576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103d19061224d565b60405180910390fd5b600160008c8152602001908152602001600020600201548a02600160008d8152602001908152602001600020600101540189101561044d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104449061206d565b60405180910390fd5b600160008c81526020019081526020016000206002015460018b0102600160008d8152602001908152602001600020600101540189106104c2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104b99061222d565b60405180910390fd5b600115156104d1868d8d6113e4565b151514610513576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161050a9061218d565b60405180910390fd5b836105218c8b8b8b8b611556565b14610561576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610558906120ed565b60405180910390fd5b60011515610570868686611609565b1515146105b2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105a9906120ad565b60405180910390fd5b600180600160008e8152602001908152602001600020600401540391509150995099975050505050505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610670576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610667906120cd565b60405180910390fd5b600015156001600087815260200190815260200160002060000160009054906101000a900460ff161515146106da576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106d19061210d565b60405180910390fd5b6040518061010001604052806001151581526020018581526020018481526020016000815260200185815260200142815260200183815260200160008152506001600087815260200190815260200160002060008201518160000160006101000a81548160ff0219169083151502179055506020820151816001015560408201518160020155606082015181600301556080820151816004015560a0820151816005015560c0820151816006015560e0820151816007015590505060019050949350505050565b6000600115156001600085815260200190815260200160002060000160009054906101000a900460ff1615151461080d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108049061224d565b60405180910390fd5b6000836040516020016108209190611f2d565b60405160208183030381529060405280519060200120836040516020016108479190611f2d565b6040516020818303038152906040528051906020012060405160200161086e929190611f01565b6040516020818303038152906040528051906020012090506002600082815260200190815260200160002060000160009054906101000a900460ff1691505092915050565b60008060008060003273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610929576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610920906121ad565b60405180910390fd5b60011515600160008b815260200190815260200160002060000160009054906101000a900460ff16151514610993576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161098a9061224d565b60405180910390fd5b600160008a815260200190815260200160002060020154600160008b815260200190815260200160002060040154018814610a03576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109fa9061216d565b60405180910390fd5b600160008a8152602001908152602001600020600201546001880102600160008b815260200190815260200160002060010154018814610a78576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a6f9061212d565b60405180910390fd5b600160008a8152602001908152602001600020600501544210610b0957600160008a815260200190815260200160002060070154600160008b815260200190815260200160002060050154420360020211610b08576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610aff9061204d565b60405180910390fd5b5b600089604051602001610b1c9190611f2d565b6040516020818303038152906040528051906020012088604051602001610b439190611f2d565b60405160208183030381529060405280519060200120604051602001610b6a929190611f01565b604051602081830303815290604052805190602001209050600015156002600083815260200190815260200160002060000160009054906101000a900460ff16151514610bec576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610be39061208d565b60405180910390fd5b6000881115610cdd5760008a604051602001610c089190611f2d565b6040516020818303038152906040528051906020012060018a03604051602001610c329190611f2d565b60405160208183030381529060405280519060200120604051602001610c59929190611f01565b604051602081830303815290604052805190602001209050600115156002600083815260200190815260200160002060000160009054906101000a900460ff16151514610cdb576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cd29061214d565b60405180910390fd5b505b3373ffffffffffffffffffffffffffffffffffffffff164173ffffffffffffffffffffffffffffffffffffffff161480610d56575073010000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff164173ffffffffffffffffffffffffffffffffffffffff16145b610d95576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d8c906121ed565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff164173ffffffffffffffffffffffffffffffffffffffff16148015610e10575073010000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff164173ffffffffffffffffffffffffffffffffffffffff1614155b15611077576001600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205401600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506040518060400160405280600115158152602001888152506002600083815260200190815260200160002060008201518160000160006101000a81548160ff0219169083151502179055506020820151816001015590505060018801600160008c81526020019081526020016000206003018190555088600160008c815260200190815260200160002060040181905550600160008b81526020019081526020016000206005015442106110195760006002600160008d8152602001908152602001600020600501544203600160008e8152602001908152602001600020600701540181610f8557fe5b049050600160008c815260200190815260200160002060060154811115610fdc57600160008c815260200190815260200160002060060154600160008d815260200190815260200160002060070181905550610ff8565b80600160008d8152602001908152602001600020600701819055505b42600160008d8152602001908152602001600020600501819055505061104b565b600160008b815260200190815260200160002060060154600160008c8152602001908152602001600020600701819055505b60018a8a600160008e8152602001908152602001600020600201548a955095509550955095505061109f565b60008a8a600160008e8152602001908152602001600020600201548a95509550955095509550505b945094509450945094565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000801515600060149054906101000a900460ff1615151461114e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611145906121cd565b60405180910390fd5b7310000000000000000000000000000000000000006000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040518061010001604052806001151581526020016303a38d8a815260200160328152602001600081526020016303a38d8a81526020014281526020016096815260200160008152506001600080815260200190815260200160002060008201518160000160006101000a81548160ff0219169083151502179055506020820151816001015560408201518160020155606082015181600301556080820151816004015560a0820151816005015560c0820151816006015560e082015181600701559050506001600060146101000a81548160ff0219169083151502179055506001905090565b6000600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b600080600080600080600115156001600089815260200190815260200160002060000160009054906101000a900460ff16151514611345576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161133c9061224d565b60405180910390fd5b60016000888152602001908152602001600020600101546001600089815260200190815260200160002060030154600160008a815260200190815260200160002060020154600160008b815260200190815260200160002060040154600160008c815260200190815260200160002060050154600160008d81526020019081526020016000206007015495509550955095509550955091939550919395565b6000600115156001600085815260200190815260200160002060000160009054906101000a900460ff16151514611450576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114479061224d565b60405180910390fd5b6000836040516020016114639190611f2d565b604051602081830303815290604052805190602001208360405160200161148a9190611f2d565b604051602081830303815290604052805190602001206040516020016114b1929190611f01565b604051602081830303815290604052805190602001209050600115156002600083815260200190815260200160002060000160009054906101000a900460ff16151514611533576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161152a9061220d565b60405180910390fd5b846002600083815260200190815260200160002060010154149150509392505050565b6000808660405160200161156a919061226d565b6040516020818303038152906040528051906020012086604051602001611591919061226d565b604051602081830303815290604052805190602001208686866040516020016115ba919061226d565b604051602081830303815290604052805190602001206040516020016115e4959493929190611ffa565b6040516020818303038152906040528051906020012090508091505095945050505050565b60008083905060005b83518110156116a557600084828151811061162957fe5b602002602001015190508083101561166b57828160405160200161164e929190611f01565b604051602081830303815290604052805190602001209250611697565b808360405160200161167e929190611f01565b6040516020818303038152906040528051906020012092505b508080600101915050611612565b508481149150509392505050565b6000813590506116c2816123d1565b92915050565b6000813590506116d7816123e8565b92915050565b600082601f8301126116ee57600080fd5b81356117016116fc8261231a565b6122e9565b9150818183526020840193506020810190508385602084028201111561172657600080fd5b60005b83811015611756578161173c8882611760565b845260208401935060208301925050600181019050611729565b5050505092915050565b60008135905061176f816123ff565b92915050565b60008135905061178481612416565b92915050565b60006020828403121561179c57600080fd5b60006117aa848285016116b3565b91505092915050565b6000602082840312156117c557600080fd5b60006117d3848285016116c8565b91505092915050565b6000602082840312156117ee57600080fd5b60006117fc84828501611775565b91505092915050565b6000806040838503121561181857600080fd5b600061182685828601611775565b925050602061183785828601611775565b9150509250929050565b6000806000806080858703121561185757600080fd5b600061186587828801611775565b945050602061187687828801611775565b935050604061188787828801611775565b925050606061189887828801611760565b91505092959194509250565b60008060008060008060008060006101208a8c0312156118c357600080fd5b60006118d18c828d01611775565b99505060206118e28c828d01611775565b98505060406118f38c828d01611775565b97505060606119048c828d01611760565b96505060806119158c828d01611760565b95505060a06119268c828d01611775565b94505060c06119378c828d01611760565b93505060e06119488c828d01611760565b9250506101008a013567ffffffffffffffff81111561196657600080fd5b6119728c828d016116dd565b9150509295985092959850929598565b6000806000806080858703121561199857600080fd5b60006119a687828801611775565b94505060206119b787828801611775565b93505060406119c887828801611775565b92505060606119d987828801611775565b91505092959194509250565b6119ee81612357565b82525050565b6119fd8161237b565b82525050565b611a0c81612387565b82525050565b611a23611a1e82612387565b6123bb565b82525050565b6000611a36602c83612346565b91507f6e6f7420656e6f7567682074696d6520656c61707365642073696e636520707260008301527f696f722066696e616c69747900000000000000000000000000000000000000006020830152604082019050919050565b6000611a9c602083612346565b91507f6c6564676572203c20636c61696d506572696f64496e64657820726567696f6e6000830152602082019050919050565b6000611adc602183612346565b91507f636c61696d506572696f644861736820616c72656164792066696e616c69736560008301527f64000000000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000611b42601483612346565b91507f5061796d656e74206e6f742076657269666965640000000000000000000000006000830152602082019050919050565b6000611b82602083612346565b91507f6d73672e73656e64657220213d20676f7665726e616e6365436f6e74726163746000830152602082019050919050565b6000611bc2601783612346565b91507f636f6e73747275637465644c65616620213d206c6561660000000000000000006000830152602082019050919050565b6000611c02601683612346565b91507f636861696e496420616c726561647920657869737473000000000000000000006000830152602082019050919050565b6000611c42601883612346565b91507f696e76616c696420636c61696d506572696f64496e64657800000000000000006000830152602082019050919050565b6000611c82602783612346565b91507f70726576696f757320636c61696d20706572696f64206e6f742079657420666960008301527f6e616c69736564000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000611ce8600e83612346565b91507f696e76616c6964206c65646765720000000000000000000000000000000000006000830152602082019050919050565b6000611d28601a83612346565b91507f436c61696d20706572696f64206e6f742066696e616c697365640000000000006000830152602082019050919050565b6000611d68601783612346565b91507f6d73672e73656e64657220213d2074782e6f726967696e0000000000000000006000830152602082019050919050565b6000611da8601483612346565b91507f696e697469616c6973656420213d2066616c73650000000000000000000000006000830152602082019050919050565b6000611de8601c83612346565b91507f496e76616c696420626c6f636b2e636f696e626173652076616c7565000000006000830152602082019050919050565b6000611e28601e83612346565b91507f636c61696d506572696f644861736820646f6573206e6f7420657869737400006000830152602082019050919050565b6000611e68602083612346565b91507f6c6564676572203e20636c61696d506572696f64496e64657820726567696f6e6000830152602082019050919050565b6000611ea8601683612346565b91507f636861696e496420646f6573206e6f74206578697374000000000000000000006000830152602082019050919050565b611ee4816123b1565b82525050565b611efb611ef6826123b1565b6123c5565b82525050565b6000611f0d8285611a12565b602082019150611f1d8284611a12565b6020820191508190509392505050565b6000611f398284611eea565b60208201915081905092915050565b6000602082019050611f5d60008301846119e5565b92915050565b6000602082019050611f7860008301846119f4565b92915050565b6000604082019050611f9360008301856119f4565b611fa06020830184611edb565b9392505050565b600060a082019050611fbc60008301886119f4565b611fc96020830187611edb565b611fd66040830186611edb565b611fe36060830185611edb565b611ff06080830184611a03565b9695505050505050565b600060a08201905061200f6000830188611a03565b61201c6020830187611a03565b6120296040830186611a03565b6120366060830185611a03565b6120436080830184611a03565b9695505050505050565b6000602082019050818103600083015261206681611a29565b9050919050565b6000602082019050818103600083015261208681611a8f565b9050919050565b600060208201905081810360008301526120a681611acf565b9050919050565b600060208201905081810360008301526120c681611b35565b9050919050565b600060208201905081810360008301526120e681611b75565b9050919050565b6000602082019050818103600083015261210681611bb5565b9050919050565b6000602082019050818103600083015261212681611bf5565b9050919050565b6000602082019050818103600083015261214681611c35565b9050919050565b6000602082019050818103600083015261216681611c75565b9050919050565b6000602082019050818103600083015261218681611cdb565b9050919050565b600060208201905081810360008301526121a681611d1b565b9050919050565b600060208201905081810360008301526121c681611d5b565b9050919050565b600060208201905081810360008301526121e681611d9b565b9050919050565b6000602082019050818103600083015261220681611ddb565b9050919050565b6000602082019050818103600083015261222681611e1b565b9050919050565b6000602082019050818103600083015261224681611e5b565b9050919050565b6000602082019050818103600083015261226681611e9b565b9050919050565b60006020820190506122826000830184611edb565b92915050565b600060c08201905061229d6000830189611edb565b6122aa6020830188611edb565b6122b76040830187611edb565b6122c46060830186611edb565b6122d16080830185611edb565b6122de60a0830184611edb565b979650505050505050565b6000604051905081810181811067ffffffffffffffff821117156123105761230f6123cf565b5b8060405250919050565b600067ffffffffffffffff821115612335576123346123cf565b5b602082029050602081019050919050565b600082825260208201905092915050565b600061236282612391565b9050919050565b600061237482612391565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6000819050919050565b6000819050919050565bfe5b6123da81612357565b81146123e557600080fd5b50565b6123f181612369565b81146123fc57600080fd5b50565b61240881612387565b811461241357600080fd5b50565b61241f816123b1565b811461242a57600080fd5b5056fea26469706673582212205013e058a1fefa00834557f815688c6c968a85014290d29c83a8513eef8bd0c264736f6c63430007030033"
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