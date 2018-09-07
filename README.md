# Ess Go SDK

This repo is development based on stellar go sdk.

## Installation

```shell
go get github.com/essblock/go-sdk
```

## Usage

Let's first decode a transaction.
Please see [examples](https://github.com/essblock/go-sdk/examples)


## Dependencies

This repository we use [dep](https://golang.github.io/dep/) to manage them.
Dep is used to populate the [vendor directory](https://golang.github.io/dep/docs/ensure-mechanics.html), ensuring that builds are reproducible even as upstream dependencies are changed.
Please see the [dep](https://golang.github.io/dep/) website for installation instructions.

You can use dep yourself in your project and add ess go as a vendor'd dependency, or you can just drop this repos as `$GOPATH/src/github.com/essblock/go-sdk` to import it the canonical way (you still need to run `dep ensure -v`).


