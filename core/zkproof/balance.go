package zkproof

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/iden3/go-iden3-crypto/poseidon"
	"math/big"
)

const (
	Balance_wasm_path = "build/circuits/BalanceHash_js/BalanceHash.wasm"
	Balance_zkey_path = "build/circuits/BalanceHash.zkey"
	Balance_vkey_path = "build/circuits/BalanceHash_verification_key.json"
)

func NewBalanceInput(balance, br *big.Int, address common.Address) (*BalanceInput, error) {

	return &BalanceInput{Balance: balance, BalanceSecret: br, Address: address}, nil
}

type BalanceInput struct {
	Balance       *big.Int       `json:"balance"`
	Address       common.Address `json:"address"`
	BalanceSecret *big.Int       `json:"balanceSecret"`
}

func (b *BalanceInput) BalanceHash() *big.Int {
	nbh, _ := poseidon.Hash([]*big.Int{b.Balance, b.BalanceSecret})
	return nbh
}

func (i *BalanceInput) MarshalJSON() ([]byte, error) {
	addrInt := new(big.Int).SetBytes(i.Address.Bytes())
	jb := fmt.Sprintf("{\"balance\":\"%s\",\"balanceSecret\":\"%s\",\"address\":\"%s\"}", i.Balance.String(), i.BalanceSecret.String(), addrInt.String())
	return []byte(jb), nil
}
