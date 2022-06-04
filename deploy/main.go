package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var url = "https://kovan.infura.io/v3/4d584c8d84714c5db4b20656605dfb04"

func main() {
	b, err := ioutil.ReadFile("./wallet/UTC--2022-06-04T10-02-26.295920618Z--550e041e51429ec280e838a0f4d22e0fae9e8702")
	if err != nil {
		log.Fatal(err)
	}

	key, err := keystore.DecryptKey(b, "password")
	if err != nil {
		log.Fatal(err)
	}

	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	add:=crypto.PubkeyToAddress(key.PrivateKey.PublicKey)

	nonce, err := client.PendingNonceAt(context.Background(),add)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err!=nil {
		log.Fatal(err)
	}
	
	chainID, err := client.NetworkID(context.Background())
	if err!= nil {
		log.Fatal(err)
	}

	auth,err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, chainID)
	if err!= nil {
		log.Fatal(err)
	}
	auth.GasPrice = gasPrice
	auth.GasLimit = uint64(3000000)
	auth.Nonce = big.NewInt(int64(nonce))
	a,tx, _, err := DeployTodo(auth, client)
	if err!= nil {
		log.Fatal(err)
	}
	fmt.Println("-----------------------------------")
	fmt.Println(a.Hex())
	fmt.Println(tx.Hash().Hex())
	fmt.Println("-----------------------------------")



}