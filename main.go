package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var infuraURL = "https://mainnet.infura.io/v3/4d584c8d84714c5db4b20656605dfb04"
// var ganachecli = "http://localhost:8545"

// var infuraURL = "https://nodeapi.test.energi.network/v1/jsonrpc"
// var infuraURL = "https://nodeapi.energi.network/v1/jsonrpc"
func main() {
    client, err := ethclient.DialContext(context.Background(),infuraURL)
	// client, err := ethclient.DialContext(context.Background(),ganachecli)
    if err != nil {
        log.Fatal(err)
    }

	defer client.Close()

   block, err := client.BlockByNumber(context.Background(),nil)
   if err != nil {
	   log.Fatalf("Error to get a block:%v", err)
   }
   fmt.Println("The block number is",block.Number())

   addr := "0x214e2128C360F8dF45f19415C95A2E67636803e4"
   address:=common.HexToAddress(addr)

   balance, err:= client.BalanceAt(context.Background(), address, nil)
   if err != nil {
	   log.Fatal("Error to get balance")
   }

   fmt.Println("The balance is ",balance)
	//    1 ether = 10^18 Wei
	fBlance := new(big.Float)
	fBlance.SetString(balance.String())
	fmt.Println(fBlance)
	balanceEther := new(big.Float).Quo(fBlance, big.NewFloat(math.Pow10(18)))
	fmt.Println("ether balance is", balanceEther)
}