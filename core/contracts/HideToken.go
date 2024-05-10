// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// HideTokenZKProof is an auto generated low-level Go binding around an user-defined struct.
type HideTokenZKProof struct {
	PA [2]*big.Int
	PB [2][2]*big.Int
	PC [2]*big.Int
}

// HideTokenMetaData contains all meta data concerning the HideToken contract.
var HideTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"balanceVerifierAddr_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"transferVerifierAddr_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"effectiveBlock\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"CancelTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"ConfirmTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RegisterWallet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"valueBytes\",\"type\":\"bytes\"}],\"name\":\"RequestTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"_pA\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"_pB\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"_pC\",\"type\":\"uint256[2]\"}],\"internalType\":\"structHideToken.ZKProof\",\"name\":\"fromProof\",\"type\":\"tuple\"},{\"internalType\":\"uint256[4]\",\"name\":\"fromPubSignals\",\"type\":\"uint256[4]\"}],\"name\":\"cancelTransfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"_pA\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"_pB\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"_pC\",\"type\":\"uint256[2]\"}],\"internalType\":\"structHideToken.ZKProof\",\"name\":\"toProof\",\"type\":\"tuple\"},{\"internalType\":\"uint256[4]\",\"name\":\"toPubSignals\",\"type\":\"uint256[4]\"}],\"name\":\"confirmTransfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals_\",\"type\":\"uint8\"}],\"name\":\"initHideToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"_pA\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"_pB\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"_pC\",\"type\":\"uint256[2]\"}],\"internalType\":\"structHideToken.ZKProof\",\"name\":\"proof\",\"type\":\"tuple\"},{\"internalType\":\"uint256[4]\",\"name\":\"_pubSignals\",\"type\":\"uint256[4]\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"queryPublicKey\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"queryTransfer\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"_pA\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"_pB\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"_pC\",\"type\":\"uint256[2]\"}],\"internalType\":\"structHideToken.ZKProof\",\"name\":\"proof\",\"type\":\"tuple\"},{\"internalType\":\"uint256[3]\",\"name\":\"_pubSignals\",\"type\":\"uint256[3]\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"registerWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"valueBytes\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"_pA\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"_pB\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"_pC\",\"type\":\"uint256[2]\"}],\"internalType\":\"structHideToken.ZKProof\",\"name\":\"fromProof\",\"type\":\"tuple\"},{\"internalType\":\"uint256[4]\",\"name\":\"fromPubSignals\",\"type\":\"uint256[4]\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162001bc138038062001bc18339810160408190526200003491620000e5565b6200003f3362000078565b600180546001600160a01b039485166001600160a01b031991821617909155600280549390941692169190911790915560095562000126565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b80516001600160a01b0381168114620000e057600080fd5b919050565b600080600060608486031215620000fb57600080fd5b6200010684620000c8565b92506200011660208501620000c8565b9150604084015190509250925092565b611a8b80620001366000396000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c806360a12c071161009757806395d89b411161006657806395d89b411461021a578063a47b580614610222578063a5819bee14610235578063f2fde38b1461024857600080fd5b806360a12c07146101ad57806370a08231146101c0578063715018a6146101f75780638da5cb5b146101ff57600080fd5b8063313ce567116100d3578063313ce5671461014d57806342bbf9e114610162578063501bb7951461017757806353ea00041461019a57600080fd5b806306fdde03146100fa5780630dbef1dc146101185780631cfcee9f1461013a575b600080fd5b61010261025b565b60405161010f91906113dc565b60405180910390f35b61012b610126366004611412565b6102ed565b60405161010f93929190611445565b61010261014836600461146a565b6103eb565b60085460405160ff909116815260200161010f565b610175610170366004611528565b610497565b005b61018a6101853660046115d0565b6104d1565b604051901515815260200161010f565b6101756101a8366004611617565b6106ec565b61018a6101bb3660046115d0565b61094c565b6101e96101ce36600461146a565b6001600160a01b031660009081526003602052604090205490565b60405190815260200161010f565b610175610c32565b6000546040516001600160a01b03909116815260200161010f565b610102610c46565b61018a61023036600461167e565b610c55565b61018a6102433660046115d0565b610f3f565b61017561025636600461146a565b611218565b60606006805461026a906116f1565b80601f0160208091040260200160405190810160405280929190818152602001828054610296906116f1565b80156102e35780601f106102b8576101008083540402835291602001916102e3565b820191906000526020600020905b8154815290600101906020018083116102c657829003601f168201915b5050505050905090565b60606000806102fc8585611291565b156103225760405162461bcd60e51b815260040161031990611725565b60405180910390fd5b6001600160a01b038581166000908152600560209081526040808320938816835292905220600181015460028201548254839061035e906116f1565b80601f016020809104026020016040519081016040528092919081815260200182805461038a906116f1565b80156103d75780601f106103ac576101008083540402835291602001916103d7565b820191906000526020600020905b8154815290600101906020018083116103ba57829003601f168201915b505050505092509250925092509250925092565b6001600160a01b0381166000908152600460205260409020805460609190610412906116f1565b80601f016020809104026020016040519081016040528092919081815260200182805461043e906116f1565b801561048b5780601f106104605761010080835404028352916020019161048b565b820191906000526020600020905b81548152906001019060200180831161046e57829003601f168201915b50505050509050919050565b61049f6112ec565b60066104ab84826117ab565b5060076104b883826117ab565b506008805460ff191660ff929092169190911790555050565b6001600160a01b03831660009081526003602052604081205481036105305760405162461bcd60e51b81526020600482015260156024820152741859191c995cdcc81b5d5cdd081c9959da5cdd1959605a1b6044820152606401610319565b6001600160a01b0384166000908152600360205260409020548235146105ac5760405162461bcd60e51b815260206004820152602b60248201527f70726f6f662062616c616e636520696e707574206d757374206265206164647260448201526a6573732062616c616e636560a81b6064820152608401610319565b60025460408051635fe8c13b60e01b81526001600160a01b0390921691635fe8c13b916105e791879182019060c08301908890600401611895565b602060405180830381865afa158015610604573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061062891906118c6565b15156001146106495760405162461bcd60e51b8152600401610319906118e8565b600160608301351461068f5760405162461bcd60e51b815260206004820152600f60248201526e185919081b5d5cdd08189948185919608a1b6044820152606401610319565b6001600160a01b0384166000818152600360209081526040808320818701359055518186013581527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35060019392505050565b33604083013581146107335760405162461bcd60e51b815260206004820152601060248201526f6164647220696e636f686572656e636560801b6044820152606401610319565b6001600160a01b038116600090815260036020526040902054156107915760405162461bcd60e51b81526020600482015260156024820152741859191c995cdcc81a5cc81c9959da5cdd195c9959605a1b6044820152606401610319565b6020830135156107e35760405162461bcd60e51b815260206004820181905260248201527f70726f6f662062616c616e636520696e707574206d757374206265207a65726f6044820152606401610319565b815160208301206001600160a01b038281169116146108445760405162461bcd60e51b815260206004820152601d60248201527f7075626c6963206b6579206d757374206265207468652073656e6465720000006044820152606401610319565b600154604080516308a3cff560e11b81526001600160a01b03909216916311479fea9161087f91889182019060c08301908990600401611915565b602060405180830381865afa15801561089c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108c091906118c6565b15156001146108e15760405162461bcd60e51b8152600401610319906118e8565b6001600160a01b0381166000908152600360209081526040808320863590556004909152902061091183826117ab565b506040516001600160a01b038216907f3de0c4f89978c8c9ac219df0e8813f0b4a3df2fad979b3e3607cbe6c0ce6c76490600090a250505050565b3360008181526005602090815260408083206001600160a01b0388168452909152812060030154909190829060ff16600281111561098c5761098c611946565b146109a95760405162461bcd60e51b815260040161031990611725565b6001600160a01b03811660009081526003602052604081205490036109e05760405162461bcd60e51b81526004016103199061195c565b6001600160a01b0385166000908152600360205260408120549003610a175760405162461bcd60e51b815260040161031990611993565b6001606084013514610a5d5760405162461bcd60e51b815260206004820152600f60248201526e30b2321036bab9ba1031329039bab160891b6044820152606401610319565b6001600160a01b038116600090815260036020526040902054833514610a955760405162461bcd60e51b8152600401610319906119ca565b6001600160a01b0381811660009081526005602090815260408083209389168352928152919020600101549084013514610b075760405162461bcd60e51b81526020600482015260136024820152721d985b1d59481b5d5cdd08189948195c5d585b606a1b6044820152606401610319565b60025460408051635fe8c13b60e01b81526001600160a01b0390921691635fe8c13b91610b4291889182019060c08301908990600401611895565b602060405180830381865afa158015610b5f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b8391906118c6565b1515600114610ba45760405162461bcd60e51b8152600401610319906118e8565b6001600160a01b038181166000818152600360208181526040808420818a0135905560058252808420958b168085529590915290912001805460ff191660021790557f152fb15aa5d80f843e1e4bd5f2fc9161714f169945024decec7e84fb910fdd518560015b6020020135604051610c1f91815260200190565b60405180910390a3506001949350505050565b610c3a6112ec565b610c446000611346565b565b60606007805461026a906116f1565b600033610c628187611291565b610cae5760405162461bcd60e51b815260206004820152601d60248201527f6c617374207472616e7366657220686173206e6f7420657870697265640000006044820152606401610319565b6001600160a01b0381166000908152600360205260408120549003610ce55760405162461bcd60e51b81526004016103199061195c565b6001600160a01b0386166000908152600360205260408120549003610d1c5760405162461bcd60e51b815260040161031990611993565b606083013515610d5c5760405162461bcd60e51b815260206004820152600b60248201526a36bab9ba1031329039bab160a91b6044820152606401610319565b6001600160a01b038116600090815260036020526040902054833514610d945760405162461bcd60e51b8152600401610319906119ca565b6001600160a01b038082166000908152600560209081526040808320938a16835292905220610dc386826117ab565b506001600160a01b038181166000908152600560209081526040808320938a16835292815291902090840135600190910155600954610e029043611a13565b6001600160a01b0382811660009081526005602090815260408083208b8516845290915290819020600280820194909455600301805460ff1916905591548251635fe8c13b60e01b8152911691635fe8c13b91610e6e9188919082019060c08301908990600401611895565b602060405180830381865afa158015610e8b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610eaf91906118c6565b1515600114610ed05760405162461bcd60e51b8152600401610319906118e8565b6001600160a01b038181166000818152600360209081526040918290208288013590559051928916927fafaeb45426cfc6351df3bf2c13164ab06bc6abbbb28f2e2602117dcb63842deb91610f2b9190880135908a90611a34565b60405180910390a350600195945050505050565b600033610f4c8582611291565b15610f695760405162461bcd60e51b815260040161031990611725565b6001600160a01b0385166000908152600360205260408120549003610fa05760405162461bcd60e51b81526004016103199061195c565b6001600160a01b0381166000908152600360205260408120549003610fd75760405162461bcd60e51b815260040161031990611993565b600160608401351461101d5760405162461bcd60e51b815260206004820152600f60248201526e30b2321036bab9ba1031329039bab160891b6044820152606401610319565b6001600160a01b0381166000908152600360205260409020548335146110955760405162461bcd60e51b815260206004820152602760248201527f70726f6f6620696e747075742062616c616e6365206d75737420626520746f2060448201526662616c616e636560c81b6064820152608401610319565b6001600160a01b03858116600090815260056020908152604080832093851683529281529190206001015490840135146111075760405162461bcd60e51b81526020600482015260136024820152721d985b1d59481b5d5cdd08189948195c5d585b606a1b6044820152606401610319565b60025460408051635fe8c13b60e01b81526001600160a01b0390921691635fe8c13b9161114291889182019060c08301908990600401611895565b602060405180830381865afa15801561115f573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061118391906118c6565b15156001146111a45760405162461bcd60e51b8152600401610319906118e8565b6001600160a01b038181166000818152600360208181526040808420818a01359055948a16808452600582528584208585529091529390912001805460ff191660019081179091559091907f7dca1b3ed21e108d0813c2b6e42f0cdd726c9abe60d2ad9ed0b44e5dc1f98c16908690610c0b565b6112206112ec565b6001600160a01b0381166112855760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610319565b61128e81611346565b50565b6001600160a01b038083166000908152600560209081526040808320938516835292905290812081600382015460ff1660028111156112d2576112d2611946565b1415806112e25750438160020154105b9150505b92915050565b6000546001600160a01b03163314610c445760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610319565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000815180845260005b818110156113bc576020818501810151868301820152016113a0565b506000602082860101526020601f19601f83011685010191505092915050565b6020815260006113ef6020830184611396565b9392505050565b80356001600160a01b038116811461140d57600080fd5b919050565b6000806040838503121561142557600080fd5b61142e836113f6565b915061143c602084016113f6565b90509250929050565b6060815260006114586060830186611396565b60208301949094525060400152919050565b60006020828403121561147c57600080fd5b6113ef826113f6565b634e487b7160e01b600052604160045260246000fd5b600082601f8301126114ac57600080fd5b813567ffffffffffffffff808211156114c7576114c7611485565b604051601f8301601f19908116603f011681019082821181831017156114ef576114ef611485565b8160405283815286602085880101111561150857600080fd5b836020870160208301376000602085830101528094505050505092915050565b60008060006060848603121561153d57600080fd5b833567ffffffffffffffff8082111561155557600080fd5b6115618783880161149b565b9450602086013591508082111561157757600080fd5b506115848682870161149b565b925050604084013560ff8116811461159b57600080fd5b809150509250925092565b600061010082840312156115b957600080fd5b50919050565b80608081018310156112e657600080fd5b60008060006101a084860312156115e657600080fd5b6115ef846113f6565b92506115fe85602086016115a6565b915061160e8561012086016115bf565b90509250925092565b6000806000610180848603121561162d57600080fd5b61163785856115a6565b925061016084018581111561164b57600080fd5b610100850192503567ffffffffffffffff81111561166857600080fd5b6116748682870161149b565b9150509250925092565b6000806000806101c0858703121561169557600080fd5b61169e856113f6565b9350602085013567ffffffffffffffff8111156116ba57600080fd5b6116c68782880161149b565b9350506116d686604087016115a6565b91506116e68661014087016115bf565b905092959194509250565b600181811c9082168061170557607f821691505b6020821081036115b957634e487b7160e01b600052602260045260246000fd5b60208082526018908201527f746865207472616e736665722068617320657870697265640000000000000000604082015260600190565b601f8211156117a657600081815260208120601f850160051c810160208610156117835750805b601f850160051c820191505b818110156117a25782815560010161178f565b5050505b505050565b815167ffffffffffffffff8111156117c5576117c5611485565b6117d9816117d384546116f1565b8461175c565b602080601f83116001811461180e57600084156117f65750858301515b600019600386901b1c1916600185901b1785556117a2565b600085815260208120601f198616915b8281101561183d5788860151825594840194600190910190840161181e565b508582101561185b5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b8060005b600281101561188f5760408083863793840193919091019060010161186f565b50505050565b610180810160408683376118ac604083018661186b565b60408460c084013760808361010084013795945050505050565b6000602082840312156118d857600080fd5b815180151581146113ef57600080fd5b6020808252601390820152721c1c9bdbd9881d995c9a599e4819985a5b1959606a1b604082015260600190565b6101608101604086833761192c604083018661186b565b60408460c084013760608361010084013795945050505050565b634e487b7160e01b600052602160045260246000fd5b6020808252601c908201527f61646472657373206066726f6d60206d75737420726567697374656400000000604082015260600190565b6020808252601a908201527f616464726573732060746f60206d757374207265676973746564000000000000604082015260600190565b60208082526029908201527f70726f6f6620696e747075742062616c616e6365206d7573742062652066726f6040820152686d2062616c616e636560b81b606082015260800190565b808201808211156112e657634e487b7160e01b600052601160045260246000fd5b828152604060208201526000611a4d6040830184611396565b94935050505056fea26469706673582212201548fd8aeab711311d14ba7bfbd2d7056fc23c150abe1acaef0d5626bada9e8264736f6c63430008130033",
}

// HideTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use HideTokenMetaData.ABI instead.
var HideTokenABI = HideTokenMetaData.ABI

// HideTokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use HideTokenMetaData.Bin instead.
var HideTokenBin = HideTokenMetaData.Bin

// DeployHideToken deploys a new Ethereum contract, binding an instance of HideToken to it.
func DeployHideToken(auth *bind.TransactOpts, backend bind.ContractBackend, balanceVerifierAddr_ common.Address, transferVerifierAddr_ common.Address, effectiveBlock *big.Int) (common.Address, *types.Transaction, *HideToken, error) {
	parsed, err := HideTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(HideTokenBin), backend, balanceVerifierAddr_, transferVerifierAddr_, effectiveBlock)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HideToken{HideTokenCaller: HideTokenCaller{contract: contract}, HideTokenTransactor: HideTokenTransactor{contract: contract}, HideTokenFilterer: HideTokenFilterer{contract: contract}}, nil
}

// HideToken is an auto generated Go binding around an Ethereum contract.
type HideToken struct {
	HideTokenCaller     // Read-only binding to the contract
	HideTokenTransactor // Write-only binding to the contract
	HideTokenFilterer   // Log filterer for contract events
}

// HideTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type HideTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HideTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HideTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HideTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HideTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HideTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HideTokenSession struct {
	Contract     *HideToken        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HideTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HideTokenCallerSession struct {
	Contract *HideTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// HideTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HideTokenTransactorSession struct {
	Contract     *HideTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// HideTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type HideTokenRaw struct {
	Contract *HideToken // Generic contract binding to access the raw methods on
}

// HideTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HideTokenCallerRaw struct {
	Contract *HideTokenCaller // Generic read-only contract binding to access the raw methods on
}

// HideTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HideTokenTransactorRaw struct {
	Contract *HideTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHideToken creates a new instance of HideToken, bound to a specific deployed contract.
func NewHideToken(address common.Address, backend bind.ContractBackend) (*HideToken, error) {
	contract, err := bindHideToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HideToken{HideTokenCaller: HideTokenCaller{contract: contract}, HideTokenTransactor: HideTokenTransactor{contract: contract}, HideTokenFilterer: HideTokenFilterer{contract: contract}}, nil
}

// NewHideTokenCaller creates a new read-only instance of HideToken, bound to a specific deployed contract.
func NewHideTokenCaller(address common.Address, caller bind.ContractCaller) (*HideTokenCaller, error) {
	contract, err := bindHideToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HideTokenCaller{contract: contract}, nil
}

// NewHideTokenTransactor creates a new write-only instance of HideToken, bound to a specific deployed contract.
func NewHideTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*HideTokenTransactor, error) {
	contract, err := bindHideToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HideTokenTransactor{contract: contract}, nil
}

// NewHideTokenFilterer creates a new log filterer instance of HideToken, bound to a specific deployed contract.
func NewHideTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*HideTokenFilterer, error) {
	contract, err := bindHideToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HideTokenFilterer{contract: contract}, nil
}

// bindHideToken binds a generic wrapper to an already deployed contract.
func bindHideToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HideTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HideToken *HideTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HideToken.Contract.HideTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HideToken *HideTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HideToken.Contract.HideTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HideToken *HideTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HideToken.Contract.HideTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HideToken *HideTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HideToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HideToken *HideTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HideToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HideToken *HideTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HideToken.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_HideToken *HideTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _HideToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_HideToken *HideTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _HideToken.Contract.BalanceOf(&_HideToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_HideToken *HideTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _HideToken.Contract.BalanceOf(&_HideToken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_HideToken *HideTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _HideToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_HideToken *HideTokenSession) Decimals() (uint8, error) {
	return _HideToken.Contract.Decimals(&_HideToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_HideToken *HideTokenCallerSession) Decimals() (uint8, error) {
	return _HideToken.Contract.Decimals(&_HideToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_HideToken *HideTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _HideToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_HideToken *HideTokenSession) Name() (string, error) {
	return _HideToken.Contract.Name(&_HideToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_HideToken *HideTokenCallerSession) Name() (string, error) {
	return _HideToken.Contract.Name(&_HideToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_HideToken *HideTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HideToken.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_HideToken *HideTokenSession) Owner() (common.Address, error) {
	return _HideToken.Contract.Owner(&_HideToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_HideToken *HideTokenCallerSession) Owner() (common.Address, error) {
	return _HideToken.Contract.Owner(&_HideToken.CallOpts)
}

// QueryPublicKey is a free data retrieval call binding the contract method 0x1cfcee9f.
//
// Solidity: function queryPublicKey(address to) view returns(bytes)
func (_HideToken *HideTokenCaller) QueryPublicKey(opts *bind.CallOpts, to common.Address) ([]byte, error) {
	var out []interface{}
	err := _HideToken.contract.Call(opts, &out, "queryPublicKey", to)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// QueryPublicKey is a free data retrieval call binding the contract method 0x1cfcee9f.
//
// Solidity: function queryPublicKey(address to) view returns(bytes)
func (_HideToken *HideTokenSession) QueryPublicKey(to common.Address) ([]byte, error) {
	return _HideToken.Contract.QueryPublicKey(&_HideToken.CallOpts, to)
}

// QueryPublicKey is a free data retrieval call binding the contract method 0x1cfcee9f.
//
// Solidity: function queryPublicKey(address to) view returns(bytes)
func (_HideToken *HideTokenCallerSession) QueryPublicKey(to common.Address) ([]byte, error) {
	return _HideToken.Contract.QueryPublicKey(&_HideToken.CallOpts, to)
}

// QueryTransfer is a free data retrieval call binding the contract method 0x0dbef1dc.
//
// Solidity: function queryTransfer(address from, address to) view returns(bytes, uint256, uint256)
func (_HideToken *HideTokenCaller) QueryTransfer(opts *bind.CallOpts, from common.Address, to common.Address) ([]byte, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _HideToken.contract.Call(opts, &out, "queryTransfer", from, to)

	if err != nil {
		return *new([]byte), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// QueryTransfer is a free data retrieval call binding the contract method 0x0dbef1dc.
//
// Solidity: function queryTransfer(address from, address to) view returns(bytes, uint256, uint256)
func (_HideToken *HideTokenSession) QueryTransfer(from common.Address, to common.Address) ([]byte, *big.Int, *big.Int, error) {
	return _HideToken.Contract.QueryTransfer(&_HideToken.CallOpts, from, to)
}

// QueryTransfer is a free data retrieval call binding the contract method 0x0dbef1dc.
//
// Solidity: function queryTransfer(address from, address to) view returns(bytes, uint256, uint256)
func (_HideToken *HideTokenCallerSession) QueryTransfer(from common.Address, to common.Address) ([]byte, *big.Int, *big.Int, error) {
	return _HideToken.Contract.QueryTransfer(&_HideToken.CallOpts, from, to)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_HideToken *HideTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _HideToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_HideToken *HideTokenSession) Symbol() (string, error) {
	return _HideToken.Contract.Symbol(&_HideToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_HideToken *HideTokenCallerSession) Symbol() (string, error) {
	return _HideToken.Contract.Symbol(&_HideToken.CallOpts)
}

// CancelTransfer is a paid mutator transaction binding the contract method 0x60a12c07.
//
// Solidity: function cancelTransfer(address to, (uint256[2],uint256[2][2],uint256[2]) fromProof, uint256[4] fromPubSignals) returns(bool)
func (_HideToken *HideTokenTransactor) CancelTransfer(opts *bind.TransactOpts, to common.Address, fromProof HideTokenZKProof, fromPubSignals [4]*big.Int) (*types.Transaction, error) {
	return _HideToken.contract.Transact(opts, "cancelTransfer", to, fromProof, fromPubSignals)
}

// CancelTransfer is a paid mutator transaction binding the contract method 0x60a12c07.
//
// Solidity: function cancelTransfer(address to, (uint256[2],uint256[2][2],uint256[2]) fromProof, uint256[4] fromPubSignals) returns(bool)
func (_HideToken *HideTokenSession) CancelTransfer(to common.Address, fromProof HideTokenZKProof, fromPubSignals [4]*big.Int) (*types.Transaction, error) {
	return _HideToken.Contract.CancelTransfer(&_HideToken.TransactOpts, to, fromProof, fromPubSignals)
}

// CancelTransfer is a paid mutator transaction binding the contract method 0x60a12c07.
//
// Solidity: function cancelTransfer(address to, (uint256[2],uint256[2][2],uint256[2]) fromProof, uint256[4] fromPubSignals) returns(bool)
func (_HideToken *HideTokenTransactorSession) CancelTransfer(to common.Address, fromProof HideTokenZKProof, fromPubSignals [4]*big.Int) (*types.Transaction, error) {
	return _HideToken.Contract.CancelTransfer(&_HideToken.TransactOpts, to, fromProof, fromPubSignals)
}

// ConfirmTransfer is a paid mutator transaction binding the contract method 0xa5819bee.
//
// Solidity: function confirmTransfer(address from, (uint256[2],uint256[2][2],uint256[2]) toProof, uint256[4] toPubSignals) returns(bool)
func (_HideToken *HideTokenTransactor) ConfirmTransfer(opts *bind.TransactOpts, from common.Address, toProof HideTokenZKProof, toPubSignals [4]*big.Int) (*types.Transaction, error) {
	return _HideToken.contract.Transact(opts, "confirmTransfer", from, toProof, toPubSignals)
}

// ConfirmTransfer is a paid mutator transaction binding the contract method 0xa5819bee.
//
// Solidity: function confirmTransfer(address from, (uint256[2],uint256[2][2],uint256[2]) toProof, uint256[4] toPubSignals) returns(bool)
func (_HideToken *HideTokenSession) ConfirmTransfer(from common.Address, toProof HideTokenZKProof, toPubSignals [4]*big.Int) (*types.Transaction, error) {
	return _HideToken.Contract.ConfirmTransfer(&_HideToken.TransactOpts, from, toProof, toPubSignals)
}

// ConfirmTransfer is a paid mutator transaction binding the contract method 0xa5819bee.
//
// Solidity: function confirmTransfer(address from, (uint256[2],uint256[2][2],uint256[2]) toProof, uint256[4] toPubSignals) returns(bool)
func (_HideToken *HideTokenTransactorSession) ConfirmTransfer(from common.Address, toProof HideTokenZKProof, toPubSignals [4]*big.Int) (*types.Transaction, error) {
	return _HideToken.Contract.ConfirmTransfer(&_HideToken.TransactOpts, from, toProof, toPubSignals)
}

// InitHideToken is a paid mutator transaction binding the contract method 0x42bbf9e1.
//
// Solidity: function initHideToken(string name_, string symbol_, uint8 decimals_) returns()
func (_HideToken *HideTokenTransactor) InitHideToken(opts *bind.TransactOpts, name_ string, symbol_ string, decimals_ uint8) (*types.Transaction, error) {
	return _HideToken.contract.Transact(opts, "initHideToken", name_, symbol_, decimals_)
}

// InitHideToken is a paid mutator transaction binding the contract method 0x42bbf9e1.
//
// Solidity: function initHideToken(string name_, string symbol_, uint8 decimals_) returns()
func (_HideToken *HideTokenSession) InitHideToken(name_ string, symbol_ string, decimals_ uint8) (*types.Transaction, error) {
	return _HideToken.Contract.InitHideToken(&_HideToken.TransactOpts, name_, symbol_, decimals_)
}

// InitHideToken is a paid mutator transaction binding the contract method 0x42bbf9e1.
//
// Solidity: function initHideToken(string name_, string symbol_, uint8 decimals_) returns()
func (_HideToken *HideTokenTransactorSession) InitHideToken(name_ string, symbol_ string, decimals_ uint8) (*types.Transaction, error) {
	return _HideToken.Contract.InitHideToken(&_HideToken.TransactOpts, name_, symbol_, decimals_)
}

// Mint is a paid mutator transaction binding the contract method 0x501bb795.
//
// Solidity: function mint(address addr, (uint256[2],uint256[2][2],uint256[2]) proof, uint256[4] _pubSignals) returns(bool)
func (_HideToken *HideTokenTransactor) Mint(opts *bind.TransactOpts, addr common.Address, proof HideTokenZKProof, _pubSignals [4]*big.Int) (*types.Transaction, error) {
	return _HideToken.contract.Transact(opts, "mint", addr, proof, _pubSignals)
}

// Mint is a paid mutator transaction binding the contract method 0x501bb795.
//
// Solidity: function mint(address addr, (uint256[2],uint256[2][2],uint256[2]) proof, uint256[4] _pubSignals) returns(bool)
func (_HideToken *HideTokenSession) Mint(addr common.Address, proof HideTokenZKProof, _pubSignals [4]*big.Int) (*types.Transaction, error) {
	return _HideToken.Contract.Mint(&_HideToken.TransactOpts, addr, proof, _pubSignals)
}

// Mint is a paid mutator transaction binding the contract method 0x501bb795.
//
// Solidity: function mint(address addr, (uint256[2],uint256[2][2],uint256[2]) proof, uint256[4] _pubSignals) returns(bool)
func (_HideToken *HideTokenTransactorSession) Mint(addr common.Address, proof HideTokenZKProof, _pubSignals [4]*big.Int) (*types.Transaction, error) {
	return _HideToken.Contract.Mint(&_HideToken.TransactOpts, addr, proof, _pubSignals)
}

// RegisterWallet is a paid mutator transaction binding the contract method 0x53ea0004.
//
// Solidity: function registerWallet((uint256[2],uint256[2][2],uint256[2]) proof, uint256[3] _pubSignals, bytes publicKey) returns()
func (_HideToken *HideTokenTransactor) RegisterWallet(opts *bind.TransactOpts, proof HideTokenZKProof, _pubSignals [3]*big.Int, publicKey []byte) (*types.Transaction, error) {
	return _HideToken.contract.Transact(opts, "registerWallet", proof, _pubSignals, publicKey)
}

// RegisterWallet is a paid mutator transaction binding the contract method 0x53ea0004.
//
// Solidity: function registerWallet((uint256[2],uint256[2][2],uint256[2]) proof, uint256[3] _pubSignals, bytes publicKey) returns()
func (_HideToken *HideTokenSession) RegisterWallet(proof HideTokenZKProof, _pubSignals [3]*big.Int, publicKey []byte) (*types.Transaction, error) {
	return _HideToken.Contract.RegisterWallet(&_HideToken.TransactOpts, proof, _pubSignals, publicKey)
}

// RegisterWallet is a paid mutator transaction binding the contract method 0x53ea0004.
//
// Solidity: function registerWallet((uint256[2],uint256[2][2],uint256[2]) proof, uint256[3] _pubSignals, bytes publicKey) returns()
func (_HideToken *HideTokenTransactorSession) RegisterWallet(proof HideTokenZKProof, _pubSignals [3]*big.Int, publicKey []byte) (*types.Transaction, error) {
	return _HideToken.Contract.RegisterWallet(&_HideToken.TransactOpts, proof, _pubSignals, publicKey)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_HideToken *HideTokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HideToken.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_HideToken *HideTokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _HideToken.Contract.RenounceOwnership(&_HideToken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_HideToken *HideTokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _HideToken.Contract.RenounceOwnership(&_HideToken.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa47b5806.
//
// Solidity: function transfer(address to, bytes valueBytes, (uint256[2],uint256[2][2],uint256[2]) fromProof, uint256[4] fromPubSignals) returns(bool)
func (_HideToken *HideTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, valueBytes []byte, fromProof HideTokenZKProof, fromPubSignals [4]*big.Int) (*types.Transaction, error) {
	return _HideToken.contract.Transact(opts, "transfer", to, valueBytes, fromProof, fromPubSignals)
}

// Transfer is a paid mutator transaction binding the contract method 0xa47b5806.
//
// Solidity: function transfer(address to, bytes valueBytes, (uint256[2],uint256[2][2],uint256[2]) fromProof, uint256[4] fromPubSignals) returns(bool)
func (_HideToken *HideTokenSession) Transfer(to common.Address, valueBytes []byte, fromProof HideTokenZKProof, fromPubSignals [4]*big.Int) (*types.Transaction, error) {
	return _HideToken.Contract.Transfer(&_HideToken.TransactOpts, to, valueBytes, fromProof, fromPubSignals)
}

// Transfer is a paid mutator transaction binding the contract method 0xa47b5806.
//
// Solidity: function transfer(address to, bytes valueBytes, (uint256[2],uint256[2][2],uint256[2]) fromProof, uint256[4] fromPubSignals) returns(bool)
func (_HideToken *HideTokenTransactorSession) Transfer(to common.Address, valueBytes []byte, fromProof HideTokenZKProof, fromPubSignals [4]*big.Int) (*types.Transaction, error) {
	return _HideToken.Contract.Transfer(&_HideToken.TransactOpts, to, valueBytes, fromProof, fromPubSignals)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_HideToken *HideTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _HideToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_HideToken *HideTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _HideToken.Contract.TransferOwnership(&_HideToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_HideToken *HideTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _HideToken.Contract.TransferOwnership(&_HideToken.TransactOpts, newOwner)
}

// HideTokenCancelTransferIterator is returned from FilterCancelTransfer and is used to iterate over the raw logs and unpacked data for CancelTransfer events raised by the HideToken contract.
type HideTokenCancelTransferIterator struct {
	Event *HideTokenCancelTransfer // Event containing the contract specifics and raw log

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
func (it *HideTokenCancelTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HideTokenCancelTransfer)
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
		it.Event = new(HideTokenCancelTransfer)
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
func (it *HideTokenCancelTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HideTokenCancelTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HideTokenCancelTransfer represents a CancelTransfer event raised by the HideToken contract.
type HideTokenCancelTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterCancelTransfer is a free log retrieval operation binding the contract event 0x152fb15aa5d80f843e1e4bd5f2fc9161714f169945024decec7e84fb910fdd51.
//
// Solidity: event CancelTransfer(address indexed from, address indexed to, uint256 value)
func (_HideToken *HideTokenFilterer) FilterCancelTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*HideTokenCancelTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _HideToken.contract.FilterLogs(opts, "CancelTransfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &HideTokenCancelTransferIterator{contract: _HideToken.contract, event: "CancelTransfer", logs: logs, sub: sub}, nil
}

// WatchCancelTransfer is a free log subscription operation binding the contract event 0x152fb15aa5d80f843e1e4bd5f2fc9161714f169945024decec7e84fb910fdd51.
//
// Solidity: event CancelTransfer(address indexed from, address indexed to, uint256 value)
func (_HideToken *HideTokenFilterer) WatchCancelTransfer(opts *bind.WatchOpts, sink chan<- *HideTokenCancelTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _HideToken.contract.WatchLogs(opts, "CancelTransfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HideTokenCancelTransfer)
				if err := _HideToken.contract.UnpackLog(event, "CancelTransfer", log); err != nil {
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

// ParseCancelTransfer is a log parse operation binding the contract event 0x152fb15aa5d80f843e1e4bd5f2fc9161714f169945024decec7e84fb910fdd51.
//
// Solidity: event CancelTransfer(address indexed from, address indexed to, uint256 value)
func (_HideToken *HideTokenFilterer) ParseCancelTransfer(log types.Log) (*HideTokenCancelTransfer, error) {
	event := new(HideTokenCancelTransfer)
	if err := _HideToken.contract.UnpackLog(event, "CancelTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HideTokenConfirmTransferIterator is returned from FilterConfirmTransfer and is used to iterate over the raw logs and unpacked data for ConfirmTransfer events raised by the HideToken contract.
type HideTokenConfirmTransferIterator struct {
	Event *HideTokenConfirmTransfer // Event containing the contract specifics and raw log

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
func (it *HideTokenConfirmTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HideTokenConfirmTransfer)
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
		it.Event = new(HideTokenConfirmTransfer)
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
func (it *HideTokenConfirmTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HideTokenConfirmTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HideTokenConfirmTransfer represents a ConfirmTransfer event raised by the HideToken contract.
type HideTokenConfirmTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterConfirmTransfer is a free log retrieval operation binding the contract event 0x7dca1b3ed21e108d0813c2b6e42f0cdd726c9abe60d2ad9ed0b44e5dc1f98c16.
//
// Solidity: event ConfirmTransfer(address indexed from, address indexed to, uint256 value)
func (_HideToken *HideTokenFilterer) FilterConfirmTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*HideTokenConfirmTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _HideToken.contract.FilterLogs(opts, "ConfirmTransfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &HideTokenConfirmTransferIterator{contract: _HideToken.contract, event: "ConfirmTransfer", logs: logs, sub: sub}, nil
}

// WatchConfirmTransfer is a free log subscription operation binding the contract event 0x7dca1b3ed21e108d0813c2b6e42f0cdd726c9abe60d2ad9ed0b44e5dc1f98c16.
//
// Solidity: event ConfirmTransfer(address indexed from, address indexed to, uint256 value)
func (_HideToken *HideTokenFilterer) WatchConfirmTransfer(opts *bind.WatchOpts, sink chan<- *HideTokenConfirmTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _HideToken.contract.WatchLogs(opts, "ConfirmTransfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HideTokenConfirmTransfer)
				if err := _HideToken.contract.UnpackLog(event, "ConfirmTransfer", log); err != nil {
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

// ParseConfirmTransfer is a log parse operation binding the contract event 0x7dca1b3ed21e108d0813c2b6e42f0cdd726c9abe60d2ad9ed0b44e5dc1f98c16.
//
// Solidity: event ConfirmTransfer(address indexed from, address indexed to, uint256 value)
func (_HideToken *HideTokenFilterer) ParseConfirmTransfer(log types.Log) (*HideTokenConfirmTransfer, error) {
	event := new(HideTokenConfirmTransfer)
	if err := _HideToken.contract.UnpackLog(event, "ConfirmTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HideTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the HideToken contract.
type HideTokenOwnershipTransferredIterator struct {
	Event *HideTokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *HideTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HideTokenOwnershipTransferred)
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
		it.Event = new(HideTokenOwnershipTransferred)
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
func (it *HideTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HideTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HideTokenOwnershipTransferred represents a OwnershipTransferred event raised by the HideToken contract.
type HideTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_HideToken *HideTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*HideTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _HideToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &HideTokenOwnershipTransferredIterator{contract: _HideToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_HideToken *HideTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *HideTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _HideToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HideTokenOwnershipTransferred)
				if err := _HideToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_HideToken *HideTokenFilterer) ParseOwnershipTransferred(log types.Log) (*HideTokenOwnershipTransferred, error) {
	event := new(HideTokenOwnershipTransferred)
	if err := _HideToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HideTokenRegisterWalletIterator is returned from FilterRegisterWallet and is used to iterate over the raw logs and unpacked data for RegisterWallet events raised by the HideToken contract.
type HideTokenRegisterWalletIterator struct {
	Event *HideTokenRegisterWallet // Event containing the contract specifics and raw log

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
func (it *HideTokenRegisterWalletIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HideTokenRegisterWallet)
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
		it.Event = new(HideTokenRegisterWallet)
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
func (it *HideTokenRegisterWalletIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HideTokenRegisterWalletIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HideTokenRegisterWallet represents a RegisterWallet event raised by the HideToken contract.
type HideTokenRegisterWallet struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRegisterWallet is a free log retrieval operation binding the contract event 0x3de0c4f89978c8c9ac219df0e8813f0b4a3df2fad979b3e3607cbe6c0ce6c764.
//
// Solidity: event RegisterWallet(address indexed sender)
func (_HideToken *HideTokenFilterer) FilterRegisterWallet(opts *bind.FilterOpts, sender []common.Address) (*HideTokenRegisterWalletIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _HideToken.contract.FilterLogs(opts, "RegisterWallet", senderRule)
	if err != nil {
		return nil, err
	}
	return &HideTokenRegisterWalletIterator{contract: _HideToken.contract, event: "RegisterWallet", logs: logs, sub: sub}, nil
}

// WatchRegisterWallet is a free log subscription operation binding the contract event 0x3de0c4f89978c8c9ac219df0e8813f0b4a3df2fad979b3e3607cbe6c0ce6c764.
//
// Solidity: event RegisterWallet(address indexed sender)
func (_HideToken *HideTokenFilterer) WatchRegisterWallet(opts *bind.WatchOpts, sink chan<- *HideTokenRegisterWallet, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _HideToken.contract.WatchLogs(opts, "RegisterWallet", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HideTokenRegisterWallet)
				if err := _HideToken.contract.UnpackLog(event, "RegisterWallet", log); err != nil {
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

// ParseRegisterWallet is a log parse operation binding the contract event 0x3de0c4f89978c8c9ac219df0e8813f0b4a3df2fad979b3e3607cbe6c0ce6c764.
//
// Solidity: event RegisterWallet(address indexed sender)
func (_HideToken *HideTokenFilterer) ParseRegisterWallet(log types.Log) (*HideTokenRegisterWallet, error) {
	event := new(HideTokenRegisterWallet)
	if err := _HideToken.contract.UnpackLog(event, "RegisterWallet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HideTokenRequestTransferIterator is returned from FilterRequestTransfer and is used to iterate over the raw logs and unpacked data for RequestTransfer events raised by the HideToken contract.
type HideTokenRequestTransferIterator struct {
	Event *HideTokenRequestTransfer // Event containing the contract specifics and raw log

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
func (it *HideTokenRequestTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HideTokenRequestTransfer)
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
		it.Event = new(HideTokenRequestTransfer)
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
func (it *HideTokenRequestTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HideTokenRequestTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HideTokenRequestTransfer represents a RequestTransfer event raised by the HideToken contract.
type HideTokenRequestTransfer struct {
	From       common.Address
	To         common.Address
	Value      *big.Int
	ValueBytes []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRequestTransfer is a free log retrieval operation binding the contract event 0xafaeb45426cfc6351df3bf2c13164ab06bc6abbbb28f2e2602117dcb63842deb.
//
// Solidity: event RequestTransfer(address indexed from, address indexed to, uint256 value, bytes valueBytes)
func (_HideToken *HideTokenFilterer) FilterRequestTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*HideTokenRequestTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _HideToken.contract.FilterLogs(opts, "RequestTransfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &HideTokenRequestTransferIterator{contract: _HideToken.contract, event: "RequestTransfer", logs: logs, sub: sub}, nil
}

// WatchRequestTransfer is a free log subscription operation binding the contract event 0xafaeb45426cfc6351df3bf2c13164ab06bc6abbbb28f2e2602117dcb63842deb.
//
// Solidity: event RequestTransfer(address indexed from, address indexed to, uint256 value, bytes valueBytes)
func (_HideToken *HideTokenFilterer) WatchRequestTransfer(opts *bind.WatchOpts, sink chan<- *HideTokenRequestTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _HideToken.contract.WatchLogs(opts, "RequestTransfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HideTokenRequestTransfer)
				if err := _HideToken.contract.UnpackLog(event, "RequestTransfer", log); err != nil {
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

// ParseRequestTransfer is a log parse operation binding the contract event 0xafaeb45426cfc6351df3bf2c13164ab06bc6abbbb28f2e2602117dcb63842deb.
//
// Solidity: event RequestTransfer(address indexed from, address indexed to, uint256 value, bytes valueBytes)
func (_HideToken *HideTokenFilterer) ParseRequestTransfer(log types.Log) (*HideTokenRequestTransfer, error) {
	event := new(HideTokenRequestTransfer)
	if err := _HideToken.contract.UnpackLog(event, "RequestTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HideTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the HideToken contract.
type HideTokenTransferIterator struct {
	Event *HideTokenTransfer // Event containing the contract specifics and raw log

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
func (it *HideTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HideTokenTransfer)
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
		it.Event = new(HideTokenTransfer)
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
func (it *HideTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HideTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HideTokenTransfer represents a Transfer event raised by the HideToken contract.
type HideTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_HideToken *HideTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*HideTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _HideToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &HideTokenTransferIterator{contract: _HideToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_HideToken *HideTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *HideTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _HideToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HideTokenTransfer)
				if err := _HideToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_HideToken *HideTokenFilterer) ParseTransfer(log types.Log) (*HideTokenTransfer, error) {
	event := new(HideTokenTransfer)
	if err := _HideToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
