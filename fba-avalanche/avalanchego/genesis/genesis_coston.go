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

	CostonParams = Params{
		TxFee:                units.MilliAvax,
		CreationTxFee:        10 * units.MilliAvax,
		UptimeRequirement:    .6, // 60%
		MinValidatorStake:    1 * units.Avax,
		MaxValidatorStake:    3 * units.MegaAvax,
		MinDelegatorStake:    1 * units.Avax,
		MinDelegationFee:     20000, // 2%
		MinStakeDuration:     24 * time.Hour,
		MaxStakeDuration:     365 * 24 * time.Hour,
		StakeMintingPeriod:   365 * 24 * time.Hour,
		EpochFirstTransition: time.Unix(10000000000, 0),
		EpochDuration:        5 * time.Minute,
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
	      "petersburgBlock": 0,
		  "istanbulBlock":0,
		  "muirGlacierBlock":0,
		  "apricotPhase1BlockTimestamp":10000000000
	    },
	    "nonce": "0x0",
	    "timestamp": "0x0",
	    "extraData": "0x00",
	    "gasLimit": "0xbebc200",
	    "difficulty": "0x0",
	    "mixHash": "0x0000000000000000000000000000000000000000000000000000000000000000",
	    "coinbase": "0x0100000000000000000000000000000000000000",
	    "alloc": {
	      "1000000000000000000000000000000000000001": {
	        "balance": "0x0",
	        "code": "0x608060405234801561001057600080fd5b50600436106101005760003560e01c8063a060b64711610097578063dbcfda2711610066578063dbcfda27146105f0578063e54caf921461063a578063ef2fa85f14610680578063ef4c169e146106a457610100565b8063a060b6471461050a578063a57d0e251461052d578063ad2bb1b3146105a7578063c1b0e574146105cd57610100565b806349d5436f116100d357806349d5436f146103425780634d7cf6d51461038a5780637f582432146103c2578063955a22561461048a57610100565b80631129753f14610105578063186d9d881461012d5780632a2434a214610153578063388492dd146101ce575b600080fd5b61012b6004803603602081101561011b57600080fd5b50356001600160a01b03166106ac565b005b61012b6004803603602081101561014357600080fd5b50356001600160a01b0316610776565b6101766004803603602081101561016957600080fd5b503563ffffffff1661083b565b60405180876001600160401b03168152602001866001600160401b031681526020018561ffff168152602001846001600160401b03168152602001838152602001828152602001965050505050505060405180910390f35b610296600480360360808110156101e457600080fd5b63ffffffff823516916020810135916001600160401b03604083013516919081019060808101606082013564010000000081111561022157600080fd5b82018360208201111561023357600080fd5b8035906020019184600183028401116401000000008311171561025557600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506108fe945050505050565b604051808663ffffffff168152602001856001600160401b03168152602001846001600160401b0316815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b838110156103035781810151838201526020016102eb565b50505050905090810190601f1680156103305780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390f35b61036e6004803603604081101561035857600080fd5b506001600160a01b038135169060200135610dd7565b604080516001600160401b039092168252519081900360200190f35b61012b600480360360608110156103a057600080fd5b5063ffffffff813516906001600160401b036020820135169060400135610e08565b610296600480360360808110156103d857600080fd5b63ffffffff823516916020810135916001600160401b03604083013516919081019060808101606082013564010000000081111561041557600080fd5b82018360208201111561042757600080fd5b8035906020019184600183028401116401000000008311171561044957600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610efc945050505050565b6104de600480360360e08110156104a057600080fd5b5063ffffffff813516906020810135906040810135906060810135906001600160401b03608082013581169160a08101359091169060c0013561144c565b604080516001600160401b03948516815292909316602083015215158183015290519081900360600190f35b61012b6004803603602081101561052057600080fd5b503563ffffffff16611674565b61056f6004803603608081101561054357600080fd5b5063ffffffff813516906001600160401b0360208201358116916040810135909116906060013561173a565b6040805163ffffffff90951685526001600160401b03909316602085015261ffff909116838301526060830152519081900360800190f35b61012b600480360360208110156105bd57600080fd5b50356001600160a01b0316611f5c565b61012b600480360360208110156105e357600080fd5b503563ffffffff1661201a565b61012b600480360360a081101561060657600080fd5b506001600160401b03813581169160208101359091169061ffff604082013581169160608101359091169060800135612151565b61066c6004803603604081101561065057600080fd5b50803563ffffffff1690602001356001600160401b031661247c565b604080519115158252519081900360200190f35b610688612536565b604080516001600160a01b039092168252519081900360200190f35b61066c612545565b6000546001600160a01b031633146106f9576040805162461bcd60e51b8152602060048201819052602482015260008051602061298c833981519152604482015290519081900360640190fd5b6001600160a01b038116610754576040805162461bcd60e51b815260206004820152601a60248201527f5f676f7665726e616e6365436f6e7472616374203d3d20307830000000000000604482015290519081900360640190fd5b600080546001600160a01b0319166001600160a01b0392909216919091179055565b6000546001600160a01b031633146107c3576040805162461bcd60e51b8152602060048201819052602482015260008051602061298c833981519152604482015290519081900360640190fd5b6001600160a01b03811660009081526007602052604090205460ff1661081a5760405162461bcd60e51b815260040180806020018281038252602a815260200180612a7f602a913960400191505060405180910390fd5b6001600160a01b03166000908152600760205260409020805460ff19169055565b63ffffffff811660009081526003602052604081205481908190819081908190879060ff1661089f576040805162461bcd60e51b81526020600482015260166024820152600080516020612ade833981519152604482015290519081900360640190fd5b5050505063ffffffff93909316600090815260036020526040902080546001820154600283015460049093015461010083046001600160401b0390811698600160a81b850482169850600160881b90940461ffff169650909116935090565b63ffffffff8416600090815260036020526040812054819081908190606090899060ff16610961576040805162461bcd60e51b81526020600482015260166024820152600080516020612ade833981519152604482015290519081900360640190fd5b6000876040516020018082805190602001908083835b602083106109965780518252601f199092019160209182019101610977565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051602081830303815290604052805190602001209050600560008c63ffffffff1663ffffffff168152602001908152602001600020600082815260200190815260200160002060020160109054906101000a900460ff1615610a5c576040805162461bcd60e51b81526020600482015260136024820152723a3c24b21030b63932b0b23c90383937bb32b760691b604482015290519081900360640190fd5b63ffffffff8b166000908152600360205260409020600101546001600160401b03908116908a1610610abf5760405162461bcd60e51b815260040180806020018281038252602e81526020018061293a602e913960400191505060405180910390fd5b63ffffffff8b166000908152600360205260409020546001600160401b036101008204811691600160481b90041615610c135763ffffffff8c16600090815260036020526040902080546001909101546001600160401b03600160481b830481166101009093048116918116919091031611610b6c5760405162461bcd60e51b81526004018080602001828103825260698152602001806129ac6069913960800191505060405180910390fd5b63ffffffff8c1660009081526003602052604090208054600190910154600160481b9091046001600160401b039081169181168c90031610610bdf5760405162461bcd60e51b81526004018080602001828103825260528152602001806128e86052913960600191505060405180910390fd5b5063ffffffff8b16600090815260036020526040902080546001909101546001600160401b03600160481b90920482169116035b41331480610c24575041600160981b145b610c75576040805162461bcd60e51b815260206004820152601c60248201527f696e76616c696420626c6f636b2e636f696e626173652076616c756500000000604482015290519081900360640190fd5b4133148015610c88575041600160981b14155b15610da4576040518060a001604052806001151581526020018c81526020018b6001600160401b03168152602001826001600160401b0316815260200160011515815250600560008e63ffffffff1663ffffffff168152602001908152602001600020600084815260200190815260200160002060008201518160000160006101000a81548160ff0219169083151502179055506020820151816001015560408201518160020160006101000a8154816001600160401b0302191690836001600160401b0316021790555060608201518160020160086101000a8154816001600160401b0302191690836001600160401b0316021790555060808201518160020160106101000a81548160ff0219169083151502179055509050505b5050505063ffffffff8816600090815260036020526040902060010154979895976001600160401b031696959350505050565b6001600160a01b0391909116600090815260066020908152604080832093835292905220546001600160401b031690565b6000546001600160a01b03163314610e55576040805162461bcd60e51b8152602060048201819052602482015260008051602061298c833981519152604482015290519081900360640190fd5b63ffffffff8316600090815260036020526040902054839060ff16610eaf576040805162461bcd60e51b81526020600482015260166024820152600080516020612ade833981519152604482015290519081900360640190fd5b5063ffffffff909216600090815260036020819052604090912080546001600160401b03909316600160481b0270ffffffffffffffff000000000000000000199093169290921782550155565b63ffffffff8416600090815260036020526040812054819081908190606090899060ff16610f5f576040805162461bcd60e51b81526020600482015260166024820152600080516020612ade833981519152604482015290519081900360640190fd5b6000876040516020018082805190602001908083835b60208310610f945780518252601f199092019160209182019101610f75565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051602081830303815290604052805190602001209050600560008c63ffffffff1663ffffffff168152602001908152602001600020600082815260200190815260200160002060020160109054906101000a900460ff161561105a576040805162461bcd60e51b81526020600482015260136024820152723a3c24b21030b63932b0b23c90383937bb32b760691b604482015290519081900360640190fd5b63ffffffff8b1660009081526005602090815260408083208484529091529020600201546001600160401b03808b169116106110c75760405162461bcd60e51b81526004018080602001828103825260348152602001806128886034913960400191505060405180910390fd5b63ffffffff8b166000908152600360205260409020600101546001600160401b03908116908a161061112a5760405162461bcd60e51b815260040180806020018281038252602e81526020018061293a602e913960400191505060405180910390fd5b63ffffffff8b166000908152600360205260409020546001600160401b036101008204811691600160481b9004161561127e5763ffffffff8c16600090815260036020526040902080546001909101546001600160401b03600160481b8304811661010090930481169181169190910316116111d75760405162461bcd60e51b81526004018080602001828103825260698152602001806129ac6069913960800191505060405180910390fd5b63ffffffff8c1660009081526003602052604090208054600190910154600160481b9091046001600160401b039081169181168c9003161061124a5760405162461bcd60e51b81526004018080602001828103825260528152602001806128e86052913960600191505060405180910390fd5b5063ffffffff8b16600090815260036020526040902080546001909101546001600160401b03600160481b90920482169116035b4133148061128f575041600160981b145b6112e0576040805162461bcd60e51b815260206004820152601c60248201527f696e76616c696420626c6f636b2e636f696e626173652076616c756500000000604482015290519081900360640190fd5b41331480156112f3575041600160981b14155b15610da4576040518060a001604052806001151581526020018c81526020018b6001600160401b03168152602001826001600160401b0316815260200160001515815250600560008e63ffffffff1663ffffffff168152602001908152602001600020600084815260200190815260200160002060008201518160000160006101000a81548160ff0219169083151502179055506020820151816001015560408201518160020160006101000a8154816001600160401b0302191690836001600160401b0316021790555060608201518160020160086101000a8154816001600160401b0302191690836001600160401b0316021790555060808201518160020160106101000a81548160ff021916908315150217905550905050505063ffffffff8a166000908152600360205260409020600101548a96508895506001600160401b0316935088925086915050945094509450945094565b63ffffffff8716600090815260036020526040812054819081908a9060ff166114aa576040805162461bcd60e51b81526020600482015260166024820152600080516020612ade833981519152604482015290519081900360640190fd5b63ffffffff8b1660009081526005602090815260408083208d845290915290205460ff16611515576040805162461bcd60e51b81526020600482015260136024820152721d1e125908191bd95cc81b9bdd08195e1a5cdd606a1b604482015290519081900360640190fd5b604080516001600160401b03808a166020808401919091528351808403820181528385018552805190820120918a166060808501919091528451808503909101815260808401855280519082012060a084018f905260c084018e905260e084018d905261010084019290925261012083019190915261014080830189905283518084039091018152610160909201835281519181019190912063ffffffff8e166000908152600583528381208e8252909252919020600101548114611617576040805162461bcd60e51b81526020600482015260136024820152720d2dcecc2d8d2c840e0c2f2dacadce890c2e6d606b1b604482015290519081900360640190fd5b5050505063ffffffff97909716600090815260056020908152604080832098835297905295909520600201546001600160401b03808216986801000000000000000083049091169750600160801b90910460ff1695509350505050565b6000546001600160a01b031633146116c1576040805162461bcd60e51b8152602060048201819052602482015260008051602061298c833981519152604482015290519081900360640190fd5b63ffffffff8116600090815260036020526040902054819060ff1661171b576040805162461bcd60e51b81526020600482015260166024820152600080516020612ade833981519152604482015290519081900360640190fd5b5063ffffffff166000908152600360205260409020805460ff19169055565b63ffffffff8416600090815260036020526040812054819081908190889060ff1661179a576040805162461bcd60e51b81526020600482015260166024820152600080516020612ade833981519152604482015290519081900360640190fd5b3360009081526007602052604090205460ff16156117e95760405162461bcd60e51b8152600401808060200182810382526022815260200180612a3c6022913960400191505060405180910390fd5b63ffffffff891660009081526003602052604090208054600190910154600160881b90910461ffff166001600160401b0391821601811690891614611866576040805162461bcd60e51b815260206004820152600e60248201526d34b73b30b634b2103632b233b2b960911b604482015290519081900360640190fd5b63ffffffff89166000908152600360205260409020546001600160401b03888116600160a81b90920416146118e2576040805162461bcd60e51b815260206004820152601860248201527f696e76616c696420636c61696d506572696f64496e6465780000000000000000604482015290519081900360640190fd5b63ffffffff891660009081526003602052604090206002015442116119385760405162461bcd60e51b8152600401808060200182810382526035815260200180612aa96035913960400191505060405180910390fd5b63ffffffff891660009081526003602081905260409091209081015460049091015460020210156119d25763ffffffff8916600090815260036020819052604090912060048101546002918201549102429190910390910210156119cd5760405162461bcd60e51b815260040180806020018281038252602c8152602001806128bc602c913960400191505060405180910390fd5b611a34565b63ffffffff8916600090815260036020526040902060048101546002909101544203600f011015611a345760405162461bcd60e51b815260040180806020018281038252602c8152602001806128bc602c913960400191505060405180910390fd5b604080516001600160e01b031960e08c901b166020808301919091526001600160c01b031960c08b901b1660248301528251600c818403018152602c9092018352815191810191909120600081815260049092529190205460ff1615611ae1576040805162461bcd60e51b815260206004820152601e60248201527f6c6f636174696f6e4861736820616c72656164792066696e616c697365640000604482015290519081900360640190fd5b6001600160401b03881615611b8b57604080516001600160e01b031960e08d901b166020808301919091526001600160c01b03196000198c0160c01b1660248301528251600c818403018152602c9092018352815191810191909120600081815260049092529190205460ff16611b895760405162461bcd60e51b8152600401808060200182810382526027815260200180612a156027913960400191505060405180910390fd5b505b41331480611b9c575041600160981b145b611bed576040805162461bcd60e51b815260206004820152601c60248201527f696e76616c696420626c6f636b2e636f696e626173652076616c756500000000604482015290519081900360640190fd5b4133148015611c00575041600160981b14155b15611f25576000611c0f612824565b905060066000336001600160a01b03166001600160a01b03168152602001908152602001600020600082815260200190815260200160002060009054906101000a90046001600160401b031660010160066000336001600160a01b03166001600160a01b03168152602001908152602001600020600083815260200190815260200160002060006101000a8154816001600160401b0302191690836001600160401b031602179055506040518060a001604052806001151581526020018981526020018b6001600160401b0316815260200160006001600160401b03168152602001600115158152506004600084815260200190815260200160002060008201518160000160006101000a81548160ff0219169083151502179055506020820151816001015560408201518160020160006101000a8154816001600160401b0302191690836001600160401b0316021790555060608201518160020160086101000a8154816001600160401b0302191690836001600160401b0316021790555060808201518160020160106101000a81548160ff02191690831515021790555090505088600101600360008d63ffffffff1663ffffffff16815260200190815260200160002060000160156101000a8154816001600160401b0302191690836001600160401b0316021790555089600360008d63ffffffff1663ffffffff16815260200190815260200160002060010160006101000a8154816001600160401b0302191690836001600160401b0316021790555060006002600360008e63ffffffff1663ffffffff168152602001908152602001600020600201544203600360008f63ffffffff1663ffffffff168152602001908152602001600020600401540181611e9757fe5b63ffffffff8e16600090815260036020819052604090912001549190049150600202811115611eea5763ffffffff8c16600090815260036020819052604090912090810154600202600490910155611f06565b63ffffffff8c1660009081526003602052604090206004018190555b505063ffffffff8a166000908152600360205260409020426002909101555b5050505063ffffffff861660009081526003602052604090205495966000199590950195600160981b900461ffff16949293505050565b6000546001600160a01b03163314611fa9576040805162461bcd60e51b8152602060048201819052602482015260008051602061298c833981519152604482015290519081900360640190fd5b6000546001600160a01b0382811691161415611ff65760405162461bcd60e51b81526004018080602001828103825260248152602001806129686024913960400191505060405180910390fd5b6001600160a01b03166000908152600760205260409020805460ff19166001179055565b6000546001600160a01b03163314612067576040805162461bcd60e51b8152602060048201819052602482015260008051602061298c833981519152604482015290519081900360640190fd5b60005463ffffffff600160a81b9091048116908216106120c5576040805162461bcd60e51b8152602060048201526014602482015273636861696e4964203e3d206e756d436861696e7360601b604482015290519081900360640190fd5b63ffffffff811660009081526003602052604090205460ff1615612130576040805162461bcd60e51b815260206004820152601e60248201527f636861696e735b636861696e49645d2e657869737473203d3d20747275650000604482015290519081900360640190fd5b63ffffffff166000908152600360205260409020805460ff19166001179055565b6000546001600160a01b0316331461219e576040805162461bcd60e51b8152602060048201819052602482015260008051602061298c833981519152604482015290519081900360640190fd5b60008054600160a81b900463ffffffff1681526003602052604090205460ff1615612209576040805162461bcd60e51b8152602060048201526016602482015275636861696e496420616c72656164792065786973747360501b604482015290519081900360640190fd5b60008361ffff161161225b576040805162461bcd60e51b81526020600482015260166024820152750636c61696d506572696f644c656e677468203d3d20360541b604482015290519081900360640190fd5b604051806101400160405280600115158152602001866001600160401b03168152602001856001600160401b031681526020018461ffff1681526020018361ffff16815260200160006001600160401b03168152602001866001600160401b031681526020014281526020018281526020016000815250600360008060159054906101000a900463ffffffff1663ffffffff1663ffffffff16815260200190815260200160002060008201518160000160006101000a81548160ff02191690831515021790555060208201518160000160016101000a8154816001600160401b0302191690836001600160401b0316021790555060408201518160000160096101000a8154816001600160401b0302191690836001600160401b0316021790555060608201518160000160116101000a81548161ffff021916908361ffff16021790555060808201518160000160136101000a81548161ffff021916908361ffff16021790555060a08201518160000160156101000a8154816001600160401b0302191690836001600160401b0316021790555060c08201518160010160006101000a8154816001600160401b0302191690836001600160401b0316021790555060e0820151816002015561010082015181600301556101208201518160040155905050600060159054906101000a900463ffffffff16600101600060156101000a81548163ffffffff021916908363ffffffff1602179055505050505050565b63ffffffff8216600090815260036020526040812054839060ff166124d6576040805162461bcd60e51b81526020600482015260166024820152600080516020612ade833981519152604482015290519081900360640190fd5b50506040805160e09390931b6001600160e01b03191660208085019190915260c09290921b6001600160c01b03191660248401528051808403600c018152602c909301815282519282019290922060009081526004909152205460ff1690565b6000546001600160a01b031690565b60008054600160a01b900460ff161561259c576040805162461bcd60e51b8152602060048201526014602482015273696e697469616c6973656420213d2066616c736560601b604482015290519081900360640190fd5b73fffec6c83c8bf5c3f4ae0ccf8c45ce20e4560bd76000806101000a8154816001600160a01b0302191690836001600160a01b031602179055506040518061014001604052806001151581526020016303bf79006001600160401b0316815260200160006001600160401b03168152602001601e61ffff168152602001600061ffff16815260200160006001600160401b031681526020016303bf79006001600160401b03168152602001428152602001607881526020016000815250600360008063ffffffff16815260200190815260200160002060008201518160000160006101000a81548160ff02191690831515021790555060208201518160000160016101000a8154816001600160401b0302191690836001600160401b0316021790555060408201518160000160096101000a8154816001600160401b0302191690836001600160401b0316021790555060608201518160000160116101000a81548161ffff021916908361ffff16021790555060808201518160000160136101000a81548161ffff021916908361ffff16021790555060a08201518160000160156101000a8154816001600160401b0302191690836001600160401b0316021790555060c08201518160010160006101000a8154816001600160401b0302191690836001600160401b0316021790555060e08201518160020155610100820151816003015561012082015181600401559050506001600060156101000a81548163ffffffff021916908363ffffffff16021790555062093a80600260006101000a8154816001600160401b0302191690836001600160401b03160217905550426001819055506001600060146101000a81548160ff0219169083151502179055506001905090565b600060015442116128665760405162461bcd60e51b8152600401808060200182810382526021815260200180612a5e6021913960400191505060405180910390fd5b6002546001546001600160401b039091169042038161288157fe5b0490509056fe66696e616c697365645061796d656e74735b636861696e49645d5b74784964486173685d2e696e646578203e3d206c65646765726e6f7420656e6f7567682074696d6520656c61707365642073696e6365207072696f722066696e616c697479636861696e735b636861696e49645d2e66696e616c697365644c6564676572496e646578202d206c6564676572203e3d20636861696e735b636861696e49645d2e6c6564676572486973746f727953697a656c6564676572203e3d20636861696e735b636861696e49645d2e66696e616c697365644c6564676572496e646578626c6f636b656441646472657373203d3d20676f7665726e616e6365436f6e74726163746d73672e73656e64657220213d20676f7665726e616e6365436f6e7472616374636861696e735b636861696e49645d2e66696e616c697365644c6564676572496e646578202d20636861696e735b636861696e49645d2e67656e657369734c6564676572203c3d20636861696e735b636861696e49645d2e6c6564676572486973746f727953697a6570726576696f757320636c61696d20706572696f64206e6f74207965742066696e616c6973656474686973206163636f756e7420697320676f7665726e616e636520626c6f636b6564626c6f636b2e74696d657374616d70203c3d20696e697469616c69736554696d6521676f7665726e616e6365426c6f636b65644163636f756e74735b626c6f636b6564416464726573735d626c6f636b2e74696d657374616d70203c3d20636861696e735b636861696e49645d2e66696e616c6973656454696d657374616d70636861696e496420646f6573206e6f7420657869737400000000000000000000a264697066735822122049cabc01c587deb83acf6c3dd9c693a835d3436cdf9df20745fae7a2e4d1f54c64736f6c63430007060033"
	      },
	      "1000000000000000000000000000000000000002": {
	        "balance": "0x0",
	        "code": "0x608060405234801561001057600080fd5b50600436106100885760003560e01c8063b172b2221161005b578063b172b222146100c1578063c49f561f146100e5578063caf874ce1461010b578063d11e4c211461012857610088565b806344ed2b151461008d578063592e6f59146100a75780637fec8d38146100b15780638be2fb86146100b9575b600080fd5b61009561014e565b60408051918252519081900360200190f35b6100af610153565b005b6100af6101e0565b6100956102da565b6100c96102e0565b604080516001600160a01b039092168252519081900360200190f35b6100af600480360360208110156100fb57600080fd5b50356001600160a01b03166102ef565b6100c96004803603602081101561012157600080fd5b50356104ae565b6100af6004803603602081101561013e57600080fd5b50356001600160a01b03166104d8565b600a81565b600254600160a01b900460ff16156101a9576040805162461bcd60e51b8152602060048201526014602482015273696e697469616c6973656420213d2066616c736560601b604482015290519081900360640190fd5b6002805460ff60a01b196001600160a01b031990911673ff50ef6f4b0568493175defa3655b10d68bf41fb1716600160a01b179055565b600154431161022b576040805162461bcd60e51b8152602060048201526012602482015271189b1bd8dacb9b9d5b58995c881cdb585b1b60721b604482015290519081900360640190fd5b4360015560008054905b818110156102d6576000818154811061024a57fe5b9060005260206000200160009054906101000a90046001600160a01b03166001600160a01b031663e4d06d826040518163ffffffff1660e01b8152600401602060405180830381600087803b1580156102a257600080fd5b505af11580156102b6573d6000803e3d6000fd5b505050506040513d60208110156102cc57600080fd5b5050600101610235565b5050565b60015481565b6002546001600160a01b031681565b6002546001600160a01b0316331461034e576040805162461bcd60e51b815260206004820181905260248201527f6d73672e73656e64657220213d20676f7665726e616e6365436f6e7472616374604482015290519081900360640190fd5b60008054905b81811015610467576000818154811061036957fe5b6000918252602090912001546001600160a01b038481169116141561045f576000600183038154811061039857fe5b600091825260208220015481546001600160a01b039091169190839081106103bc57fe5b6000918252602082200180546001600160a01b0319166001600160a01b0393909316929092179091558054806103ee57fe5b60008281526020808220830160001990810180546001600160a01b0319169055909201909255604080516001600160a01b03871681529182019290925281517f7b11c8af33e77c52fff95f7c830b6b76307fde6ed54c82a4aa96457ac07d2c72929181900390910190a150506104ab565b600101610354565b506040805162461bcd60e51b815260206004820152601360248201527210d85b89dd08199a5b990818dbdb9d1c9858dd606a1b604482015290519081900360640190fd5b50565b600081815481106104be57600080fd5b6000918252602090912001546001600160a01b0316905081565b6002546001600160a01b03163314610537576040805162461bcd60e51b815260206004820181905260248201527f6d73672e73656e64657220213d20676f7665726e616e6365436f6e7472616374604482015290519081900360640190fd5b600054600a6001820110610587576040805162461bcd60e51b8152602060048201526012602482015271546f6f206d616e7920636f6e74726163747360701b604482015290519081900360640190fd5b60005b818110156105cd576000818154811061059f57fe5b6000918252602090912001546001600160a01b03848116911614156105c55750506104ab565b60010161058a565b5060008054600180820183559180527f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e5630180546001600160a01b0319166001600160a01b03851690811790915560408051918252602082019290925281517f7b11c8af33e77c52fff95f7c830b6b76307fde6ed54c82a4aa96457ac07d2c72929181900390910190a1505056fea2646970667358221220d67e71ef2de0c0da48193efe122f5ee6fbd627355682cf5f0dfe9249b19319c064736f6c63430007060033"
	      },
	      "ff50eF6F4b0568493175defa3655b10d68Bf41FB": {
	        "balance": "0x314dc6448d9338c15B0a00000000"
	      },
	      "ff898D83DE2F1E07ad44f9Ff34bB1ABDBCfcBB00": {
	        "balance": "0x314dc6448d9338c15B0a00000000"
	      },
	      "ff31f7568813E69991fAeCA13907141cc4874723": {
	        "balance": "0x314dc6448d9338c15B0a00000000"
	      },
	      "ffF9AcF70B7aFaFAe6C495aEEDC0eD5B0EF4011e": {
	        "balance": "0x314dc6448d9338c15B0a00000000"
	      },
	      "ff89975844E384a1798b0cD24D7611F44Dd17040": {
	        "balance": "0x314dc6448d9338c15B0a00000000"
	      },
		  "ff65397290C660596bFf1564E333f870E577F4D2": {
	        "balance": "0x314dc6448d9338c15B0a00000000"
	      },
	      "ff57CaF5B871db64F2a7F4C5bc2d17A5E666F7E8": {
	        "balance": "0x314dc6448d9338c15B0a00000000"
	      },
	      "ffC11262622D5069aBad729efe84a95C169d9c06": {
	        "balance": "0x314dc6448d9338c15B0a00000000"
	      }
	    },
	    "number": "0x0",
	    "gasUsed": "0x0",
	    "parentHash": "0x0000000000000000000000000000000000000000000000000000000000000000"
	  }`
)
