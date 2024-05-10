package key

import (
	"encoding/hex"
	"fmt"
	"github.com/UDPN/zk-token/core/testdata"
	"github.com/UDPN/zk-token/core/utils"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGenKey(t *testing.T) {

	key, _ := crypto.GenerateKey()

	address := crypto.PubkeyToAddress(key.PublicKey)
	fmt.Println(address)

	keyB := crypto.FromECDSA(key)

	keyHex := hex.EncodeToString(keyB)

	//fmt.Println(keyHex)
	utils.WriteFile([]byte(keyHex), fmt.Sprintf("./account/%s.key", address.String()), true)

}

func TestPubKey(t *testing.T) {
	hexKey := testdata.User1_HexKey
	keyBytes, err := hexutil.Decode(hexKey)
	require.NoError(t, err)

	pk, err := crypto.ToECDSA(keyBytes)
	require.NoError(t, err)

	pubkeyBytes := crypto.FromECDSAPub(&pk.PublicKey)[1:]
	fmt.Println(len(pubkeyBytes))
	fmt.Println(hexutil.Encode(pubkeyBytes))

	var pubb [65]byte
	pubb[0] = 4
	copy(pubb[1:], pubkeyBytes)
	pub2, err := crypto.UnmarshalPubkey(pubb[:])
	require.NoError(t, err)

	fmt.Println(crypto.PubkeyToAddress(pk.PublicKey))
	fmt.Println(crypto.PubkeyToAddress(*pub2))
}

func TestECDHKey(t *testing.T) {

	hexKey := testdata.User1_HexKey
	hexKey2 := testdata.User2_HexKey

	keyBytes, err := hexutil.Decode(hexKey)
	require.NoError(t, err)

	pk, err := crypto.ToECDSA(keyBytes)
	require.NoError(t, err)

	keyBytes2, err := hexutil.Decode(hexKey2)
	require.NoError(t, err)

	pk2, err := crypto.ToECDSA(keyBytes2)
	require.NoError(t, err)

	dh, err := EcdhKeyAgreement(&pk.PublicKey, pk2)
	require.NoError(t, err)
	fmt.Println(hexutil.Encode(dh), len(dh))

	dh2, err := EcdhKeyAgreement(&pk2.PublicKey, pk)
	require.NoError(t, err)
	fmt.Println(hexutil.Encode(dh2))

	data := []byte("100000")

	s, err := AESEncrypt(dh, data)
	require.NoError(t, err)
	fmt.Println(hexutil.Encode(s))

	d, err := AESDecrypt(dh2, s)
	require.NoError(t, err)
	fmt.Println(string(d))

}
