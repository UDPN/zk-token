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

// TransferVerifierMetaData contains all meta data concerning the TransferVerifier contract.
var TransferVerifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"_pA\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"_pB\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"_pC\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[4]\",\"name\":\"_pubSignals\",\"type\":\"uint256[4]\"}],\"name\":\"verifyProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610687806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80635fe8c13b14610030575b600080fd5b61004361003e3660046105f7565b610057565b604051901515815260200160405180910390f35b600061056a565b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47811061008f576000805260206000f35b50565b600060405183815284602082015285604082015260408160608360076107d05a03fa9150816100c5576000805260206000f35b825160408201526020830151606082015260408360808360066107d05a03fa915050806100f6576000805260206000f35b5050505050565b7f2cfb0c3e444fe816f517167f32a04d56d8a7cf1c62c038c3ef61efc23006842485527f01799b3ba7ebc1407e6d5fbf23323b8326ab72337f395242a037c857bba0e9e560208601526000608086018661019a87357f09ec0b554c5595d2c9cba5848e0d8d0a2f5115880072ea5cb313fb55551b462f7f2c52bcf68f830fafb0ec0c7533ead312dec43c17a663c0c2fc1e8d2f8f4d6a1684610092565b6101ea60208801357f050e73bf24a4a72b020cc6329579f2cc575400c813e7f85705a3f5b1050df2a47f19955b1f1e1bc0aa5c66334f4075e3a69db69070b9c120ab2a5f4816a54d4a2884610092565b61023a60408801357f16d9edf8cc1e417f82af2adf81eb290e613b7b6eb8493450a19988719b1aa2757f17bc5abc9c2e6084d5991a7369afb64c911b83334b35d0b932c4a079100bd29284610092565b61028a60608801357f08aabbf7883d29801d10c77b5b37fc71b98e9b5190ac9510738e02f7b01126187f199ef126856608cf182e2d887def2e6ca5cc2e0062ce2172381747adb74b1e6d84610092565b50823581527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4760208401357f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4703066020820152833560408201526020840135606082015260408401356080820152606084013560a08201527f2d4d9aa7e302d9df41749d5507949d05dbea33fbb16c643b22f599a2be6df2e260c08201527f14bedd503c37ceb061d8ec60209fe345ce89830a19230301f076caff004d192660e08201527f0967032fcbf776d1afc985f88877f182d38480a653f2decaa9794cbc3bf3060c6101008201527f0e187847ad4c798374d0d6732bf501847dd68bc0e071241e0213bc7fc13db7ab6101208201527f304cfbd1e08a704a99f5e847d93f8c3caafddec46b7a0d379da69a4d112346a76101408201527f1739c1b1a457a8c7313123d24d2f9192f896b7c63eea05a9d57f06547ad0cec8610160820152600087015161018082015260206000018701516101a08201527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c26101c08201527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6101e08201527f090689d0585ff075ec9e99ad690c3395bc4b313370b38ef355acdadcd122975b6102008201527f12c85ea5db8c6deb4aab71808dcb408fe3d1e7690c43d37b4ce6cc0166fa7daa610220820152843561024082015260208501356102608201527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c26102808201527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6102a08201527f090689d0585ff075ec9e99ad690c3395bc4b313370b38ef355acdadcd122975b6102c08201527f12c85ea5db8c6deb4aab71808dcb408fe3d1e7690c43d37b4ce6cc0166fa7daa6102e08201526020816103008360086107d05a03fa9051169695505050505050565b6040516103808101604052610582600084013561005e565b61058f602084013561005e565b61059c604084013561005e565b6105a9606084013561005e565b6105b6608084013561005e565b6105c3818486888a6100fd565b90508060005260206000f35b80604081018310156105e057600080fd5b92915050565b80608081018310156105e057600080fd5b600080600080610180858703121561060e57600080fd5b61061886866105cf565b935061062786604087016105e6565b92506106368660c087016105cf565b91506106468661010087016105e6565b90509295919450925056fea2646970667358221220402439e7fa7d67a5db0e88a65bf8406e7ec422dcb9d9fccc4d3ea432a0554f7864736f6c63430008130033",
}

// TransferVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use TransferVerifierMetaData.ABI instead.
var TransferVerifierABI = TransferVerifierMetaData.ABI

// TransferVerifierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TransferVerifierMetaData.Bin instead.
var TransferVerifierBin = TransferVerifierMetaData.Bin

// DeployTransferVerifier deploys a new Ethereum contract, binding an instance of TransferVerifier to it.
func DeployTransferVerifier(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TransferVerifier, error) {
	parsed, err := TransferVerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TransferVerifierBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TransferVerifier{TransferVerifierCaller: TransferVerifierCaller{contract: contract}, TransferVerifierTransactor: TransferVerifierTransactor{contract: contract}, TransferVerifierFilterer: TransferVerifierFilterer{contract: contract}}, nil
}

// TransferVerifier is an auto generated Go binding around an Ethereum contract.
type TransferVerifier struct {
	TransferVerifierCaller     // Read-only binding to the contract
	TransferVerifierTransactor // Write-only binding to the contract
	TransferVerifierFilterer   // Log filterer for contract events
}

// TransferVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type TransferVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TransferVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TransferVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TransferVerifierSession struct {
	Contract     *TransferVerifier // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TransferVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TransferVerifierCallerSession struct {
	Contract *TransferVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// TransferVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TransferVerifierTransactorSession struct {
	Contract     *TransferVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// TransferVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type TransferVerifierRaw struct {
	Contract *TransferVerifier // Generic contract binding to access the raw methods on
}

// TransferVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TransferVerifierCallerRaw struct {
	Contract *TransferVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// TransferVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TransferVerifierTransactorRaw struct {
	Contract *TransferVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTransferVerifier creates a new instance of TransferVerifier, bound to a specific deployed contract.
func NewTransferVerifier(address common.Address, backend bind.ContractBackend) (*TransferVerifier, error) {
	contract, err := bindTransferVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TransferVerifier{TransferVerifierCaller: TransferVerifierCaller{contract: contract}, TransferVerifierTransactor: TransferVerifierTransactor{contract: contract}, TransferVerifierFilterer: TransferVerifierFilterer{contract: contract}}, nil
}

// NewTransferVerifierCaller creates a new read-only instance of TransferVerifier, bound to a specific deployed contract.
func NewTransferVerifierCaller(address common.Address, caller bind.ContractCaller) (*TransferVerifierCaller, error) {
	contract, err := bindTransferVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TransferVerifierCaller{contract: contract}, nil
}

// NewTransferVerifierTransactor creates a new write-only instance of TransferVerifier, bound to a specific deployed contract.
func NewTransferVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*TransferVerifierTransactor, error) {
	contract, err := bindTransferVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TransferVerifierTransactor{contract: contract}, nil
}

// NewTransferVerifierFilterer creates a new log filterer instance of TransferVerifier, bound to a specific deployed contract.
func NewTransferVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*TransferVerifierFilterer, error) {
	contract, err := bindTransferVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TransferVerifierFilterer{contract: contract}, nil
}

// bindTransferVerifier binds a generic wrapper to an already deployed contract.
func bindTransferVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TransferVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TransferVerifier *TransferVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TransferVerifier.Contract.TransferVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TransferVerifier *TransferVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransferVerifier.Contract.TransferVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TransferVerifier *TransferVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TransferVerifier.Contract.TransferVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TransferVerifier *TransferVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TransferVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TransferVerifier *TransferVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransferVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TransferVerifier *TransferVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TransferVerifier.Contract.contract.Transact(opts, method, params...)
}

// VerifyProof is a free data retrieval call binding the contract method 0x5fe8c13b.
//
// Solidity: function verifyProof(uint256[2] _pA, uint256[2][2] _pB, uint256[2] _pC, uint256[4] _pubSignals) view returns(bool)
func (_TransferVerifier *TransferVerifierCaller) VerifyProof(opts *bind.CallOpts, _pA [2]*big.Int, _pB [2][2]*big.Int, _pC [2]*big.Int, _pubSignals [4]*big.Int) (bool, error) {
	var out []interface{}
	err := _TransferVerifier.contract.Call(opts, &out, "verifyProof", _pA, _pB, _pC, _pubSignals)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyProof is a free data retrieval call binding the contract method 0x5fe8c13b.
//
// Solidity: function verifyProof(uint256[2] _pA, uint256[2][2] _pB, uint256[2] _pC, uint256[4] _pubSignals) view returns(bool)
func (_TransferVerifier *TransferVerifierSession) VerifyProof(_pA [2]*big.Int, _pB [2][2]*big.Int, _pC [2]*big.Int, _pubSignals [4]*big.Int) (bool, error) {
	return _TransferVerifier.Contract.VerifyProof(&_TransferVerifier.CallOpts, _pA, _pB, _pC, _pubSignals)
}

// VerifyProof is a free data retrieval call binding the contract method 0x5fe8c13b.
//
// Solidity: function verifyProof(uint256[2] _pA, uint256[2][2] _pB, uint256[2] _pC, uint256[4] _pubSignals) view returns(bool)
func (_TransferVerifier *TransferVerifierCallerSession) VerifyProof(_pA [2]*big.Int, _pB [2][2]*big.Int, _pC [2]*big.Int, _pubSignals [4]*big.Int) (bool, error) {
	return _TransferVerifier.Contract.VerifyProof(&_TransferVerifier.CallOpts, _pA, _pB, _pC, _pubSignals)
}
