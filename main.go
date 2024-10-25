package main

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	// generate private key(256 bits(32 bytes))
	privateKeyECDSA, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}

	// get public key from private key(256 bits(32 bytes)) by ECDSA algorithm
	publicKey := privateKeyECDSA.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("cannot convert public key to ECDSA")
	}

	// get address from public key(160 bits(20 bytes))
	// keccak256 hash of the end of 20 bytes public key
	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	// convert ECDSA private key to private key
	privateKeyHex := fmt.Sprintf("%x", crypto.FromECDSA(privateKeyECDSA))

	// output private key and address, separated by |
	fmt.Printf("%s|%s\n", privateKeyHex, address.Hex())
}
