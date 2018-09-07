package main

import (
	"fmt"

	"github.com/essblock/go-sdk/build"
	"github.com/essblock/go-sdk/clients/horizon"
)

func main() {
	from := "SB2CSRVWBBYMB6IY2PUYUB7PTSEJ32G6375SEMY76YLTKI2XSGCF43MX"

	tx, err := build.Transaction(
		build.SourceAccount{AddressOrSeed: from},
		build.AutoSequence{SequenceProvider: horizon.DefaultTestNetClient},
		build.TestNetwork,
		build.Payment(
			build.Destination{AddressOrSeed: "GBXLI7NVHPCDYFFHISSLQEZH7L6CGNROYMJY4W7VPWOEUGNAM2JJVP4E"},
			build.CreditAmount{"USD","GBXLI7NVHPCDYFFHISSLQEZH7L6CGNROYMJY4W7VPWOEUGNAM2JJVP4E","2"},
			build.PayWith(build.CreditAsset("EUR","GBXLI7NVHPCDYFFHISSLQEZH7L6CGNROYMJY4W7VPWOEUGNAM2JJVP4E"),"2").
				Through(build.Asset{Native: true}).
				Through(build.CreditAsset("BTC", "GAHJZHVKFLATAATJH46C7OK2ZOVRD47GZBGQ7P6OCVF6RJDCEG5JMQBQ")),
		),
	)

	if err != nil {
		panic(err)
	}

	txe, err := tx.Sign(from)
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
