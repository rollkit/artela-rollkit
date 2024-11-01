// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package prediction

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ElectionPredictionMarketMetaData contains all meta data concerning the ElectionPredictionMarket contract.
var ElectionPredictionMarketMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_betToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_electionName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_initialDemocratOdds\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ELECTION_END_TIME\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"betToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimWinnings\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"electionMarket\",\"inputs\":[],\"outputs\":[{\"name\":\"electionName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"totalPoolSize\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isResolved\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"winner\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"democratOdds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"republicanOdds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lastOddsUpdate\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMarketInfo\",\"inputs\":[],\"outputs\":[{\"name\":\"electionName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"endTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalPoolSize\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isResolved\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"democratPool\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"republicanPool\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"democratOdds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"republicanOdds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lastOddsUpdate\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUserBets\",\"inputs\":[{\"name\":\"_user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"parties\",\"type\":\"uint8[]\",\"internalType\":\"uint8[]\"},{\"name\":\"amounts\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"claimed\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"placeBet\",\"inputs\":[{\"name\":\"_party\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"resolveMarket\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateOdds\",\"inputs\":[{\"name\":\"_democratOdds\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"userBets\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"party\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"claimed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"BetPlaced\",\"inputs\":[{\"name\":\"bettor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"party\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MarketResolved\",\"inputs\":[{\"name\":\"winner\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OddsUpdated\",\"inputs\":[{\"name\":\"democratOdds\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"republicanOdds\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WinningsClaimed\",\"inputs\":[{\"name\":\"bettor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]}]",
	Bin: "0x60a060405234801561001057600080fd5b5060405161176738038061176783398101604081905261002f916101f5565b6001600055338061005b57604051631e4fbdf760e01b8152600060048201526024015b60405180910390fd5b6100648161018d565b506001600160a01b0383166100bb5760405162461bcd60e51b815260206004820152601560248201527f496e76616c696420746f6b656e206164647265737300000000000000000000006044820152606401610052565b61271081111561010d5760405162461bcd60e51b815260206004820152601560248201527f4f646473206d757374206265203c3d20313030303000000000000000000000006044820152606401610052565b6001600160a01b03831660805260026101268382610376565b50600681905561013881612710610434565b600755426008557f401f5799446884f85b1967999c05db7393e8ce1fd302dc57ad43d6c1fafe9e298161016d81612710610434565b6040805192835260208301919091520160405180910390a150505061045b565b600180546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b634e487b7160e01b600052604160045260246000fd5b60008060006060848603121561020a57600080fd5b83516001600160a01b038116811461022157600080fd5b60208501519093506001600160401b0381111561023d57600080fd5b8401601f8101861361024e57600080fd5b80516001600160401b03811115610267576102676101df565b604051601f8201601f19908116603f011681016001600160401b0381118282101715610295576102956101df565b6040528181528282016020018810156102ad57600080fd5b60005b828110156102cc576020818501810151838301820152016102b0565b50600091810160200191909152604095909501519396949550929392505050565b600181811c9082168061030157607f821691505b60208210810361032157634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111561037157806000526020600020601f840160051c8101602085101561034e5750805b601f840160051c820191505b8181101561036e576000815560010161035a565b50505b505050565b81516001600160401b0381111561038f5761038f6101df565b6103a38161039d84546102ed565b84610327565b6020601f8211600181146103d757600083156103bf5750848201515b600019600385901b1c1916600184901b17845561036e565b600084815260208120601f198516915b8281101561040757878501518255602094850194600190920191016103e7565b50848210156104255786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b8181038181111561045557634e487b7160e01b600052601160045260246000fd5b92915050565b6080516112e3610484600039600081816101630152818161093f0152610b6b01526112e36000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c806378691f161161008c578063b401faf111610066578063b401faf1146101e2578063dcb3b30e146101ea578063eb50f6bc146101fd578063f2fde38b1461021057600080fd5b806378691f161461015e5780638068aa681461019d5780638da5cb5b146101d157600080fd5b806303b7338e146100d457806323341a05146100f25780632b9b55b31461010f5780632d56657d146101195780635536f2621461013b578063715018a614610156575b600080fd5b6100df63671ed3ff81565b6040519081526020015b60405180910390f35b6100fa610223565b6040516100e999989796959493929190610f5d565b610117610355565b005b61012c610127366004610fcd565b610446565b6040516100e99392919061102d565b61014361063d565b6040516100e997969594939291906110c5565b6101176106fd565b6101857f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016100e9565b6101b06101ab36600461110c565b610711565b6040805160ff909416845260208401929092521515908201526060016100e9565b6001546001600160a01b0316610185565b610117610758565b6101176101f8366004611136565b610a39565b61011761020b36600461115a565b610d1c565b61011761021e366004610fcd565b610e30565b60035460045460056020527f05b8ccbb9d4d8fb16ea74ce3c29a41f1b461fbdaff4714a0d9a8eb05499746bc54600160009081527f1471eb6eb2c5e789fc3de43f8ce62938c7d1836ec861730447e2ada8fd81017b5460065460075460085460028054606099969889988998899889988998899889989763671ed3ff97959660ff90961695949392919089906102b890611173565b80601f01602080910402602001604051908101604052809291908181526020018280546102e490611173565b80156103315780601f1061030657610100808354040283529160200191610331565b820191906000526020600020905b81548152906001019060200180831161031457829003601f168201915b50505050509850985098509850985098509850985098509850909192939495969798565b61035d610e6e565b63671ed3ff4210156103ab5760405162461bcd60e51b8152602060048201526012602482015271115b1958dd1a5bdb881b9bdd08195b99195960721b60448201526064015b60405180910390fd5b60045460ff16156103ce5760405162461bcd60e51b81526004016103a2906111ad565b600754600654600091106103e35760016103e6565b60005b6004805460ff83166101000261ffff199091161760011790556040519091507f94e5f71d5001613eb5e4ad876acb1da127582270c692c34fbcf810be21728b1d9061043b90839060ff91909116815260200190565b60405180910390a150565b6001600160a01b03811660009081526009602052604090208054606091829182919067ffffffffffffffff811115610480576104806111e4565b6040519080825280602002602001820160405280156104a9578160200160208202803683370190505b50815490945067ffffffffffffffff8111156104c7576104c76111e4565b6040519080825280602002602001820160405280156104f0578160200160208202803683370190505b50815490935067ffffffffffffffff81111561050e5761050e6111e4565b604051908082528060200260200182016040528015610537578160200160208202803683370190505b50915060005b815481101561063457818181548110610558576105586111fa565b6000918252602090912060039091020154855160ff90911690869083908110610583576105836111fa565b602002602001019060ff16908160ff16815250508181815481106105a9576105a96111fa565b9060005260206000209060030201600101548482815181106105cd576105cd6111fa565b6020026020010181815250508181815481106105eb576105eb6111fa565b906000526020600020906003020160020160009054906101000a900460ff1683828151811061061c5761061c6111fa565b9115156020928302919091019091015260010161053d565b50509193909250565b60028054819061064c90611173565b80601f016020809104026020016040519081016040528092919081815260200182805461067890611173565b80156106c55780601f1061069a576101008083540402835291602001916106c5565b820191906000526020600020905b8154815290600101906020018083116106a857829003601f168201915b50505060018401546002850154600486015460058701546006909701549596929560ff80841696506101009093049092169350919087565b610705610e6e565b61070f6000610e9b565b565b6009602052816000526040600020818154811061072d57600080fd5b600091825260209091206003909102018054600182015460029092015460ff91821694509192501683565b610760610eed565b60045460ff166107a85760405162461bcd60e51b815260206004820152601360248201527213585c9ad95d081b9bdd081c995cdbdb1d9959606a1b60448201526064016103a2565b336000908152600960205260408120815b81548110156108db578181815481106107d4576107d46111fa565b600091825260209091206002600390920201015460ff1615801561082a5750600454825461010090910460ff1690839083908110610814576108146111fa565b600091825260209091206003909102015460ff16145b156108d357600454610100900460ff166000908152600560205260408120546003548454919291839190869086908110610866576108666111fa565b9060005260206000209060030201600101546108829190611226565b61088c9190611243565b90506108988186611265565b945060018484815481106108ae576108ae6111fa565b60009182526020909120600390910201600201805460ff191691151591909117905550505b6001016107b9565b50600082116109235760405162461bcd60e51b81526020600482015260146024820152734e6f2077696e6e696e677320746f20636c61696d60601b60448201526064016103a2565b60405163a9059cbb60e01b8152336004820152602481018390527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063a9059cbb906044016020604051808303816000875af1158015610990573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109b49190611278565b6109f85760405162461bcd60e51b8152602060048201526015602482015274151bdad95b881d1c985b9cd9995c8819985a5b1959605a1b60448201526064016103a2565b60405182815233907f1a31e733a172afcf46074b3106c17f0c298e226442682a03c1e99ce256139ec29060200160405180910390a2505061070f6001600055565b610a41610eed565b60008111610a915760405162461bcd60e51b815260206004820152601b60248201527f42657420616d6f756e74206d75737420626520706f736974697665000000000060448201526064016103a2565b60018260ff161115610adc5760405162461bcd60e51b8152602060048201526014602482015273496e76616c69642070617274792063686f69636560601b60448201526064016103a2565b63671ed3ff4210610b265760405162461bcd60e51b815260206004820152601460248201527310995d1d1a5b99c81c195c9a5bd908195b99195960621b60448201526064016103a2565b60045460ff1615610b495760405162461bcd60e51b81526004016103a2906111ad565b6040516323b872dd60e01b8152336004820152306024820152604481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906323b872dd906064016020604051808303816000875af1158015610bbc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610be09190611278565b610c245760405162461bcd60e51b8152602060048201526015602482015274151bdad95b881d1c985b9cd9995c8819985a5b1959605a1b60448201526064016103a2565b8060026001016000828254610c399190611265565b909155505060ff821660009081526005602052604081208054839290610c60908490611265565b9091555050336000818152600960209081526040808320815160608101835260ff888116808352828601898152838601888152855460018082018855968a52988890209451600390990290940180549890931660ff199889161783555193820193909355905160029091018054911515919095161790935580519283529082018490527f9f6b4cced64244a64465098f682029975b1614400b96bc6bd12c668083719166910160405180910390a2610d186001600055565b5050565b610d24610e6e565b612710811115610d6e5760405162461bcd60e51b815260206004820152601560248201527404f646473206d757374206265203c3d20313030303605c1b60448201526064016103a2565b63671ed3ff4210610db25760405162461bcd60e51b815260206004820152600e60248201526d115b1958dd1a5bdb88195b99195960921b60448201526064016103a2565b60045460ff1615610dd55760405162461bcd60e51b81526004016103a2906111ad565b6006819055610de68161271061129a565b600755426008557f401f5799446884f85b1967999c05db7393e8ce1fd302dc57ad43d6c1fafe9e2981610e1b8161271061129a565b6040805192835260208301919091520161043b565b610e38610e6e565b6001600160a01b038116610e6257604051631e4fbdf760e01b8152600060048201526024016103a2565b610e6b81610e9b565b50565b6001546001600160a01b0316331461070f5760405163118cdaa760e01b81523360048201526024016103a2565b600180546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600260005403610f1057604051633ee5aeb560e01b815260040160405180910390fd5b6002600055565b6000815180845260005b81811015610f3d57602081850181015186830182015201610f21565b506000602082860101526020601f19601f83011685010191505092915050565b61012081526000610f7261012083018c610f17565b602083019a909a525060408101979097529415156060870152608086019390935260a085019190915260c084015260e083015261010090910152919050565b80356001600160a01b0381168114610fc857600080fd5b919050565b600060208284031215610fdf57600080fd5b610fe882610fb1565b9392505050565b600081518084526020840193506020830160005b828110156110235781511515865260209586019590910190600101611003565b5093949350505050565b6060808252845190820181905260009060208601906080840190835b8181101561106a57835160ff16835260209384019390920191600101611049565b50508381036020808601919091528651808352918101925086019060005b818110156110a6578251845260209384019390920191600101611088565b50505082810360408401526110bb8185610fef565b9695505050505050565b60e0815260006110d860e083018a610f17565b602083019890985250941515604086015260ff939093166060850152608084019190915260a083015260c090910152919050565b6000806040838503121561111f57600080fd5b61112883610fb1565b946020939093013593505050565b6000806040838503121561114957600080fd5b823560ff8116811461112857600080fd5b60006020828403121561116c57600080fd5b5035919050565b600181811c9082168061118757607f821691505b6020821081036111a757634e487b7160e01b600052602260045260246000fd5b50919050565b60208082526017908201527f4d61726b657420616c7265616479207265736f6c766564000000000000000000604082015260600190565b634e487b7160e01b600052604160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b808202811582820484141761123d5761123d611210565b92915050565b60008261126057634e487b7160e01b600052601260045260246000fd5b500490565b8082018082111561123d5761123d611210565b60006020828403121561128a57600080fd5b81518015158114610fe857600080fd5b8181038181111561123d5761123d61121056fea2646970667358221220eafb7e6122d85cf13da72b56656079fb175480286435f7b74b3ab293b50ba55c64736f6c634300081a0033",
}

// ElectionPredictionMarketABI is the input ABI used to generate the binding from.
// Deprecated: Use ElectionPredictionMarketMetaData.ABI instead.
var ElectionPredictionMarketABI = ElectionPredictionMarketMetaData.ABI

// ElectionPredictionMarketBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ElectionPredictionMarketMetaData.Bin instead.
var ElectionPredictionMarketBin = ElectionPredictionMarketMetaData.Bin

// DeployElectionPredictionMarket deploys a new Ethereum contract, binding an instance of ElectionPredictionMarket to it.
func DeployElectionPredictionMarket(auth *bind.TransactOpts, backend bind.ContractBackend, _betToken common.Address, _electionName string, _initialDemocratOdds *big.Int) (common.Address, *types.Transaction, *ElectionPredictionMarket, error) {
	parsed, err := ElectionPredictionMarketMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ElectionPredictionMarketBin), backend, _betToken, _electionName, _initialDemocratOdds)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ElectionPredictionMarket{ElectionPredictionMarketCaller: ElectionPredictionMarketCaller{contract: contract}, ElectionPredictionMarketTransactor: ElectionPredictionMarketTransactor{contract: contract}, ElectionPredictionMarketFilterer: ElectionPredictionMarketFilterer{contract: contract}}, nil
}

// ElectionPredictionMarket is an auto generated Go binding around an Ethereum contract.
type ElectionPredictionMarket struct {
	ElectionPredictionMarketCaller     // Read-only binding to the contract
	ElectionPredictionMarketTransactor // Write-only binding to the contract
	ElectionPredictionMarketFilterer   // Log filterer for contract events
}

// ElectionPredictionMarketCaller is an auto generated read-only Go binding around an Ethereum contract.
type ElectionPredictionMarketCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ElectionPredictionMarketTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ElectionPredictionMarketTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ElectionPredictionMarketFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ElectionPredictionMarketFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ElectionPredictionMarketSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ElectionPredictionMarketSession struct {
	Contract     *ElectionPredictionMarket // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ElectionPredictionMarketCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ElectionPredictionMarketCallerSession struct {
	Contract *ElectionPredictionMarketCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// ElectionPredictionMarketTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ElectionPredictionMarketTransactorSession struct {
	Contract     *ElectionPredictionMarketTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// ElectionPredictionMarketRaw is an auto generated low-level Go binding around an Ethereum contract.
type ElectionPredictionMarketRaw struct {
	Contract *ElectionPredictionMarket // Generic contract binding to access the raw methods on
}

// ElectionPredictionMarketCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ElectionPredictionMarketCallerRaw struct {
	Contract *ElectionPredictionMarketCaller // Generic read-only contract binding to access the raw methods on
}

// ElectionPredictionMarketTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ElectionPredictionMarketTransactorRaw struct {
	Contract *ElectionPredictionMarketTransactor // Generic write-only contract binding to access the raw methods on
}

// NewElectionPredictionMarket creates a new instance of ElectionPredictionMarket, bound to a specific deployed contract.
func NewElectionPredictionMarket(address common.Address, backend bind.ContractBackend) (*ElectionPredictionMarket, error) {
	contract, err := bindElectionPredictionMarket(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ElectionPredictionMarket{ElectionPredictionMarketCaller: ElectionPredictionMarketCaller{contract: contract}, ElectionPredictionMarketTransactor: ElectionPredictionMarketTransactor{contract: contract}, ElectionPredictionMarketFilterer: ElectionPredictionMarketFilterer{contract: contract}}, nil
}

// NewElectionPredictionMarketCaller creates a new read-only instance of ElectionPredictionMarket, bound to a specific deployed contract.
func NewElectionPredictionMarketCaller(address common.Address, caller bind.ContractCaller) (*ElectionPredictionMarketCaller, error) {
	contract, err := bindElectionPredictionMarket(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ElectionPredictionMarketCaller{contract: contract}, nil
}

// NewElectionPredictionMarketTransactor creates a new write-only instance of ElectionPredictionMarket, bound to a specific deployed contract.
func NewElectionPredictionMarketTransactor(address common.Address, transactor bind.ContractTransactor) (*ElectionPredictionMarketTransactor, error) {
	contract, err := bindElectionPredictionMarket(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ElectionPredictionMarketTransactor{contract: contract}, nil
}

// NewElectionPredictionMarketFilterer creates a new log filterer instance of ElectionPredictionMarket, bound to a specific deployed contract.
func NewElectionPredictionMarketFilterer(address common.Address, filterer bind.ContractFilterer) (*ElectionPredictionMarketFilterer, error) {
	contract, err := bindElectionPredictionMarket(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ElectionPredictionMarketFilterer{contract: contract}, nil
}

// bindElectionPredictionMarket binds a generic wrapper to an already deployed contract.
func bindElectionPredictionMarket(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ElectionPredictionMarketMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ElectionPredictionMarket *ElectionPredictionMarketRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ElectionPredictionMarket.Contract.ElectionPredictionMarketCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ElectionPredictionMarket *ElectionPredictionMarketRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ElectionPredictionMarket.Contract.ElectionPredictionMarketTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ElectionPredictionMarket *ElectionPredictionMarketRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ElectionPredictionMarket.Contract.ElectionPredictionMarketTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ElectionPredictionMarket *ElectionPredictionMarketCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ElectionPredictionMarket.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ElectionPredictionMarket *ElectionPredictionMarketTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ElectionPredictionMarket.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ElectionPredictionMarket *ElectionPredictionMarketTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ElectionPredictionMarket.Contract.contract.Transact(opts, method, params...)
}

// ELECTIONENDTIME is a free data retrieval call binding the contract method 0x03b7338e.
//
// Solidity: function ELECTION_END_TIME() view returns(uint256)
func (_ElectionPredictionMarket *ElectionPredictionMarketCaller) ELECTIONENDTIME(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ElectionPredictionMarket.contract.Call(opts, &out, "ELECTION_END_TIME")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ELECTIONENDTIME is a free data retrieval call binding the contract method 0x03b7338e.
//
// Solidity: function ELECTION_END_TIME() view returns(uint256)
func (_ElectionPredictionMarket *ElectionPredictionMarketSession) ELECTIONENDTIME() (*big.Int, error) {
	return _ElectionPredictionMarket.Contract.ELECTIONENDTIME(&_ElectionPredictionMarket.CallOpts)
}

// ELECTIONENDTIME is a free data retrieval call binding the contract method 0x03b7338e.
//
// Solidity: function ELECTION_END_TIME() view returns(uint256)
func (_ElectionPredictionMarket *ElectionPredictionMarketCallerSession) ELECTIONENDTIME() (*big.Int, error) {
	return _ElectionPredictionMarket.Contract.ELECTIONENDTIME(&_ElectionPredictionMarket.CallOpts)
}

// BetToken is a free data retrieval call binding the contract method 0x78691f16.
//
// Solidity: function betToken() view returns(address)
func (_ElectionPredictionMarket *ElectionPredictionMarketCaller) BetToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ElectionPredictionMarket.contract.Call(opts, &out, "betToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BetToken is a free data retrieval call binding the contract method 0x78691f16.
//
// Solidity: function betToken() view returns(address)
func (_ElectionPredictionMarket *ElectionPredictionMarketSession) BetToken() (common.Address, error) {
	return _ElectionPredictionMarket.Contract.BetToken(&_ElectionPredictionMarket.CallOpts)
}

// BetToken is a free data retrieval call binding the contract method 0x78691f16.
//
// Solidity: function betToken() view returns(address)
func (_ElectionPredictionMarket *ElectionPredictionMarketCallerSession) BetToken() (common.Address, error) {
	return _ElectionPredictionMarket.Contract.BetToken(&_ElectionPredictionMarket.CallOpts)
}

// ElectionMarket is a free data retrieval call binding the contract method 0x5536f262.
//
// Solidity: function electionMarket() view returns(string electionName, uint256 totalPoolSize, bool isResolved, uint8 winner, uint256 democratOdds, uint256 republicanOdds, uint256 lastOddsUpdate)
func (_ElectionPredictionMarket *ElectionPredictionMarketCaller) ElectionMarket(opts *bind.CallOpts) (struct {
	ElectionName   string
	TotalPoolSize  *big.Int
	IsResolved     bool
	Winner         uint8
	DemocratOdds   *big.Int
	RepublicanOdds *big.Int
	LastOddsUpdate *big.Int
}, error) {
	var out []interface{}
	err := _ElectionPredictionMarket.contract.Call(opts, &out, "electionMarket")

	outstruct := new(struct {
		ElectionName   string
		TotalPoolSize  *big.Int
		IsResolved     bool
		Winner         uint8
		DemocratOdds   *big.Int
		RepublicanOdds *big.Int
		LastOddsUpdate *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ElectionName = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.TotalPoolSize = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.IsResolved = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.Winner = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.DemocratOdds = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.RepublicanOdds = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.LastOddsUpdate = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ElectionMarket is a free data retrieval call binding the contract method 0x5536f262.
//
// Solidity: function electionMarket() view returns(string electionName, uint256 totalPoolSize, bool isResolved, uint8 winner, uint256 democratOdds, uint256 republicanOdds, uint256 lastOddsUpdate)
func (_ElectionPredictionMarket *ElectionPredictionMarketSession) ElectionMarket() (struct {
	ElectionName   string
	TotalPoolSize  *big.Int
	IsResolved     bool
	Winner         uint8
	DemocratOdds   *big.Int
	RepublicanOdds *big.Int
	LastOddsUpdate *big.Int
}, error) {
	return _ElectionPredictionMarket.Contract.ElectionMarket(&_ElectionPredictionMarket.CallOpts)
}

// ElectionMarket is a free data retrieval call binding the contract method 0x5536f262.
//
// Solidity: function electionMarket() view returns(string electionName, uint256 totalPoolSize, bool isResolved, uint8 winner, uint256 democratOdds, uint256 republicanOdds, uint256 lastOddsUpdate)
func (_ElectionPredictionMarket *ElectionPredictionMarketCallerSession) ElectionMarket() (struct {
	ElectionName   string
	TotalPoolSize  *big.Int
	IsResolved     bool
	Winner         uint8
	DemocratOdds   *big.Int
	RepublicanOdds *big.Int
	LastOddsUpdate *big.Int
}, error) {
	return _ElectionPredictionMarket.Contract.ElectionMarket(&_ElectionPredictionMarket.CallOpts)
}

// GetMarketInfo is a free data retrieval call binding the contract method 0x23341a05.
//
// Solidity: function getMarketInfo() view returns(string electionName, uint256 endTime, uint256 totalPoolSize, bool isResolved, uint256 democratPool, uint256 republicanPool, uint256 democratOdds, uint256 republicanOdds, uint256 lastOddsUpdate)
func (_ElectionPredictionMarket *ElectionPredictionMarketCaller) GetMarketInfo(opts *bind.CallOpts) (struct {
	ElectionName   string
	EndTime        *big.Int
	TotalPoolSize  *big.Int
	IsResolved     bool
	DemocratPool   *big.Int
	RepublicanPool *big.Int
	DemocratOdds   *big.Int
	RepublicanOdds *big.Int
	LastOddsUpdate *big.Int
}, error) {
	var out []interface{}
	err := _ElectionPredictionMarket.contract.Call(opts, &out, "getMarketInfo")

	outstruct := new(struct {
		ElectionName   string
		EndTime        *big.Int
		TotalPoolSize  *big.Int
		IsResolved     bool
		DemocratPool   *big.Int
		RepublicanPool *big.Int
		DemocratOdds   *big.Int
		RepublicanOdds *big.Int
		LastOddsUpdate *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ElectionName = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.EndTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TotalPoolSize = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.IsResolved = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.DemocratPool = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.RepublicanPool = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.DemocratOdds = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.RepublicanOdds = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.LastOddsUpdate = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetMarketInfo is a free data retrieval call binding the contract method 0x23341a05.
//
// Solidity: function getMarketInfo() view returns(string electionName, uint256 endTime, uint256 totalPoolSize, bool isResolved, uint256 democratPool, uint256 republicanPool, uint256 democratOdds, uint256 republicanOdds, uint256 lastOddsUpdate)
func (_ElectionPredictionMarket *ElectionPredictionMarketSession) GetMarketInfo() (struct {
	ElectionName   string
	EndTime        *big.Int
	TotalPoolSize  *big.Int
	IsResolved     bool
	DemocratPool   *big.Int
	RepublicanPool *big.Int
	DemocratOdds   *big.Int
	RepublicanOdds *big.Int
	LastOddsUpdate *big.Int
}, error) {
	return _ElectionPredictionMarket.Contract.GetMarketInfo(&_ElectionPredictionMarket.CallOpts)
}

// GetMarketInfo is a free data retrieval call binding the contract method 0x23341a05.
//
// Solidity: function getMarketInfo() view returns(string electionName, uint256 endTime, uint256 totalPoolSize, bool isResolved, uint256 democratPool, uint256 republicanPool, uint256 democratOdds, uint256 republicanOdds, uint256 lastOddsUpdate)
func (_ElectionPredictionMarket *ElectionPredictionMarketCallerSession) GetMarketInfo() (struct {
	ElectionName   string
	EndTime        *big.Int
	TotalPoolSize  *big.Int
	IsResolved     bool
	DemocratPool   *big.Int
	RepublicanPool *big.Int
	DemocratOdds   *big.Int
	RepublicanOdds *big.Int
	LastOddsUpdate *big.Int
}, error) {
	return _ElectionPredictionMarket.Contract.GetMarketInfo(&_ElectionPredictionMarket.CallOpts)
}

// GetUserBets is a free data retrieval call binding the contract method 0x2d56657d.
//
// Solidity: function getUserBets(address _user) view returns(uint8[] parties, uint256[] amounts, bool[] claimed)
func (_ElectionPredictionMarket *ElectionPredictionMarketCaller) GetUserBets(opts *bind.CallOpts, _user common.Address) (struct {
	Parties []uint8
	Amounts []*big.Int
	Claimed []bool
}, error) {
	var out []interface{}
	err := _ElectionPredictionMarket.contract.Call(opts, &out, "getUserBets", _user)

	outstruct := new(struct {
		Parties []uint8
		Amounts []*big.Int
		Claimed []bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Parties = *abi.ConvertType(out[0], new([]uint8)).(*[]uint8)
	outstruct.Amounts = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.Claimed = *abi.ConvertType(out[2], new([]bool)).(*[]bool)

	return *outstruct, err

}

// GetUserBets is a free data retrieval call binding the contract method 0x2d56657d.
//
// Solidity: function getUserBets(address _user) view returns(uint8[] parties, uint256[] amounts, bool[] claimed)
func (_ElectionPredictionMarket *ElectionPredictionMarketSession) GetUserBets(_user common.Address) (struct {
	Parties []uint8
	Amounts []*big.Int
	Claimed []bool
}, error) {
	return _ElectionPredictionMarket.Contract.GetUserBets(&_ElectionPredictionMarket.CallOpts, _user)
}

// GetUserBets is a free data retrieval call binding the contract method 0x2d56657d.
//
// Solidity: function getUserBets(address _user) view returns(uint8[] parties, uint256[] amounts, bool[] claimed)
func (_ElectionPredictionMarket *ElectionPredictionMarketCallerSession) GetUserBets(_user common.Address) (struct {
	Parties []uint8
	Amounts []*big.Int
	Claimed []bool
}, error) {
	return _ElectionPredictionMarket.Contract.GetUserBets(&_ElectionPredictionMarket.CallOpts, _user)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ElectionPredictionMarket *ElectionPredictionMarketCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ElectionPredictionMarket.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ElectionPredictionMarket *ElectionPredictionMarketSession) Owner() (common.Address, error) {
	return _ElectionPredictionMarket.Contract.Owner(&_ElectionPredictionMarket.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ElectionPredictionMarket *ElectionPredictionMarketCallerSession) Owner() (common.Address, error) {
	return _ElectionPredictionMarket.Contract.Owner(&_ElectionPredictionMarket.CallOpts)
}

// UserBets is a free data retrieval call binding the contract method 0x8068aa68.
//
// Solidity: function userBets(address , uint256 ) view returns(uint8 party, uint256 amount, bool claimed)
func (_ElectionPredictionMarket *ElectionPredictionMarketCaller) UserBets(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Party   uint8
	Amount  *big.Int
	Claimed bool
}, error) {
	var out []interface{}
	err := _ElectionPredictionMarket.contract.Call(opts, &out, "userBets", arg0, arg1)

	outstruct := new(struct {
		Party   uint8
		Amount  *big.Int
		Claimed bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Party = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Claimed = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// UserBets is a free data retrieval call binding the contract method 0x8068aa68.
//
// Solidity: function userBets(address , uint256 ) view returns(uint8 party, uint256 amount, bool claimed)
func (_ElectionPredictionMarket *ElectionPredictionMarketSession) UserBets(arg0 common.Address, arg1 *big.Int) (struct {
	Party   uint8
	Amount  *big.Int
	Claimed bool
}, error) {
	return _ElectionPredictionMarket.Contract.UserBets(&_ElectionPredictionMarket.CallOpts, arg0, arg1)
}

// UserBets is a free data retrieval call binding the contract method 0x8068aa68.
//
// Solidity: function userBets(address , uint256 ) view returns(uint8 party, uint256 amount, bool claimed)
func (_ElectionPredictionMarket *ElectionPredictionMarketCallerSession) UserBets(arg0 common.Address, arg1 *big.Int) (struct {
	Party   uint8
	Amount  *big.Int
	Claimed bool
}, error) {
	return _ElectionPredictionMarket.Contract.UserBets(&_ElectionPredictionMarket.CallOpts, arg0, arg1)
}

// ClaimWinnings is a paid mutator transaction binding the contract method 0xb401faf1.
//
// Solidity: function claimWinnings() returns()
func (_ElectionPredictionMarket *ElectionPredictionMarketTransactor) ClaimWinnings(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ElectionPredictionMarket.contract.Transact(opts, "claimWinnings")
}

// ClaimWinnings is a paid mutator transaction binding the contract method 0xb401faf1.
//
// Solidity: function claimWinnings() returns()
func (_ElectionPredictionMarket *ElectionPredictionMarketSession) ClaimWinnings() (*types.Transaction, error) {
	return _ElectionPredictionMarket.Contract.ClaimWinnings(&_ElectionPredictionMarket.TransactOpts)
}

// ClaimWinnings is a paid mutator transaction binding the contract method 0xb401faf1.
//
// Solidity: function claimWinnings() returns()
func (_ElectionPredictionMarket *ElectionPredictionMarketTransactorSession) ClaimWinnings() (*types.Transaction, error) {
	return _ElectionPredictionMarket.Contract.ClaimWinnings(&_ElectionPredictionMarket.TransactOpts)
}

// PlaceBet is a paid mutator transaction binding the contract method 0xdcb3b30e.
//
// Solidity: function placeBet(uint8 _party, uint256 _amount) returns()
func (_ElectionPredictionMarket *ElectionPredictionMarketTransactor) PlaceBet(opts *bind.TransactOpts, _party uint8, _amount *big.Int) (*types.Transaction, error) {
	return _ElectionPredictionMarket.contract.Transact(opts, "placeBet", _party, _amount)
}

// PlaceBet is a paid mutator transaction binding the contract method 0xdcb3b30e.
//
// Solidity: function placeBet(uint8 _party, uint256 _amount) returns()
func (_ElectionPredictionMarket *ElectionPredictionMarketSession) PlaceBet(_party uint8, _amount *big.Int) (*types.Transaction, error) {
	return _ElectionPredictionMarket.Contract.PlaceBet(&_ElectionPredictionMarket.TransactOpts, _party, _amount)
}

// PlaceBet is a paid mutator transaction binding the contract method 0xdcb3b30e.
//
// Solidity: function placeBet(uint8 _party, uint256 _amount) returns()
func (_ElectionPredictionMarket *ElectionPredictionMarketTransactorSession) PlaceBet(_party uint8, _amount *big.Int) (*types.Transaction, error) {
	return _ElectionPredictionMarket.Contract.PlaceBet(&_ElectionPredictionMarket.TransactOpts, _party, _amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ElectionPredictionMarket *ElectionPredictionMarketTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ElectionPredictionMarket.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ElectionPredictionMarket *ElectionPredictionMarketSession) RenounceOwnership() (*types.Transaction, error) {
	return _ElectionPredictionMarket.Contract.RenounceOwnership(&_ElectionPredictionMarket.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ElectionPredictionMarket *ElectionPredictionMarketTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ElectionPredictionMarket.Contract.RenounceOwnership(&_ElectionPredictionMarket.TransactOpts)
}

// ResolveMarket is a paid mutator transaction binding the contract method 0x2b9b55b3.
//
// Solidity: function resolveMarket() returns()
func (_ElectionPredictionMarket *ElectionPredictionMarketTransactor) ResolveMarket(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ElectionPredictionMarket.contract.Transact(opts, "resolveMarket")
}

// ResolveMarket is a paid mutator transaction binding the contract method 0x2b9b55b3.
//
// Solidity: function resolveMarket() returns()
func (_ElectionPredictionMarket *ElectionPredictionMarketSession) ResolveMarket() (*types.Transaction, error) {
	return _ElectionPredictionMarket.Contract.ResolveMarket(&_ElectionPredictionMarket.TransactOpts)
}

// ResolveMarket is a paid mutator transaction binding the contract method 0x2b9b55b3.
//
// Solidity: function resolveMarket() returns()
func (_ElectionPredictionMarket *ElectionPredictionMarketTransactorSession) ResolveMarket() (*types.Transaction, error) {
	return _ElectionPredictionMarket.Contract.ResolveMarket(&_ElectionPredictionMarket.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ElectionPredictionMarket *ElectionPredictionMarketTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ElectionPredictionMarket.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ElectionPredictionMarket *ElectionPredictionMarketSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ElectionPredictionMarket.Contract.TransferOwnership(&_ElectionPredictionMarket.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ElectionPredictionMarket *ElectionPredictionMarketTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ElectionPredictionMarket.Contract.TransferOwnership(&_ElectionPredictionMarket.TransactOpts, newOwner)
}

// UpdateOdds is a paid mutator transaction binding the contract method 0xeb50f6bc.
//
// Solidity: function updateOdds(uint256 _democratOdds) returns()
func (_ElectionPredictionMarket *ElectionPredictionMarketTransactor) UpdateOdds(opts *bind.TransactOpts, _democratOdds *big.Int) (*types.Transaction, error) {
	return _ElectionPredictionMarket.contract.Transact(opts, "updateOdds", _democratOdds)
}

// UpdateOdds is a paid mutator transaction binding the contract method 0xeb50f6bc.
//
// Solidity: function updateOdds(uint256 _democratOdds) returns()
func (_ElectionPredictionMarket *ElectionPredictionMarketSession) UpdateOdds(_democratOdds *big.Int) (*types.Transaction, error) {
	return _ElectionPredictionMarket.Contract.UpdateOdds(&_ElectionPredictionMarket.TransactOpts, _democratOdds)
}

// UpdateOdds is a paid mutator transaction binding the contract method 0xeb50f6bc.
//
// Solidity: function updateOdds(uint256 _democratOdds) returns()
func (_ElectionPredictionMarket *ElectionPredictionMarketTransactorSession) UpdateOdds(_democratOdds *big.Int) (*types.Transaction, error) {
	return _ElectionPredictionMarket.Contract.UpdateOdds(&_ElectionPredictionMarket.TransactOpts, _democratOdds)
}

// ElectionPredictionMarketBetPlacedIterator is returned from FilterBetPlaced and is used to iterate over the raw logs and unpacked data for BetPlaced events raised by the ElectionPredictionMarket contract.
type ElectionPredictionMarketBetPlacedIterator struct {
	Event *ElectionPredictionMarketBetPlaced // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ElectionPredictionMarketBetPlacedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ElectionPredictionMarketBetPlaced)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ElectionPredictionMarketBetPlaced)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ElectionPredictionMarketBetPlacedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ElectionPredictionMarketBetPlacedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ElectionPredictionMarketBetPlaced represents a BetPlaced event raised by the ElectionPredictionMarket contract.
type ElectionPredictionMarketBetPlaced struct {
	Bettor common.Address
	Party  uint8
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBetPlaced is a free log retrieval operation binding the contract event 0x9f6b4cced64244a64465098f682029975b1614400b96bc6bd12c668083719166.
//
// Solidity: event BetPlaced(address indexed bettor, uint8 party, uint256 amount)
func (_ElectionPredictionMarket *ElectionPredictionMarketFilterer) FilterBetPlaced(opts *bind.FilterOpts, bettor []common.Address) (*ElectionPredictionMarketBetPlacedIterator, error) {

	var bettorRule []interface{}
	for _, bettorItem := range bettor {
		bettorRule = append(bettorRule, bettorItem)
	}

	logs, sub, err := _ElectionPredictionMarket.contract.FilterLogs(opts, "BetPlaced", bettorRule)
	if err != nil {
		return nil, err
	}
	return &ElectionPredictionMarketBetPlacedIterator{contract: _ElectionPredictionMarket.contract, event: "BetPlaced", logs: logs, sub: sub}, nil
}

// WatchBetPlaced is a free log subscription operation binding the contract event 0x9f6b4cced64244a64465098f682029975b1614400b96bc6bd12c668083719166.
//
// Solidity: event BetPlaced(address indexed bettor, uint8 party, uint256 amount)
func (_ElectionPredictionMarket *ElectionPredictionMarketFilterer) WatchBetPlaced(opts *bind.WatchOpts, sink chan<- *ElectionPredictionMarketBetPlaced, bettor []common.Address) (event.Subscription, error) {

	var bettorRule []interface{}
	for _, bettorItem := range bettor {
		bettorRule = append(bettorRule, bettorItem)
	}

	logs, sub, err := _ElectionPredictionMarket.contract.WatchLogs(opts, "BetPlaced", bettorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ElectionPredictionMarketBetPlaced)
				if err := _ElectionPredictionMarket.contract.UnpackLog(event, "BetPlaced", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBetPlaced is a log parse operation binding the contract event 0x9f6b4cced64244a64465098f682029975b1614400b96bc6bd12c668083719166.
//
// Solidity: event BetPlaced(address indexed bettor, uint8 party, uint256 amount)
func (_ElectionPredictionMarket *ElectionPredictionMarketFilterer) ParseBetPlaced(log types.Log) (*ElectionPredictionMarketBetPlaced, error) {
	event := new(ElectionPredictionMarketBetPlaced)
	if err := _ElectionPredictionMarket.contract.UnpackLog(event, "BetPlaced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ElectionPredictionMarketMarketResolvedIterator is returned from FilterMarketResolved and is used to iterate over the raw logs and unpacked data for MarketResolved events raised by the ElectionPredictionMarket contract.
type ElectionPredictionMarketMarketResolvedIterator struct {
	Event *ElectionPredictionMarketMarketResolved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ElectionPredictionMarketMarketResolvedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ElectionPredictionMarketMarketResolved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ElectionPredictionMarketMarketResolved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ElectionPredictionMarketMarketResolvedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ElectionPredictionMarketMarketResolvedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ElectionPredictionMarketMarketResolved represents a MarketResolved event raised by the ElectionPredictionMarket contract.
type ElectionPredictionMarketMarketResolved struct {
	Winner uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMarketResolved is a free log retrieval operation binding the contract event 0x94e5f71d5001613eb5e4ad876acb1da127582270c692c34fbcf810be21728b1d.
//
// Solidity: event MarketResolved(uint8 winner)
func (_ElectionPredictionMarket *ElectionPredictionMarketFilterer) FilterMarketResolved(opts *bind.FilterOpts) (*ElectionPredictionMarketMarketResolvedIterator, error) {

	logs, sub, err := _ElectionPredictionMarket.contract.FilterLogs(opts, "MarketResolved")
	if err != nil {
		return nil, err
	}
	return &ElectionPredictionMarketMarketResolvedIterator{contract: _ElectionPredictionMarket.contract, event: "MarketResolved", logs: logs, sub: sub}, nil
}

// WatchMarketResolved is a free log subscription operation binding the contract event 0x94e5f71d5001613eb5e4ad876acb1da127582270c692c34fbcf810be21728b1d.
//
// Solidity: event MarketResolved(uint8 winner)
func (_ElectionPredictionMarket *ElectionPredictionMarketFilterer) WatchMarketResolved(opts *bind.WatchOpts, sink chan<- *ElectionPredictionMarketMarketResolved) (event.Subscription, error) {

	logs, sub, err := _ElectionPredictionMarket.contract.WatchLogs(opts, "MarketResolved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ElectionPredictionMarketMarketResolved)
				if err := _ElectionPredictionMarket.contract.UnpackLog(event, "MarketResolved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMarketResolved is a log parse operation binding the contract event 0x94e5f71d5001613eb5e4ad876acb1da127582270c692c34fbcf810be21728b1d.
//
// Solidity: event MarketResolved(uint8 winner)
func (_ElectionPredictionMarket *ElectionPredictionMarketFilterer) ParseMarketResolved(log types.Log) (*ElectionPredictionMarketMarketResolved, error) {
	event := new(ElectionPredictionMarketMarketResolved)
	if err := _ElectionPredictionMarket.contract.UnpackLog(event, "MarketResolved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ElectionPredictionMarketOddsUpdatedIterator is returned from FilterOddsUpdated and is used to iterate over the raw logs and unpacked data for OddsUpdated events raised by the ElectionPredictionMarket contract.
type ElectionPredictionMarketOddsUpdatedIterator struct {
	Event *ElectionPredictionMarketOddsUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ElectionPredictionMarketOddsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ElectionPredictionMarketOddsUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ElectionPredictionMarketOddsUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ElectionPredictionMarketOddsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ElectionPredictionMarketOddsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ElectionPredictionMarketOddsUpdated represents a OddsUpdated event raised by the ElectionPredictionMarket contract.
type ElectionPredictionMarketOddsUpdated struct {
	DemocratOdds   *big.Int
	RepublicanOdds *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOddsUpdated is a free log retrieval operation binding the contract event 0x401f5799446884f85b1967999c05db7393e8ce1fd302dc57ad43d6c1fafe9e29.
//
// Solidity: event OddsUpdated(uint256 democratOdds, uint256 republicanOdds)
func (_ElectionPredictionMarket *ElectionPredictionMarketFilterer) FilterOddsUpdated(opts *bind.FilterOpts) (*ElectionPredictionMarketOddsUpdatedIterator, error) {

	logs, sub, err := _ElectionPredictionMarket.contract.FilterLogs(opts, "OddsUpdated")
	if err != nil {
		return nil, err
	}
	return &ElectionPredictionMarketOddsUpdatedIterator{contract: _ElectionPredictionMarket.contract, event: "OddsUpdated", logs: logs, sub: sub}, nil
}

// WatchOddsUpdated is a free log subscription operation binding the contract event 0x401f5799446884f85b1967999c05db7393e8ce1fd302dc57ad43d6c1fafe9e29.
//
// Solidity: event OddsUpdated(uint256 democratOdds, uint256 republicanOdds)
func (_ElectionPredictionMarket *ElectionPredictionMarketFilterer) WatchOddsUpdated(opts *bind.WatchOpts, sink chan<- *ElectionPredictionMarketOddsUpdated) (event.Subscription, error) {

	logs, sub, err := _ElectionPredictionMarket.contract.WatchLogs(opts, "OddsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ElectionPredictionMarketOddsUpdated)
				if err := _ElectionPredictionMarket.contract.UnpackLog(event, "OddsUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOddsUpdated is a log parse operation binding the contract event 0x401f5799446884f85b1967999c05db7393e8ce1fd302dc57ad43d6c1fafe9e29.
//
// Solidity: event OddsUpdated(uint256 democratOdds, uint256 republicanOdds)
func (_ElectionPredictionMarket *ElectionPredictionMarketFilterer) ParseOddsUpdated(log types.Log) (*ElectionPredictionMarketOddsUpdated, error) {
	event := new(ElectionPredictionMarketOddsUpdated)
	if err := _ElectionPredictionMarket.contract.UnpackLog(event, "OddsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ElectionPredictionMarketOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ElectionPredictionMarket contract.
type ElectionPredictionMarketOwnershipTransferredIterator struct {
	Event *ElectionPredictionMarketOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ElectionPredictionMarketOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ElectionPredictionMarketOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ElectionPredictionMarketOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ElectionPredictionMarketOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ElectionPredictionMarketOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ElectionPredictionMarketOwnershipTransferred represents a OwnershipTransferred event raised by the ElectionPredictionMarket contract.
type ElectionPredictionMarketOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ElectionPredictionMarket *ElectionPredictionMarketFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ElectionPredictionMarketOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ElectionPredictionMarket.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ElectionPredictionMarketOwnershipTransferredIterator{contract: _ElectionPredictionMarket.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ElectionPredictionMarket *ElectionPredictionMarketFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ElectionPredictionMarketOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ElectionPredictionMarket.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ElectionPredictionMarketOwnershipTransferred)
				if err := _ElectionPredictionMarket.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ElectionPredictionMarket *ElectionPredictionMarketFilterer) ParseOwnershipTransferred(log types.Log) (*ElectionPredictionMarketOwnershipTransferred, error) {
	event := new(ElectionPredictionMarketOwnershipTransferred)
	if err := _ElectionPredictionMarket.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ElectionPredictionMarketWinningsClaimedIterator is returned from FilterWinningsClaimed and is used to iterate over the raw logs and unpacked data for WinningsClaimed events raised by the ElectionPredictionMarket contract.
type ElectionPredictionMarketWinningsClaimedIterator struct {
	Event *ElectionPredictionMarketWinningsClaimed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ElectionPredictionMarketWinningsClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ElectionPredictionMarketWinningsClaimed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ElectionPredictionMarketWinningsClaimed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ElectionPredictionMarketWinningsClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ElectionPredictionMarketWinningsClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ElectionPredictionMarketWinningsClaimed represents a WinningsClaimed event raised by the ElectionPredictionMarket contract.
type ElectionPredictionMarketWinningsClaimed struct {
	Bettor common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWinningsClaimed is a free log retrieval operation binding the contract event 0x1a31e733a172afcf46074b3106c17f0c298e226442682a03c1e99ce256139ec2.
//
// Solidity: event WinningsClaimed(address indexed bettor, uint256 amount)
func (_ElectionPredictionMarket *ElectionPredictionMarketFilterer) FilterWinningsClaimed(opts *bind.FilterOpts, bettor []common.Address) (*ElectionPredictionMarketWinningsClaimedIterator, error) {

	var bettorRule []interface{}
	for _, bettorItem := range bettor {
		bettorRule = append(bettorRule, bettorItem)
	}

	logs, sub, err := _ElectionPredictionMarket.contract.FilterLogs(opts, "WinningsClaimed", bettorRule)
	if err != nil {
		return nil, err
	}
	return &ElectionPredictionMarketWinningsClaimedIterator{contract: _ElectionPredictionMarket.contract, event: "WinningsClaimed", logs: logs, sub: sub}, nil
}

// WatchWinningsClaimed is a free log subscription operation binding the contract event 0x1a31e733a172afcf46074b3106c17f0c298e226442682a03c1e99ce256139ec2.
//
// Solidity: event WinningsClaimed(address indexed bettor, uint256 amount)
func (_ElectionPredictionMarket *ElectionPredictionMarketFilterer) WatchWinningsClaimed(opts *bind.WatchOpts, sink chan<- *ElectionPredictionMarketWinningsClaimed, bettor []common.Address) (event.Subscription, error) {

	var bettorRule []interface{}
	for _, bettorItem := range bettor {
		bettorRule = append(bettorRule, bettorItem)
	}

	logs, sub, err := _ElectionPredictionMarket.contract.WatchLogs(opts, "WinningsClaimed", bettorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ElectionPredictionMarketWinningsClaimed)
				if err := _ElectionPredictionMarket.contract.UnpackLog(event, "WinningsClaimed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWinningsClaimed is a log parse operation binding the contract event 0x1a31e733a172afcf46074b3106c17f0c298e226442682a03c1e99ce256139ec2.
//
// Solidity: event WinningsClaimed(address indexed bettor, uint256 amount)
func (_ElectionPredictionMarket *ElectionPredictionMarketFilterer) ParseWinningsClaimed(log types.Log) (*ElectionPredictionMarketWinningsClaimed, error) {
	event := new(ElectionPredictionMarketWinningsClaimed)
	if err := _ElectionPredictionMarket.contract.UnpackLog(event, "WinningsClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
