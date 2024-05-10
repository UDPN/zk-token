package zkproof

import (
	"fmt"
	"github.com/iden3/go-iden3-crypto/poseidon"
	"math/big"
)

const (
	Transfer_wasm_path = "build/circuits/Transfer_js/Transfer.wasm"
	Transfer_zkey_path = "build/circuits/Transfer.zkey"
	Transfer_vkey_path = "build/circuits/Transfer_verification_key.json"
)

func NewTransferInput(balance, balanceSecret, value, valueSecret *big.Int, add bool) (*TransferInput, error) {

	input := &TransferInput{
		Balance:       balance,
		BalanceSecret: balanceSecret,
		Value:         value,
		ValueSecret:   valueSecret,
	}
	if add {
		input.Add = "1"
	} else {
		input.Add = "0"
	}

	err := input.buildHash()
	if err != nil {
		return nil, err
	}
	return input, nil
}

type TransferInput struct {
	Balance       *big.Int
	BalanceSecret *big.Int

	Value       *big.Int
	ValueSecret *big.Int

	Add string

	newBalance     *big.Int
	newBalanceHash *big.Int
	oldBalanceHash *big.Int
	valueHash      *big.Int
}

func (i *TransferInput) NewBalance() *big.Int {
	return i.newBalance
}

func (i *TransferInput) OldBalanceHash() *big.Int {
	return i.oldBalanceHash
}

func (i *TransferInput) NewBalanceHash() *big.Int {
	return i.newBalanceHash
}

func (i *TransferInput) ValueHash() *big.Int {
	return i.valueHash
}

func (i *TransferInput) buildHash() error {

	bh, err := poseidon.Hash([]*big.Int{i.Balance, i.BalanceSecret})
	if err != nil {
		return err
	}
	i.oldBalanceHash = bh

	vh, err := poseidon.Hash([]*big.Int{i.Value, i.ValueSecret})
	if err != nil {
		return err
	}
	i.valueHash = vh

	if i.Add == "1" {
		i.newBalance = new(big.Int).Add(i.Balance, i.Value)
	} else {
		i.newBalance = new(big.Int).Sub(i.Balance, i.Value)
	}
	nbh, err := poseidon.Hash([]*big.Int{i.newBalance, i.BalanceSecret})
	if err != nil {
		return err
	}
	i.newBalanceHash = nbh

	return nil
}

func (i *TransferInput) MarshalJSON() ([]byte, error) {
	inputFormat := `{
    "balance": "%s",
    "balanceSecret":"%s",
    "value": "%s",
    "valueSecret":"%s",
    "add":"%s"
}`
	//,
	//    "balanceHash":"%s",
	//    "valueHash":"%s"
	jb := fmt.Sprintf(inputFormat,
		i.Balance.String(),
		i.BalanceSecret.String(),
		i.Value.String(),
		i.ValueSecret.String(),
		i.Add)
	return []byte(jb), nil
}
