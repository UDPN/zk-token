package token

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"github.com/UDPN/zk-token/core/client"
	"github.com/UDPN/zk-token/core/contracts"
	"github.com/UDPN/zk-token/core/key"
	"github.com/UDPN/zk-token/core/utils"
	"github.com/UDPN/zk-token/core/zkproof"
	"github.com/pkg/errors"
	"io"

	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/iden3/go-iden3-crypto/poseidon"
	"github.com/iden3/go-rapidsnark/types"

	"math/big"
	"sync"
	"time"
)

const defTimeout = 50 * time.Second

func NewTokenCli(url, hexKey string, addr common.Address) (*TokenClient, error) {
	cli, err := client.NewETHClientByUrlAndHexKey(url, hexKey)
	if err != nil {
		return nil, err
	}

	s, err := newSession(cli, addr)
	if err != nil {
		return nil, err
	}

	return &TokenClient{
		cli:          cli,
		session:      s,
		tokenAddress: addr,
		timeout:      defTimeout,
	}, nil
}

func newSession(cli *client.ETHClient, addr common.Address) (*contracts.HideTokenSession, error) {
	cd, err := contracts.NewHideToken(addr, cli.Client())
	if err != nil {
		return nil, err
	}
	transactOpts := cli.GetTrandOpts()
	callOpts := bind.CallOpts{From: transactOpts.From}
	s := &contracts.HideTokenSession{Contract: cd, CallOpts: callOpts, TransactOpts: *transactOpts}
	return s, nil
}

type TokenClient struct {
	cli          *client.ETHClient
	session      *contracts.HideTokenSession
	tokenAddress common.Address

	balanceProof  *zkproof.ZKProof
	transferProof *zkproof.ZKProof

	timeout time.Duration

	abi *abi.ABI

	stop        chan struct{}
	balanceLock sync.Mutex
}

func (t *TokenClient) WaitTx(txId string) error {
	tick := time.NewTicker(t.timeout)
	for {
		select {
		case <-tick.C:
			return errors.New("tx pending , wait time out")
		default:

			time.Sleep(time.Second)
			_, p, err := t.cli.Client().TransactionByHash(context.Background(), common.HexToHash(txId))
			if err != nil {
				return err
			}
			if !p {
				return nil
			}

		}
	}
}

func (t *TokenClient) GetSession() *contracts.HideTokenSession {
	return t.session
}

func (t *TokenClient) SetBalanceProofConfig(conf *zkproof.ProofConfig) error {
	proof, err := zkproof.NewZKProof(conf)
	if err != nil {
		return err
	}
	t.balanceProof = proof
	return nil
}

func (t *TokenClient) SetTransferProofConfig(conf *zkproof.ProofConfig) error {
	proof, err := zkproof.NewZKProof(conf)
	if err != nil {
		return err
	}
	t.transferProof = proof
	return nil
}

func (t *TokenClient) Owner() (common.Address, error) {
	return t.session.Owner()
}

func (t *TokenClient) RegisterWallet(secret *big.Int) (*zkproof.BalanceInput, string, error) {
	if t.balanceProof == nil {
		return nil, "", errors.New("must set balance zk proof config")
	}
	input, err := zkproof.NewBalanceInput(new(big.Int).SetInt64(0), secret, t.cli.OwnerAddress())
	if err != nil {
		return nil, "", err
	}
	inputBytes, _ := input.MarshalJSON()

	proof, err := t.balanceProof.GenProof(inputBytes)
	if err != nil {
		return nil, "", err
	}
	bz, _ := json.Marshal(proof)
	fmt.Println("proof：", string(bz))
	_proof, _ps := BalanceSolidityData(proof)

	pubkeyBytes := crypto.FromECDSAPub(t.cli.PublicKey())[1:]
	tx, err := t.session.RegisterWallet(_proof, _ps, pubkeyBytes)
	if err != nil {
		return nil, "", err
	}
	return input, tx.Hash().String(), nil
}

func (t *TokenClient) Balance() (*big.Int, error) {
	fmt.Println("Address:", t.cli.OwnerAddress())
	return t.session.BalanceOf(t.cli.OwnerAddress())
}

func (t *TokenClient) BalanceOfAddress(addr common.Address) (*big.Int, error) {
	return t.session.BalanceOf(addr)
}

func (t *TokenClient) MintToken(addr common.Address, proof *types.ZKProof, value, valueSecret string) (string, error) {

	valueBig, _ := new(big.Int).SetString(value, 10)
	valueSecretBig, _ := new(big.Int).SetString(valueSecret, 10)

	vh, err := poseidon.Hash([]*big.Int{valueBig, valueSecretBig})
	if vh.String() != proof.PubSignals[1] {

	}

	bz, _ := json.Marshal(proof)
	fmt.Println("proof：", string(bz))
	_proof, _ps := TransferSolidityData(proof)
	tx, err := t.session.Mint(addr, _proof, _ps)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

func (t *TokenClient) CheckBalance(addr common.Address, balance, balanceSecret *big.Int) (bool, error) {
	balanceHash, err := t.session.BalanceOf(addr)
	if err != nil {
		return false, err
	}

	bh, err := BalanceHash(balance, balanceSecret)
	if err != nil {
		return false, err
	}

	return balanceHash.String() == bh.String(), nil
}

func (t *TokenClient) Mint(addr common.Address, balance, balanceSecret, value, valueSecret *big.Int) (*zkproof.TransferInput, string, error) {

	input, proof, err := t.ReceiveTokenProof(balance, balanceSecret, value, valueSecret)
	if err != nil {
		return nil, "", err
	}
	bz, _ := json.Marshal(proof)
	fmt.Println("proof：", string(bz))
	_proof, _ps := TransferSolidityData(proof)
	tx, err := t.session.Mint(addr, _proof, _ps)
	if err != nil {
		return input, "", err
	}
	return input, tx.Hash().String(), nil
}

func (t *TokenClient) TransferToken(
	toAddr common.Address,
	balance, balanceSecret *big.Int,
	value, valueSecret *big.Int) (*zkproof.TransferInput, string, error) {

	input, proof, err := t.SendTokenProof(balance, balanceSecret, value, valueSecret)
	if err != nil {
		return nil, "", err
	}
	bz, _ := json.Marshal(proof)
	fmt.Println("send proof：", string(bz))

	_fromProof, _fromPs := TransferSolidityData(proof)

	valueBytes, err := t.encryptValue(toAddr, input.Value, input.ValueSecret)
	if err != nil {
		return nil, "", err
	}
	tx, err := t.session.Transfer(toAddr, valueBytes, _fromProof, _fromPs)
	if err != nil {
		return nil, "", err
	}
	return input, tx.Hash().String(), nil
}

func (t *TokenClient) ConfirmTransferToken(from common.Address, balance, balanceSecret *big.Int) (*zkproof.TransferInput, string, error) {
	valueBytes, valueHash, _, err := t.session.QueryTransfer(from, t.cli.OwnerAddress())
	if err != nil {
		return nil, "", errors.WithMessage(err, "query transfer has err, may be have expired ")
	}

	value, valueSecret, err := t.decryptValue(from, valueBytes)
	if err != nil {
		return nil, "", err
	}
	input, proof, err := t.ReceiveTokenProof(balance, balanceSecret, value, valueSecret)
	if err != nil {
		return nil, "", err
	}
	balanceHash, err := t.session.BalanceOf(t.cli.OwnerAddress())
	if err != nil {
		return nil, "", err
	}
	if balanceHash.String() != input.OldBalanceHash().String() {
		return nil, "", errors.New("Inconsistent old balance hash ")
	}

	if valueHash.String() != input.ValueHash().String() {
		return nil, "", errors.New("Inconsistent value hash ")
	}

	_toProof, _toPs := TransferSolidityData(proof)
	tx, err := t.session.ConfirmTransfer(from, _toProof, _toPs)
	if err != nil {
		return nil, "", err
	}
	return input, tx.Hash().String(), nil
}

func (t *TokenClient) CancelTransferToken(to common.Address, balance, balanceSecret *big.Int) (*zkproof.TransferInput, string, error) {
	valueBytes, valueHash, _, err := t.session.QueryTransfer(t.cli.OwnerAddress(), to)
	if err != nil {
		return nil, "", errors.WithMessage(err, "query transfer has err, may be have expired ")
	}

	value, valueSecret, err := t.decryptValue(to, valueBytes)
	if err != nil {
		return nil, "", err
	}
	input, proof, err := t.ReceiveTokenProof(balance, balanceSecret, value, valueSecret)
	if err != nil {
		return nil, "", err
	}
	balanceHash, err := t.session.BalanceOf(t.cli.OwnerAddress())
	if err != nil {
		return nil, "", err
	}
	if balanceHash.String() != input.OldBalanceHash().String() {
		return nil, "", errors.New("Inconsistent old balance hash ")
	}

	if valueHash.String() != input.ValueHash().String() {
		return nil, "", errors.New("Inconsistent value hash ")
	}

	_proof, _ps := TransferSolidityData(proof)
	tx, err := t.session.CancelTransfer(to, _proof, _ps)
	if err != nil {
		return nil, "", err
	}
	return input, tx.Hash().String(), nil
}

func (t *TokenClient) decryptValue(from common.Address, valueBytes []byte) (value *big.Int, valueSecret *big.Int, err error) {
	dh, err := t.ECDHKey(from)
	if err != nil {
		return
	}
	d, err := key.AESDecrypt(dh, valueBytes)
	if err != nil {
		return
	}

	if len(d) != 64 {
		return nil, nil, errors.New("valueBytes length must be 64 ")
	}
	value, valueSecret = fromValueBytes(d)
	return
}

func (t *TokenClient) encryptValue(to common.Address, value *big.Int, valueSecret *big.Int) (valueBytes []byte, err error) {
	dh, err := t.ECDHKey(to)
	if err != nil {
		return
	}
	d := toValueBytes(value, valueSecret)
	valueBytes, err = key.AESEncrypt(dh, d)
	return
}

func (t *TokenClient) ECDHKey(address common.Address) ([]byte, error) {
	publicKeyBytes, err := t.session.QueryPublicKey(address)
	if err != nil {
		return nil, err
	}
	var s256PubKeyBytes [65]byte
	s256PubKeyBytes[0] = 4
	copy(s256PubKeyBytes[1:], publicKeyBytes)
	publicKey, err := crypto.UnmarshalPubkey(s256PubKeyBytes[:])

	return key.EcdhKeyAgreement(publicKey, t.cli.PrivateKey())
}

func TransferSolidityData(proof *types.ZKProof) (contracts.HideTokenZKProof, [4]*big.Int) {
	pA, pB, pC, ps := utils.SolidityCallData(proof)
	var _ps [4]*big.Int
	copy(_ps[0:], ps)
	_proof := contracts.HideTokenZKProof{
		PA: pA,
		PB: pB,
		PC: pC,
	}

	return _proof, _ps

}

func BalanceSolidityData(proof *types.ZKProof) (contracts.HideTokenZKProof, [3]*big.Int) {
	pA, pB, pC, ps := utils.SolidityCallData(proof)
	var _ps [3]*big.Int
	copy(_ps[0:], ps)
	_proof := contracts.HideTokenZKProof{
		PA: pA,
		PB: pB,
		PC: pC,
	}

	return _proof, _ps

}

func (t *TokenClient) ReceiveTokenProof(balance, balanceSecret, value, valueSecret *big.Int) (*zkproof.TransferInput, *types.ZKProof, error) {
	input, err := zkproof.NewTransferInput(
		balance,
		balanceSecret,
		value,
		valueSecret,
		true)
	if err != nil {
		return input, nil, err
	}

	inputBytes, _ := input.MarshalJSON()

	proof, err := t.transferProof.GenProof(inputBytes)
	if err != nil {
		return input, nil, err
	}

	return input, proof, nil
}

func (t *TokenClient) SendTokenProof(balance, balanceSecret, value, valueSecret *big.Int) (*zkproof.TransferInput, *types.ZKProof, error) {
	input, err := zkproof.NewTransferInput(
		balance,
		balanceSecret,
		value,
		valueSecret,
		false)
	if err != nil {
		return input, nil, err
	}

	inputBytes, _ := input.MarshalJSON()

	proof, err := t.transferProof.GenProof(inputBytes)
	if err != nil {
		return input, nil, err
	}

	return input, proof, nil
}

func BalanceHash(balance, secret *big.Int) (*big.Int, error) {

	h, err := poseidon.Hash([]*big.Int{balance, secret})
	if err != nil {
		return nil, err
	}

	return h, nil
}

func toValueBytes(value, valueSecret *big.Int) []byte {
	var valueBytes [64]byte
	vb := value.Bytes()
	copy(valueBytes[32-len(vb):32], vb)
	vsb := valueSecret.Bytes()
	copy(valueBytes[64-len(vsb):64], vsb)

	return valueBytes[:]
}

func fromValueBytes(valueBytes []byte) (*big.Int, *big.Int) {
	return new(big.Int).SetBytes(valueBytes[:32]), new(big.Int).SetBytes(valueBytes[32:])
}

func RandBigInt() *big.Int {
	randomBytes := make([]byte, 32)
	io.ReadFull(rand.Reader, randomBytes)
	return new(big.Int).SetBytes(randomBytes)
}
