package main

import (
	"fmt"

	"github.com/essblock/go-sdk/build"
	"github.com/essblock/go-sdk/clients/horizon"
	"github.com/essblock/go-sdk/keypair"
)

func main() {
	pair, err := keypair.Random()
	if err != nil {
		fmt.Println("创建公私钥失败：", err)
		return
	}
	fmt.Println("创建的种子：", pair.Seed())
	fmt.Println("创建的公钥：", pair.Address())

	tx, err := build.Transaction(
		build.SourceAccount{AddressOrSeed: "GC24J75M54M2T6KSF2P3WDWKIMXR2KO6NYEOR6NKI4RNQXDKW5QKKNHV"},
		build.AutoSequence{SequenceProvider: horizon.DefaultTestNetClient},
		build.TestNetwork,
		build.CreateAccount(
			build.Destination{AddressOrSeed: pair.Address()},
		 	build.NativeAmount{Amount: "30"},
		),
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	txe, err := tx.Sign("SB5E6YYZAYMQDZBTI7ECNADEXJLMGWBC4VR5FSF7A5KUWPR64B76Q57X")
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
		fmt.Println(string(err.(*horizon.Error).Problem.Extras["result_codes"]))
		panic(err)
	}
	fmt.Println("transaction posted in ledger:", resp.Ledger)
}
