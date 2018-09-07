package main

import (
	"fmt"

	"github.com/essblock/go-sdk/build"
	"github.com/essblock/go-sdk/clients/horizon"
)

func main() {
	seed := "SB2CSRVWBBYMB6IY2PUYUB7PTSEJ32G6375SEMY76YLTKI2XSGCF43MX"

	rate := build.Rate{
		Selling: build.NativeAsset(),
		Buying:  build.CreditAsset("HAHA", "GDC3OTAEJ2QXWOAKMGAY4PBSOHHG7HOF5BICLLOIIXDCHZLZPTEDVSN4"),
		Price:   build.Price("1"),
	}

	tx, err := build.Transaction(
		build.SourceAccount{AddressOrSeed: seed},
		build.AutoSequence{SequenceProvider: horizon.DefaultTestNetClient},
		build.TestNetwork,


		//create offer
		build.CreateOffer(rate, "2"),

		////update offer
		//build.UpdateOffer(rate, "4", build.OfferID(2), ),

		////delete offer
		//build.DeleteOffer(rate, build.OfferID(2)),
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
