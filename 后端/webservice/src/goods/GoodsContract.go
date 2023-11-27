// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package goods

import (
	"math/big"
	"strings"

	"github.com/FISCO-BCOS/go-sdk/abi"
	"github.com/FISCO-BCOS/go-sdk/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
)

// GoodsContractBasicInfo is an auto generated low-level Go binding around an user-defined struct.
type GoodsContractBasicInfo struct {
	ImgUrl       string
	Owner        common.Address
	GoodsName    string
	GoodsDescrip string
}

// GoodsContractDealInfo is an auto generated low-level Go binding around an user-defined struct.
type GoodsContractDealInfo struct {
	DealPrice *big.Int
	DealTime  *big.Int
	Buyer     common.Address
	Seller    common.Address
}

// EshouGoodsABI is the input ABI used to generate the binding from.
const EshouGoodsABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"story\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"imgUrl\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"goodsName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"goodsDescrip\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"createGoods\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"dealPrice\",\"type\":\"uint256\"}],\"name\":\"createTrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"id2Goods\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"story\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"goodsId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"imgUrl\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"goodsName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"goodsDescrip\",\"type\":\"string\"}],\"internalType\":\"structGoodsContract.BasicInfo\",\"name\":\"basicInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"dealPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dealTime\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"}],\"internalType\":\"structGoodsContract.DealInfo\",\"name\":\"dealInfo\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"isSold\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"}],\"name\":\"indexByBuyer\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"}],\"name\":\"indexBySeller\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gId\",\"type\":\"uint256\"}],\"name\":\"indexBygid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"indexGoodsonList\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"indexStoryById\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gId\",\"type\":\"uint256\"}],\"name\":\"isGoodExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// EshouGoodsBin is the compiled bytecode used for deploying new contracts.
var EshouGoodsBin = "0x608060405234801561001057600080fd5b5060008081905550611858806100276000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c80637b0c8a91116100665780637b0c8a91146101365780638299b99214610166578063af26806e14610196578063f770d8ec146101b4578063f8277d9d146101d057610093565b8063094569f2146100985780633827a1a2146100b45780634de0f19b146100e45780635bbeb12814610102575b600080fd5b6100b260048036038101906100ad9190610f25565b610200565b005b6100ce60048036038101906100c99190610f78565b6103a1565b6040516100db9190610fc0565b60405180910390f35b6100ec6103f9565b6040516100f99190611099565b60405180910390f35b61011c60048036038101906101179190610f78565b6104e8565b60405161012d959493929190611282565b60405180910390f35b610150600480360381019061014b91906112e4565b610893565b60405161015d9190611099565b60405180910390f35b610180600480360381019061017b91906112e4565b6109c1565b60405161018d9190611099565b60405180910390f35b61019e610aef565b6040516101ab91906113d3565b60405180910390f35b6101ce60048036038101906101c9919061152a565b610c29565b005b6101ea60048036038101906101e59190610f78565b610d83565b6040516101f79190611627565b60405180910390f35b610209836103a1565b610248576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161023f9061168e565b60405180910390fd5b600061025384610d83565b90506000600160008381526020019081526020016000209050838160060160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508060020160010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168160060160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550828160060160000181905550428160060160010181905550838160020160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600181600a0160006101000a81548160ff0219169083151502179055505050505050565b600080600090505b6000548110156103ee5782600160008381526020019081526020016000206001015414156103db5760019150506103f4565b80806103e6906116dd565b9150506103a9565b50600090505b919050565b60606000805467ffffffffffffffff811115610418576104176113ff565b5b6040519080825280602002602001820160405280156104465781602001602082028036833780820191505090505b5090506000805b6000548110156104df576000151560016000838152602001908152602001600020600a0160009054906101000a900460ff16151514156104cc5760016000828152602001908152602001600020600101548383815181106104b1576104b0611726565b5b60200260200101818152505081806104c8906116dd565b9250505b80806104d7906116dd565b91505061044d565b50819250505090565b600160205280600052604060002060009150905080600001805461050b90611784565b80601f016020809104026020016040519081016040528092919081815260200182805461053790611784565b80156105845780601f1061055957610100808354040283529160200191610584565b820191906000526020600020905b81548152906001019060200180831161056757829003601f168201915b505050505090806001015490806002016040518060800160405290816000820180546105af90611784565b80601f01602080910402602001604051908101604052809291908181526020018280546105db90611784565b80156106285780601f106105fd57610100808354040283529160200191610628565b820191906000526020600020905b81548152906001019060200180831161060b57829003601f168201915b505050505081526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160028201805461069790611784565b80601f01602080910402602001604051908101604052809291908181526020018280546106c390611784565b80156107105780601f106106e557610100808354040283529160200191610710565b820191906000526020600020905b8154815290600101906020018083116106f357829003601f168201915b5050505050815260200160038201805461072990611784565b80601f016020809104026020016040519081016040528092919081815260200182805461075590611784565b80156107a25780601f10610777576101008083540402835291602001916107a2565b820191906000526020600020905b81548152906001019060200180831161078557829003601f168201915b505050505081525050908060060160405180608001604052908160008201548152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815250509080600a0160009054906101000a900460ff16905085565b60606000805467ffffffffffffffff8111156108b2576108b16113ff565b5b6040519080825280602002602001820160405280156108e05781602001602082028036833780820191505090505b5090506000805b6000548110156109b6578473ffffffffffffffffffffffffffffffffffffffff166001600083815260200190815260200160002060060160030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614156109a357600160008281526020019081526020016000206001015483838151811061098857610987611726565b5b602002602001018181525050818061099f906116dd565b9250505b80806109ae906116dd565b9150506108e7565b508192505050919050565b60606000805467ffffffffffffffff8111156109e0576109df6113ff565b5b604051908082528060200260200182016040528015610a0e5781602001602082028036833780820191505090505b5090506000805b600054811015610ae4578473ffffffffffffffffffffffffffffffffffffffff166001600083815260200190815260200160002060060160020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415610ad1576001600082815260200190815260200160002060010154838381518110610ab657610ab5611726565b5b6020026020010181815250508180610acd906116dd565b9250505b8080610adc906116dd565b915050610a15565b508192505050919050565b60606000805467ffffffffffffffff811115610b0e57610b0d6113ff565b5b604051908082528060200260200182016040528015610b4157816020015b6060815260200190600190039081610b2c5790505b50905060005b600054811015610c2157600160008281526020019081526020016000206000018054610b7290611784565b80601f0160208091040260200160405190810160405280929190818152602001828054610b9e90611784565b8015610beb5780601f10610bc057610100808354040283529160200191610beb565b820191906000526020600020905b815481529060010190602001808311610bce57829003601f168201915b5050505050828281518110610c0357610c02611726565b5b60200260200101819052508080610c19906116dd565b915050610b47565b508091505090565b610c32866103a1565b15610c72576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c6990611802565b60405180910390fd5b600060016000805481526020019081526020016000209050600080815480929190610c9c906116dd565b919050555085816000019080519060200190610cb9929190610dda565b5086816001018190555084816002016000019080519060200190610cde929190610dda565b50818160020160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555083816002016002019080519060200190610d40929190610dda565b5082816002016003019080519060200190610d5c929190610dda565b50600081600a0160006101000a81548160ff02191690831515021790555050505050505050565b600080600090505b600054811015610dcf578260016000838152602001908152602001600020600101541415610dbc5780915050610dd5565b8080610dc7906116dd565b915050610d8b565b50600090505b919050565b828054610de690611784565b90600052602060002090601f016020900481019282610e085760008555610e4f565b82601f10610e2157805160ff1916838001178555610e4f565b82800160010185558215610e4f579182015b82811115610e4e578251825591602001919060010190610e33565b5b509050610e5c9190610e60565b5090565b5b80821115610e79576000816000905550600101610e61565b5090565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b610ea481610e91565b8114610eaf57600080fd5b50565b600081359050610ec181610e9b565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610ef282610ec7565b9050919050565b610f0281610ee7565b8114610f0d57600080fd5b50565b600081359050610f1f81610ef9565b92915050565b600080600060608486031215610f3e57610f3d610e87565b5b6000610f4c86828701610eb2565b9350506020610f5d86828701610f10565b9250506040610f6e86828701610eb2565b9150509250925092565b600060208284031215610f8e57610f8d610e87565b5b6000610f9c84828501610eb2565b91505092915050565b60008115159050919050565b610fba81610fa5565b82525050565b6000602082019050610fd56000830184610fb1565b92915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b61101081610e91565b82525050565b60006110228383611007565b60208301905092915050565b6000602082019050919050565b600061104682610fdb565b6110508185610fe6565b935061105b83610ff7565b8060005b8381101561108c5781516110738882611016565b975061107e8361102e565b92505060018101905061105f565b5085935050505092915050565b600060208201905081810360008301526110b3818461103b565b905092915050565b600081519050919050565b600082825260208201905092915050565b60005b838110156110f55780820151818401526020810190506110da565b83811115611104576000848401525b50505050565b6000601f19601f8301169050919050565b6000611126826110bb565b61113081856110c6565b93506111408185602086016110d7565b6111498161110a565b840191505092915050565b61115d81610e91565b82525050565b600082825260208201905092915050565b600061117f826110bb565b6111898185611163565b93506111998185602086016110d7565b6111a28161110a565b840191505092915050565b6111b681610ee7565b82525050565b600060808301600083015184820360008601526111d98282611174565b91505060208301516111ee60208601826111ad565b50604083015184820360408601526112068282611174565b915050606083015184820360608601526112208282611174565b9150508091505092915050565b6080820160008201516112436000850182611007565b5060208201516112566020850182611007565b50604082015161126960408501826111ad565b50606082015161127c60608501826111ad565b50505050565b600061010082019050818103600083015261129d818861111b565b90506112ac6020830187611154565b81810360408301526112be81866111bc565b90506112cd606083018561122d565b6112da60e0830184610fb1565b9695505050505050565b6000602082840312156112fa576112f9610e87565b5b600061130884828501610f10565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b60006113498383611174565b905092915050565b6000602082019050919050565b600061136982611311565b611373818561131c565b9350836020820285016113858561132d565b8060005b858110156113c157848403895281516113a2858261133d565b94506113ad83611351565b925060208a01995050600181019050611389565b50829750879550505050505092915050565b600060208201905081810360008301526113ed818461135e565b905092915050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6114378261110a565b810181811067ffffffffffffffff82111715611456576114556113ff565b5b80604052505050565b6000611469610e7d565b9050611475828261142e565b919050565b600067ffffffffffffffff821115611495576114946113ff565b5b61149e8261110a565b9050602081019050919050565b82818337600083830152505050565b60006114cd6114c88461147a565b61145f565b9050828152602081018484840111156114e9576114e86113fa565b5b6114f48482856114ab565b509392505050565b600082601f830112611511576115106113f5565b5b81356115218482602086016114ba565b91505092915050565b60008060008060008060c0878903121561154757611546610e87565b5b600061155589828a01610eb2565b965050602087013567ffffffffffffffff81111561157657611575610e8c565b5b61158289828a016114fc565b955050604087013567ffffffffffffffff8111156115a3576115a2610e8c565b5b6115af89828a016114fc565b945050606087013567ffffffffffffffff8111156115d0576115cf610e8c565b5b6115dc89828a016114fc565b935050608087013567ffffffffffffffff8111156115fd576115fc610e8c565b5b61160989828a016114fc565b92505060a061161a89828a01610f10565b9150509295509295509295565b600060208201905061163c6000830184611154565b92915050565b7f496e76616c696420676f6f647349640000000000000000000000000000000000600082015250565b6000611678600f836110c6565b915061168382611642565b602082019050919050565b600060208201905081810360008301526116a78161166b565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006116e882610e91565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82141561171b5761171a6116ae565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061179c57607f821691505b602082108114156117b0576117af611755565b5b50919050565b7f546f6b656e496420616c72656164792065786973747300000000000000000000600082015250565b60006117ec6016836110c6565b91506117f7826117b6565b602082019050919050565b6000602082019050818103600083015261181b816117df565b905091905056fea26469706673582212206fa14dadc538cdb9d46c7d37f28b9d5b75bd1abd9667e4b04bdfe8b0381242b464736f6c634300080b0033"

// DeployEshouGoods deploys a new contract, binding an instance of EshouGoods to it.
func DeployEshouGoods(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Receipt, *EshouGoods, error) {
	parsed, err := abi.JSON(strings.NewReader(EshouGoodsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, receipt, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EshouGoodsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, receipt, &EshouGoods{EshouGoodsCaller: EshouGoodsCaller{contract: contract}, EshouGoodsTransactor: EshouGoodsTransactor{contract: contract}, EshouGoodsFilterer: EshouGoodsFilterer{contract: contract}}, nil
}

func AsyncDeployEshouGoods(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(EshouGoodsABI))
	if err != nil {
		return nil, err
	}

	tx, err := bind.AsyncDeployContract(auth, handler, parsed, common.FromHex(EshouGoodsBin), backend)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// EshouGoods is an auto generated Go binding around a Solidity contract.
type EshouGoods struct {
	EshouGoodsCaller     // Read-only binding to the contract
	EshouGoodsTransactor // Write-only binding to the contract
	EshouGoodsFilterer   // Log filterer for contract events
}

// EshouGoodsCaller is an auto generated read-only Go binding around a Solidity contract.
type EshouGoodsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EshouGoodsTransactor is an auto generated write-only Go binding around a Solidity contract.
type EshouGoodsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EshouGoodsFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type EshouGoodsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EshouGoodsSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type EshouGoodsSession struct {
	Contract     *EshouGoods       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EshouGoodsCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type EshouGoodsCallerSession struct {
	Contract *EshouGoodsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// EshouGoodsTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type EshouGoodsTransactorSession struct {
	Contract     *EshouGoodsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// EshouGoodsRaw is an auto generated low-level Go binding around a Solidity contract.
type EshouGoodsRaw struct {
	Contract *EshouGoods // Generic contract binding to access the raw methods on
}

// EshouGoodsCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type EshouGoodsCallerRaw struct {
	Contract *EshouGoodsCaller // Generic read-only contract binding to access the raw methods on
}

// EshouGoodsTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type EshouGoodsTransactorRaw struct {
	Contract *EshouGoodsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEshouGoods creates a new instance of EshouGoods, bound to a specific deployed contract.
func NewEshouGoods(address common.Address, backend bind.ContractBackend) (*EshouGoods, error) {
	contract, err := bindEshouGoods(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EshouGoods{EshouGoodsCaller: EshouGoodsCaller{contract: contract}, EshouGoodsTransactor: EshouGoodsTransactor{contract: contract}, EshouGoodsFilterer: EshouGoodsFilterer{contract: contract}}, nil
}

// NewEshouGoodsCaller creates a new read-only instance of EshouGoods, bound to a specific deployed contract.
func NewEshouGoodsCaller(address common.Address, caller bind.ContractCaller) (*EshouGoodsCaller, error) {
	contract, err := bindEshouGoods(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EshouGoodsCaller{contract: contract}, nil
}

// NewEshouGoodsTransactor creates a new write-only instance of EshouGoods, bound to a specific deployed contract.
func NewEshouGoodsTransactor(address common.Address, transactor bind.ContractTransactor) (*EshouGoodsTransactor, error) {
	contract, err := bindEshouGoods(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EshouGoodsTransactor{contract: contract}, nil
}

// NewEshouGoodsFilterer creates a new log filterer instance of EshouGoods, bound to a specific deployed contract.
func NewEshouGoodsFilterer(address common.Address, filterer bind.ContractFilterer) (*EshouGoodsFilterer, error) {
	contract, err := bindEshouGoods(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EshouGoodsFilterer{contract: contract}, nil
}

// bindEshouGoods binds a generic wrapper to an already deployed contract.
func bindEshouGoods(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EshouGoodsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EshouGoods *EshouGoodsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EshouGoods.Contract.EshouGoodsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EshouGoods *EshouGoodsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _EshouGoods.Contract.EshouGoodsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EshouGoods *EshouGoodsRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _EshouGoods.Contract.EshouGoodsTransactor.contract.TransactWithResult(opts, result, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EshouGoods *EshouGoodsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EshouGoods.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EshouGoods *EshouGoodsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _EshouGoods.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EshouGoods *EshouGoodsTransactorRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _EshouGoods.Contract.contract.TransactWithResult(opts, result, method, params...)
}

// Id2Goods is a free data retrieval call binding the contract method 0x5bbeb128.
//
// Solidity: function id2Goods(uint256 ) constant returns(string story, uint256 goodsId, GoodsContractBasicInfo basicInfo, GoodsContractDealInfo dealInfo, bool isSold)
func (_EshouGoods *EshouGoodsCaller) Id2Goods(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Story     string
	GoodsId   *big.Int
	BasicInfo GoodsContractBasicInfo
	DealInfo  GoodsContractDealInfo
	IsSold    bool
}, error) {
	ret := new(struct {
		Story     string
		GoodsId   *big.Int
		BasicInfo GoodsContractBasicInfo
		DealInfo  GoodsContractDealInfo
		IsSold    bool
	})
	out := ret
	err := _EshouGoods.contract.Call(opts, out, "id2Goods", arg0)
	return *ret, err
}

// Id2Goods is a free data retrieval call binding the contract method 0x5bbeb128.
//
// Solidity: function id2Goods(uint256 ) constant returns(string story, uint256 goodsId, GoodsContractBasicInfo basicInfo, GoodsContractDealInfo dealInfo, bool isSold)
func (_EshouGoods *EshouGoodsSession) Id2Goods(arg0 *big.Int) (struct {
	Story     string
	GoodsId   *big.Int
	BasicInfo GoodsContractBasicInfo
	DealInfo  GoodsContractDealInfo
	IsSold    bool
}, error) {
	return _EshouGoods.Contract.Id2Goods(&_EshouGoods.CallOpts, arg0)
}

// Id2Goods is a free data retrieval call binding the contract method 0x5bbeb128.
//
// Solidity: function id2Goods(uint256 ) constant returns(string story, uint256 goodsId, GoodsContractBasicInfo basicInfo, GoodsContractDealInfo dealInfo, bool isSold)
func (_EshouGoods *EshouGoodsCallerSession) Id2Goods(arg0 *big.Int) (struct {
	Story     string
	GoodsId   *big.Int
	BasicInfo GoodsContractBasicInfo
	DealInfo  GoodsContractDealInfo
	IsSold    bool
}, error) {
	return _EshouGoods.Contract.Id2Goods(&_EshouGoods.CallOpts, arg0)
}

// IndexByBuyer is a free data retrieval call binding the contract method 0x8299b992.
//
// Solidity: function indexByBuyer(address buyer) constant returns(uint256[])
func (_EshouGoods *EshouGoodsCaller) IndexByBuyer(opts *bind.CallOpts, buyer common.Address) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _EshouGoods.contract.Call(opts, out, "indexByBuyer", buyer)
	return *ret0, err
}

// IndexByBuyer is a free data retrieval call binding the contract method 0x8299b992.
//
// Solidity: function indexByBuyer(address buyer) constant returns(uint256[])
func (_EshouGoods *EshouGoodsSession) IndexByBuyer(buyer common.Address) ([]*big.Int, error) {
	return _EshouGoods.Contract.IndexByBuyer(&_EshouGoods.CallOpts, buyer)
}

// IndexByBuyer is a free data retrieval call binding the contract method 0x8299b992.
//
// Solidity: function indexByBuyer(address buyer) constant returns(uint256[])
func (_EshouGoods *EshouGoodsCallerSession) IndexByBuyer(buyer common.Address) ([]*big.Int, error) {
	return _EshouGoods.Contract.IndexByBuyer(&_EshouGoods.CallOpts, buyer)
}

// IndexBySeller is a free data retrieval call binding the contract method 0x7b0c8a91.
//
// Solidity: function indexBySeller(address seller) constant returns(uint256[])
func (_EshouGoods *EshouGoodsCaller) IndexBySeller(opts *bind.CallOpts, seller common.Address) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _EshouGoods.contract.Call(opts, out, "indexBySeller", seller)
	return *ret0, err
}

// IndexBySeller is a free data retrieval call binding the contract method 0x7b0c8a91.
//
// Solidity: function indexBySeller(address seller) constant returns(uint256[])
func (_EshouGoods *EshouGoodsSession) IndexBySeller(seller common.Address) ([]*big.Int, error) {
	return _EshouGoods.Contract.IndexBySeller(&_EshouGoods.CallOpts, seller)
}

// IndexBySeller is a free data retrieval call binding the contract method 0x7b0c8a91.
//
// Solidity: function indexBySeller(address seller) constant returns(uint256[])
func (_EshouGoods *EshouGoodsCallerSession) IndexBySeller(seller common.Address) ([]*big.Int, error) {
	return _EshouGoods.Contract.IndexBySeller(&_EshouGoods.CallOpts, seller)
}

// IndexBygid is a free data retrieval call binding the contract method 0xf8277d9d.
//
// Solidity: function indexBygid(uint256 gId) constant returns(uint256)
func (_EshouGoods *EshouGoodsCaller) IndexBygid(opts *bind.CallOpts, gId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EshouGoods.contract.Call(opts, out, "indexBygid", gId)
	return *ret0, err
}

// IndexBygid is a free data retrieval call binding the contract method 0xf8277d9d.
//
// Solidity: function indexBygid(uint256 gId) constant returns(uint256)
func (_EshouGoods *EshouGoodsSession) IndexBygid(gId *big.Int) (*big.Int, error) {
	return _EshouGoods.Contract.IndexBygid(&_EshouGoods.CallOpts, gId)
}

// IndexBygid is a free data retrieval call binding the contract method 0xf8277d9d.
//
// Solidity: function indexBygid(uint256 gId) constant returns(uint256)
func (_EshouGoods *EshouGoodsCallerSession) IndexBygid(gId *big.Int) (*big.Int, error) {
	return _EshouGoods.Contract.IndexBygid(&_EshouGoods.CallOpts, gId)
}

// IndexGoodsonList is a free data retrieval call binding the contract method 0x4de0f19b.
//
// Solidity: function indexGoodsonList() constant returns(uint256[])
func (_EshouGoods *EshouGoodsCaller) IndexGoodsonList(opts *bind.CallOpts) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _EshouGoods.contract.Call(opts, out, "indexGoodsonList")
	return *ret0, err
}

// IndexGoodsonList is a free data retrieval call binding the contract method 0x4de0f19b.
//
// Solidity: function indexGoodsonList() constant returns(uint256[])
func (_EshouGoods *EshouGoodsSession) IndexGoodsonList() ([]*big.Int, error) {
	return _EshouGoods.Contract.IndexGoodsonList(&_EshouGoods.CallOpts)
}

// IndexGoodsonList is a free data retrieval call binding the contract method 0x4de0f19b.
//
// Solidity: function indexGoodsonList() constant returns(uint256[])
func (_EshouGoods *EshouGoodsCallerSession) IndexGoodsonList() ([]*big.Int, error) {
	return _EshouGoods.Contract.IndexGoodsonList(&_EshouGoods.CallOpts)
}

// IndexStoryById is a free data retrieval call binding the contract method 0xaf26806e.
//
// Solidity: function indexStoryById() constant returns(string[])
func (_EshouGoods *EshouGoodsCaller) IndexStoryById(opts *bind.CallOpts) ([]string, error) {
	var (
		ret0 = new([]string)
	)
	out := ret0
	err := _EshouGoods.contract.Call(opts, out, "indexStoryById")
	return *ret0, err
}

// IndexStoryById is a free data retrieval call binding the contract method 0xaf26806e.
//
// Solidity: function indexStoryById() constant returns(string[])
func (_EshouGoods *EshouGoodsSession) IndexStoryById() ([]string, error) {
	return _EshouGoods.Contract.IndexStoryById(&_EshouGoods.CallOpts)
}

// IndexStoryById is a free data retrieval call binding the contract method 0xaf26806e.
//
// Solidity: function indexStoryById() constant returns(string[])
func (_EshouGoods *EshouGoodsCallerSession) IndexStoryById() ([]string, error) {
	return _EshouGoods.Contract.IndexStoryById(&_EshouGoods.CallOpts)
}

// IsGoodExist is a free data retrieval call binding the contract method 0x3827a1a2.
//
// Solidity: function isGoodExist(uint256 gId) constant returns(bool)
func (_EshouGoods *EshouGoodsCaller) IsGoodExist(opts *bind.CallOpts, gId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _EshouGoods.contract.Call(opts, out, "isGoodExist", gId)
	return *ret0, err
}

// IsGoodExist is a free data retrieval call binding the contract method 0x3827a1a2.
//
// Solidity: function isGoodExist(uint256 gId) constant returns(bool)
func (_EshouGoods *EshouGoodsSession) IsGoodExist(gId *big.Int) (bool, error) {
	return _EshouGoods.Contract.IsGoodExist(&_EshouGoods.CallOpts, gId)
}

// IsGoodExist is a free data retrieval call binding the contract method 0x3827a1a2.
//
// Solidity: function isGoodExist(uint256 gId) constant returns(bool)
func (_EshouGoods *EshouGoodsCallerSession) IsGoodExist(gId *big.Int) (bool, error) {
	return _EshouGoods.Contract.IsGoodExist(&_EshouGoods.CallOpts, gId)
}

// CreateGoods is a paid mutator transaction binding the contract method 0xf770d8ec.
//
// Solidity: function createGoods(uint256 gId, string story, string imgUrl, string goodsName, string goodsDescrip, address owner) returns()
func (_EshouGoods *EshouGoodsTransactor) CreateGoods(opts *bind.TransactOpts, gId *big.Int, story string, imgUrl string, goodsName string, goodsDescrip string, owner common.Address) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _EshouGoods.contract.TransactWithResult(opts, out, "createGoods", gId, story, imgUrl, goodsName, goodsDescrip, owner)
	return transaction, receipt, err
}

func (_EshouGoods *EshouGoodsTransactor) AsyncCreateGoods(handler func(*types.Receipt, error), opts *bind.TransactOpts, gId *big.Int, story string, imgUrl string, goodsName string, goodsDescrip string, owner common.Address) (*types.Transaction, error) {
	return _EshouGoods.contract.AsyncTransact(opts, handler, "createGoods", gId, story, imgUrl, goodsName, goodsDescrip, owner)
}

// CreateGoods is a paid mutator transaction binding the contract method 0xf770d8ec.
//
// Solidity: function createGoods(uint256 gId, string story, string imgUrl, string goodsName, string goodsDescrip, address owner) returns()
func (_EshouGoods *EshouGoodsSession) CreateGoods(gId *big.Int, story string, imgUrl string, goodsName string, goodsDescrip string, owner common.Address) (*types.Transaction, *types.Receipt, error) {
	return _EshouGoods.Contract.CreateGoods(&_EshouGoods.TransactOpts, gId, story, imgUrl, goodsName, goodsDescrip, owner)
}

func (_EshouGoods *EshouGoodsSession) AsyncCreateGoods(handler func(*types.Receipt, error), gId *big.Int, story string, imgUrl string, goodsName string, goodsDescrip string, owner common.Address) (*types.Transaction, error) {
	return _EshouGoods.Contract.AsyncCreateGoods(handler, &_EshouGoods.TransactOpts, gId, story, imgUrl, goodsName, goodsDescrip, owner)
}

// CreateGoods is a paid mutator transaction binding the contract method 0xf770d8ec.
//
// Solidity: function createGoods(uint256 gId, string story, string imgUrl, string goodsName, string goodsDescrip, address owner) returns()
func (_EshouGoods *EshouGoodsTransactorSession) CreateGoods(gId *big.Int, story string, imgUrl string, goodsName string, goodsDescrip string, owner common.Address) (*types.Transaction, *types.Receipt, error) {
	return _EshouGoods.Contract.CreateGoods(&_EshouGoods.TransactOpts, gId, story, imgUrl, goodsName, goodsDescrip, owner)
}

func (_EshouGoods *EshouGoodsTransactorSession) AsyncCreateGoods(handler func(*types.Receipt, error), gId *big.Int, story string, imgUrl string, goodsName string, goodsDescrip string, owner common.Address) (*types.Transaction, error) {
	return _EshouGoods.Contract.AsyncCreateGoods(handler, &_EshouGoods.TransactOpts, gId, story, imgUrl, goodsName, goodsDescrip, owner)
}

// CreateTrade is a paid mutator transaction binding the contract method 0x094569f2.
//
// Solidity: function createTrade(uint256 gId, address buyer, uint256 dealPrice) returns()
func (_EshouGoods *EshouGoodsTransactor) CreateTrade(opts *bind.TransactOpts, gId *big.Int, buyer common.Address, dealPrice *big.Int) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _EshouGoods.contract.TransactWithResult(opts, out, "createTrade", gId, buyer, dealPrice)
	return transaction, receipt, err
}

func (_EshouGoods *EshouGoodsTransactor) AsyncCreateTrade(handler func(*types.Receipt, error), opts *bind.TransactOpts, gId *big.Int, buyer common.Address, dealPrice *big.Int) (*types.Transaction, error) {
	return _EshouGoods.contract.AsyncTransact(opts, handler, "createTrade", gId, buyer, dealPrice)
}

// CreateTrade is a paid mutator transaction binding the contract method 0x094569f2.
//
// Solidity: function createTrade(uint256 gId, address buyer, uint256 dealPrice) returns()
func (_EshouGoods *EshouGoodsSession) CreateTrade(gId *big.Int, buyer common.Address, dealPrice *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _EshouGoods.Contract.CreateTrade(&_EshouGoods.TransactOpts, gId, buyer, dealPrice)
}

func (_EshouGoods *EshouGoodsSession) AsyncCreateTrade(handler func(*types.Receipt, error), gId *big.Int, buyer common.Address, dealPrice *big.Int) (*types.Transaction, error) {
	return _EshouGoods.Contract.AsyncCreateTrade(handler, &_EshouGoods.TransactOpts, gId, buyer, dealPrice)
}

// CreateTrade is a paid mutator transaction binding the contract method 0x094569f2.
//
// Solidity: function createTrade(uint256 gId, address buyer, uint256 dealPrice) returns()
func (_EshouGoods *EshouGoodsTransactorSession) CreateTrade(gId *big.Int, buyer common.Address, dealPrice *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _EshouGoods.Contract.CreateTrade(&_EshouGoods.TransactOpts, gId, buyer, dealPrice)
}

func (_EshouGoods *EshouGoodsTransactorSession) AsyncCreateTrade(handler func(*types.Receipt, error), gId *big.Int, buyer common.Address, dealPrice *big.Int) (*types.Transaction, error) {
	return _EshouGoods.Contract.AsyncCreateTrade(handler, &_EshouGoods.TransactOpts, gId, buyer, dealPrice)
}
