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
	        "code": "0x608060405234801561001057600080fd5b50600436106100a95760003560e01c8063a57d0e2511610071578063a57d0e2514610384578063e54caf92146103fe578063ef2fa85f14610444578063ef4c169e14610468578063efe7827214610470578063ff695ef0146104b2576100a9565b80631129753f146100ae57806313bb431c146100d65780631d4ded8e1461023f5780632a2434a2146102985780634cf5913314610313575b600080fd5b6100d4600480360360208110156100c457600080fd5b50356001600160a01b03166104db565b005b6101a3600480360360a08110156100ec57600080fd5b63ffffffff823516916001600160401b036020820135169160408201359160608101359181019060a08101608082013564010000000081111561012e57600080fd5b82018360208201111561014057600080fd5b8035906020019184600183028401116401000000008311171561016257600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061055c945050505050565b604051808563ffffffff168152602001846001600160401b0316815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b838110156102015781810151838201526020016101e9565b50505050905090810190601f16801561022e5780820380516001836020036101000a031916815260200191505b509550505050505060405180910390f35b61027f6004803603608081101561025557600080fd5b506001600160401b038135169061ffff60208201358116916040810135909116906060013561080b565b6040805163ffffffff9092168252519081900360200190f35b6102bb600480360360208110156102ae57600080fd5b503563ffffffff16610b45565b60405180876001600160401b03168152602001866001600160401b031681526020018561ffff168152602001846001600160401b03168152602001838152602001828152602001965050505050505060405180910390f35b610369600480360360e081101561032957600080fd5b5063ffffffff813516906020810135906001600160401b03604082013581169160608101359160808201359160a081013582169160c09091013516610c0f565b60408051921515835260208301919091528051918290030190f35b6103c66004803603608081101561039a57600080fd5b5063ffffffff813516906001600160401b03602082013581169160408101359091169060600135610e3c565b6040805163ffffffff90951685526001600160401b03909316602085015261ffff909116838301526060830152519081900360800190f35b6104306004803603604081101561041457600080fd5b50803563ffffffff1690602001356001600160401b0316611559565b604080519115158252519081900360200190f35b61044c61161a565b604080516001600160a01b039092168252519081900360200190f35b610430611629565b6104966004803603602081101561048657600080fd5b50356001600160a01b0316611889565b604080516001600160401b039092168252519081900360200190f35b6100d4600480360360408110156104c857600080fd5b5063ffffffff81351690602001356118ad565b6000546001600160a01b0316331461053a576040805162461bcd60e51b815260206004820181905260248201527f6d73672e73656e64657220213d20676f7665726e616e6365436f6e7472616374604482015290519081900360640190fd5b600080546001600160a01b0319166001600160a01b0392909216919091179055565b63ffffffff85166000908152600160208190526040822054829182916060918a9160ff161515146105c2576040805162461bcd60e51b81526020600482015260166024820152600080516020611bc3833981519152604482015290519081900360640190fd5b6000866040516020018082805190602001908083835b602083106105f75780518252601f1990920191602091820191016105d8565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051602081830303815290604052805190602001209050600360008c63ffffffff1663ffffffff168152602001908152602001600020600082815260200190815260200160002060000160009054906101000a900460ff16151560001515146106c3576040805162461bcd60e51b81526020600482015260136024820152723a3c24b21030b63932b0b23c90383937bb32b760691b604482015290519081900360640190fd5b60006106d08a8d8d61198b565b9050413314806106e3575041600160981b145b610734576040805162461bcd60e51b815260206004820152601c60248201527f496e76616c696420626c6f636b2e636f696e626173652076616c756500000000604482015290519081900360640190fd5b4133148015610747575041600160981b14155b156107d25760405180606001604052806001151581526020018a815260200182815250600360008e63ffffffff1663ffffffff168152602001908152602001600020600084815260200190815260200160002060008201518160000160006101000a81548160ff02191690831515021790555060208201518160010155604082015181600201559050505b5050505063ffffffff88166000908152600160205260409020549798600160a81b9098046001600160401b031697949650929450505050565b600080546001600160a01b0316331461086b576040805162461bcd60e51b815260206004820181905260248201527f6d73672e73656e64657220213d20676f7665726e616e6365436f6e7472616374604482015290519081900360640190fd5b60008461ffff16116108bd576040805162461bcd60e51b81526020600482015260166024820152750636c61696d506572696f644c656e677468203d3d20360541b604482015290519081900360640190fd5b6000546001600160a01b03164114806108d9575041600160981b145b61092a576040805162461bcd60e51b815260206004820152601c60248201527f496e76616c696420626c6f636b2e636f696e626173652076616c756500000000604482015290519081900360640190fd5b600054600160a81b810463ffffffff16906001600160a01b031641148015610956575041600160981b14155b15610b3c57604051806101200160405280600115158152602001876001600160401b031681526020018661ffff1681526020018561ffff16815260200160006001600160401b03168152602001876001600160401b031681526020014281526020018481526020016000815250600160008060159054906101000a900463ffffffff1660010163ffffffff1663ffffffff16815260200190815260200160002060008201518160000160006101000a81548160ff02191690831515021790555060208201518160000160016101000a8154816001600160401b0302191690836001600160401b0316021790555060408201518160000160096101000a81548161ffff021916908361ffff160217905550606082015181600001600b6101000a81548161ffff021916908361ffff160217905550608082015181600001600d6101000a8154816001600160401b0302191690836001600160401b0316021790555060a08201518160000160156101000a8154816001600160401b0302191690836001600160401b0316021790555060c0820151816001015560e082015181600201556101008201518160030155905050600060159054906101000a900463ffffffff16600101600060156101000a81548163ffffffff021916908363ffffffff1602179055505b95945050505050565b63ffffffff8116600090815260016020819052604082205482918291829182918291889160ff16151514610bae576040805162461bcd60e51b81526020600482015260166024820152600080516020611bc3833981519152604482015290519081900360640190fd5b5050505063ffffffff9390931660009081526001602081905260409091208054918101546003909101546001600160401b036101008404811697600160681b85048216975061ffff600160481b8604169650600160a81b9094041693509091565b63ffffffff8716600090815260016020819052604082205482918a9160ff16151514610c70576040805162461bcd60e51b81526020600482015260166024820152600080516020611bc3833981519152604482015290519081900360640190fd5b63ffffffff8a1660009081526003602090815260408083208c845290915290205460ff161515600114610ce0576040805162461bcd60e51b81526020600482015260136024820152721d1e125908191bd95cc81b9bdd08195e1a5cdd606a1b604482015290519081900360640190fd5b604080516001600160401b03808b1660208084019190915283518084038201815283850185528051908201208983166060808601919091528551808603909101815260808501865280519083012092891660a0808601919091528551808603909101815260c08501865280519083012060e085018f905261010085019190915261012084018c905261014084018b905261016084019290925261018080840192909252835180840390920182526101a0909201835280519082012063ffffffff8d166000908152600383528381208d8252909252919020600101548114610e04576040805162461bcd60e51b81526020600482015260136024820152720d2dcecc2d8d2c840e0c2f2dacadce890c2e6d606b1b604482015290519081900360640190fd5b5063ffffffff8a1660009081526003602090815260408083208c84529091529020600201546001935091505097509795505050505050565b63ffffffff84166000908152600160208190526040822054829182918291899160ff16151514610ea1576040805162461bcd60e51b81526020600482015260166024820152600080516020611bc3833981519152604482015290519081900360640190fd5b63ffffffff8916600090815260016020526040902054600160481b810461ffff16600160a81b9091046001600160401b0390811691909101811690891614610f21576040805162461bcd60e51b815260206004820152600e60248201526d34b73b30b634b2103632b233b2b960911b604482015290519081900360640190fd5b63ffffffff89166000908152600160205260409020546001600160401b03888116600160681b9092041614610f9d576040805162461bcd60e51b815260206004820152601860248201527f696e76616c696420636c61696d506572696f64496e6465780000000000000000604482015290519081900360640190fd5b63ffffffff8916600090815260016020819052604090912001544211610ff45760405162461bcd60e51b8152600401808060200182810382526035815260200180611b8e6035913960400191505060405180910390fd5b63ffffffff8916600090815260016020526040902060028082015460039092015402101561108e5763ffffffff8916600090815260016020819052604090912060038082015491909201546002909102429190910390910210156110895760405162461bcd60e51b815260040180806020018281038252602c815260200180611b09602c913960400191505060405180910390fd5b6110f0565b63ffffffff8916600090815260016020819052604090912060038101549101544203600f0110156110f05760405162461bcd60e51b815260040180806020018281038252602c815260200180611b09602c913960400191505060405180910390fd5b604080516001600160e01b031960e08c901b166020808301919091526001600160c01b031960c08b901b1660248301528251600c818403018152602c9092018352815191810191909120600081815260029092529190205460ff161561119d576040805162461bcd60e51b815260206004820152601e60248201527f6c6f636174696f6e4861736820616c72656164792066696e616c697365640000604482015290519081900360640190fd5b6001600160401b0388161561124c57604080516001600160e01b031960e08d901b166020808301919091526001600160c01b03196000198c0160c01b1660248301528251600c818403018152602c9092018352815191810191909120600081815260029092529190205460ff16151560011461124a5760405162461bcd60e51b8152600401808060200182810382526027815260200180611b676027913960400191505060405180910390fd5b505b4133148061125d575041600160981b145b6112ae576040805162461bcd60e51b815260206004820152601c60248201527f496e76616c696420626c6f636b2e636f696e626173652076616c756500000000604482015290519081900360640190fd5b41331480156112c1575041600160981b14155b156115225760046000336001600160a01b03166001600160a01b0316815260200190815260200160002060009054906101000a90046001600160401b031660010160046000336001600160a01b03166001600160a01b0316815260200190815260200160002060006101000a8154816001600160401b0302191690836001600160401b031602179055506040518060600160405280600115158152602001888152602001428152506002600083815260200190815260200160002060008201518160000160006101000a81548160ff021916908315150217905550602082015181600101556040820151816002015590505087600101600160008c63ffffffff1663ffffffff168152602001908152602001600020600001600d6101000a8154816001600160401b0302191690836001600160401b0316021790555088600160008c63ffffffff1663ffffffff16815260200190815260200160002060000160156101000a8154816001600160401b0302191690836001600160401b0316021790555060006002600160008d63ffffffff1663ffffffff168152602001908152602001600020600101544203600160008e63ffffffff1663ffffffff16815260200190815260200160002060030154018161149857fe5b63ffffffff8d166000908152600160205260409020600290810154929091049250028111156114e85763ffffffff8b16600090815260016020526040902060028082015402600390910155611504565b63ffffffff8b1660009081526001602052604090206003018190555b5063ffffffff8a166000908152600160208190526040909120429101555b5050505063ffffffff861660009081526001602052604090205495966000199590950195600160581b900461ffff16949293505050565b63ffffffff82166000908152600160208190526040822054849160ff9091161515146115ba576040805162461bcd60e51b81526020600482015260166024820152600080516020611bc3833981519152604482015290519081900360640190fd5b50506040805160e09390931b6001600160e01b03191660208085019190915260c09290921b6001600160c01b03191660248401528051808403600c018152602c909301815282519282019290922060009081526002909152205460ff1690565b6000546001600160a01b031690565b60008054600160a01b900460ff1615611680576040805162461bcd60e51b8152602060048201526014602482015273696e697469616c6973656420213d2066616c736560601b604482015290519081900360640190fd5b6001609c1b6000806101000a8154816001600160a01b0302191690836001600160a01b031602179055506040518061012001604052806001151581526020016303b589a46001600160401b03168152602001601e61ffff168152602001600061ffff16815260200160006001600160401b031681526020016303b589a46001600160401b03168152602001428152602001607881526020016000815250600160008063ffffffff16815260200190815260200160002060008201518160000160006101000a81548160ff02191690831515021790555060208201518160000160016101000a8154816001600160401b0302191690836001600160401b0316021790555060408201518160000160096101000a81548161ffff021916908361ffff160217905550606082015181600001600b6101000a81548161ffff021916908361ffff160217905550608082015181600001600d6101000a8154816001600160401b0302191690836001600160401b0316021790555060a08201518160000160156101000a8154816001600160401b0302191690836001600160401b0316021790555060c0820151816001015560e0820151816002015561010082015181600301559050506001600060156101000a81548163ffffffff021916908363ffffffff1602179055506001600060146101000a81548160ff0219169083151502179055506001905090565b6001600160a01b03166000908152600460205260409020546001600160401b031690565b6000546001600160a01b0316331461190c576040805162461bcd60e51b815260206004820181905260248201527f6d73672e73656e64657220213d20676f7665726e616e6365436f6e7472616374604482015290519081900360640190fd5b63ffffffff8216600090815260016020819052604090912054839160ff90911615151461196e576040805162461bcd60e51b81526020600482015260166024820152600080516020611bc3833981519152604482015290519081900360640190fd5b5063ffffffff909116600090815260016020526040902060020155565b63ffffffff82166000908152600160208190526040822054849160ff9091161515146119ec576040805162461bcd60e51b81526020600482015260166024820152600080516020611bc3833981519152604482015290519081900360640190fd5b604080516001600160e01b031960e087901b166020808301919091526001600160c01b031960c087901b1660248301528251600c818403018152602c9092018352815191810191909120600081815260029092529190205460ff161515600114611a875760405162461bcd60e51b8152600401808060200182810382526032815260200180611b356032913960400191505060405180910390fd5b6000818152600260205260409020600101548614611aec576040805162461bcd60e51b815260206004820152601760248201527f496e76616c696420636c61696d506572696f6448617368000000000000000000604482015290519081900360640190fd5b600090815260026020819052604090912001549594505050505056fe6e6f7420656e6f7567682074696d6520656c61707365642073696e6365207072696f722066696e616c69747966696e616c69736564436c61696d506572696f64735b6c6f636174696f6e486173685d20646f6573206e6f7420657869737470726576696f757320636c61696d20706572696f64206e6f74207965742066696e616c69736564626c6f636b2e74696d657374616d70203c3d20636861696e735b636861696e49645d2e66696e616c6973656454696d657374616d70636861696e496420646f6573206e6f7420657869737400000000000000000000a2646970667358221220fa21a99d8c78614c01460f1ec0965e7d841c578db4f188515ac2d019a12c1fb164736f6c63430007060033"
	      },
	      "1000000000000000000000000000000000000002": {
	        "balance": "0x0",
	        "code": "0x608060405234801561001057600080fd5b50600436106100885760003560e01c8063b172b2221161005b578063b172b222146100df578063c49f561f14610103578063caf874ce14610129578063d11e4c211461014657610088565b806344ed2b151461008d5780637fec8d38146100a75780638be2fb86146100b15780639d6a890f146100b9575b600080fd5b61009561016c565b60408051918252519081900360200190f35b6100af610171565b005b61009561026b565b6100af600480360360208110156100cf57600080fd5b50356001600160a01b0316610271565b6100e76102f7565b604080516001600160a01b039092168252519081900360200190f35b6100af6004803603602081101561011957600080fd5b50356001600160a01b0316610306565b6100e76004803603602081101561013f57600080fd5b50356104c5565b6100af6004803603602081101561015c57600080fd5b50356001600160a01b03166104ef565b600a81565b60015443116101bc576040805162461bcd60e51b8152602060048201526012602482015271189b1bd8dacb9b9d5b58995c881cdb585b1b60721b604482015290519081900360640190fd5b4360015560008054905b8181101561026757600081815481106101db57fe5b9060005260206000200160009054906101000a90046001600160a01b03166001600160a01b031663e4d06d826040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561023357600080fd5b505af1158015610247573d6000803e3d6000fd5b505050506040513d602081101561025d57600080fd5b50506001016101c6565b5050565b60015481565b600254600160a01b900460ff16156102c7576040805162461bcd60e51b8152602060048201526014602482015273696e697469616c6973656420213d2066616c736560601b604482015290519081900360640190fd5b6002805460ff60a01b196001600160a01b039093166001600160a01b03199091161791909116600160a01b179055565b6002546001600160a01b031681565b6002546001600160a01b03163314610365576040805162461bcd60e51b815260206004820181905260248201527f6d73672e73656e64657220213d20676f7665726e616e6365436f6e7472616374604482015290519081900360640190fd5b60008054905b8181101561047e576000818154811061038057fe5b6000918252602090912001546001600160a01b038481169116141561047657600060018303815481106103af57fe5b600091825260208220015481546001600160a01b039091169190839081106103d357fe5b6000918252602082200180546001600160a01b0319166001600160a01b03939093169290921790915580548061040557fe5b60008281526020808220830160001990810180546001600160a01b0319169055909201909255604080516001600160a01b03871681529182019290925281517f7b11c8af33e77c52fff95f7c830b6b76307fde6ed54c82a4aa96457ac07d2c72929181900390910190a150506104c2565b60010161036b565b506040805162461bcd60e51b815260206004820152601360248201527210d85b89dd08199a5b990818dbdb9d1c9858dd606a1b604482015290519081900360640190fd5b50565b600081815481106104d557600080fd5b6000918252602090912001546001600160a01b0316905081565b6002546001600160a01b0316331461054e576040805162461bcd60e51b815260206004820181905260248201527f6d73672e73656e64657220213d20676f7665726e616e6365436f6e7472616374604482015290519081900360640190fd5b600054600a600182011061059e576040805162461bcd60e51b8152602060048201526012602482015271546f6f206d616e7920636f6e74726163747360701b604482015290519081900360640190fd5b60005b818110156105e457600081815481106105b657fe5b6000918252602090912001546001600160a01b03848116911614156105dc5750506104c2565b6001016105a1565b5060008054600180820183559180527f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e5630180546001600160a01b0319166001600160a01b03851690811790915560408051918252602082019290925281517f7b11c8af33e77c52fff95f7c830b6b76307fde6ed54c82a4aa96457ac07d2c72929181900390910190a1505056fea26469706673582212209d8987b4893f97f9874d79bc52657593a74d06b674c357c48083b893e5d96bbd64736f6c63430007060033"
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
