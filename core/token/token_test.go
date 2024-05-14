package token

import (
	"context"
	"fmt"
	"github.com/UDPN/zk-token/core/key"
	"github.com/UDPN/zk-token/core/testdata"
	"github.com/UDPN/zk-token/core/zkproof"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"

	"math/big"
	"os"
	"testing"
)

var (
	NodeUrl = testdata.NodeUrl

	User1_HexKey  = testdata.User1_HexKey
	User1_address = testdata.User1_address

	User2_HexKey  = testdata.User2_HexKey
	User2_address = testdata.User2_address

	KeySecret = new(big.Int).SetInt64(123456789)

	Balance_wasm_path = "build/circuits/BalanceHash_js/BalanceHash.wasm"
	Balance_zkey_path = "build/circuits/BalanceHash.zkey"
	Balance_vkey_path = "build/circuits/BalanceHash_verification_key.json"

	Transfer_wasm_path = "build/circuits/Transfer_js/Transfer.wasm"
	Transfer_zkey_path = "build/circuits/Transfer.zkey"
	Transfer_vkey_path = "build/circuits/Transfer_verification_key.json"

	RootPath = "../../"

	TokenAddress = "0xe52e1445dd703d2Bb120423DE49E83D03cF0de8C"
)

func TestCheckAddress(t *testing.T) {
	cli, err := NewTokenCli(NodeUrl, User1_HexKey, common.HexToAddress(TokenAddress))
	require.NoError(t, err)

	fmt.Println(cli.Owner())

	keyBytes, err := hexutil.Decode(User1_HexKey)
	require.NoError(t, err)

	pk, err := crypto.ToECDSA(keyBytes)
	require.NoError(t, err)

	pubkeyBytes := crypto.FromECDSAPub(&pk.PublicKey)[1:]

	fmt.Println(hexutil.Encode(pubkeyBytes))
	add, err := cli.session.QueryPublicKey(cli.cli.OwnerAddress())
	require.NoError(t, err)
	fmt.Println(hexutil.Encode(add))
}

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

func TestSepoliaBalance(t *testing.T) {
	sepolia := NodeUrl

	cli, err := NewTokenCli(sepolia, User1_HexKey, common.HexToAddress(TokenAddress))
	require.NoError(t, err)

	balance, err := cli.cli.Client().BalanceAt(context.Background(), cli.cli.OwnerAddress(), nil)
	require.NoError(t, err)

	fmt.Println(balance)
}

func TestRegisterWallet(t *testing.T) {

	cli, err := NewTokenCli(NodeUrl, User1_HexKey, common.HexToAddress(TokenAddress))
	require.NoError(t, err)

	err = cli.SetBalanceProofConfig(GetBalanceProofConf())
	require.NoError(t, err)

	err = cli.SetTransferProofConfig(GetTransferProofConf())
	require.NoError(t, err)

	input, txid, err := cli.RegisterWallet(KeySecret)
	require.NoError(t, err)

	fmt.Println("Txid:", txid)
	fmt.Println("balance hash:", input.BalanceHash().String())

}

func TestReceiveTokenProof(t *testing.T) {
	cli, err := NewTokenCli(NodeUrl, User1_HexKey, common.HexToAddress(TokenAddress))
	require.NoError(t, err)

	err = cli.SetBalanceProofConfig(GetBalanceProofConf())
	require.NoError(t, err)

	err = cli.SetTransferProofConfig(GetTransferProofConf())
	require.NoError(t, err)

	input, proof, err := cli.ReceiveTokenProof(new(big.Int).SetInt64(100), KeySecret, new(big.Int).SetInt64(50), KeySecret)
	fmt.Println(proof.PubSignals)
	fmt.Println("New Balance :", input.NewBalance().String())
	fmt.Println("New Balance Hash:", input.NewBalanceHash().String())
	fmt.Println("Value Hash :", input.ValueHash().String())
}

func TestSendTokenProof(t *testing.T) {
	cli, err := NewTokenCli(NodeUrl, User1_HexKey, common.HexToAddress(TokenAddress))
	require.NoError(t, err)

	err = cli.SetBalanceProofConfig(GetBalanceProofConf())
	require.NoError(t, err)

	err = cli.SetTransferProofConfig(GetTransferProofConf())
	require.NoError(t, err)

	input, proof, err := cli.SendTokenProof(new(big.Int).SetInt64(100), KeySecret, new(big.Int).SetInt64(50), KeySecret)
	fmt.Println(proof.PubSignals)
	fmt.Println("New Balance :", input.NewBalance().String())
	fmt.Println("New Balance Hash:", input.NewBalanceHash().String())
	fmt.Println("Value Hash :", input.ValueHash().String())
}

func TestMint(t *testing.T) {

	cli, err := NewTokenCli(NodeUrl, User1_HexKey, common.HexToAddress(TokenAddress))
	require.NoError(t, err)

	err = cli.SetBalanceProofConfig(GetBalanceProofConf())
	require.NoError(t, err)

	err = cli.SetTransferProofConfig(GetTransferProofConf())
	require.NoError(t, err)

	balance := new(big.Int).SetInt64(0)
	value := new(big.Int).SetInt64(100)

	input, txid, err := cli.Mint(cli.cli.OwnerAddress(), balance, KeySecret, value, KeySecret)
	require.NoError(t, err)

	fmt.Println("New Balance :", input.NewBalance().String())
	fmt.Println("New Balance Hash:", input.NewBalanceHash().String())
	fmt.Println("Value Hash :", input.ValueHash().String())

	fmt.Println("Txid:", txid)

}

func TestCheckBalance(t *testing.T) {
	cli, err := NewTokenCli(NodeUrl, User1_HexKey, common.HexToAddress(TokenAddress))
	require.NoError(t, err)

	addr := common.HexToAddress(User1_address)
	ok, err := cli.CheckBalance(addr, new(big.Int).SetUint64(90), KeySecret)
	require.NoError(t, err)
	require.Equal(t, ok, true)
}

func TestTransferToken(t *testing.T) {
	fromKey := User1_HexKey
	fromBalance := new(big.Int).SetInt64(90)
	fromBalanceSecret := KeySecret

	toAddress := common.HexToAddress(User2_address)

	value := new(big.Int).SetInt64(10)
	valueSecret := RandBigInt()

	cli, err := NewTokenCli(NodeUrl, fromKey, common.HexToAddress(TokenAddress))
	require.NoError(t, err)

	err = cli.SetBalanceProofConfig(GetBalanceProofConf())
	require.NoError(t, err)

	err = cli.SetTransferProofConfig(GetTransferProofConf())
	require.NoError(t, err)

	input, txId, err := cli.TransferToken(toAddress, fromBalance, fromBalanceSecret, value, valueSecret)
	require.NoError(t, err)

	fmt.Println("FROM New Balance :", input.NewBalance().String())
	fmt.Println("FROM New Balance Hash:", input.NewBalanceHash().String())
	fmt.Println("FROM Value Hash :", input.ValueHash().String())

	fmt.Println("TxId:", txId)
}

func TestConfirmTransferToken(t *testing.T) {

	Balance := new(big.Int).SetInt64(0)
	BalanceSecret := KeySecret

	fromAddress := common.HexToAddress(User1_address)

	cli, err := NewTokenCli(NodeUrl, User2_HexKey, common.HexToAddress(TokenAddress))
	require.NoError(t, err)

	err = cli.SetBalanceProofConfig(GetBalanceProofConf())
	require.NoError(t, err)

	err = cli.SetTransferProofConfig(GetTransferProofConf())
	require.NoError(t, err)

	input, txId, err := cli.ConfirmTransferToken(fromAddress, Balance, BalanceSecret)

	require.NoError(t, err)

	fmt.Println(" New Balance :", input.NewBalance().String())
	fmt.Println(" New Balance Hash:", input.NewBalanceHash().String())
	fmt.Println(" Value Hash :", input.ValueHash().String())

	fmt.Println("TxId:", txId)
}

func TestCancelTransferToken(t *testing.T) {
	fromKey := User1_HexKey
	fromBalance := new(big.Int).SetInt64(80)
	fromBalanceSecret := KeySecret

	toAddress := common.HexToAddress(User2_address)

	cli, err := NewTokenCli(NodeUrl, fromKey, common.HexToAddress(TokenAddress))
	require.NoError(t, err)

	err = cli.SetBalanceProofConfig(GetBalanceProofConf())
	require.NoError(t, err)

	err = cli.SetTransferProofConfig(GetTransferProofConf())
	require.NoError(t, err)

	input, txId, err := cli.CancelTransferToken(toAddress, fromBalance, fromBalanceSecret)
	require.NoError(t, err)

	fmt.Println(" New Balance :", input.NewBalance().String())
	fmt.Println(" New Balance Hash:", input.NewBalanceHash().String())
	fmt.Println(" Value Hash :", input.ValueHash().String())

	fmt.Println("TxId:", txId)
}

func TestToValue(t *testing.T) {
	v := new(big.Int).SetInt64(1245677)
	vs := new(big.Int).SetInt64(789654123000)

	vb := toValueBytes(v, vs)
	fmt.Println(hexutil.Encode(vb))

	v1, vs1 := fromValueBytes(vb)
	fmt.Println(v1, vs1)
}

func TestToAESValue(t *testing.T) {
	v := new(big.Int).SetInt64(1245677)
	vs := new(big.Int).SetInt64(789654123000)

	vb := toValueBytes(v, vs)
	fmt.Println(hexutil.Encode(vb))
	k := crypto.Keccak256([]byte("abc123"))

	s, err := key.AESEncrypt(k, vb)
	require.NoError(t, err)
	fmt.Println(hexutil.Encode(s))

	vb1, err := key.AESDecrypt(k, s)
	require.NoError(t, err)
	fmt.Println(hexutil.Encode(vb1))
	v1, vs1 := fromValueBytes(vb1)
	fmt.Println(v1, vs1)
}
