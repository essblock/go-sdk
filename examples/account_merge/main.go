package main

import (
	"fmt"

	"github.com/essblock/go-sdk/build"
	"github.com/essblock/go-sdk/clients/horizon"
)

func main() {
	seed := "SBSG3LCBCRJQTAQPTXUOJ627YT7JDJOXPL3HVTVLJBPXJL6JOW32B2HB"
	tx, err := build.Transaction(
		build.SourceAccount{AddressOrSeed: seed},
		build.AutoSequence{SequenceProvider: horizon.DefaultTestNetClient},
		build.TestNetwork,
		build.AccountMerge(
			build.Destination{AddressOrSeed: "GBT7CLZHHRORQFNJ6QMDOELFRTN5IYW6QZPHKC5KVUX5JS2N62UKZGA6"},
		),
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	txe, err := tx.Sign(seed)
	if err != nil {
		fmt.Println(err)
		return
	}

	txeB64, err := txe.Base64()

	if err != nil {
		fmt.Println(err)
		return
	}


	fmt.Printf("tx base64: %s", txeB64)
	fmt.Println()

	resp, err := horizon.DefaultTestNetClient.SubmitTransaction(txeB64)
	if err != nil {
		panic(err)
	}

	fmt.Println("transaction posted in ledger:", resp.Ledger)
}
