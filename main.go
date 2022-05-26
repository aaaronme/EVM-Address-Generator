package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

var prefix = "0x000000"
var suffix = "00"

func main() {
	var i int
	for {
		i++
		fmt.Printf("\033[2K\r%d", i)
		// go gen()
		// go gen()
		// go gen()
		gen()
	}
}

func gen() {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	// publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	// fmt.Println("Public Key:", hexutil.Encode(publicKeyBytes))

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	if strings.HasPrefix(address, prefix) && strings.HasSuffix(address, suffix) {
		fmt.Println("Address:", address)
		privateKeyBytes := crypto.FromECDSA(privateKey)
		fmt.Println("SAVE BUT DO NOT SHARE THIS (Private Key):", hexutil.Encode(privateKeyBytes))
	}

}
