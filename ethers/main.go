package main

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"log"
	"math/big"

	//	"time"

	"fmt"
	"math"

	token "go-test/ethers/token"

	//	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"golang.org/x/crypto/sha3"
)

type nClient struct {
	clientUrl  string
	privateKey string
}

type sToken struct {
	name    string
	address string
}

// // mainnet
const (
	WETH  string = "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"
	WBTC         = "0x2260fac5e5542a773aa44fbcfedf7c193bc2c599"
	CWBTC        = "0xC11b1268C1A384e55C48c2391d8d480264A3A7F4"
	USDC         = "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"
	CUSDC        = "0x39AA39c021dfbaE8faC545936693aC917d5E7563"
	CETH         = "0x4Ddc2D193948926D02f9B1fE9e1daa0718270ED5"
	DAI          = "0x6B175474E89094C44Da98b954EedeAC495271d0F"
	CDAI         = "0x5d3a536E4D6DbD6114cc1Ead35777bAB948E3643"
	USDT         = "0xdAC17F958D2ee523a2206206994597C13D831ec7"
	CUSDT        = "0xf650C3d88D12dB855b8bf7D11Be6C55A4e07dCC9"

	OSQTH       = "0xf1B99e3E573A1a9C5E6B2Ce818b617F0E664E86B"
	callee      = "0xEcA3eDfD09435C2C7D2583124ca9a44f82aF1e8b"
	univ2router = "0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D"
	POOL        = "0x88e6A0c2dDD26FEEb64F039a2c41296FcB3f5640"
)

// const (
// 	WETH        string = "0xB4FBF271143F4FBf7B91A5ded31805e42b2208d6"
// 	CETH               = "0x20572e4c090f15667cF7378e16FaD2eA0e2f3EfF"
// 	USDC               = "0xD87Ba7A50B2E7E660f678A895E4B72E7CB4CCd9C"
// 	CUSDC              = "0xCEC4a43eBB02f9B80916F1c718338169d6d5C1F0"
// 	CWBTC              = "0x6CE27497A64fFFb5517AA4aeE908b1E7EB63B9fF"
// 	WBTC               = "0xC04B0d3107736C32e19F1c62b2aF67BE61d63a05"
// 	CDAI               = "0x822397d9a55d0fefd20F5c4bCaB33C5F65bd28Eb"
// 	DAI                = "0xdc31Ee1784292379Fbb2964b3B9C4124D8F89C60"
// 	callee             = "0xE97f1488F053251032ef358dE5b4188cD960D413"
// 	univ2router        = "0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D"
// 	POOL               = "0x88e6A0c2dDD26FEEb64F039a2c41296FcB3f5640"
// )

func main() { ///-----------

	networks := [...]nClient{

		{ ///0 mainnet
			"https://mainnet.infura.io/v3/68070d464ba04080a428aeef1b9803c6",
			""},
		{ ///1  geth nodes
			"http://192.168.0.223:8546",
			"f12e87b71f16f24f192941ecc3aff77e8ebcc7be01dd8d7d1245d7644ce8c93b", // 0x5D2Dd8F38a74294DD1FAF7B8c806446A3379f330 GOERLI
		}, //
		{ ///2  fork from ganache
			"http://127.0.0.1:8545", // ganache-cli windows
			//"5a2777bc3f36b7bae5c23f30a476158b871a5e471cc9fdb8d050ada6cb65b07c"
			//"http://192.168.0.223:8545", // ganache-cli centos
			"b8c1b5c1d81f9475fdf2e334517d29f733bdfa40682207571b12fc1142cbf329",
		},
	}

	nid := 2
	TokenSelect := WETH

	client, err := ethclient.Dial(networks[nid].clientUrl)
	privateKey, err := crypto.HexToECDSA(networks[nid].privateKey)

	if err != nil {
		panic(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("invalid key")
	}

	accountAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), accountAddress)
	if err != nil {
		panic(err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		panic(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)                 // in wei
	auth.GasLimit = uint64(6721975)            // in units
	auth.GasPrice = big.NewInt(8 * 1000000000) // 1 gwei = 1000000000

	tokenAddress := common.HexToAddress(TokenSelect)

	tokenInstance, err := token.NewApi(tokenAddress, client)
	if err != nil {
		panic(err)
	}

	fmt.Println("network:", networks[nid].clientUrl)

	fmt.Println("token address:", tokenAddress)

	totalSupply, err := tokenInstance.TotalSupply(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}
	fmt.Println("totalsupply:", totalSupply)

	decimals, err := tokenInstance.Decimals(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}
	fmt.Println("decimals:", decimals)
	name, err := tokenInstance.Name(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}
	fmt.Println("name:", name)
	symbol, err := tokenInstance.Symbol(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}
	fmt.Println("symbol:", symbol)

	bal, err := tokenInstance.BalanceOf(&bind.CallOpts{}, accountAddress)
	if err != nil {
		panic(err)
	}
	//convert to float64

	balanceOf := new(big.Float).SetInt(bal)

	res := new(big.Float).Quo(balanceOf, big.NewFloat(math.Pow(10, float64(decimals))))

	fmt.Println("token balanceOf accountAddress:", balanceOf)
	fmt.Println("readable:", res)

	ethbalance, err := client.BalanceAt(context.Background(), accountAddress, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ethbalance:", ethbalance) // 25893180161173005034
	fmt.Println("account address:", accountAddress)
	//return

	///# create raw transaction
	fmt.Println("Creating raw tx:")

	value := big.NewInt(0)    // in wei (1 eth)
	gasLimit := uint64(61000) // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("gass price:", gasPrice)

	//univ2router 0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D
	//swapExactETHForTokens(uint amountOutMin, address[] calldata path, address to, uint deadline)

	//callee 0xEcA3eDfD09435C2C7D2583124ca9a44f82aF1e8b
	// swap
	// osqth controller 0x64187ae08781B09368e6253F9E94951243A493D5

	//_contract = "0xEcA3eDfD09435C2C7D2583124ca9a44f82aF1e8b"
	_contract := "0x64187ae08781B09368e6253F9E94951243A493D5"
	contractAddress := common.HexToAddress(_contract)

	//	zeroForOne := false
	//	amount := big.NewInt(1e18)
	// path := WETH + "," + USDC //[2]string{WETH, USDC}
	// to := accountAddress.Hex()
	// timestamp := big.NewInt(time.Now().Unix())

	//transferFnSignature := []byte("swapExactETHForTokens(uint, address[], address, uint)")
	//	transferFnSignature := []byte("swap(address, bool, int256)")
	transferFnSignature := []byte("getExpectedNormalizationFactor()")
	hash := sha3.NewLegacyKeccak256() //.NewKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Println(hexutil.Encode(methodID))

	// paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	// fmt.Println("padded amount:", hexutil.Encode(paddedAmount))
	// paddedAddresses := common.LeftPadBytes([]byte(path), 32)
	// fmt.Println("padded addresses:", hexutil.Encode(paddedAddresses))
	// paddedAddress := common.LeftPadBytes([]byte(to), 32)
	// fmt.Println("padded address:", hexutil.Encode(paddedAddress))
	// paddedTimestamp := common.LeftPadBytes(timestamp.Bytes(), 32)
	// fmt.Println("padded timestamp:", timestamp.Bytes())

	// paddedAddress := common.LeftPadBytes([]byte(_pool), 32)
	// fmt.Println("padded address:", hexutil.Encode(paddedAddress))
	// paddedBool := common.LeftPadBytes([]byte(FormatBool(zeroForOne)), 32)
	// fmt.Println("padded bool:", hexutil.Encode(paddedBool))
	// paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	// fmt.Println("padded amount:", hexutil.Encode(paddedAmount))

	var data []byte
	data = append(data, methodID...)
	//	data = append(data, paddedAmount...)
	//	data = append(data, paddedAddresses...)
	// data = append(data, paddedAddress...)
	// data = append(data, paddedTimestamp...)

	// data = append(data, paddedAddress...)
	// data = append(data, paddedBool...)
	// data = append(data, paddedAmount...)

	tx := types.NewTransaction(nonce, contractAddress, value, gasLimit, gasPrice, data)

	chainID, err = client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	ts := types.Transactions{signedTx}
	// rawTxBytes := ts.GetRlp(0)
	// rawTxHex := hex.EncodeToString(rawTxBytes)

	b := new(bytes.Buffer)
	ts.EncodeIndex(0, b)
	rawTxBytes := b.Bytes()
	rawTxHex := hex.EncodeToString(rawTxBytes)

	fmt.Println(rawTxHex) // f86...772

	// # send raw transaction
	rawTxBytes, err = hex.DecodeString(rawTxHex)

	tx = new(types.Transaction)
	rlp.DecodeBytes(rawTxBytes, &tx)

	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s\n", tx.Hash().Hex())
}

func FormatBool(b bool) string {
	if b {
		return "true"
	}
	return "false"
}

// // encodeState encodes the state as with abi.encode() in the smart contracts.
// func computePoolAddress(a, b []byte) ([]byte, error) {
// 	args := abi.Arguments{
// 		{Type: abiUint256},
// 		{Type: abiUint256},
// 	}
// 	enc, err := args.Pack(
// 		a,
// 		b,
// 	)
// 	// now you have abi.encode(a,b)
// 	h := crypto.Keccak256(enc)
// 	// now you have keccak(abi.encode(a,b))
// 	// return the last 20 bytes
// 	return h[12:], err
// }

/*
func abiEncode() {
uint256Ty, _ := abi.NewType("uint256")
bytes32Ty, _ := abi.NewType("bytes32")
addressTy, _ := abi.NewType("address")

arguments := abi.Arguments{
{
Type: addressTy,
},
{
Type: bytes32Ty,
},
{
Type: uint256Ty,
},
}

bytes, _ := arguments.Pack(
common.HexToAddress("0x0000000000000000000000000000000000000000"),
[32]byte{'I', 'D', '1'},
big.NewInt(42),
)

var buf []byte
hash := sha3.NewKeccak256()
hash.Write(bytes)
buf = hash.Sum(buf)

log.Println(hexutil.Encode(buf))
// output:
// 0x1f214438d7c061ad56f98540db9a082d372df1ba9a3c96367f0103aa16c2fe9a
}
*/
