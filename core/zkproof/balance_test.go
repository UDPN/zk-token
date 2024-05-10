package zkproof

import (
	"fmt"
	"github.com/UDPN/zk-token/core/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"math/big"
	"os"
	"testing"
)

const RootPath = "../../"

func TestBalcnceInput(t *testing.T) {
	balance := new(big.Int).SetInt64(0)
	balanceSecret := new(big.Int).SetInt64(7896513512654)
	input, err := NewBalanceInput(balance, balanceSecret, common.HexToAddress("0x5B38Da6a701c568545dCfcB03FcB875f56beddC4"))
	require.NoError(t, err)

	jb, _ := input.MarshalJSON()
	fmt.Println(string(jb))
	assert.Equal(t, "{\"balance\":\"0\",\"balanceSecret\":\"7896513512654\",\"address\":\"520786028573371803640530888255888666801131675076\"}", string(jb))

}

func TestBalanceProof(t *testing.T) {

	wasmBytes, _ := os.ReadFile(RootPath + Balance_wasm_path)
	zkeyBytes, _ := os.ReadFile(RootPath + Balance_zkey_path)
	vkeyBytes, _ := os.ReadFile(RootPath + Balance_vkey_path)

	conf := &ProofConfig{
		WasmBytes: wasmBytes,
		ZkeyBytes: zkeyBytes,
		VkeyBytes: vkeyBytes,
	}
	zk, err := NewZKProof(conf)
	require.NoError(t, err)

	balance := new(big.Int).SetInt64(0)
	balanceSecret := new(big.Int).SetInt64(7896513512654)

	input, err := NewBalanceInput(balance, balanceSecret, common.HexToAddress("0xd67306D4C79AE8eC22579BDF6932517B8720affC"))
	require.NoError(t, err)

	inputBytes, _ := input.MarshalJSON()

	proof, err := zk.GenProof(inputBytes)
	require.NoError(t, err)

	fmt.Println("PubSignals", proof.PubSignals)
	assert.Equal(t, "21475938361495711009381873304728615467454200248911554954262344219601474942996", proof.PubSignals[0])

	fmt.Println(zk.VerifyProof(proof))
	j, _ := utils.StringCllData(proof)
	fmt.Println(string(j))
}
