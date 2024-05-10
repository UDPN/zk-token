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

// BalanceVerifierMetaData contains all meta data concerning the BalanceVerifier contract.
var BalanceVerifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"_pA\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"_pB\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"_pC\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[3]\",\"name\":\"_pubSignals\",\"type\":\"uint256[3]\"}],\"name\":\"verifyProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610625806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c806311479fea14610030575b600080fd5b61004361003e366004610589565b610057565b604051901515815260200160405180910390f35b600061051a565b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47811061008f576000805260206000f35b50565b600060405183815284602082015285604082015260408160608360076107d05a03fa9150816100c5576000805260206000f35b825160408201526020830151606082015260408360808360066107d05a03fa915050806100f6576000805260206000f35b5050505050565b7f27efad12b381ca7736f23c709ee19ff35f5faffce1edd7642d18ea66bcbd2f5b85527f01adec766cbedcd315a333a2a54534ade90b358b7a8b0a0fc96461410ab37bdf60208601526000608086018661019a87357f1b2898fdfaed6add52c82153f37d35a4b34fe40a63644db30365ad1c30fe088e7f0fbbaaa4ee0f8745d295f6cda10b8692f5b569e28fb72ee006aa5170ddc1786f84610092565b6101ea60208801357f257f51b3376ee85a55c4f7ea8c7cdd661ffac91209071c0270d94ebc8a3473737f2b9e7b72dc30d0642273b3a69d61a5db7f605cc849f1f1530f171a28a87ff8fe84610092565b61023a60408801357f0278c8606b37c2cf09895470b68765168cf6a43f7eb08d3430ae372d58fae4c07f14334415027ea7f0d48e6df88a408f91fe17da11ab57a603109101b45912b1b384610092565b50823581527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4760208401357f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4703066020820152833560408201526020840135606082015260408401356080820152606084013560a08201527f2d4d9aa7e302d9df41749d5507949d05dbea33fbb16c643b22f599a2be6df2e260c08201527f14bedd503c37ceb061d8ec60209fe345ce89830a19230301f076caff004d192660e08201527f0967032fcbf776d1afc985f88877f182d38480a653f2decaa9794cbc3bf3060c6101008201527f0e187847ad4c798374d0d6732bf501847dd68bc0e071241e0213bc7fc13db7ab6101208201527f304cfbd1e08a704a99f5e847d93f8c3caafddec46b7a0d379da69a4d112346a76101408201527f1739c1b1a457a8c7313123d24d2f9192f896b7c63eea05a9d57f06547ad0cec8610160820152600087015161018082015260206000018701516101a08201527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c26101c08201527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6101e08201527f090689d0585ff075ec9e99ad690c3395bc4b313370b38ef355acdadcd122975b6102008201527f12c85ea5db8c6deb4aab71808dcb408fe3d1e7690c43d37b4ce6cc0166fa7daa610220820152843561024082015260208501356102608201527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c26102808201527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6102a08201527f090689d0585ff075ec9e99ad690c3395bc4b313370b38ef355acdadcd122975b6102c08201527f12c85ea5db8c6deb4aab71808dcb408fe3d1e7690c43d37b4ce6cc0166fa7daa6102e08201526020816103008360086107d05a03fa9051169695505050505050565b6040516103808101604052610532600084013561005e565b61053f602084013561005e565b61054c604084013561005e565b610559606084013561005e565b610566818486888a6100fd565b90508060005260206000f35b806040810183101561058357600080fd5b92915050565b6000806000806101608086880312156105a157600080fd5b6105ab8787610572565b945060c08601878111156105be57600080fd5b6040870194506105ce8882610572565b9350508681870111156105e057600080fd5b5092959194509261010001915056fea2646970667358221220ce251e590af94a2261676b4cb665b7fdcdd18dc394c675116d051b4a775d867764736f6c63430008130033",
}

// BalanceVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use BalanceVerifierMetaData.ABI instead.
var BalanceVerifierABI = BalanceVerifierMetaData.ABI

// BalanceVerifierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BalanceVerifierMetaData.Bin instead.
var BalanceVerifierBin = BalanceVerifierMetaData.Bin

// DeployBalanceVerifier deploys a new Ethereum contract, binding an instance of BalanceVerifier to it.
func DeployBalanceVerifier(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BalanceVerifier, error) {
	parsed, err := BalanceVerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BalanceVerifierBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BalanceVerifier{BalanceVerifierCaller: BalanceVerifierCaller{contract: contract}, BalanceVerifierTransactor: BalanceVerifierTransactor{contract: contract}, BalanceVerifierFilterer: BalanceVerifierFilterer{contract: contract}}, nil
}

// BalanceVerifier is an auto generated Go binding around an Ethereum contract.
type BalanceVerifier struct {
	BalanceVerifierCaller     // Read-only binding to the contract
	BalanceVerifierTransactor // Write-only binding to the contract
	BalanceVerifierFilterer   // Log filterer for contract events
}

// BalanceVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type BalanceVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalanceVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BalanceVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalanceVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BalanceVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalanceVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BalanceVerifierSession struct {
	Contract     *BalanceVerifier  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BalanceVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BalanceVerifierCallerSession struct {
	Contract *BalanceVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// BalanceVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BalanceVerifierTransactorSession struct {
	Contract     *BalanceVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// BalanceVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type BalanceVerifierRaw struct {
	Contract *BalanceVerifier // Generic contract binding to access the raw methods on
}

// BalanceVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BalanceVerifierCallerRaw struct {
	Contract *BalanceVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// BalanceVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BalanceVerifierTransactorRaw struct {
	Contract *BalanceVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBalanceVerifier creates a new instance of BalanceVerifier, bound to a specific deployed contract.
func NewBalanceVerifier(address common.Address, backend bind.ContractBackend) (*BalanceVerifier, error) {
	contract, err := bindBalanceVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BalanceVerifier{BalanceVerifierCaller: BalanceVerifierCaller{contract: contract}, BalanceVerifierTransactor: BalanceVerifierTransactor{contract: contract}, BalanceVerifierFilterer: BalanceVerifierFilterer{contract: contract}}, nil
}

// NewBalanceVerifierCaller creates a new read-only instance of BalanceVerifier, bound to a specific deployed contract.
func NewBalanceVerifierCaller(address common.Address, caller bind.ContractCaller) (*BalanceVerifierCaller, error) {
	contract, err := bindBalanceVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BalanceVerifierCaller{contract: contract}, nil
}

// NewBalanceVerifierTransactor creates a new write-only instance of BalanceVerifier, bound to a specific deployed contract.
func NewBalanceVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*BalanceVerifierTransactor, error) {
	contract, err := bindBalanceVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BalanceVerifierTransactor{contract: contract}, nil
}

// NewBalanceVerifierFilterer creates a new log filterer instance of BalanceVerifier, bound to a specific deployed contract.
func NewBalanceVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*BalanceVerifierFilterer, error) {
	contract, err := bindBalanceVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BalanceVerifierFilterer{contract: contract}, nil
}

// bindBalanceVerifier binds a generic wrapper to an already deployed contract.
func bindBalanceVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BalanceVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BalanceVerifier *BalanceVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BalanceVerifier.Contract.BalanceVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BalanceVerifier *BalanceVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalanceVerifier.Contract.BalanceVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BalanceVerifier *BalanceVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BalanceVerifier.Contract.BalanceVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BalanceVerifier *BalanceVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BalanceVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BalanceVerifier *BalanceVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalanceVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BalanceVerifier *BalanceVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BalanceVerifier.Contract.contract.Transact(opts, method, params...)
}

// VerifyProof is a free data retrieval call binding the contract method 0x11479fea.
//
// Solidity: function verifyProof(uint256[2] _pA, uint256[2][2] _pB, uint256[2] _pC, uint256[3] _pubSignals) view returns(bool)
func (_BalanceVerifier *BalanceVerifierCaller) VerifyProof(opts *bind.CallOpts, _pA [2]*big.Int, _pB [2][2]*big.Int, _pC [2]*big.Int, _pubSignals [3]*big.Int) (bool, error) {
	var out []interface{}
	err := _BalanceVerifier.contract.Call(opts, &out, "verifyProof", _pA, _pB, _pC, _pubSignals)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyProof is a free data retrieval call binding the contract method 0x11479fea.
//
// Solidity: function verifyProof(uint256[2] _pA, uint256[2][2] _pB, uint256[2] _pC, uint256[3] _pubSignals) view returns(bool)
func (_BalanceVerifier *BalanceVerifierSession) VerifyProof(_pA [2]*big.Int, _pB [2][2]*big.Int, _pC [2]*big.Int, _pubSignals [3]*big.Int) (bool, error) {
	return _BalanceVerifier.Contract.VerifyProof(&_BalanceVerifier.CallOpts, _pA, _pB, _pC, _pubSignals)
}

// VerifyProof is a free data retrieval call binding the contract method 0x11479fea.
//
// Solidity: function verifyProof(uint256[2] _pA, uint256[2][2] _pB, uint256[2] _pC, uint256[3] _pubSignals) view returns(bool)
func (_BalanceVerifier *BalanceVerifierCallerSession) VerifyProof(_pA [2]*big.Int, _pB [2][2]*big.Int, _pC [2]*big.Int, _pubSignals [3]*big.Int) (bool, error) {
	return _BalanceVerifier.Contract.VerifyProof(&_BalanceVerifier.CallOpts, _pA, _pB, _pC, _pubSignals)
}
