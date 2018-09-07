package main

import (
	"github.com/essblock/go-sdk/build"
	"fmt"
	"github.com/essblock/go-sdk/clients/horizon"
)

func main() {
	seed := "SB2CSRVWBBYMB6IY2PUYUB7PTSEJ32G6375SEMY76YLTKI2XSGCF43MX"
	tx, err := build.Transaction(
		build.SourceAccount{AddressOrSeed: seed},
		build.AutoSequence{SequenceProvider: horizon.DefaultTestNetClient},
		build.TestNetwork,
		build.SetOptions(
			build.InflationDest("GDC3OTAEJ2QXWOAKMGAY4PBSOHHG7HOF5BICLLOIIXDCHZLZPTEDVSN4"),
			build.SetAuthRequired(),
			build.SetAuthRevocable(),
			build.SetAuthImmutable(),
			build.ClearAuthRequired(),
			build.ClearAuthRevocable(),
			build.ClearAuthImmutable(),
			build.MasterWeight(1),
			build.SetThresholds(2, 3, 4),
			build.HomeDomain("ess.org"),
			build.AddSigner("GBXLI7NVHPCDYFFHISSLQEZH7L6CGNROYMJY4W7VPWOEUGNAM2JJVP4E", 5),
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
		fmt.Println(err.(*horizon.Error).Problem)
		panic(err)
	}

	fmt.Println("transaction posted in ledger:", resp.Ledger)
}
