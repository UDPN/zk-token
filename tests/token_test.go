package tests

import (
	"github.com/UDPN/zk-token/core/testdata"
	"github.com/UDPN/zk-token/core/token"
	"github.com/UDPN/zk-token/core/zkproof"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
	"math/big"
	"os"
	"testing"
)

const (
	Balance_wasm_path = "build/circuits/BalanceHash_js/BalanceHash.wasm"
	Balance_zkey_path = "build/circuits/BalanceHash.zkey"
	Balance_vkey_path = "build/circuits/BalanceHash_verification_key.json"

	Transfer_wasm_path = "build/circuits/Transfer_js/Transfer.wasm"
	Transfer_zkey_path = "build/circuits/Transfer.zkey"
	Transfer_vkey_path = "build/circuits/Transfer_verification_key.json"

	RootPath = "../"
)

func GetBalanceProofConf() *zkproof.ProofConfig {
	wasmBytes, _ := os.ReadFile(RootPath + Balance_wasm_path)
	zkeyBytes, _ := os.ReadFile(RootPath + Balance_zkey_path)
	vkeyBytes, _ := os.ReadFile(RootPath + Balance_vkey_path)

	conf := &zkproof.ProofConfig{
		WasmBytes: wasmBytes,
		ZkeyBytes: zkeyBytes,
		VkeyBytes: vkeyBytes,
	}
	return conf
}

func GetTransferProofConf() *zkproof.ProofConfig {

	wasmBytes, _ := os.ReadFile(RootPath + Transfer_wasm_path)
	zkeyBytes, _ := os.ReadFile(RootPath + Transfer_zkey_path)
	vkeyBytes, _ := os.ReadFile(RootPath + Transfer_vkey_path)

	conf := &zkproof.ProofConfig{
		WasmBytes: wasmBytes,
		ZkeyBytes: zkeyBytes,
		VkeyBytes: vkeyBytes,
	}
	return conf
}

type UserInfo struct {
	key           string
	address       common.Address
	balance       *big.Int
	balanceSecret *big.Int
}

func TestToken(t *testing.T) {

	NodeUrl := testdata.NodeUrl
	TokenAddress := "0xe52e1445dd703d2Bb120423DE49E83D03cF0de8C"

	fromUser := &UserInfo{
		key:           testdata.User1_HexKey,
		address:       common.HexToAddress(testdata.User1_address),
		balance:       new(big.Int).SetUint64(0),
		balanceSecret: new(big.Int).SetUint64(123456789),
	}

	toUser := &UserInfo{
		key:           testdata.User2_HexKey,
		address:       common.HexToAddress(testdata.User2_address),
		balance:       new(big.Int).SetUint64(0),
		balanceSecret: new(big.Int).SetUint64(123456789123456789),
	}

	mintNumber := new(big.Int).SetInt64(100)
	value := new(big.Int).SetInt64(20)
	valueSecret := new(big.Int).SetInt64(987654321)

	fromCli := newCli(t, NodeUrl, fromUser.key, TokenAddress)
	toCli := newCli(t, NodeUrl, toUser.key, TokenAddress)

	registerWallet(t, fromCli, fromUser)
	registerWallet(t, toCli, toUser)

	mintToken(t, fromCli, fromUser, mintNumber)

	input, txId, err := fromCli.TransferToken(toUser.address, fromUser.balance, fromUser.balanceSecret, value, valueSecret)
	require.NoError(t, err)
	fromCli.WaitTx(txId)

	fromUser.balance = input.NewBalance()
	checkBalance(t, fromCli, fromUser)

	input, txId, err = toCli.ConfirmTransferToken(fromUser.address, toUser.balance, toUser.balanceSecret)
	require.NoError(t, err)
	fromCli.WaitTx(txId)

	toUser.balance = input.NewBalance()
	checkBalance(t, toCli, toUser)

	input, txId, err = fromCli.TransferToken(toUser.address, fromUser.balance, fromUser.balanceSecret, value, valueSecret)
	require.NoError(t, err)
	fromCli.WaitTx(txId)

	fromUser.balance = input.NewBalance()
	checkBalance(t, fromCli, fromUser)

	input, txId, err = fromCli.CancelTransferToken(toUser.address, fromUser.balance, fromUser.balanceSecret)
	require.NoError(t, err)
	fromCli.WaitTx(txId)

	fromUser.balance = input.NewBalance()
	checkBalance(t, toCli, toUser)

	require.Equal(t, fromUser.balance.String(), "80")
	require.Equal(t, toUser.balance.String(), "20")

}

func checkBalance(t *testing.T, cli *token.TokenClient, user *UserInfo) {
	res, err := cli.CheckBalance(user.address, user.balance, user.balanceSecret)
	require.NoError(t, err)
	require.Equal(t, res, true)
}

func registerWallet(t *testing.T, cli *token.TokenClient, user *UserInfo) {
	input, txId, err := cli.RegisterWallet(user.balanceSecret)
	require.NoError(t, err)
	cli.WaitTx(txId)
	user.balance = input.Balance
	checkBalance(t, cli, user)

}

func mintToken(t *testing.T, cli *token.TokenClient, user *UserInfo, balance *big.Int) {
	input, txId, err := cli.Mint(user.address, user.balance, user.balanceSecret, balance, user.balanceSecret)
	require.NoError(t, err)
	cli.WaitTx(txId)
	user.balance = input.NewBalance()
	checkBalance(t, cli, user)
}

func newCli(t *testing.T, nodeUrl, key, tokenAddress string) *token.TokenClient {
	cli, err := token.NewTokenCli(nodeUrl, key, common.HexToAddress(tokenAddress))
	require.NoError(t, err)

	err = cli.SetBalanceProofConfig(GetBalanceProofConf())
	require.NoError(t, err)

	err = cli.SetTransferProofConfig(GetTransferProofConf())
	require.NoError(t, err)
	return cli
}
