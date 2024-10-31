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

// PredictionMarketOracleData is an auto generated low-level Go binding around an user-defined struct.
type PredictionMarketOracleData struct {
	Odds      *big.Int
	Timestamp *big.Int
}

// PredictionMarketMetaData contains all meta data concerning the PredictionMarket contract.
var PredictionMarketMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_betToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"MAX_HISTORICAL_DATA_POINTS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MIN_UPDATE_INTERVAL\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"betToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createMarket\",\"inputs\":[{\"name\":\"_description\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_resolutionTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_minimumUpdateInterval\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getHistoricalOdds\",\"inputs\":[{\"name\":\"_marketId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"odds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getHistoricalOddsCount\",\"inputs\":[{\"name\":\"_marketId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLatestOdds\",\"inputs\":[{\"name\":\"_marketId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"odds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMarketDetails\",\"inputs\":[{\"name\":\"_marketId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"totalPoolSize\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"resolutionTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"resolved\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"cancelled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"outcomePools\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"currentOdds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lastUpdateBlock\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"marketCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"markets\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"totalPoolSize\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"resolutionTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"resolved\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"winningOutcome\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"cancelled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"currentOdds\",\"type\":\"tuple\",\"internalType\":\"structPredictionMarket.OracleData\",\"components\":[{\"name\":\"odds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"lastUpdateBlock\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minimumUpdateInterval\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"oracle\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"placeBet\",\"inputs\":[{\"name\":\"_marketId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_outcome\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOracle\",\"inputs\":[{\"name\":\"_oracle\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateOracleData\",\"inputs\":[{\"name\":\"_marketId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_odds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"userBets\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"outcome\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"oddsAtBet\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"claimed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"BetClaimed\",\"inputs\":[{\"name\":\"marketId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"bettor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BetPlaced\",\"inputs\":[{\"name\":\"marketId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"bettor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"outcome\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"oddsAtBet\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MarketCancelled\",\"inputs\":[{\"name\":\"marketId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MarketCreated\",\"inputs\":[{\"name\":\"marketId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"description\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"resolutionTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MarketResolved\",\"inputs\":[{\"name\":\"marketId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"winningOutcome\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OracleAddressUpdated\",\"inputs\":[{\"name\":\"newOracle\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OracleUpdated\",\"inputs\":[{\"name\":\"marketId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"odds\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]}]",
	Bin: "0x60a060405234801561001057600080fd5b5060405161166738038061166783398101604081905261002f91610151565b6001600055806001600160a01b03811661006457604051631e4fbdf760e01b8152600060048201526024015b60405180910390fd5b61006d816100e3565b506001805460ff60a01b191690556001600160a01b0382166100d15760405162461bcd60e51b815260206004820152601560248201527f496e76616c696420746f6b656e20616464726573730000000000000000000000604482015260640161005b565b506001600160a01b0316608052610184565b600180546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b80516001600160a01b038116811461014c57600080fd5b919050565b6000806040838503121561016457600080fd5b61016d83610135565b915061017b60208401610135565b90509250929050565b6080516114c16101a6600039600081816102010152610a8001526114c16000f3fe608060405234801561001057600080fd5b50600436106101215760003560e01c8063b06c1ba3116100ad578063da866c4811610071578063da866c48146102ef578063e7b7101214610302578063ec97908214610315578063ecfa0ddf1461031e578063f2fde38b1461032757600080fd5b8063b06c1ba314610272578063b1283e7714610299578063bd3be27c146102c1578063cfa51b16146102c9578063d92683a9146102dc57600080fd5b8063715018a6116100f4578063715018a6146101f257806378691f16146101fc5780637adbf9731461023b5780637dc0d1d01461024e5780638da5cb5b1461026157600080fd5b80633809ddda146101265780635c975abb1461016b5780636000cdee1461018857806361c4cc15146101b9575b600080fd5b610151610134366004610f88565b600090815260026020526040902060078101546008909101549091565b604080519283526020830191909152015b60405180910390f35b600154600160a01b900460ff166040519015158152602001610162565b6101ab610196366004610f88565b60009081526002602052604090206006015490565b604051908152602001610162565b6101cc6101c7366004610fbd565b61033a565b6040805160ff909516855260208501939093529183015215156060820152608001610162565b6101fa610395565b005b6102237f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b039091168152602001610162565b6101fa610249366004610ff3565b6103a9565b600554610223906001600160a01b031681565b6001546001600160a01b0316610223565b610285610280366004610f88565b610459565b60405161016298979695949392919061105b565b6102ac6102a7366004610f88565b6105d5565b604051610162999897969594939291906110e1565b6101ab603c81565b6101fa6102d7366004611160565b6106c5565b6101516102ea366004611227565b610858565b6101fa6102fd366004611249565b610904565b6101fa610310366004611276565b610c5c565b6101ab60035481565b6101ab6103e881565b6101fa610335366004610ff3565b610e76565b6004602052826000526040600020602052816000526040600020818154811061036257600080fd5b6000918252602090912060049091020180546001820154600283015460039093015460ff92831696509094509192501684565b61039d610eb4565b6103a76000610ee1565b565b6103b1610eb4565b6001600160a01b0381166104055760405162461bcd60e51b8152602060048201526016602482015275496e76616c6964206f7261636c65206164647265737360501b60448201526064015b60405180910390fd5b600580546001600160a01b0319166001600160a01b0383169081179091556040519081527f107a9fafffb7ac890f780879e423760c9ffea8dcee8045681f40f542aede2cb89060200160405180910390a150565b600081815260026020818152604080842081518481526060808201845295948594859485948994869485949293919290919083019080368337505060008080526004840160205260408120548351939750928792506104ba576104ba6112a2565b602002602001018181525050806004016000600160ff16815260200190815260200160002054846001815181106104f3576104f36112a2565b602090810291909101015260018101546002820154600383015460058401546007850154600986015486548796959460ff9081169416928b9290918890610539906112b8565b80601f0160208091040260200160405190810160405280929190818152602001828054610565906112b8565b80156105b25780601f10610587576101008083540402835291602001916105b2565b820191906000526020600020905b81548152906001019060200180831161059557829003601f168201915b505050505097509850985098509850985098509850985050919395975091939597565b6002602052600090815260409020805481906105f0906112b8565b80601f016020809104026020016040519081016040528092919081815260200182805461061c906112b8565b80156106695780601f1061063e57610100808354040283529160200191610669565b820191906000526020600020905b81548152906001019060200180831161064c57829003601f168201915b50505060018401546002850154600386015460058701546040805180820190915260078901548152600889015460208201526009890154600a909901549798949793965060ff8084169650610100909304831694919092169289565b6106cd610eb4565b6106da4262015180611308565b82116107285760405162461bcd60e51b815260206004820152601860248201527f5265736f6c7574696f6e2074696d6520746f6f20736f6f6e000000000000000060448201526064016103fc565b610736426301e13380611308565b82106107845760405162461bcd60e51b815260206004820152601760248201527f5265736f6c7574696f6e2074696d6520746f6f2066617200000000000000000060448201526064016103fc565b603c8110156107d55760405162461bcd60e51b815260206004820152601960248201527f55706461746520696e74657276616c20746f6f2073686f72740000000000000060448201526064016103fc565b60038054600091826107e683611321565b909155506000818152600260205260409020909150806108068682611388565b5060028101849055600a810183905560405182907f2d1e9ad45dbe8e898e0374deebf2a661b2ae7c1855d2a29507d584c35d287c75906108499088908890611447565b60405180910390a25050505050565b6000828152600260205260408120600681015482919084106108b25760405162461bcd60e51b8152602060048201526013602482015272496e646578206f7574206f6620626f756e647360681b60448201526064016103fc565b60008160060185815481106108c9576108c96112a2565b600091825260209182902060408051808201909152600290920201805480835260019091015491909201819052909450925050509250929050565b61090c610f33565b610914610f5d565b600081116109645760405162461bcd60e51b815260206004820152601b60248201527f42657420616d6f756e74206d75737420626520706f736974697665000000000060448201526064016103fc565b60018260ff1611156109aa5760405162461bcd60e51b815260206004820152600f60248201526e496e76616c6964206f7574636f6d6560881b60448201526064016103fc565b6000838152600260205260409020600381015460ff161580156109d25750600581015460ff16155b610a185760405162461bcd60e51b81526020600482015260176024820152764d61726b6574206e6f206c6f6e6765722061637469766560481b60448201526064016103fc565b80600201544210610a5e5760405162461bcd60e51b815260206004820152601060248201526f13585c9ad95d081a185cc8195b99195960821b60448201526064016103fc565b6040516323b872dd60e01b8152336004820152306024820152604481018390527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906323b872dd906064016020604051808303816000875af1158015610ad1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610af59190611469565b610b395760405162461bcd60e51b8152602060048201526015602482015274151bdad95b881d1c985b9cd9995c8819985a5b1959605a1b60448201526064016103fc565b81816001016000828254610b4d9190611308565b909155505060ff8316600090815260048201602052604081208054849290610b76908490611308565b9091555050600084815260046020818152604080842033808652908352818520825160808101845260ff8a81168083528287018b815260078b01805485890190815260608087018d8152885460018082018b55998f529d8c902097519d909c0290960180549c90951660ff199c8d161785559151958401959095555160028301559651600390910180549115159190981617909655548251948552928401879052908301919091529186917fbfcd8f4e47cc529f86cdc7eab8fb80e015a298e021fcaee72df1f76dd57cf487910160405180910390a350610c576001600055565b505050565b6005546001600160a01b03163314610cb65760405162461bcd60e51b815260206004820152601860248201527f43616c6c6572206973206e6f7420746865206f7261636c65000000000000000060448201526064016103fc565b610cbe610f5d565b6000838152600260205260409020600381015460ff16158015610ce65750600581015460ff16155b610d2c5760405162461bcd60e51b81526020600482015260176024820152764d61726b6574206e6f206c6f6e6765722061637469766560481b60448201526064016103fc565b80600a01548160090154610d409190611308565b421015610d845760405162461bcd60e51b8152602060048201526012602482015271546f6f20736f6f6e20746f2075706461746560701b60448201526064016103fc565b42821115610dd45760405162461bcd60e51b815260206004820152601c60248201527f4675747572652074696d657374616d70206e6f7420616c6c6f7765640000000060448201526064016103fc565b60068101546103e81115610e10576006810180546001818101835560009283526020909220600784015460029092020190815560088301549101555b60408051808201825284815260209081018490526007830185905560088301849055426009840155815185815290810184905285917f8d6131ead05d17bb37913bdff362964b1be7fc0762dfbf84399d91316c35f2ac910160405180910390a250505050565b610e7e610eb4565b6001600160a01b038116610ea857604051631e4fbdf760e01b8152600060048201526024016103fc565b610eb181610ee1565b50565b6001546001600160a01b031633146103a75760405163118cdaa760e01b81523360048201526024016103fc565b600180546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600260005403610f5657604051633ee5aeb560e01b815260040160405180910390fd5b6002600055565b600154600160a01b900460ff16156103a75760405163d93c066560e01b815260040160405180910390fd5b600060208284031215610f9a57600080fd5b5035919050565b80356001600160a01b0381168114610fb857600080fd5b919050565b600080600060608486031215610fd257600080fd5b83359250610fe260208501610fa1565b929592945050506040919091013590565b60006020828403121561100557600080fd5b61100e82610fa1565b9392505050565b6000815180845260005b8181101561103b5760208185018101518683018201520161101f565b506000602082860101526020601f19601f83011685010191505092915050565b6101008152600061107061010083018b611015565b8960208401528860408401528715156060840152861515608084015282810360a084015280865180835260208301915060208801925060005b818110156110c75783518352602093840193909201916001016110a9565b505060c0840195909552505060e001529695505050505050565b610140815260006110f661014083018c611015565b9050896020830152886040830152871515606083015260ff8716608083015285151560a0830152845160c0830152602085015160e083015283610100830152826101208301529a9950505050505050505050565b634e487b7160e01b600052604160045260246000fd5b60008060006060848603121561117557600080fd5b833567ffffffffffffffff81111561118c57600080fd5b8401601f8101861361119d57600080fd5b803567ffffffffffffffff8111156111b7576111b761114a565b604051601f8201601f19908116603f0116810167ffffffffffffffff811182821017156111e6576111e661114a565b6040528181528282016020018810156111fe57600080fd5b816020840160208301376000602092820183015297908601359650604090950135949350505050565b6000806040838503121561123a57600080fd5b50508035926020909101359150565b60008060006060848603121561125e57600080fd5b83359250602084013560ff81168114610fe257600080fd5b60008060006060848603121561128b57600080fd5b505081359360208301359350604090920135919050565b634e487b7160e01b600052603260045260246000fd5b600181811c908216806112cc57607f821691505b6020821081036112ec57634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052601160045260246000fd5b8082018082111561131b5761131b6112f2565b92915050565b600060018201611333576113336112f2565b5060010190565b601f821115610c5757806000526020600020601f840160051c810160208510156113615750805b601f840160051c820191505b81811015611381576000815560010161136d565b5050505050565b815167ffffffffffffffff8111156113a2576113a261114a565b6113b6816113b084546112b8565b8461133a565b6020601f8211600181146113ea57600083156113d25750848201515b600019600385901b1c1916600184901b178455611381565b600084815260208120601f198516915b8281101561141a57878501518255602094850194600190920191016113fa565b50848210156114385786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b60408152600061145a6040830185611015565b90508260208301529392505050565b60006020828403121561147b57600080fd5b8151801515811461100e57600080fdfea26469706673582212207d696cb51eeda8cebb249cbef8498d25ef8a12c257ac1651270e1b999b0bacd464736f6c634300081a0033",
}

// PredictionMarketABI is the input ABI used to generate the binding from.
// Deprecated: Use PredictionMarketMetaData.ABI instead.
var PredictionMarketABI = PredictionMarketMetaData.ABI

// PredictionMarketBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PredictionMarketMetaData.Bin instead.
var PredictionMarketBin = PredictionMarketMetaData.Bin

// DeployPredictionMarket deploys a new Ethereum contract, binding an instance of PredictionMarket to it.
func DeployPredictionMarket(auth *bind.TransactOpts, backend bind.ContractBackend, _betToken common.Address, initialOwner common.Address) (common.Address, *types.Transaction, *PredictionMarket, error) {
	parsed, err := PredictionMarketMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PredictionMarketBin), backend, _betToken, initialOwner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PredictionMarket{PredictionMarketCaller: PredictionMarketCaller{contract: contract}, PredictionMarketTransactor: PredictionMarketTransactor{contract: contract}, PredictionMarketFilterer: PredictionMarketFilterer{contract: contract}}, nil
}

// PredictionMarket is an auto generated Go binding around an Ethereum contract.
type PredictionMarket struct {
	PredictionMarketCaller     // Read-only binding to the contract
	PredictionMarketTransactor // Write-only binding to the contract
	PredictionMarketFilterer   // Log filterer for contract events
}

// PredictionMarketCaller is an auto generated read-only Go binding around an Ethereum contract.
type PredictionMarketCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PredictionMarketTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PredictionMarketTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PredictionMarketFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PredictionMarketFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PredictionMarketSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PredictionMarketSession struct {
	Contract     *PredictionMarket // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PredictionMarketCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PredictionMarketCallerSession struct {
	Contract *PredictionMarketCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// PredictionMarketTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PredictionMarketTransactorSession struct {
	Contract     *PredictionMarketTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// PredictionMarketRaw is an auto generated low-level Go binding around an Ethereum contract.
type PredictionMarketRaw struct {
	Contract *PredictionMarket // Generic contract binding to access the raw methods on
}

// PredictionMarketCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PredictionMarketCallerRaw struct {
	Contract *PredictionMarketCaller // Generic read-only contract binding to access the raw methods on
}

// PredictionMarketTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PredictionMarketTransactorRaw struct {
	Contract *PredictionMarketTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPredictionMarket creates a new instance of PredictionMarket, bound to a specific deployed contract.
func NewPredictionMarket(address common.Address, backend bind.ContractBackend) (*PredictionMarket, error) {
	contract, err := bindPredictionMarket(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PredictionMarket{PredictionMarketCaller: PredictionMarketCaller{contract: contract}, PredictionMarketTransactor: PredictionMarketTransactor{contract: contract}, PredictionMarketFilterer: PredictionMarketFilterer{contract: contract}}, nil
}

// NewPredictionMarketCaller creates a new read-only instance of PredictionMarket, bound to a specific deployed contract.
func NewPredictionMarketCaller(address common.Address, caller bind.ContractCaller) (*PredictionMarketCaller, error) {
	contract, err := bindPredictionMarket(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PredictionMarketCaller{contract: contract}, nil
}

// NewPredictionMarketTransactor creates a new write-only instance of PredictionMarket, bound to a specific deployed contract.
func NewPredictionMarketTransactor(address common.Address, transactor bind.ContractTransactor) (*PredictionMarketTransactor, error) {
	contract, err := bindPredictionMarket(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PredictionMarketTransactor{contract: contract}, nil
}

// NewPredictionMarketFilterer creates a new log filterer instance of PredictionMarket, bound to a specific deployed contract.
func NewPredictionMarketFilterer(address common.Address, filterer bind.ContractFilterer) (*PredictionMarketFilterer, error) {
	contract, err := bindPredictionMarket(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PredictionMarketFilterer{contract: contract}, nil
}

// bindPredictionMarket binds a generic wrapper to an already deployed contract.
func bindPredictionMarket(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PredictionMarketMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PredictionMarket *PredictionMarketRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PredictionMarket.Contract.PredictionMarketCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PredictionMarket *PredictionMarketRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PredictionMarket.Contract.PredictionMarketTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PredictionMarket *PredictionMarketRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PredictionMarket.Contract.PredictionMarketTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PredictionMarket *PredictionMarketCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PredictionMarket.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PredictionMarket *PredictionMarketTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PredictionMarket.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PredictionMarket *PredictionMarketTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PredictionMarket.Contract.contract.Transact(opts, method, params...)
}

// MAXHISTORICALDATAPOINTS is a free data retrieval call binding the contract method 0xecfa0ddf.
//
// Solidity: function MAX_HISTORICAL_DATA_POINTS() view returns(uint256)
func (_PredictionMarket *PredictionMarketCaller) MAXHISTORICALDATAPOINTS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PredictionMarket.contract.Call(opts, &out, "MAX_HISTORICAL_DATA_POINTS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXHISTORICALDATAPOINTS is a free data retrieval call binding the contract method 0xecfa0ddf.
//
// Solidity: function MAX_HISTORICAL_DATA_POINTS() view returns(uint256)
func (_PredictionMarket *PredictionMarketSession) MAXHISTORICALDATAPOINTS() (*big.Int, error) {
	return _PredictionMarket.Contract.MAXHISTORICALDATAPOINTS(&_PredictionMarket.CallOpts)
}

// MAXHISTORICALDATAPOINTS is a free data retrieval call binding the contract method 0xecfa0ddf.
//
// Solidity: function MAX_HISTORICAL_DATA_POINTS() view returns(uint256)
func (_PredictionMarket *PredictionMarketCallerSession) MAXHISTORICALDATAPOINTS() (*big.Int, error) {
	return _PredictionMarket.Contract.MAXHISTORICALDATAPOINTS(&_PredictionMarket.CallOpts)
}

// MINUPDATEINTERVAL is a free data retrieval call binding the contract method 0xbd3be27c.
//
// Solidity: function MIN_UPDATE_INTERVAL() view returns(uint256)
func (_PredictionMarket *PredictionMarketCaller) MINUPDATEINTERVAL(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PredictionMarket.contract.Call(opts, &out, "MIN_UPDATE_INTERVAL")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINUPDATEINTERVAL is a free data retrieval call binding the contract method 0xbd3be27c.
//
// Solidity: function MIN_UPDATE_INTERVAL() view returns(uint256)
func (_PredictionMarket *PredictionMarketSession) MINUPDATEINTERVAL() (*big.Int, error) {
	return _PredictionMarket.Contract.MINUPDATEINTERVAL(&_PredictionMarket.CallOpts)
}

// MINUPDATEINTERVAL is a free data retrieval call binding the contract method 0xbd3be27c.
//
// Solidity: function MIN_UPDATE_INTERVAL() view returns(uint256)
func (_PredictionMarket *PredictionMarketCallerSession) MINUPDATEINTERVAL() (*big.Int, error) {
	return _PredictionMarket.Contract.MINUPDATEINTERVAL(&_PredictionMarket.CallOpts)
}

// BetToken is a free data retrieval call binding the contract method 0x78691f16.
//
// Solidity: function betToken() view returns(address)
func (_PredictionMarket *PredictionMarketCaller) BetToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PredictionMarket.contract.Call(opts, &out, "betToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BetToken is a free data retrieval call binding the contract method 0x78691f16.
//
// Solidity: function betToken() view returns(address)
func (_PredictionMarket *PredictionMarketSession) BetToken() (common.Address, error) {
	return _PredictionMarket.Contract.BetToken(&_PredictionMarket.CallOpts)
}

// BetToken is a free data retrieval call binding the contract method 0x78691f16.
//
// Solidity: function betToken() view returns(address)
func (_PredictionMarket *PredictionMarketCallerSession) BetToken() (common.Address, error) {
	return _PredictionMarket.Contract.BetToken(&_PredictionMarket.CallOpts)
}

// GetHistoricalOdds is a free data retrieval call binding the contract method 0xd92683a9.
//
// Solidity: function getHistoricalOdds(uint256 _marketId, uint256 _index) view returns(uint256 odds, uint256 timestamp)
func (_PredictionMarket *PredictionMarketCaller) GetHistoricalOdds(opts *bind.CallOpts, _marketId *big.Int, _index *big.Int) (struct {
	Odds      *big.Int
	Timestamp *big.Int
}, error) {
	var out []interface{}
	err := _PredictionMarket.contract.Call(opts, &out, "getHistoricalOdds", _marketId, _index)

	outstruct := new(struct {
		Odds      *big.Int
		Timestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Odds = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetHistoricalOdds is a free data retrieval call binding the contract method 0xd92683a9.
//
// Solidity: function getHistoricalOdds(uint256 _marketId, uint256 _index) view returns(uint256 odds, uint256 timestamp)
func (_PredictionMarket *PredictionMarketSession) GetHistoricalOdds(_marketId *big.Int, _index *big.Int) (struct {
	Odds      *big.Int
	Timestamp *big.Int
}, error) {
	return _PredictionMarket.Contract.GetHistoricalOdds(&_PredictionMarket.CallOpts, _marketId, _index)
}

// GetHistoricalOdds is a free data retrieval call binding the contract method 0xd92683a9.
//
// Solidity: function getHistoricalOdds(uint256 _marketId, uint256 _index) view returns(uint256 odds, uint256 timestamp)
func (_PredictionMarket *PredictionMarketCallerSession) GetHistoricalOdds(_marketId *big.Int, _index *big.Int) (struct {
	Odds      *big.Int
	Timestamp *big.Int
}, error) {
	return _PredictionMarket.Contract.GetHistoricalOdds(&_PredictionMarket.CallOpts, _marketId, _index)
}

// GetHistoricalOddsCount is a free data retrieval call binding the contract method 0x6000cdee.
//
// Solidity: function getHistoricalOddsCount(uint256 _marketId) view returns(uint256)
func (_PredictionMarket *PredictionMarketCaller) GetHistoricalOddsCount(opts *bind.CallOpts, _marketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PredictionMarket.contract.Call(opts, &out, "getHistoricalOddsCount", _marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetHistoricalOddsCount is a free data retrieval call binding the contract method 0x6000cdee.
//
// Solidity: function getHistoricalOddsCount(uint256 _marketId) view returns(uint256)
func (_PredictionMarket *PredictionMarketSession) GetHistoricalOddsCount(_marketId *big.Int) (*big.Int, error) {
	return _PredictionMarket.Contract.GetHistoricalOddsCount(&_PredictionMarket.CallOpts, _marketId)
}

// GetHistoricalOddsCount is a free data retrieval call binding the contract method 0x6000cdee.
//
// Solidity: function getHistoricalOddsCount(uint256 _marketId) view returns(uint256)
func (_PredictionMarket *PredictionMarketCallerSession) GetHistoricalOddsCount(_marketId *big.Int) (*big.Int, error) {
	return _PredictionMarket.Contract.GetHistoricalOddsCount(&_PredictionMarket.CallOpts, _marketId)
}

// GetLatestOdds is a free data retrieval call binding the contract method 0x3809ddda.
//
// Solidity: function getLatestOdds(uint256 _marketId) view returns(uint256 odds, uint256 timestamp)
func (_PredictionMarket *PredictionMarketCaller) GetLatestOdds(opts *bind.CallOpts, _marketId *big.Int) (struct {
	Odds      *big.Int
	Timestamp *big.Int
}, error) {
	var out []interface{}
	err := _PredictionMarket.contract.Call(opts, &out, "getLatestOdds", _marketId)

	outstruct := new(struct {
		Odds      *big.Int
		Timestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Odds = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetLatestOdds is a free data retrieval call binding the contract method 0x3809ddda.
//
// Solidity: function getLatestOdds(uint256 _marketId) view returns(uint256 odds, uint256 timestamp)
func (_PredictionMarket *PredictionMarketSession) GetLatestOdds(_marketId *big.Int) (struct {
	Odds      *big.Int
	Timestamp *big.Int
}, error) {
	return _PredictionMarket.Contract.GetLatestOdds(&_PredictionMarket.CallOpts, _marketId)
}

// GetLatestOdds is a free data retrieval call binding the contract method 0x3809ddda.
//
// Solidity: function getLatestOdds(uint256 _marketId) view returns(uint256 odds, uint256 timestamp)
func (_PredictionMarket *PredictionMarketCallerSession) GetLatestOdds(_marketId *big.Int) (struct {
	Odds      *big.Int
	Timestamp *big.Int
}, error) {
	return _PredictionMarket.Contract.GetLatestOdds(&_PredictionMarket.CallOpts, _marketId)
}

// GetMarketDetails is a free data retrieval call binding the contract method 0xb06c1ba3.
//
// Solidity: function getMarketDetails(uint256 _marketId) view returns(string description, uint256 totalPoolSize, uint256 resolutionTime, bool resolved, bool cancelled, uint256[] outcomePools, uint256 currentOdds, uint256 lastUpdateBlock)
func (_PredictionMarket *PredictionMarketCaller) GetMarketDetails(opts *bind.CallOpts, _marketId *big.Int) (struct {
	Description     string
	TotalPoolSize   *big.Int
	ResolutionTime  *big.Int
	Resolved        bool
	Cancelled       bool
	OutcomePools    []*big.Int
	CurrentOdds     *big.Int
	LastUpdateBlock *big.Int
}, error) {
	var out []interface{}
	err := _PredictionMarket.contract.Call(opts, &out, "getMarketDetails", _marketId)

	outstruct := new(struct {
		Description     string
		TotalPoolSize   *big.Int
		ResolutionTime  *big.Int
		Resolved        bool
		Cancelled       bool
		OutcomePools    []*big.Int
		CurrentOdds     *big.Int
		LastUpdateBlock *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Description = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.TotalPoolSize = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ResolutionTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Resolved = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.Cancelled = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.OutcomePools = *abi.ConvertType(out[5], new([]*big.Int)).(*[]*big.Int)
	outstruct.CurrentOdds = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.LastUpdateBlock = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetMarketDetails is a free data retrieval call binding the contract method 0xb06c1ba3.
//
// Solidity: function getMarketDetails(uint256 _marketId) view returns(string description, uint256 totalPoolSize, uint256 resolutionTime, bool resolved, bool cancelled, uint256[] outcomePools, uint256 currentOdds, uint256 lastUpdateBlock)
func (_PredictionMarket *PredictionMarketSession) GetMarketDetails(_marketId *big.Int) (struct {
	Description     string
	TotalPoolSize   *big.Int
	ResolutionTime  *big.Int
	Resolved        bool
	Cancelled       bool
	OutcomePools    []*big.Int
	CurrentOdds     *big.Int
	LastUpdateBlock *big.Int
}, error) {
	return _PredictionMarket.Contract.GetMarketDetails(&_PredictionMarket.CallOpts, _marketId)
}

// GetMarketDetails is a free data retrieval call binding the contract method 0xb06c1ba3.
//
// Solidity: function getMarketDetails(uint256 _marketId) view returns(string description, uint256 totalPoolSize, uint256 resolutionTime, bool resolved, bool cancelled, uint256[] outcomePools, uint256 currentOdds, uint256 lastUpdateBlock)
func (_PredictionMarket *PredictionMarketCallerSession) GetMarketDetails(_marketId *big.Int) (struct {
	Description     string
	TotalPoolSize   *big.Int
	ResolutionTime  *big.Int
	Resolved        bool
	Cancelled       bool
	OutcomePools    []*big.Int
	CurrentOdds     *big.Int
	LastUpdateBlock *big.Int
}, error) {
	return _PredictionMarket.Contract.GetMarketDetails(&_PredictionMarket.CallOpts, _marketId)
}

// MarketCount is a free data retrieval call binding the contract method 0xec979082.
//
// Solidity: function marketCount() view returns(uint256)
func (_PredictionMarket *PredictionMarketCaller) MarketCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PredictionMarket.contract.Call(opts, &out, "marketCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MarketCount is a free data retrieval call binding the contract method 0xec979082.
//
// Solidity: function marketCount() view returns(uint256)
func (_PredictionMarket *PredictionMarketSession) MarketCount() (*big.Int, error) {
	return _PredictionMarket.Contract.MarketCount(&_PredictionMarket.CallOpts)
}

// MarketCount is a free data retrieval call binding the contract method 0xec979082.
//
// Solidity: function marketCount() view returns(uint256)
func (_PredictionMarket *PredictionMarketCallerSession) MarketCount() (*big.Int, error) {
	return _PredictionMarket.Contract.MarketCount(&_PredictionMarket.CallOpts)
}

// Markets is a free data retrieval call binding the contract method 0xb1283e77.
//
// Solidity: function markets(uint256 ) view returns(string description, uint256 totalPoolSize, uint256 resolutionTime, bool resolved, uint8 winningOutcome, bool cancelled, (uint256,uint256) currentOdds, uint256 lastUpdateBlock, uint256 minimumUpdateInterval)
func (_PredictionMarket *PredictionMarketCaller) Markets(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Description           string
	TotalPoolSize         *big.Int
	ResolutionTime        *big.Int
	Resolved              bool
	WinningOutcome        uint8
	Cancelled             bool
	CurrentOdds           PredictionMarketOracleData
	LastUpdateBlock       *big.Int
	MinimumUpdateInterval *big.Int
}, error) {
	var out []interface{}
	err := _PredictionMarket.contract.Call(opts, &out, "markets", arg0)

	outstruct := new(struct {
		Description           string
		TotalPoolSize         *big.Int
		ResolutionTime        *big.Int
		Resolved              bool
		WinningOutcome        uint8
		Cancelled             bool
		CurrentOdds           PredictionMarketOracleData
		LastUpdateBlock       *big.Int
		MinimumUpdateInterval *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Description = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.TotalPoolSize = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ResolutionTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Resolved = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.WinningOutcome = *abi.ConvertType(out[4], new(uint8)).(*uint8)
	outstruct.Cancelled = *abi.ConvertType(out[5], new(bool)).(*bool)
	outstruct.CurrentOdds = *abi.ConvertType(out[6], new(PredictionMarketOracleData)).(*PredictionMarketOracleData)
	outstruct.LastUpdateBlock = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.MinimumUpdateInterval = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Markets is a free data retrieval call binding the contract method 0xb1283e77.
//
// Solidity: function markets(uint256 ) view returns(string description, uint256 totalPoolSize, uint256 resolutionTime, bool resolved, uint8 winningOutcome, bool cancelled, (uint256,uint256) currentOdds, uint256 lastUpdateBlock, uint256 minimumUpdateInterval)
func (_PredictionMarket *PredictionMarketSession) Markets(arg0 *big.Int) (struct {
	Description           string
	TotalPoolSize         *big.Int
	ResolutionTime        *big.Int
	Resolved              bool
	WinningOutcome        uint8
	Cancelled             bool
	CurrentOdds           PredictionMarketOracleData
	LastUpdateBlock       *big.Int
	MinimumUpdateInterval *big.Int
}, error) {
	return _PredictionMarket.Contract.Markets(&_PredictionMarket.CallOpts, arg0)
}

// Markets is a free data retrieval call binding the contract method 0xb1283e77.
//
// Solidity: function markets(uint256 ) view returns(string description, uint256 totalPoolSize, uint256 resolutionTime, bool resolved, uint8 winningOutcome, bool cancelled, (uint256,uint256) currentOdds, uint256 lastUpdateBlock, uint256 minimumUpdateInterval)
func (_PredictionMarket *PredictionMarketCallerSession) Markets(arg0 *big.Int) (struct {
	Description           string
	TotalPoolSize         *big.Int
	ResolutionTime        *big.Int
	Resolved              bool
	WinningOutcome        uint8
	Cancelled             bool
	CurrentOdds           PredictionMarketOracleData
	LastUpdateBlock       *big.Int
	MinimumUpdateInterval *big.Int
}, error) {
	return _PredictionMarket.Contract.Markets(&_PredictionMarket.CallOpts, arg0)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_PredictionMarket *PredictionMarketCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PredictionMarket.contract.Call(opts, &out, "oracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_PredictionMarket *PredictionMarketSession) Oracle() (common.Address, error) {
	return _PredictionMarket.Contract.Oracle(&_PredictionMarket.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_PredictionMarket *PredictionMarketCallerSession) Oracle() (common.Address, error) {
	return _PredictionMarket.Contract.Oracle(&_PredictionMarket.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PredictionMarket *PredictionMarketCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PredictionMarket.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PredictionMarket *PredictionMarketSession) Owner() (common.Address, error) {
	return _PredictionMarket.Contract.Owner(&_PredictionMarket.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PredictionMarket *PredictionMarketCallerSession) Owner() (common.Address, error) {
	return _PredictionMarket.Contract.Owner(&_PredictionMarket.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PredictionMarket *PredictionMarketCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PredictionMarket.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PredictionMarket *PredictionMarketSession) Paused() (bool, error) {
	return _PredictionMarket.Contract.Paused(&_PredictionMarket.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PredictionMarket *PredictionMarketCallerSession) Paused() (bool, error) {
	return _PredictionMarket.Contract.Paused(&_PredictionMarket.CallOpts)
}

// UserBets is a free data retrieval call binding the contract method 0x61c4cc15.
//
// Solidity: function userBets(uint256 , address , uint256 ) view returns(uint8 outcome, uint256 amount, uint256 oddsAtBet, bool claimed)
func (_PredictionMarket *PredictionMarketCaller) UserBets(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address, arg2 *big.Int) (struct {
	Outcome   uint8
	Amount    *big.Int
	OddsAtBet *big.Int
	Claimed   bool
}, error) {
	var out []interface{}
	err := _PredictionMarket.contract.Call(opts, &out, "userBets", arg0, arg1, arg2)

	outstruct := new(struct {
		Outcome   uint8
		Amount    *big.Int
		OddsAtBet *big.Int
		Claimed   bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Outcome = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.OddsAtBet = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Claimed = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// UserBets is a free data retrieval call binding the contract method 0x61c4cc15.
//
// Solidity: function userBets(uint256 , address , uint256 ) view returns(uint8 outcome, uint256 amount, uint256 oddsAtBet, bool claimed)
func (_PredictionMarket *PredictionMarketSession) UserBets(arg0 *big.Int, arg1 common.Address, arg2 *big.Int) (struct {
	Outcome   uint8
	Amount    *big.Int
	OddsAtBet *big.Int
	Claimed   bool
}, error) {
	return _PredictionMarket.Contract.UserBets(&_PredictionMarket.CallOpts, arg0, arg1, arg2)
}

// UserBets is a free data retrieval call binding the contract method 0x61c4cc15.
//
// Solidity: function userBets(uint256 , address , uint256 ) view returns(uint8 outcome, uint256 amount, uint256 oddsAtBet, bool claimed)
func (_PredictionMarket *PredictionMarketCallerSession) UserBets(arg0 *big.Int, arg1 common.Address, arg2 *big.Int) (struct {
	Outcome   uint8
	Amount    *big.Int
	OddsAtBet *big.Int
	Claimed   bool
}, error) {
	return _PredictionMarket.Contract.UserBets(&_PredictionMarket.CallOpts, arg0, arg1, arg2)
}

// CreateMarket is a paid mutator transaction binding the contract method 0xcfa51b16.
//
// Solidity: function createMarket(string _description, uint256 _resolutionTime, uint256 _minimumUpdateInterval) returns()
func (_PredictionMarket *PredictionMarketTransactor) CreateMarket(opts *bind.TransactOpts, _description string, _resolutionTime *big.Int, _minimumUpdateInterval *big.Int) (*types.Transaction, error) {
	return _PredictionMarket.contract.Transact(opts, "createMarket", _description, _resolutionTime, _minimumUpdateInterval)
}

// CreateMarket is a paid mutator transaction binding the contract method 0xcfa51b16.
//
// Solidity: function createMarket(string _description, uint256 _resolutionTime, uint256 _minimumUpdateInterval) returns()
func (_PredictionMarket *PredictionMarketSession) CreateMarket(_description string, _resolutionTime *big.Int, _minimumUpdateInterval *big.Int) (*types.Transaction, error) {
	return _PredictionMarket.Contract.CreateMarket(&_PredictionMarket.TransactOpts, _description, _resolutionTime, _minimumUpdateInterval)
}

// CreateMarket is a paid mutator transaction binding the contract method 0xcfa51b16.
//
// Solidity: function createMarket(string _description, uint256 _resolutionTime, uint256 _minimumUpdateInterval) returns()
func (_PredictionMarket *PredictionMarketTransactorSession) CreateMarket(_description string, _resolutionTime *big.Int, _minimumUpdateInterval *big.Int) (*types.Transaction, error) {
	return _PredictionMarket.Contract.CreateMarket(&_PredictionMarket.TransactOpts, _description, _resolutionTime, _minimumUpdateInterval)
}

// PlaceBet is a paid mutator transaction binding the contract method 0xda866c48.
//
// Solidity: function placeBet(uint256 _marketId, uint8 _outcome, uint256 _amount) returns()
func (_PredictionMarket *PredictionMarketTransactor) PlaceBet(opts *bind.TransactOpts, _marketId *big.Int, _outcome uint8, _amount *big.Int) (*types.Transaction, error) {
	return _PredictionMarket.contract.Transact(opts, "placeBet", _marketId, _outcome, _amount)
}

// PlaceBet is a paid mutator transaction binding the contract method 0xda866c48.
//
// Solidity: function placeBet(uint256 _marketId, uint8 _outcome, uint256 _amount) returns()
func (_PredictionMarket *PredictionMarketSession) PlaceBet(_marketId *big.Int, _outcome uint8, _amount *big.Int) (*types.Transaction, error) {
	return _PredictionMarket.Contract.PlaceBet(&_PredictionMarket.TransactOpts, _marketId, _outcome, _amount)
}

// PlaceBet is a paid mutator transaction binding the contract method 0xda866c48.
//
// Solidity: function placeBet(uint256 _marketId, uint8 _outcome, uint256 _amount) returns()
func (_PredictionMarket *PredictionMarketTransactorSession) PlaceBet(_marketId *big.Int, _outcome uint8, _amount *big.Int) (*types.Transaction, error) {
	return _PredictionMarket.Contract.PlaceBet(&_PredictionMarket.TransactOpts, _marketId, _outcome, _amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PredictionMarket *PredictionMarketTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PredictionMarket.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PredictionMarket *PredictionMarketSession) RenounceOwnership() (*types.Transaction, error) {
	return _PredictionMarket.Contract.RenounceOwnership(&_PredictionMarket.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PredictionMarket *PredictionMarketTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PredictionMarket.Contract.RenounceOwnership(&_PredictionMarket.TransactOpts)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _oracle) returns()
func (_PredictionMarket *PredictionMarketTransactor) SetOracle(opts *bind.TransactOpts, _oracle common.Address) (*types.Transaction, error) {
	return _PredictionMarket.contract.Transact(opts, "setOracle", _oracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _oracle) returns()
func (_PredictionMarket *PredictionMarketSession) SetOracle(_oracle common.Address) (*types.Transaction, error) {
	return _PredictionMarket.Contract.SetOracle(&_PredictionMarket.TransactOpts, _oracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _oracle) returns()
func (_PredictionMarket *PredictionMarketTransactorSession) SetOracle(_oracle common.Address) (*types.Transaction, error) {
	return _PredictionMarket.Contract.SetOracle(&_PredictionMarket.TransactOpts, _oracle)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PredictionMarket *PredictionMarketTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PredictionMarket.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PredictionMarket *PredictionMarketSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PredictionMarket.Contract.TransferOwnership(&_PredictionMarket.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PredictionMarket *PredictionMarketTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PredictionMarket.Contract.TransferOwnership(&_PredictionMarket.TransactOpts, newOwner)
}

// UpdateOracleData is a paid mutator transaction binding the contract method 0xe7b71012.
//
// Solidity: function updateOracleData(uint256 _marketId, uint256 _odds, uint256 _timestamp) returns()
func (_PredictionMarket *PredictionMarketTransactor) UpdateOracleData(opts *bind.TransactOpts, _marketId *big.Int, _odds *big.Int, _timestamp *big.Int) (*types.Transaction, error) {
	return _PredictionMarket.contract.Transact(opts, "updateOracleData", _marketId, _odds, _timestamp)
}

// UpdateOracleData is a paid mutator transaction binding the contract method 0xe7b71012.
//
// Solidity: function updateOracleData(uint256 _marketId, uint256 _odds, uint256 _timestamp) returns()
func (_PredictionMarket *PredictionMarketSession) UpdateOracleData(_marketId *big.Int, _odds *big.Int, _timestamp *big.Int) (*types.Transaction, error) {
	return _PredictionMarket.Contract.UpdateOracleData(&_PredictionMarket.TransactOpts, _marketId, _odds, _timestamp)
}

// UpdateOracleData is a paid mutator transaction binding the contract method 0xe7b71012.
//
// Solidity: function updateOracleData(uint256 _marketId, uint256 _odds, uint256 _timestamp) returns()
func (_PredictionMarket *PredictionMarketTransactorSession) UpdateOracleData(_marketId *big.Int, _odds *big.Int, _timestamp *big.Int) (*types.Transaction, error) {
	return _PredictionMarket.Contract.UpdateOracleData(&_PredictionMarket.TransactOpts, _marketId, _odds, _timestamp)
}

// PredictionMarketBetClaimedIterator is returned from FilterBetClaimed and is used to iterate over the raw logs and unpacked data for BetClaimed events raised by the PredictionMarket contract.
type PredictionMarketBetClaimedIterator struct {
	Event *PredictionMarketBetClaimed // Event containing the contract specifics and raw log

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
func (it *PredictionMarketBetClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PredictionMarketBetClaimed)
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
		it.Event = new(PredictionMarketBetClaimed)
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
func (it *PredictionMarketBetClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PredictionMarketBetClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PredictionMarketBetClaimed represents a BetClaimed event raised by the PredictionMarket contract.
type PredictionMarketBetClaimed struct {
	MarketId *big.Int
	Bettor   common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBetClaimed is a free log retrieval operation binding the contract event 0xad65ef4ba28f4ce1aaee7e8f86645eb4ce0ddc8b6c8fa2821564a94dbd2485fe.
//
// Solidity: event BetClaimed(uint256 indexed marketId, address indexed bettor, uint256 amount)
func (_PredictionMarket *PredictionMarketFilterer) FilterBetClaimed(opts *bind.FilterOpts, marketId []*big.Int, bettor []common.Address) (*PredictionMarketBetClaimedIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var bettorRule []interface{}
	for _, bettorItem := range bettor {
		bettorRule = append(bettorRule, bettorItem)
	}

	logs, sub, err := _PredictionMarket.contract.FilterLogs(opts, "BetClaimed", marketIdRule, bettorRule)
	if err != nil {
		return nil, err
	}
	return &PredictionMarketBetClaimedIterator{contract: _PredictionMarket.contract, event: "BetClaimed", logs: logs, sub: sub}, nil
}

// WatchBetClaimed is a free log subscription operation binding the contract event 0xad65ef4ba28f4ce1aaee7e8f86645eb4ce0ddc8b6c8fa2821564a94dbd2485fe.
//
// Solidity: event BetClaimed(uint256 indexed marketId, address indexed bettor, uint256 amount)
func (_PredictionMarket *PredictionMarketFilterer) WatchBetClaimed(opts *bind.WatchOpts, sink chan<- *PredictionMarketBetClaimed, marketId []*big.Int, bettor []common.Address) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var bettorRule []interface{}
	for _, bettorItem := range bettor {
		bettorRule = append(bettorRule, bettorItem)
	}

	logs, sub, err := _PredictionMarket.contract.WatchLogs(opts, "BetClaimed", marketIdRule, bettorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PredictionMarketBetClaimed)
				if err := _PredictionMarket.contract.UnpackLog(event, "BetClaimed", log); err != nil {
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

// ParseBetClaimed is a log parse operation binding the contract event 0xad65ef4ba28f4ce1aaee7e8f86645eb4ce0ddc8b6c8fa2821564a94dbd2485fe.
//
// Solidity: event BetClaimed(uint256 indexed marketId, address indexed bettor, uint256 amount)
func (_PredictionMarket *PredictionMarketFilterer) ParseBetClaimed(log types.Log) (*PredictionMarketBetClaimed, error) {
	event := new(PredictionMarketBetClaimed)
	if err := _PredictionMarket.contract.UnpackLog(event, "BetClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PredictionMarketBetPlacedIterator is returned from FilterBetPlaced and is used to iterate over the raw logs and unpacked data for BetPlaced events raised by the PredictionMarket contract.
type PredictionMarketBetPlacedIterator struct {
	Event *PredictionMarketBetPlaced // Event containing the contract specifics and raw log

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
func (it *PredictionMarketBetPlacedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PredictionMarketBetPlaced)
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
		it.Event = new(PredictionMarketBetPlaced)
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
func (it *PredictionMarketBetPlacedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PredictionMarketBetPlacedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PredictionMarketBetPlaced represents a BetPlaced event raised by the PredictionMarket contract.
type PredictionMarketBetPlaced struct {
	MarketId  *big.Int
	Bettor    common.Address
	Outcome   uint8
	Amount    *big.Int
	OddsAtBet *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBetPlaced is a free log retrieval operation binding the contract event 0xbfcd8f4e47cc529f86cdc7eab8fb80e015a298e021fcaee72df1f76dd57cf487.
//
// Solidity: event BetPlaced(uint256 indexed marketId, address indexed bettor, uint8 outcome, uint256 amount, uint256 oddsAtBet)
func (_PredictionMarket *PredictionMarketFilterer) FilterBetPlaced(opts *bind.FilterOpts, marketId []*big.Int, bettor []common.Address) (*PredictionMarketBetPlacedIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var bettorRule []interface{}
	for _, bettorItem := range bettor {
		bettorRule = append(bettorRule, bettorItem)
	}

	logs, sub, err := _PredictionMarket.contract.FilterLogs(opts, "BetPlaced", marketIdRule, bettorRule)
	if err != nil {
		return nil, err
	}
	return &PredictionMarketBetPlacedIterator{contract: _PredictionMarket.contract, event: "BetPlaced", logs: logs, sub: sub}, nil
}

// WatchBetPlaced is a free log subscription operation binding the contract event 0xbfcd8f4e47cc529f86cdc7eab8fb80e015a298e021fcaee72df1f76dd57cf487.
//
// Solidity: event BetPlaced(uint256 indexed marketId, address indexed bettor, uint8 outcome, uint256 amount, uint256 oddsAtBet)
func (_PredictionMarket *PredictionMarketFilterer) WatchBetPlaced(opts *bind.WatchOpts, sink chan<- *PredictionMarketBetPlaced, marketId []*big.Int, bettor []common.Address) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var bettorRule []interface{}
	for _, bettorItem := range bettor {
		bettorRule = append(bettorRule, bettorItem)
	}

	logs, sub, err := _PredictionMarket.contract.WatchLogs(opts, "BetPlaced", marketIdRule, bettorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PredictionMarketBetPlaced)
				if err := _PredictionMarket.contract.UnpackLog(event, "BetPlaced", log); err != nil {
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

// ParseBetPlaced is a log parse operation binding the contract event 0xbfcd8f4e47cc529f86cdc7eab8fb80e015a298e021fcaee72df1f76dd57cf487.
//
// Solidity: event BetPlaced(uint256 indexed marketId, address indexed bettor, uint8 outcome, uint256 amount, uint256 oddsAtBet)
func (_PredictionMarket *PredictionMarketFilterer) ParseBetPlaced(log types.Log) (*PredictionMarketBetPlaced, error) {
	event := new(PredictionMarketBetPlaced)
	if err := _PredictionMarket.contract.UnpackLog(event, "BetPlaced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PredictionMarketMarketCancelledIterator is returned from FilterMarketCancelled and is used to iterate over the raw logs and unpacked data for MarketCancelled events raised by the PredictionMarket contract.
type PredictionMarketMarketCancelledIterator struct {
	Event *PredictionMarketMarketCancelled // Event containing the contract specifics and raw log

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
func (it *PredictionMarketMarketCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PredictionMarketMarketCancelled)
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
		it.Event = new(PredictionMarketMarketCancelled)
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
func (it *PredictionMarketMarketCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PredictionMarketMarketCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PredictionMarketMarketCancelled represents a MarketCancelled event raised by the PredictionMarket contract.
type PredictionMarketMarketCancelled struct {
	MarketId *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMarketCancelled is a free log retrieval operation binding the contract event 0x2ca440fb7fca85d7f55d395a4abd94817330b83a62f3502efbb4770144e4ca97.
//
// Solidity: event MarketCancelled(uint256 indexed marketId)
func (_PredictionMarket *PredictionMarketFilterer) FilterMarketCancelled(opts *bind.FilterOpts, marketId []*big.Int) (*PredictionMarketMarketCancelledIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _PredictionMarket.contract.FilterLogs(opts, "MarketCancelled", marketIdRule)
	if err != nil {
		return nil, err
	}
	return &PredictionMarketMarketCancelledIterator{contract: _PredictionMarket.contract, event: "MarketCancelled", logs: logs, sub: sub}, nil
}

// WatchMarketCancelled is a free log subscription operation binding the contract event 0x2ca440fb7fca85d7f55d395a4abd94817330b83a62f3502efbb4770144e4ca97.
//
// Solidity: event MarketCancelled(uint256 indexed marketId)
func (_PredictionMarket *PredictionMarketFilterer) WatchMarketCancelled(opts *bind.WatchOpts, sink chan<- *PredictionMarketMarketCancelled, marketId []*big.Int) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _PredictionMarket.contract.WatchLogs(opts, "MarketCancelled", marketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PredictionMarketMarketCancelled)
				if err := _PredictionMarket.contract.UnpackLog(event, "MarketCancelled", log); err != nil {
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

// ParseMarketCancelled is a log parse operation binding the contract event 0x2ca440fb7fca85d7f55d395a4abd94817330b83a62f3502efbb4770144e4ca97.
//
// Solidity: event MarketCancelled(uint256 indexed marketId)
func (_PredictionMarket *PredictionMarketFilterer) ParseMarketCancelled(log types.Log) (*PredictionMarketMarketCancelled, error) {
	event := new(PredictionMarketMarketCancelled)
	if err := _PredictionMarket.contract.UnpackLog(event, "MarketCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PredictionMarketMarketCreatedIterator is returned from FilterMarketCreated and is used to iterate over the raw logs and unpacked data for MarketCreated events raised by the PredictionMarket contract.
type PredictionMarketMarketCreatedIterator struct {
	Event *PredictionMarketMarketCreated // Event containing the contract specifics and raw log

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
func (it *PredictionMarketMarketCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PredictionMarketMarketCreated)
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
		it.Event = new(PredictionMarketMarketCreated)
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
func (it *PredictionMarketMarketCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PredictionMarketMarketCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PredictionMarketMarketCreated represents a MarketCreated event raised by the PredictionMarket contract.
type PredictionMarketMarketCreated struct {
	MarketId       *big.Int
	Description    string
	ResolutionTime *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterMarketCreated is a free log retrieval operation binding the contract event 0x2d1e9ad45dbe8e898e0374deebf2a661b2ae7c1855d2a29507d584c35d287c75.
//
// Solidity: event MarketCreated(uint256 indexed marketId, string description, uint256 resolutionTime)
func (_PredictionMarket *PredictionMarketFilterer) FilterMarketCreated(opts *bind.FilterOpts, marketId []*big.Int) (*PredictionMarketMarketCreatedIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _PredictionMarket.contract.FilterLogs(opts, "MarketCreated", marketIdRule)
	if err != nil {
		return nil, err
	}
	return &PredictionMarketMarketCreatedIterator{contract: _PredictionMarket.contract, event: "MarketCreated", logs: logs, sub: sub}, nil
}

// WatchMarketCreated is a free log subscription operation binding the contract event 0x2d1e9ad45dbe8e898e0374deebf2a661b2ae7c1855d2a29507d584c35d287c75.
//
// Solidity: event MarketCreated(uint256 indexed marketId, string description, uint256 resolutionTime)
func (_PredictionMarket *PredictionMarketFilterer) WatchMarketCreated(opts *bind.WatchOpts, sink chan<- *PredictionMarketMarketCreated, marketId []*big.Int) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _PredictionMarket.contract.WatchLogs(opts, "MarketCreated", marketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PredictionMarketMarketCreated)
				if err := _PredictionMarket.contract.UnpackLog(event, "MarketCreated", log); err != nil {
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

// ParseMarketCreated is a log parse operation binding the contract event 0x2d1e9ad45dbe8e898e0374deebf2a661b2ae7c1855d2a29507d584c35d287c75.
//
// Solidity: event MarketCreated(uint256 indexed marketId, string description, uint256 resolutionTime)
func (_PredictionMarket *PredictionMarketFilterer) ParseMarketCreated(log types.Log) (*PredictionMarketMarketCreated, error) {
	event := new(PredictionMarketMarketCreated)
	if err := _PredictionMarket.contract.UnpackLog(event, "MarketCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PredictionMarketMarketResolvedIterator is returned from FilterMarketResolved and is used to iterate over the raw logs and unpacked data for MarketResolved events raised by the PredictionMarket contract.
type PredictionMarketMarketResolvedIterator struct {
	Event *PredictionMarketMarketResolved // Event containing the contract specifics and raw log

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
func (it *PredictionMarketMarketResolvedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PredictionMarketMarketResolved)
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
		it.Event = new(PredictionMarketMarketResolved)
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
func (it *PredictionMarketMarketResolvedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PredictionMarketMarketResolvedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PredictionMarketMarketResolved represents a MarketResolved event raised by the PredictionMarket contract.
type PredictionMarketMarketResolved struct {
	MarketId       *big.Int
	WinningOutcome uint8
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterMarketResolved is a free log retrieval operation binding the contract event 0x739f283563fb51ab6b89ee95d937b2e63a6cfcb83c385dbebb629f9d97bd43e6.
//
// Solidity: event MarketResolved(uint256 indexed marketId, uint8 winningOutcome)
func (_PredictionMarket *PredictionMarketFilterer) FilterMarketResolved(opts *bind.FilterOpts, marketId []*big.Int) (*PredictionMarketMarketResolvedIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _PredictionMarket.contract.FilterLogs(opts, "MarketResolved", marketIdRule)
	if err != nil {
		return nil, err
	}
	return &PredictionMarketMarketResolvedIterator{contract: _PredictionMarket.contract, event: "MarketResolved", logs: logs, sub: sub}, nil
}

// WatchMarketResolved is a free log subscription operation binding the contract event 0x739f283563fb51ab6b89ee95d937b2e63a6cfcb83c385dbebb629f9d97bd43e6.
//
// Solidity: event MarketResolved(uint256 indexed marketId, uint8 winningOutcome)
func (_PredictionMarket *PredictionMarketFilterer) WatchMarketResolved(opts *bind.WatchOpts, sink chan<- *PredictionMarketMarketResolved, marketId []*big.Int) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _PredictionMarket.contract.WatchLogs(opts, "MarketResolved", marketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PredictionMarketMarketResolved)
				if err := _PredictionMarket.contract.UnpackLog(event, "MarketResolved", log); err != nil {
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

// ParseMarketResolved is a log parse operation binding the contract event 0x739f283563fb51ab6b89ee95d937b2e63a6cfcb83c385dbebb629f9d97bd43e6.
//
// Solidity: event MarketResolved(uint256 indexed marketId, uint8 winningOutcome)
func (_PredictionMarket *PredictionMarketFilterer) ParseMarketResolved(log types.Log) (*PredictionMarketMarketResolved, error) {
	event := new(PredictionMarketMarketResolved)
	if err := _PredictionMarket.contract.UnpackLog(event, "MarketResolved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PredictionMarketOracleAddressUpdatedIterator is returned from FilterOracleAddressUpdated and is used to iterate over the raw logs and unpacked data for OracleAddressUpdated events raised by the PredictionMarket contract.
type PredictionMarketOracleAddressUpdatedIterator struct {
	Event *PredictionMarketOracleAddressUpdated // Event containing the contract specifics and raw log

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
func (it *PredictionMarketOracleAddressUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PredictionMarketOracleAddressUpdated)
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
		it.Event = new(PredictionMarketOracleAddressUpdated)
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
func (it *PredictionMarketOracleAddressUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PredictionMarketOracleAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PredictionMarketOracleAddressUpdated represents a OracleAddressUpdated event raised by the PredictionMarket contract.
type PredictionMarketOracleAddressUpdated struct {
	NewOracle common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOracleAddressUpdated is a free log retrieval operation binding the contract event 0x107a9fafffb7ac890f780879e423760c9ffea8dcee8045681f40f542aede2cb8.
//
// Solidity: event OracleAddressUpdated(address newOracle)
func (_PredictionMarket *PredictionMarketFilterer) FilterOracleAddressUpdated(opts *bind.FilterOpts) (*PredictionMarketOracleAddressUpdatedIterator, error) {

	logs, sub, err := _PredictionMarket.contract.FilterLogs(opts, "OracleAddressUpdated")
	if err != nil {
		return nil, err
	}
	return &PredictionMarketOracleAddressUpdatedIterator{contract: _PredictionMarket.contract, event: "OracleAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchOracleAddressUpdated is a free log subscription operation binding the contract event 0x107a9fafffb7ac890f780879e423760c9ffea8dcee8045681f40f542aede2cb8.
//
// Solidity: event OracleAddressUpdated(address newOracle)
func (_PredictionMarket *PredictionMarketFilterer) WatchOracleAddressUpdated(opts *bind.WatchOpts, sink chan<- *PredictionMarketOracleAddressUpdated) (event.Subscription, error) {

	logs, sub, err := _PredictionMarket.contract.WatchLogs(opts, "OracleAddressUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PredictionMarketOracleAddressUpdated)
				if err := _PredictionMarket.contract.UnpackLog(event, "OracleAddressUpdated", log); err != nil {
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

// ParseOracleAddressUpdated is a log parse operation binding the contract event 0x107a9fafffb7ac890f780879e423760c9ffea8dcee8045681f40f542aede2cb8.
//
// Solidity: event OracleAddressUpdated(address newOracle)
func (_PredictionMarket *PredictionMarketFilterer) ParseOracleAddressUpdated(log types.Log) (*PredictionMarketOracleAddressUpdated, error) {
	event := new(PredictionMarketOracleAddressUpdated)
	if err := _PredictionMarket.contract.UnpackLog(event, "OracleAddressUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PredictionMarketOracleUpdatedIterator is returned from FilterOracleUpdated and is used to iterate over the raw logs and unpacked data for OracleUpdated events raised by the PredictionMarket contract.
type PredictionMarketOracleUpdatedIterator struct {
	Event *PredictionMarketOracleUpdated // Event containing the contract specifics and raw log

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
func (it *PredictionMarketOracleUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PredictionMarketOracleUpdated)
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
		it.Event = new(PredictionMarketOracleUpdated)
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
func (it *PredictionMarketOracleUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PredictionMarketOracleUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PredictionMarketOracleUpdated represents a OracleUpdated event raised by the PredictionMarket contract.
type PredictionMarketOracleUpdated struct {
	MarketId  *big.Int
	Odds      *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOracleUpdated is a free log retrieval operation binding the contract event 0x8d6131ead05d17bb37913bdff362964b1be7fc0762dfbf84399d91316c35f2ac.
//
// Solidity: event OracleUpdated(uint256 indexed marketId, uint256 odds, uint256 timestamp)
func (_PredictionMarket *PredictionMarketFilterer) FilterOracleUpdated(opts *bind.FilterOpts, marketId []*big.Int) (*PredictionMarketOracleUpdatedIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _PredictionMarket.contract.FilterLogs(opts, "OracleUpdated", marketIdRule)
	if err != nil {
		return nil, err
	}
	return &PredictionMarketOracleUpdatedIterator{contract: _PredictionMarket.contract, event: "OracleUpdated", logs: logs, sub: sub}, nil
}

// WatchOracleUpdated is a free log subscription operation binding the contract event 0x8d6131ead05d17bb37913bdff362964b1be7fc0762dfbf84399d91316c35f2ac.
//
// Solidity: event OracleUpdated(uint256 indexed marketId, uint256 odds, uint256 timestamp)
func (_PredictionMarket *PredictionMarketFilterer) WatchOracleUpdated(opts *bind.WatchOpts, sink chan<- *PredictionMarketOracleUpdated, marketId []*big.Int) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _PredictionMarket.contract.WatchLogs(opts, "OracleUpdated", marketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PredictionMarketOracleUpdated)
				if err := _PredictionMarket.contract.UnpackLog(event, "OracleUpdated", log); err != nil {
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

// ParseOracleUpdated is a log parse operation binding the contract event 0x8d6131ead05d17bb37913bdff362964b1be7fc0762dfbf84399d91316c35f2ac.
//
// Solidity: event OracleUpdated(uint256 indexed marketId, uint256 odds, uint256 timestamp)
func (_PredictionMarket *PredictionMarketFilterer) ParseOracleUpdated(log types.Log) (*PredictionMarketOracleUpdated, error) {
	event := new(PredictionMarketOracleUpdated)
	if err := _PredictionMarket.contract.UnpackLog(event, "OracleUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PredictionMarketOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PredictionMarket contract.
type PredictionMarketOwnershipTransferredIterator struct {
	Event *PredictionMarketOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PredictionMarketOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PredictionMarketOwnershipTransferred)
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
		it.Event = new(PredictionMarketOwnershipTransferred)
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
func (it *PredictionMarketOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PredictionMarketOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PredictionMarketOwnershipTransferred represents a OwnershipTransferred event raised by the PredictionMarket contract.
type PredictionMarketOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PredictionMarket *PredictionMarketFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PredictionMarketOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PredictionMarket.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PredictionMarketOwnershipTransferredIterator{contract: _PredictionMarket.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PredictionMarket *PredictionMarketFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PredictionMarketOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PredictionMarket.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PredictionMarketOwnershipTransferred)
				if err := _PredictionMarket.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_PredictionMarket *PredictionMarketFilterer) ParseOwnershipTransferred(log types.Log) (*PredictionMarketOwnershipTransferred, error) {
	event := new(PredictionMarketOwnershipTransferred)
	if err := _PredictionMarket.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PredictionMarketPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the PredictionMarket contract.
type PredictionMarketPausedIterator struct {
	Event *PredictionMarketPaused // Event containing the contract specifics and raw log

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
func (it *PredictionMarketPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PredictionMarketPaused)
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
		it.Event = new(PredictionMarketPaused)
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
func (it *PredictionMarketPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PredictionMarketPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PredictionMarketPaused represents a Paused event raised by the PredictionMarket contract.
type PredictionMarketPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PredictionMarket *PredictionMarketFilterer) FilterPaused(opts *bind.FilterOpts) (*PredictionMarketPausedIterator, error) {

	logs, sub, err := _PredictionMarket.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PredictionMarketPausedIterator{contract: _PredictionMarket.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PredictionMarket *PredictionMarketFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PredictionMarketPaused) (event.Subscription, error) {

	logs, sub, err := _PredictionMarket.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PredictionMarketPaused)
				if err := _PredictionMarket.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PredictionMarket *PredictionMarketFilterer) ParsePaused(log types.Log) (*PredictionMarketPaused, error) {
	event := new(PredictionMarketPaused)
	if err := _PredictionMarket.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PredictionMarketUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the PredictionMarket contract.
type PredictionMarketUnpausedIterator struct {
	Event *PredictionMarketUnpaused // Event containing the contract specifics and raw log

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
func (it *PredictionMarketUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PredictionMarketUnpaused)
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
		it.Event = new(PredictionMarketUnpaused)
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
func (it *PredictionMarketUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PredictionMarketUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PredictionMarketUnpaused represents a Unpaused event raised by the PredictionMarket contract.
type PredictionMarketUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PredictionMarket *PredictionMarketFilterer) FilterUnpaused(opts *bind.FilterOpts) (*PredictionMarketUnpausedIterator, error) {

	logs, sub, err := _PredictionMarket.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PredictionMarketUnpausedIterator{contract: _PredictionMarket.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PredictionMarket *PredictionMarketFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PredictionMarketUnpaused) (event.Subscription, error) {

	logs, sub, err := _PredictionMarket.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PredictionMarketUnpaused)
				if err := _PredictionMarket.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PredictionMarket *PredictionMarketFilterer) ParseUnpaused(log types.Log) (*PredictionMarketUnpaused, error) {
	event := new(PredictionMarketUnpaused)
	if err := _PredictionMarket.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
