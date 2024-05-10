package zkproof

import (
	"fmt"
	"github.com/UDPN/zk-token/core/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"math/big"
	"os"
	"testing"
)

func TestTransferInput(t *testing.T) {

	balance := new(big.Int).SetInt64(100)
	balanceSecret := new(big.Int).SetInt64(255)
	value := new(big.Int).SetInt64(20)
	valueSecret := new(big.Int).SetInt64(50696493)

	input, err := NewTransferInput(
		balance,
		balanceSecret,
		value,
		valueSecret,
		true)
	require.NoError(t, err)

	jb, _ := input.MarshalJSON()

	exp := `{
    "balance": "100",
    "balanceSecret":"255",
    "value": "20",
    "valueSecret":"50696493",
    "add":"1"
}`
	assert.Equal(t, exp, string(jb))
}

func TestTransferProof(t *testing.T) {

	wasmBytes, _ := os.ReadFile(RootPath + Transfer_wasm_path)
	zkeyBytes, _ := os.ReadFile(RootPath + Transfer_zkey_path)
	vkeyBytes, _ := os.ReadFile(RootPath + Transfer_vkey_path)

	conf := &ProofConfig{
		WasmBytes: wasmBytes,
		ZkeyBytes: zkeyBytes,
		VkeyBytes: vkeyBytes,
	}
	zk, err := NewZKProof(conf)
	require.NoError(t, err)

	balance := new(big.Int).SetInt64(100)
	balanceSecret := new(big.Int).SetInt64(255)
	value := new(big.Int).SetInt64(20)
	valueSecret := new(big.Int).SetInt64(50696493)

	input, err := NewTransferInput(
		balance,
		balanceSecret,
		value,
		valueSecret,
		true)
	require.NoError(t, err)

	inputBytes, _ := input.MarshalJSON()

	proof, err := zk.GenProof(inputBytes)
	require.NoError(t, err)

	fmt.Println("PubSignals", proof.PubSignals)
	assert.Equal(t, "10223455449150275677113714670804258163030003758703361642470232280418323128899", proof.PubSignals[0])
	assert.Equal(t, "2843265801781970540986260575391792481309526658608150407121968371442642513149", proof.PubSignals[1])
	assert.Equal(t, "8179988688509462396711174912065383564071088031391330809900547465136026704168", proof.PubSignals[2])
	assert.Equal(t, "1", proof.PubSignals[3])

	err = zk.VerifyProof(proof)
	require.NoError(t, err)

	j, _ := utils.StringCllData(proof)
	fmt.Println(string(j))
}
