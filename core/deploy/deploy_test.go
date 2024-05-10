package deploy

import (
	"context"
	"fmt"
	"github.com/UDPN/zk-token/core/testdata"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
	"testing"
)

var (
	NodeUrl = testdata.NodeUrl
	HexKey  = testdata.User1_HexKey
)

func TestBalance(t *testing.T) {
	d, err := NewDeployClient(NodeUrl, HexKey)
	require.NoError(t, err)

	b, err := d.cli.Client().BalanceAt(context.Background(), common.HexToAddress("0x0012d00Dd5A23b74556740f6087070372D23f6D0"), nil)
	require.NoError(t, err)

	fmt.Println(b)
}

func TestDeploy(t *testing.T) {

	d, err := NewDeployClient(NodeUrl, HexKey)
	require.NoError(t, err)

	err = d.GoDeploy()
	require.NoError(t, err)

	//Deploy BalanceVerifler Contracts Success!, Address : 0x64f5e6C04AD7eecD878DBdE21D6E11CE46202F5a , TxId : 0x94bf45a5b4ea992752ae90d5fec3816f432292ea405323c154e1ab57e171a4eb .
	//Deploy TransferVerifier Contracts Success!, Address : 0x11D546b9Daf0269CA1941682C18aeE23f6Eb765e , TxId : 0xd7e9535c8f715ba58045fb82272ffac681378becd5d747e8e8d5466eb3fc7e4f .
	//Deploy Token Contracts Success!, Address : 0xe52e1445dd703d2Bb120423DE49E83D03cF0de8C , TxId : 0x502ef5a91fce1c4595b604723dc2af8ed894f62a50bf71033ece5a30c32ceadb .
}
