package zkproof

import (
	"github.com/iden3/go-rapidsnark/prover"
	"github.com/iden3/go-rapidsnark/types"
	"github.com/iden3/go-rapidsnark/verifier"
	"github.com/iden3/go-rapidsnark/witness/v2"
	"github.com/iden3/go-rapidsnark/witness/wazero"
)

func NewZKProof(conf *ProofConfig) (*ZKProof, error) {
	proof := &ZKProof{
		conf: conf,
	}
	if err := proof.initCalc(); err != nil {
		return nil, err
	}
	return proof, nil
}

type ProofConfig struct {
	WasmBytes []byte
	ZkeyBytes []byte
	VkeyBytes []byte
}

type ZKProof struct {
	conf *ProofConfig
	calc witness.Calculator
}

func (zk *ZKProof) initCalc() error {
	var ops []witness.Option
	//ops = append(ops, witness.WithWasmEngine(wasmer.NewCircom2WitnessCalculator))
	ops = append(ops, witness.WithWasmEngine(wazero.NewCircom2WZWitnessCalculator))
	calc, err := witness.NewCalculator(zk.conf.WasmBytes, ops...)
	if err != nil {
		return err
	}
	zk.calc = calc
	return nil
}

func (zk *ZKProof) GenWitness(inputBytes []byte) ([]byte, error) {

	inputs, err := witness.ParseInputs(inputBytes)
	if err != nil {
		return nil, err
	}

	//wtns, err2 := calc.CalculateWitness(inputs, true)
	//wtns, err2 := calc.CalculateBinWitness(inputs, true)

	return zk.calc.CalculateWTNSBin(inputs, true)
}

func (zk *ZKProof) GenProof(inputBytes []byte) (*types.ZKProof, error) {

	witnessBytes, err := zk.GenWitness(inputBytes)
	if err != nil {
		return nil, err
	}

	proof, err := prover.Groth16Prover(zk.conf.ZkeyBytes, witnessBytes)
	if err != nil {
		return nil, err
	}

	return proof, nil
}

func (zk *ZKProof) VerifyProof(proof *types.ZKProof) error {
	return verifier.VerifyGroth16(*proof, zk.conf.VkeyBytes)
}
