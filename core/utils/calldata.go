package utils

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/iden3/go-rapidsnark/types"
	"math/big"
)

func StringToBigInt(s string) *big.Int {
	bigInt, _ := new(big.Int).SetString(s, 10)
	return bigInt
}

func SolidityCallData(proof *types.ZKProof) (pA [2]*big.Int, pB [2][2]*big.Int, pC [2]*big.Int, publicSignals []*big.Int) {
	if proof == nil {
		return
	}

	pA[0] = StringToBigInt(proof.Proof.A[0])
	pA[1] = StringToBigInt(proof.Proof.A[1])

	pB[0][0] = StringToBigInt(proof.Proof.B[0][1])
	pB[0][1] = StringToBigInt(proof.Proof.B[0][0])
	pB[1][0] = StringToBigInt(proof.Proof.B[1][1])
	pB[1][1] = StringToBigInt(proof.Proof.B[1][0])

	pC[0] = StringToBigInt(proof.Proof.C[0])
	pC[1] = StringToBigInt(proof.Proof.C[1])

	for _, signal := range proof.PubSignals {
		publicSignals = append(publicSignals, StringToBigInt(signal))
	}

	return
}

func StringToHexString(s string) string {
	bigInt, _ := new(big.Int).SetString(s, 10)
	b := bigInt.Bytes()

	var a [32]byte
	if len(b) > len(a) {
		b = b[len(b)-32:]
	}

	copy(a[32-len(b):], b)
	return hexutil.Encode(a[:])
}

func StringCllData(proof *types.ZKProof) ([]byte, error) {
	var s = struct {
		Pa            [2]string
		Pb            [2][2]string
		Pc            [2]string
		PublicSignals []string
	}{}

	s.Pa[0] = StringToHexString(proof.Proof.A[0])
	s.Pa[1] = StringToHexString(proof.Proof.A[1])

	s.Pb[0][0] = StringToHexString(proof.Proof.B[0][1])
	s.Pb[0][1] = StringToHexString(proof.Proof.B[0][0])
	s.Pb[1][0] = StringToHexString(proof.Proof.B[1][1])
	s.Pb[1][1] = StringToHexString(proof.Proof.B[1][0])

	s.Pc[0] = StringToHexString(proof.Proof.C[0])
	s.Pc[1] = StringToHexString(proof.Proof.C[1])

	for _, signal := range proof.PubSignals {
		s.PublicSignals = append(s.PublicSignals, StringToHexString(signal))
	}

	return json.Marshal(&s)
}
