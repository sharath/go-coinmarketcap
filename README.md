![CMC Logo](coinmarketcap.svg)

[![GoDoc](https://godoc.org/github.com/sharath/go-coinmarketcap/github?status.svg)](https://godoc.org/github.com/sharath/go-coinmarketcap) [![Build Status](https://travis-ci.org/sharath/go-coinmarketcap.svg?branch=master)](https://travis-ci.org/sharath/go-coinmarketcap) [![Coverage Status](https://coveralls.io/repos/github/sharath/go-coinmarketcap/badge.svg?branch=master)](https://coveralls.io/github/sharath/go-coinmarketcap?branch=master)

go-coinmarketcap is a Go client library for accessing the CoinMarketCap Professional API

go-coinmarketcap requires Go version 1.8 or greater.

# Usage

```go
package main

import (
	"fmt"
	cmc "github.com/sharath/go-coinmarketcap"
)

func main() {
	client := cmc.NewClient("my_api_key")
	fmt.Println(client.GetInfoBySymbol("BTC", "ETH", "LTC"))
}
```

# Roadmap

Implemented Endpoints
- [ ] 