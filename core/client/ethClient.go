package client

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strconv"
)

func NewETHClientByUrlAndHexKey(url, hexKey string) (*ETHClient, error) {
	keyBytes, err := hexutil.Decode(hexKey)
	if err != nil {
		return nil, err
	}

	pk, err := crypto.ToECDSA(keyBytes)
	if err != nil {
		return nil, err
	}
	ethCli, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}

	return NewETHClient(ethCli, pk)
}

func NewETHClient(cli *ethclient.Client, pk *ecdsa.PrivateKey) (*ETHClient, error) {
	contractCli := &ETHClient{
		client: cli,
		key:    pk,
	}

	chainId, err := cli.ChainID(context.Background())
	if err != nil {
		return nil, err
	}
	contractCli.chainId = chainId

	return contractCli, nil
}

type ETHClient struct {
	client       *ethclient.Client
	session      interface{}
	key          *ecdsa.PrivateKey
	chainId      *big.Int
	transactOpts *bind.TransactOpts

	GasPrice string
	GasLimit string
}

func (d *ETHClient) Nonce() uint64 {
	nonce, _ := d.client.NonceAt(context.Background(), d.OwnerAddress(), nil)
	return nonce
}
func (d *ETHClient) PrivateKey() *ecdsa.PrivateKey {
	return d.key
}
func (d *ETHClient) PublicKey() *ecdsa.PublicKey {
	return &d.key.PublicKey
}

func (d *ETHClient) OwnerAddress() common.Address {
	return crypto.PubkeyToAddress(d.key.PublicKey)
}

type DeployContract func(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, error)

func (d *ETHClient) GetTrandOpts() *bind.TransactOpts {
	if d.transactOpts == nil {
		gasPrice, _ := d.ChainGasPrice()
		d.transactOpts = NewKeyedTransactor(d.key, d.chainId, gasPrice.String(), d.GasLimit)
	}
	return d.transactOpts

}
func (cli *ETHClient) ChainGasPrice() (*big.Int, error) {
	if cli.GasPrice != "" {
		gp, ok := new(big.Int).SetString(cli.GasPrice, 10)
		if ok {
			return gp, nil
		}
	}

	return cli.client.SuggestGasPrice(context.Background())

}
func (d *ETHClient) Deploy(deploy DeployContract) (addr common.Address, txId string, err error) {

	addr, tx, err := deploy(d.GetTrandOpts(), d.client)
	if err != nil {
		fmt.Println("deploy contract has err :", err.Error())
		return
	}
	return addr, tx.Hash().String(), err
}

type NewSession func(auth *bind.TransactOpts, opts bind.CallOpts, client *ethclient.Client) (interface{}, error)

func (d *ETHClient) Connection(session NewSession) error {
	transactOpts := d.GetTrandOpts()
	callOpts := bind.CallOpts{From: transactOpts.From}
	var err error
	d.session, err = session(transactOpts, callOpts, d.client)

	return err

}

func (d *ETHClient) Client() *ethclient.Client {
	return d.client
}

func (d *ETHClient) Session() interface{} {
	return d.session
}

//var gasPrice = new(big.Int).SetInt64(params.GWei)

//var gasPrice = new(big.Int).SetInt64(1900000000)

// NewKeyedTransactor is a utility method to easily create a transaction signer
// from a single private key.
func NewKeyedTransactor(key *ecdsa.PrivateKey, chainId *big.Int, gasPrice string, gaslimit string) *bind.TransactOpts {
	keyAddr := crypto.PubkeyToAddress(key.PublicKey)

	to := &bind.TransactOpts{
		//GasPrice: gasPrice,
		//GasLimit: 30000000,

		From: keyAddr,
		Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {

			signer := types.NewEIP155Signer(chainId)
			if address != keyAddr {
				return nil, errors.New("not authorized to sign this account")
			}
			signature, err := crypto.Sign(signer.Hash(tx).Bytes(), key)
			if err != nil {
				return nil, err
			}
			return tx.WithSignature(signer, signature)
		},
	}

	if gasPrice != "" {
		gp, ok := new(big.Int).SetString(gasPrice, 10)
		if ok {
			to.GasPrice = gp
		}
	}

	if gaslimit != "" {
		gl, err := strconv.ParseUint(gaslimit, 10, 64)
		if err != nil {
			to.GasLimit = gl
		}

	}

	return to
}

func (cli *ETHClient) Transfer(to common.Address, amount *big.Int) (common.Hash, error) {
	tranOpt := cli.GetTrandOpts()
	gasPrice, _ := cli.ChainGasPrice()
	data := &types.LegacyTx{
		Nonce:    cli.Nonce(),
		To:       &to,
		Value:    amount,
		Gas:      21000,
		GasPrice: gasPrice,
		Data:     nil,
	}

	tx := types.NewTx(data)

	signTx, err := tranOpt.Signer(cli.OwnerAddress(), tx)
	if err != nil {
		return common.Hash{}, err
	}
	err = cli.Client().SendTransaction(context.Background(), signTx)
	if err != nil {
		return common.Hash{}, err
	}
	return signTx.Hash(), nil
}
