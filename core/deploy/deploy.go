package deploy

import (
	"fmt"
	"github.com/UDPN/zk-token/core/client"
	"github.com/UDPN/zk-token/core/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"math/big"
	"time"
)

func NewDeployClient(url, hexKey string) (*DeployToken, error) {
	cli, err := client.NewETHClientByUrlAndHexKey(url, hexKey)
	if err != nil {
		return nil, err
	}
	return &DeployToken{
		cli: cli,
	}, nil
}

type DeployToken struct {
	cli *client.ETHClient

	BalanceVerifierAddress  common.Address
	TransferVerifierAddress common.Address
	TokenAddress            common.Address
}

func (t *DeployToken) GoDeploy() error {

	if err := t.deployBalanceVerifier(); err != nil {
		return err
	}
	time.Sleep(20 * time.Second)
	if err := t.deployTransferVerifier(); err != nil {
		return err
	}
	time.Sleep(20 * time.Second)
	if err := t.deployToken(); err != nil {
		return err
	}

	return nil
}

func (t *DeployToken) deployBalanceVerifier() error {
	addr, txId, err := t.cli.Deploy(func(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, error) {
		addr, txId, _, err := contracts.DeployBalanceVerifier(auth, backend)
		return addr, txId, err
	})

	if err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("Deploy BalanceVerifler Contracts Success!, Address : %s , TxId : %s .", addr.String(), txId))
	t.BalanceVerifierAddress = addr
	return nil
}

func (t *DeployToken) deployTransferVerifier() error {
	addr, txId, err := t.cli.Deploy(func(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, error) {
		addr, txId, _, err := contracts.DeployTransferVerifier(auth, backend)
		return addr, txId, err
	})

	if err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("Deploy TransferVerifier Contracts Success!, Address : %s , TxId : %s .", addr.String(), txId))
	t.TransferVerifierAddress = addr
	return nil
}

func (t *DeployToken) deployToken() error {
	addr, txId, err := t.cli.Deploy(func(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, error) {
		addr, txId, _, err := contracts.DeployHideToken(auth, backend, t.BalanceVerifierAddress, t.TransferVerifierAddress, new(big.Int).SetUint64(1000))
		return addr, txId, err
	})

	if err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("Deploy Token Contracts Success!, Address : %s , TxId : %s .", addr.String(), txId))
	t.TokenAddress = addr
	return nil
}
