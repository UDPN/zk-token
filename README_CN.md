# zk-token

## 概述

该合约是基于零知识证明zkSNARK技术实现的隐藏账户余额的Token合约。
该合约中，账户持有人自己保管账户余额信息，链上存储隐藏的余额信息。
转账的双方通过提交零知识证明可以在余额保密的情况下完成交易。
其主要的流程是在交易中，由转账发起者向链上提交自己的余额变动证明和交易额度的加密信息，
转账的接收方监听到交易后，获取加密信息，解密后生成自己的余额变动证明向链上提交，确认交易。
在这个过程中转账发起者也可以撤回交易。  
合约的源码位于 `contracts`文件夹中 ，零知识证明电路文件位于`circuits`文件夹中。
我们在`Sepolia`测试网络部署了合约，地址是`0xe52e1445dd703d2Bb120423DE49E83D03cF0de8C`。
可以在[这里](https://sepolia.etherscan.io/address/0xe52e1445dd703d2bb120423de49e83d03cf0de8c) 看到详细信息。
你可以在[这里](https://github.com/UDPN/zk-token) 获取Go语言的SDK并参考下面的说明尝试调用该合约。

## 调用合约
1. 首先，你需要获取零知识证明电路的编译文件，他们包含 wasm、zkey和vkey文件。
你可以从代码仓库的`build`文件夹中获取这些。在初始化调用Client时，需要配置这个文件内容。如果你想学习zkSNARK的详细知识，可以在[snarkjs](https://github.com/iden3/snarkjs)和[circom](https://docs.circom.io/)获取帮助。
2. 在你的Go项目中引用SDK，并准备开始调用合约；
```shell
go get github.com/UDPN/zk-token
```
在你的文件中 import 路径
```go
import (
"github.com/UDPN/zk-token/core/token"
"github.com/UDPN/zk-token/core/zkproof"
)
```

3. 在调用之前，你需要准备一些调用参数包含节点地址、合约地址、调用ETH的账户私钥和一个用于保密余额的随机数字。
```go
nodeUrl := "你的sepolia测试节点连接"
key := "你的钱包地址hex格式私钥"
tokenAddress :="0xe52e1445dd703d2Bb120423DE49E83D03cF0de8C" //sepolia测试网络中部署的合约地址
balance :=new(big.Int).SetUint64(0), //初始的Balance信息
balanceSecret := new(big.Int).SetUint64(123456789) // 你的余额保密数字，随机的，bytes长度不大于32位

```

4. 创建TokenClient对象用于接下来的调用，并且在Client中配置电路文件，代码示例如下：
```go
cli, err := token.NewTokenCli(nodeUrl, key, common.HexToAddress(tokenAddress))
require.NoError(t, err)

err = cli.SetBalanceProofConfig(GetBalanceProofConf())
require.NoError(t, err)

err = cli.SetTransferProofConfig(GetTransferProofConf())
require.NoError(t, err)
```
关于电路文件的读取，参考代码如下：
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
5. 你可以通过下面的代码发起注册钱包的调用，将注册一个余额为0的钱包账户。
```go
input, txId, err := cli.RegisterWallet(balanceSecret)
require.NoError(t, err)
```
6. 你可以通过下面的代码为注册的钱包mint一些余额，用于后面的转账交易，在示例的合约中没有对mint操作进行限制，意味着你可以为任何已经注册的钱包mint余额，前提是你需要知道该钱包的现有余额和保密数字。请牢记你的钱包余额和保密数字，一旦丢失将无法交易。
```go
 addr :=common.HexToAddress("0x00000101")
 value := new(big.Int).SetInt64(100) //mint的余额
 valueSecret := new(big.Int).SetInt64(987654321) // value的保密数字，随机的

input, txId, err := cli.Mint(addr, balance, balanceSecret, value, valueSecret)
require.NoError(t, err)

balance = input.NewBalance() // 新的余额
```
7. 你可以通过下面的代码向一个其他已经注册的钱包发送余额，这些发送的余额将会立即从你的钱包中扣除，暂时存储在合约中，对方可以通过事件获取这次转账，并且通过确认交易收取余额。当然你也可以在确认之前随时取消这次转账。在示例合约的调用中，你只能向同一个账户同时发起一次交易，在交易被确认和撤销前无法再发起转账交易。
```go
toAddr :=common.HexToAddress("0x00000101")
value := new(big.Int).SetInt64(100) //发送的余额
valueSecret := new(big.Int).SetInt64(987654321) // value的保密数字，随机的

input, txId, err := cli.TransferToken(toAddr, balance, balanceSecret, value, valueSecret)
require.NoError(t, err)

balance = input.NewBalance()
```
8. 你可以通过下面的代码取消交易，将扣除的余额返还到自己的钱包账户中。请注意过期的交易不会自动返还余额，你仍然需要主动调用合约。
```go
toAddr :=common.HexToAddress("0x00000101")

input, txId, err = cli.CancelTransferToken(toAddr, balance, balanceSecret)
require.NoError(t, err)

balance = input.NewBalance()
```
9. 当你监听到有账户向你转账时，可以通过下面的代码确认交易，并且收取余额。交易将在1000个块后过期，必须在此之前确认收款。
```go
fromAddr :=common.HexToAddress("0x00000102")
input, txId, err = cli.ConfirmTransferToken(fromAddr, balance, balanceSecret)
require.NoError(t, err)

balance = input.NewBalance()
```
10. 请注意，在上面的交易中你需要在每次交易后记录自己的余额和保密数字，一旦丢失这些信息将无法再发起交易。你可以通过下面的代码随时检查你的余额是否和链上的信息保持一致。
```go
addr :=common.HexToAddress("0x00000101")
res, err := cli.CheckBalance(addr, balance, balanceSecret)
require.NoError(t, err)
require.Equal(t, res, true)
``` 

以上这些示例的代码可以在完整的流程测试中找到，他们在测试文件(tests/token_test.go)中。



