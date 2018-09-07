package main

import (
"fmt"

"github.com/essblock/go-sdk/build"
"github.com/essblock/go-sdk/clients/horizon"
)

func main() {
	seed := "SB2CSRVWBBYMB6IY2PUYUB7PTSEJ32G6375SEMY76YLTKI2XSGCF43MX"

	tx, err := build.Transaction(
		build.SourceAccount{AddressOrSeed: seed},
		build.AutoSequence{SequenceProvider: horizon.DefaultTestNetClient},
		build.TestNetwork,
		build.AllowTrust(
			"HAHA",
			"GBXLI7NVHPCDYFFHISSLQEZH7L6CGNROYMJY4W7VPWOEUGNAM2JJVP4E",
		),
	)

	if err != nil {
		panic(err)
	}

	txe, err := tx.Sign(seed)
	if err != nil {
		panic(err)
	}

	txeB64, err := txe.Base64()
	if err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}

	fmt.Printf("tx base64: %s", txeB64)
	fmt.Println()

	resp, err := horizon.DefaultTestNetClient.SubmitTransaction(txeB64)
	if err != nil {
		panic(err)
	}

	fmt.Println("transaction posted in ledger:", resp.Ledger)
}

