# zk-token

## Overview

This contract is a Token contract for hiding account balances based on zero knowledge proof (zkSNARK) technology.
In this contract, the account holder holds their account balance information and stores hidden balance information on the chain.
Both parties involved in the transfer can complete the transaction while keeping the balance confidential by submitting a zero knowledge proof request.
The main process is that in the transaction, the initiator of the transfer submits their own proof of balance change and encrypted information of the transaction limit to the blockchain.
After the recipient of the transfer get the transaction, they can obtain the encrypted information. After decrypting the information and generating their proof of balance change, they can submit it to the blockchain to confirm this transaction.
During this process, the initiator of the transfer can also cancel the transaction.  

The source code of the contract is located in the 'contracts' folder, and the zero knowledge proof circuit file is located in the' circuits' folder.
We have deployed a contract with the address  0xe52e1445dd703d2Bb120423DE49E83D03cF0de8C on the Sepolia testnet.
You can find the detailed information [here](https://sepolia.etherscan.io/address/0xe52e1445dd703d2bb120423de49e83d03cf0de8c )
You can also obtain the SDK programmed by Go language [here](https://github.com/UDPN/zk-token) and try calling the contract by the instructions below.

## Calling Contract
1. Get files
   You need to obtain the compiled files of zero knowledge proof circuits, which include wasm, zkey, and vkey files.
   You can obtain these from the 'build' folder in the code repository. When initializing the call to the client, it is necessary to configure the contents of this file. To learn more about zkSNARK, you can go to [signaljs]( https://github.com/iden3/snarkjs) and [circle](https://docs.circom.io/).
2. Import
   Import the SDK into your Go project and prepare to call the contract:
```shell
go get github.com/UDPN/zk-token
```
Import the path into your folder:
```go
import (
"github.com/UDPN/zk-token/core/token"
"github.com/UDPN/zk-token/core/zkproof"
)
```

3. Preparation
   Before calling the contract, you need to prepare some parameters including node address, contract address, account private key for calling Ethereum network, and a random number for keeping the balance confidential.
```go
nodeUrl := "Your Sepolia testnet node url"
key := "The private key in hex format that corresponds to your wallet address"
tokenAddress := "0xe52e1445dd703d2Bb120423DE49E83D03cF0de8C" //The contract address deployed in the Sepolia testnet
balance :=new (big.Int).SetUint64(0), //Initial Balance information
balanceSecret :=new (big.Int).SetUint64(123456789) //Your balance is a random confidential number, with a byte length of no more than 32 bits
```

4. Configuration
   Create a TokenClient object for subsequent calls and configure the circuit file in the Client.
   The code example is as follows:
```go
cli, err := token.NewTokenCli(nodeUrl, key, common.HexToAddress(tokenAddress))
require.NoError(t, err)

err = cli.SetBalanceProofConfig(GetBalanceProofConf())
require.NoError(t, err)

err = cli.SetTransferProofConfig(GetTransferProofConf())
require.NoError(t, err)
```

Example code for reading the circuit file:

```go
func GetBalanceProofConf() *zkproof.ProofConfig {
    wasmBytes, _ := os.ReadFile(RootPath + Balance_wasm_path)
    zkeyBytes, _ := os.ReadFile(RootPath + Balance_zkey_path)
    vkeyBytes, _ := os.ReadFile(RootPath + Balance_vkey_path)

    conf := &zkproof.ProofConfig{
       WasmBytes: wasmBytes,
       ZkeyBytes: zkeyBytes,
       VkeyBytes: vkeyBytes,
    }
    return conf
}

func GetTransferProofConf() *zkproof.ProofConfig {

    wasmBytes, _ := os.ReadFile(RootPath + Transfer_wasm_path)
    zkeyBytes, _ := os.ReadFile(RootPath + Transfer_zkey_path)
    vkeyBytes, _ := os.ReadFile(RootPath + Transfer_vkey_path)

    conf := &zkproof.ProofConfig{
       WasmBytes: wasmBytes,
       ZkeyBytes: zkeyBytes,
       VkeyBytes: vkeyBytes,
    }
    return conf
}
```
5. Register wallet
   You can initiate a call to register a wallet with a balance of 0 by using the following code:

```go
input, txId, err := cli.RegisterWallet(balanceSecret)
require.NoError(t, err)
```

6. Mint balance
   You can mint some balance for the registered wallet through the following code for subsequent transfer transactions. There are no restrictions on mint operations in the example contract, which means you can mint the balance for any registered wallet of which you should know the existing balance and confidential number. Please remember your wallet balance and confidential numbers, once lost, you will not be able to send transactions.

```go
 addr :=common.HexToAddress("0x00000101")
 value := new(big.Int).SetInt64(100) //The minted balance
 valueSecret := new(big.Int).SetInt64(987654321) //The confidential number of value, which is a random number

input, txId, err := cli.Mint(addr, balance, balanceSecret, value, valueSecret)
require.NoError(t, err)

balance = input.NewBalance() //New balance
```
7. Send balance
   You can send balance to another registered wallet using the following code. The sent balance will be immediately deducted from your wallet and temporarily stored in the contract. The recipient can obtain this transfer by monitoring the event and collect the balance by confirming the transaction. You can also cancel this transfer at any time before confirmation. In the call of the example contract, you can only initiate a transaction to the same wallet once, and cannot initiate a transfer transaction again until the transaction is confirmed or cancelled.

```go
toAddr :=common.HexToAddress("0x00000101")
value := new(big.Int).SetInt64(100) //Sent balance
valueSecret := new(big.Int).SetInt64(987654321) //The confidential number of value, which is a random number

input, txId, err := cli.TransferToken(toAddr, balance, balanceSecret, value, valueSecret)
require.NoError(t, err)

balance = input.NewBalance()
```

8. Cancel transaction
   You can cancel the transaction and return the deducted balance to your wallet using the following code. Please note that expired transactions will not automatically refund the balance, and you need to call the contract again.

```go
toAddr :=common.HexToAddress("0x00000101")

input, txId, err = cli.CancelTransferToken(toAddr, balance, balanceSecret)
require.NoError(t, err)

balance = input.NewBalance()
```

9. Confirmation
   When you get a transfer transaction, you can confirm it and collect the balance using the following code. The transaction will expire in 1000 blocks and the confirmation must be before that.

```go
fromAddr :=common.HexToAddress("0x00000102")
input, txId, err = cli.ConfirmTransferToken(fromAddr, balance, balanceSecret)
require.NoError(t, err)

balance = input.NewBalance()
```

10. Check balance
    Please note that in the above transaction, you need to record your balance and confidential number after each transaction. Once this information is lost, you can no longer initiate the transaction again. You can check whether your balance matches the information on the chain at any time through the following code.

```go
addr :=common.HexToAddress("0x00000101")
res, err := cli.CheckBalance(addr, balance, balanceSecret)
require.NoError(t, err)
require.Equal(t, res, true)
```

The code for these examples can be found in the process testing, which is located in the testing file (tests/token_test. go).



