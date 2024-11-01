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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_betToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"MAX_FEE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_HISTORICAL_DATA_POINTS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"betToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cancelMarket\",\"inputs\":[{\"name\":\"_marketId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimBet\",\"inputs\":[{\"name\":\"_marketId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_betIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createMarket\",\"inputs\":[{\"name\":\"_description\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getHistoricalOdds\",\"inputs\":[{\"name\":\"_marketId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"odds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getHistoricalOddsCount\",\"inputs\":[{\"name\":\"_marketId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLatestOdds\",\"inputs\":[{\"name\":\"_marketId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"odds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMarketDetails\",\"inputs\":[{\"name\":\"_marketId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"totalPoolSize\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"resolutionTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"resolved\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"cancelled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"outcomePools\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"currentOdds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lastUpdateBlock\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUserBetCount\",\"inputs\":[{\"name\":\"_marketId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"marketCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"markets\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"totalPoolSize\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"resolutionTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"resolved\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"winningOutcome\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"cancelled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"currentOdds\",\"type\":\"tuple\",\"internalType\":\"structPredictionMarket.OracleData\",\"components\":[{\"name\":\"odds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"lastUpdateBlock\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minimumUpdateInterval\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"oracle\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"placeBet\",\"inputs\":[{\"name\":\"_marketId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_outcome\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"platformFee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"resolveMarket\",\"inputs\":[{\"name\":\"_marketId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_winningOutcome\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOracle\",\"inputs\":[{\"name\":\"_oracle\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPlatformFee\",\"inputs\":[{\"name\":\"_newFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateOracleData\",\"inputs\":[{\"name\":\"_marketId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_odds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"userBets\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"outcome\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"oddsAtBet\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"claimed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"BetClaimed\",\"inputs\":[{\"name\":\"marketId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"bettor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BetPlaced\",\"inputs\":[{\"name\":\"marketId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"bettor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"outcome\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"oddsAtBet\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MarketCancelled\",\"inputs\":[{\"name\":\"marketId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MarketCreated\",\"inputs\":[{\"name\":\"marketId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"description\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"resolutionTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MarketResolved\",\"inputs\":[{\"name\":\"marketId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"winningOutcome\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OracleAddressUpdated\",\"inputs\":[{\"name\":\"newOracle\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OracleUpdated\",\"inputs\":[{\"name\":\"marketId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"odds\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PlatformFeeUpdated\",\"inputs\":[{\"name\":\"newFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]}]",
	Bin: "0x60a0604052601960055534801561001557600080fd5b506040516119a73803806119a783398101604081905261003491610149565b6001600055806001600160a01b03811661006957604051631e4fbdf760e01b8152600060048201526024015b60405180910390fd5b610072816100db565b506001600160a01b0382166100c95760405162461bcd60e51b815260206004820152601560248201527f496e76616c696420746f6b656e206164647265737300000000000000000000006044820152606401610060565b506001600160a01b031660805261017c565b600180546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b80516001600160a01b038116811461014457600080fd5b919050565b6000806040838503121561015c57600080fd5b6101658361012d565b91506101736020840161012d565b90509250929050565b6080516118026101a56000396000818161024a01528181610b35015261104a01526118026000f3fe608060405234801561001057600080fd5b50600436106101585760003560e01c8063b06c1ba3116100c3578063df55406e1161007c578063df55406e1461035f578063e7b7101214610372578063ec97908214610385578063ecfa0ddf1461038e578063f2fde38b14610397578063fc63a350146103aa57600080fd5b8063b06c1ba3146102ce578063b1283e77146102f5578063bc063e1a1461031d578063cb39709a14610326578063d92683a914610339578063da866c481461034c57600080fd5b8063715018a611610115578063715018a61461023d57806378691f16146102455780637adbf973146102845780637c16cd9e146102975780637dc0d1d0146102aa5780638da5cb5b146102bd57600080fd5b806312e8e2c31461015d57806326232a2e146101725780633809ddda1461018e57806354888f55146101ce5780636000cdee146101e157806361c4cc1514610204575b600080fd5b61017061016b3660046111f5565b6103bd565b005b61017b60055481565b6040519081526020015b60405180910390f35b6101b961019c3660046111f5565b600090815260026020526040902060078101546008909101549091565b60408051928352602083019190915201610185565b6101706101dc366004611224565b610447565b61017b6101ef3660046111f5565b60009081526002602052604090206006015490565b6102176102123660046112f9565b6104db565b6040805160ff909516855260208501939093529183015215156060820152608001610185565b610170610536565b61026c7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b039091168152602001610185565b61017061029236600461132f565b61054a565b6101706102a53660046111f5565b6105ef565b60065461026c906001600160a01b031681565b6001546001600160a01b031661026c565b6102e16102dc3660046111f5565b610679565b604051610185989796959493929190611397565b6103086103033660046111f5565b6107f5565b6040516101859998979695949392919061141d565b61017b6101f481565b61017b610334366004611486565b6108e5565b6101b96103473660046114b2565b61090f565b61017061035a3660046114e5565b6109bb565b61017061036d36600461150a565b610d0c565b61017061038036600461152d565b505050565b61017b60035481565b61017b6103e881565b6101706103a536600461132f565b610e3f565b6101706103b83660046114b2565b610e7d565b6103c561114c565b6101f481111561040b5760405162461bcd60e51b815260206004820152600c60248201526b08ccaca40e8dede40d0d2ced60a31b60448201526064015b60405180910390fd5b60058190556040518181527f45610d581145924dd7090a5017e5f2b1d6f42213bb2e95707ff86846bbfcb1ca906020015b60405180910390a150565b61044f61114c565b60038054600091826104608361156f565b909155506000818152600260205260409020909150806104808482611610565b5061048e42620151806116cf565b60028201819055603c600a83015560405183917f2d1e9ad45dbe8e898e0374deebf2a661b2ae7c1855d2a29507d584c35d287c75916104ce9187916116e2565b60405180910390a2505050565b6004602052826000526040600020602052816000526040600020818154811061050357600080fd5b6000918252602090912060049091020180546001820154600283015460039093015460ff92831696509094509192501684565b61053e61114c565b6105486000611179565b565b61055261114c565b6001600160a01b0381166105a15760405162461bcd60e51b8152602060048201526016602482015275496e76616c6964206f7261636c65206164647265737360501b6044820152606401610402565b600680546001600160a01b0319166001600160a01b0383169081179091556040519081527f107a9fafffb7ac890f780879e423760c9ffea8dcee8045681f40f542aede2cb89060200161043c565b6105f761114c565b6000818152600260205260409020600381015460ff1615801561061f5750600581015460ff16155b61063b5760405162461bcd60e51b815260040161040290611704565b60058101805460ff1916600117905560405182907f2ca440fb7fca85d7f55d395a4abd94817330b83a62f3502efbb4770144e4ca9790600090a25050565b600081815260026020818152604080842081518481526060808201845295948594859485948994869485949293919290919083019080368337505060008080526004840160205260408120548351939750928792506106da576106da611748565b602002602001018181525050806004016000600160ff168152602001908152602001600020548460018151811061071357610713611748565b602090810291909101015260018101546002820154600383015460058401546007850154600986015486548796959460ff9081169416928b929091889061075990611588565b80601f016020809104026020016040519081016040528092919081815260200182805461078590611588565b80156107d25780601f106107a7576101008083540402835291602001916107d2565b820191906000526020600020905b8154815290600101906020018083116107b557829003601f168201915b505050505097509850985098509850985098509850985050919395975091939597565b60026020526000908152604090208054819061081090611588565b80601f016020809104026020016040519081016040528092919081815260200182805461083c90611588565b80156108895780601f1061085e57610100808354040283529160200191610889565b820191906000526020600020905b81548152906001019060200180831161086c57829003601f168201915b50505060018401546002850154600386015460058701546040805180820190915260078901548152600889015460208201526009890154600a909901549798949793965060ff8084169650610100909304831694919092169289565b60008281526004602090815260408083206001600160a01b03851684529091529020545b92915050565b6000828152600260205260408120600681015482919084106109695760405162461bcd60e51b8152602060048201526013602482015272496e646578206f7574206f6620626f756e647360681b6044820152606401610402565b600081600601858154811061098057610980611748565b600091825260209182902060408051808201909152600290920201805480835260019091015491909201819052909450925050509250929050565b6109c36111cb565b60008111610a135760405162461bcd60e51b815260206004820152601b60248201527f42657420616d6f756e74206d75737420626520706f73697469766500000000006044820152606401610402565b60018260ff161115610a595760405162461bcd60e51b815260206004820152600f60248201526e496e76616c6964206f7574636f6d6560881b6044820152606401610402565b6000838152600260205260409020600381015460ff16158015610a815750600581015460ff16155b610acd5760405162461bcd60e51b815260206004820152601760248201527f4d61726b6574206e6f206c6f6e676572206163746976650000000000000000006044820152606401610402565b80600201544210610b135760405162461bcd60e51b815260206004820152601060248201526f13585c9ad95d081a185cc8195b99195960821b6044820152606401610402565b6040516323b872dd60e01b8152336004820152306024820152604481018390527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906323b872dd906064016020604051808303816000875af1158015610b86573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610baa919061175e565b610bee5760405162461bcd60e51b8152602060048201526015602482015274151bdad95b881d1c985b9cd9995c8819985a5b1959605a1b6044820152606401610402565b81816001016000828254610c0291906116cf565b909155505060ff8316600090815260048201602052604081208054849290610c2b9084906116cf565b9091555050600084815260046020818152604080842033808652908352818520825160808101845260ff8a81168083528287018b815260078b01805485890190815260608087018d8152885460018082018b55998f529d8c902097519d909c0290960180549c90951660ff199c8d161785559151958401959095555160028301559651600390910180549115159190981617909655548251948552928401879052908301919091529186917fbfcd8f4e47cc529f86cdc7eab8fb80e015a298e021fcaee72df1f76dd57cf487910160405180910390a3506103806001600055565b610d1461114c565b6000828152600260205260409020600381015460ff16158015610d3c5750600581015460ff16155b610d585760405162461bcd60e51b815260040161040290611704565b8060020154421015610da35760405162461bcd60e51b8152602060048201526014602482015273546f6f206561726c7920746f207265736f6c766560601b6044820152606401610402565b60018260ff161115610de95760405162461bcd60e51b815260206004820152600f60248201526e496e76616c6964206f7574636f6d6560881b6044820152606401610402565b60038101805460ff84166101000261ffff1990911617600117905560405183907f739f283563fb51ab6b89ee95d937b2e63a6cfcb83c385dbebb629f9d97bd43e6906104ce90859060ff91909116815260200190565b610e4761114c565b6001600160a01b038116610e7157604051631e4fbdf760e01b815260006004820152602401610402565b610e7a81611179565b50565b610e856111cb565b6000828152600260205260409020600381015460ff1680610eaa5750600581015460ff165b610ef65760405162461bcd60e51b815260206004820181905260248201527f4d61726b6574206e6f74207265736f6c766564206f722063616e63656c6c65646044820152606401610402565b60008381526004602090815260408083203384529091528120805484908110610f2157610f21611748565b60009182526020909120600490910201600381015490915060ff1615610f7f5760405162461bcd60e51b815260206004820152601360248201527210995d08185b1c9958591e4818db185a5b5959606a1b6044820152606401610402565b60038101805460ff19166001179055600582015460009060ff1615610fa957506001810154611028565b6003830154825460ff61010090920482169116036110285760018301546003840154610100900460ff166000908152600485016020526040902054610ff081612710611780565b600554610fff90612710611797565b83866001015461100f9190611780565b6110199190611780565b61102391906117aa565b925050505b801561113b5760405163a9059cbb60e01b8152336004820152602481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063a9059cbb906044016020604051808303816000875af115801561109b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110bf919061175e565b6111035760405162461bcd60e51b8152602060048201526015602482015274151bdad95b881d1c985b9cd9995c8819985a5b1959605a1b6044820152606401610402565b604051818152339086907fad65ef4ba28f4ce1aaee7e8f86645eb4ce0ddc8b6c8fa2821564a94dbd2485fe9060200160405180910390a35b5050506111486001600055565b5050565b6001546001600160a01b031633146105485760405163118cdaa760e01b8152336004820152602401610402565b600180546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6002600054036111ee57604051633ee5aeb560e01b815260040160405180910390fd5b6002600055565b60006020828403121561120757600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b60006020828403121561123657600080fd5b813567ffffffffffffffff81111561124d57600080fd5b8201601f8101841361125e57600080fd5b803567ffffffffffffffff8111156112785761127861120e565b604051601f8201601f19908116603f0116810167ffffffffffffffff811182821017156112a7576112a761120e565b6040528181528282016020018610156112bf57600080fd5b81602084016020830137600091810160200191909152949350505050565b80356001600160a01b03811681146112f457600080fd5b919050565b60008060006060848603121561130e57600080fd5b8335925061131e602085016112dd565b929592945050506040919091013590565b60006020828403121561134157600080fd5b61134a826112dd565b9392505050565b6000815180845260005b818110156113775760208185018101518683018201520161135b565b506000602082860101526020601f19601f83011685010191505092915050565b610100815260006113ac61010083018b611351565b8960208401528860408401528715156060840152861515608084015282810360a084015280865180835260208301915060208801925060005b818110156114035783518352602093840193909201916001016113e5565b505060c0840195909552505060e001529695505050505050565b6101408152600061143261014083018c611351565b9050896020830152886040830152871515606083015260ff8716608083015285151560a0830152845160c0830152602085015160e083015283610100830152826101208301529a9950505050505050505050565b6000806040838503121561149957600080fd5b823591506114a9602084016112dd565b90509250929050565b600080604083850312156114c557600080fd5b50508035926020909101359150565b803560ff811681146112f457600080fd5b6000806000606084860312156114fa57600080fd5b8335925061131e602085016114d4565b6000806040838503121561151d57600080fd5b823591506114a9602084016114d4565b60008060006060848603121561154257600080fd5b505081359360208301359350604090920135919050565b634e487b7160e01b600052601160045260246000fd5b60006001820161158157611581611559565b5060010190565b600181811c9082168061159c57607f821691505b6020821081036115bc57634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111561038057806000526020600020601f840160051c810160208510156115e95750805b601f840160051c820191505b8181101561160957600081556001016115f5565b5050505050565b815167ffffffffffffffff81111561162a5761162a61120e565b61163e816116388454611588565b846115c2565b6020601f821160018114611672576000831561165a5750848201515b600019600385901b1c1916600184901b178455611609565b600084815260208120601f198516915b828110156116a25787850151825560209485019460019092019101611682565b50848210156116c05786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b8082018082111561090957610909611559565b6040815260006116f56040830185611351565b90508260208301529392505050565b60208082526024908201527f4d61726b657420616c7265616479207265736f6c766564206f722063616e63656040820152631b1b195960e21b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b60006020828403121561177057600080fd5b8151801515811461134a57600080fd5b808202811582820484141761090957610909611559565b8181038181111561090957610909611559565b6000826117c757634e487b7160e01b600052601260045260246000fd5b50049056fea264697066735822122056072b37cfba7cd3152ba9c1a4e2228d7d9f70e1389a82a46bb1679bc73575f164736f6c634300081a0033",
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

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_PredictionMarket *PredictionMarketCaller) MAXFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PredictionMarket.contract.Call(opts, &out, "MAX_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_PredictionMarket *PredictionMarketSession) MAXFEE() (*big.Int, error) {
	return _PredictionMarket.Contract.MAXFEE(&_PredictionMarket.CallOpts)
}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_PredictionMarket *PredictionMarketCallerSession) MAXFEE() (*big.Int, error) {
	return _PredictionMarket.Contract.MAXFEE(&_PredictionMarket.CallOpts)
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

// GetUserBetCount is a free data retrieval call binding the contract method 0xcb39709a.
//
// Solidity: function getUserBetCount(uint256 _marketId, address _user) view returns(uint256)
func (_PredictionMarket *PredictionMarketCaller) GetUserBetCount(opts *bind.CallOpts, _marketId *big.Int, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PredictionMarket.contract.Call(opts, &out, "getUserBetCount", _marketId, _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserBetCount is a free data retrieval call binding the contract method 0xcb39709a.
//
// Solidity: function getUserBetCount(uint256 _marketId, address _user) view returns(uint256)
func (_PredictionMarket *PredictionMarketSession) GetUserBetCount(_marketId *big.Int, _user common.Address) (*big.Int, error) {
	return _PredictionMarket.Contract.GetUserBetCount(&_PredictionMarket.CallOpts, _marketId, _user)
}

// GetUserBetCount is a free data retrieval call binding the contract method 0xcb39709a.
//
// Solidity: function getUserBetCount(uint256 _marketId, address _user) view returns(uint256)
func (_PredictionMarket *PredictionMarketCallerSession) GetUserBetCount(_marketId *big.Int, _user common.Address) (*big.Int, error) {
	return _PredictionMarket.Contract.GetUserBetCount(&_PredictionMarket.CallOpts, _marketId, _user)
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

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint256)
func (_PredictionMarket *PredictionMarketCaller) PlatformFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PredictionMarket.contract.Call(opts, &out, "platformFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint256)
func (_PredictionMarket *PredictionMarketSession) PlatformFee() (*big.Int, error) {
	return _PredictionMarket.Contract.PlatformFee(&_PredictionMarket.CallOpts)
}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint256)
func (_PredictionMarket *PredictionMarketCallerSession) PlatformFee() (*big.Int, error) {
	return _PredictionMarket.Contract.PlatformFee(&_PredictionMarket.CallOpts)
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

// CancelMarket is a paid mutator transaction binding the contract method 0x7c16cd9e.
//
// Solidity: function cancelMarket(uint256 _marketId) returns()
func (_PredictionMarket *PredictionMarketTransactor) CancelMarket(opts *bind.TransactOpts, _marketId *big.Int) (*types.Transaction, error) {
	return _PredictionMarket.contract.Transact(opts, "cancelMarket", _marketId)
}

// CancelMarket is a paid mutator transaction binding the contract method 0x7c16cd9e.
//
// Solidity: function cancelMarket(uint256 _marketId) returns()
func (_PredictionMarket *PredictionMarketSession) CancelMarket(_marketId *big.Int) (*types.Transaction, error) {
	return _PredictionMarket.Contract.CancelMarket(&_PredictionMarket.TransactOpts, _marketId)
}

// CancelMarket is a paid mutator transaction binding the contract method 0x7c16cd9e.
//
// Solidity: function cancelMarket(uint256 _marketId) returns()
func (_PredictionMarket *PredictionMarketTransactorSession) CancelMarket(_marketId *big.Int) (*types.Transaction, error) {
	return _PredictionMarket.Contract.CancelMarket(&_PredictionMarket.TransactOpts, _marketId)
}

// ClaimBet is a paid mutator transaction binding the contract method 0xfc63a350.
//
// Solidity: function claimBet(uint256 _marketId, uint256 _betIndex) returns()
func (_PredictionMarket *PredictionMarketTransactor) ClaimBet(opts *bind.TransactOpts, _marketId *big.Int, _betIndex *big.Int) (*types.Transaction, error) {
	return _PredictionMarket.contract.Transact(opts, "claimBet", _marketId, _betIndex)
}

// ClaimBet is a paid mutator transaction binding the contract method 0xfc63a350.
//
// Solidity: function claimBet(uint256 _marketId, uint256 _betIndex) returns()
func (_PredictionMarket *PredictionMarketSession) ClaimBet(_marketId *big.Int, _betIndex *big.Int) (*types.Transaction, error) {
	return _PredictionMarket.Contract.ClaimBet(&_PredictionMarket.TransactOpts, _marketId, _betIndex)
}

// ClaimBet is a paid mutator transaction binding the contract method 0xfc63a350.
//
// Solidity: function claimBet(uint256 _marketId, uint256 _betIndex) returns()
func (_PredictionMarket *PredictionMarketTransactorSession) ClaimBet(_marketId *big.Int, _betIndex *big.Int) (*types.Transaction, error) {
	return _PredictionMarket.Contract.ClaimBet(&_PredictionMarket.TransactOpts, _marketId, _betIndex)
}

// CreateMarket is a paid mutator transaction binding the contract method 0x54888f55.
//
// Solidity: function createMarket(string _description) returns()
func (_PredictionMarket *PredictionMarketTransactor) CreateMarket(opts *bind.TransactOpts, _description string) (*types.Transaction, error) {
	return _PredictionMarket.contract.Transact(opts, "createMarket", _description)
}

// CreateMarket is a paid mutator transaction binding the contract method 0x54888f55.
//
// Solidity: function createMarket(string _description) returns()
func (_PredictionMarket *PredictionMarketSession) CreateMarket(_description string) (*types.Transaction, error) {
	return _PredictionMarket.Contract.CreateMarket(&_PredictionMarket.TransactOpts, _description)
}

// CreateMarket is a paid mutator transaction binding the contract method 0x54888f55.
//
// Solidity: function createMarket(string _description) returns()
func (_PredictionMarket *PredictionMarketTransactorSession) CreateMarket(_description string) (*types.Transaction, error) {
	return _PredictionMarket.Contract.CreateMarket(&_PredictionMarket.TransactOpts, _description)
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

// ResolveMarket is a paid mutator transaction binding the contract method 0xdf55406e.
//
// Solidity: function resolveMarket(uint256 _marketId, uint8 _winningOutcome) returns()
func (_PredictionMarket *PredictionMarketTransactor) ResolveMarket(opts *bind.TransactOpts, _marketId *big.Int, _winningOutcome uint8) (*types.Transaction, error) {
	return _PredictionMarket.contract.Transact(opts, "resolveMarket", _marketId, _winningOutcome)
}

// ResolveMarket is a paid mutator transaction binding the contract method 0xdf55406e.
//
// Solidity: function resolveMarket(uint256 _marketId, uint8 _winningOutcome) returns()
func (_PredictionMarket *PredictionMarketSession) ResolveMarket(_marketId *big.Int, _winningOutcome uint8) (*types.Transaction, error) {
	return _PredictionMarket.Contract.ResolveMarket(&_PredictionMarket.TransactOpts, _marketId, _winningOutcome)
}

// ResolveMarket is a paid mutator transaction binding the contract method 0xdf55406e.
//
// Solidity: function resolveMarket(uint256 _marketId, uint8 _winningOutcome) returns()
func (_PredictionMarket *PredictionMarketTransactorSession) ResolveMarket(_marketId *big.Int, _winningOutcome uint8) (*types.Transaction, error) {
	return _PredictionMarket.Contract.ResolveMarket(&_PredictionMarket.TransactOpts, _marketId, _winningOutcome)
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

// SetPlatformFee is a paid mutator transaction binding the contract method 0x12e8e2c3.
//
// Solidity: function setPlatformFee(uint256 _newFee) returns()
func (_PredictionMarket *PredictionMarketTransactor) SetPlatformFee(opts *bind.TransactOpts, _newFee *big.Int) (*types.Transaction, error) {
	return _PredictionMarket.contract.Transact(opts, "setPlatformFee", _newFee)
}

// SetPlatformFee is a paid mutator transaction binding the contract method 0x12e8e2c3.
//
// Solidity: function setPlatformFee(uint256 _newFee) returns()
func (_PredictionMarket *PredictionMarketSession) SetPlatformFee(_newFee *big.Int) (*types.Transaction, error) {
	return _PredictionMarket.Contract.SetPlatformFee(&_PredictionMarket.TransactOpts, _newFee)
}

// SetPlatformFee is a paid mutator transaction binding the contract method 0x12e8e2c3.
//
// Solidity: function setPlatformFee(uint256 _newFee) returns()
func (_PredictionMarket *PredictionMarketTransactorSession) SetPlatformFee(_newFee *big.Int) (*types.Transaction, error) {
	return _PredictionMarket.Contract.SetPlatformFee(&_PredictionMarket.TransactOpts, _newFee)
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

// PredictionMarketPlatformFeeUpdatedIterator is returned from FilterPlatformFeeUpdated and is used to iterate over the raw logs and unpacked data for PlatformFeeUpdated events raised by the PredictionMarket contract.
type PredictionMarketPlatformFeeUpdatedIterator struct {
	Event *PredictionMarketPlatformFeeUpdated // Event containing the contract specifics and raw log

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
func (it *PredictionMarketPlatformFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PredictionMarketPlatformFeeUpdated)
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
		it.Event = new(PredictionMarketPlatformFeeUpdated)
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
func (it *PredictionMarketPlatformFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PredictionMarketPlatformFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PredictionMarketPlatformFeeUpdated represents a PlatformFeeUpdated event raised by the PredictionMarket contract.
type PredictionMarketPlatformFeeUpdated struct {
	NewFee *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPlatformFeeUpdated is a free log retrieval operation binding the contract event 0x45610d581145924dd7090a5017e5f2b1d6f42213bb2e95707ff86846bbfcb1ca.
//
// Solidity: event PlatformFeeUpdated(uint256 newFee)
func (_PredictionMarket *PredictionMarketFilterer) FilterPlatformFeeUpdated(opts *bind.FilterOpts) (*PredictionMarketPlatformFeeUpdatedIterator, error) {

	logs, sub, err := _PredictionMarket.contract.FilterLogs(opts, "PlatformFeeUpdated")
	if err != nil {
		return nil, err
	}
	return &PredictionMarketPlatformFeeUpdatedIterator{contract: _PredictionMarket.contract, event: "PlatformFeeUpdated", logs: logs, sub: sub}, nil
}

// WatchPlatformFeeUpdated is a free log subscription operation binding the contract event 0x45610d581145924dd7090a5017e5f2b1d6f42213bb2e95707ff86846bbfcb1ca.
//
// Solidity: event PlatformFeeUpdated(uint256 newFee)
func (_PredictionMarket *PredictionMarketFilterer) WatchPlatformFeeUpdated(opts *bind.WatchOpts, sink chan<- *PredictionMarketPlatformFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _PredictionMarket.contract.WatchLogs(opts, "PlatformFeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PredictionMarketPlatformFeeUpdated)
				if err := _PredictionMarket.contract.UnpackLog(event, "PlatformFeeUpdated", log); err != nil {
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

// ParsePlatformFeeUpdated is a log parse operation binding the contract event 0x45610d581145924dd7090a5017e5f2b1d6f42213bb2e95707ff86846bbfcb1ca.
//
// Solidity: event PlatformFeeUpdated(uint256 newFee)
func (_PredictionMarket *PredictionMarketFilterer) ParsePlatformFeeUpdated(log types.Log) (*PredictionMarketPlatformFeeUpdated, error) {
	event := new(PredictionMarketPlatformFeeUpdated)
	if err := _PredictionMarket.contract.UnpackLog(event, "PlatformFeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
