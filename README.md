![CMC Logo](coinmarketcap.svg)

[![GoDoc](https://godoc.org/github.com/sharath/go-coinmarketcap/github?status.svg)](https://godoc.org/github.com/sharath/go-coinmarketcap)
[![Build Status](https://travis-ci.org/sharath/go-coinmarketcap.svg?branch=master)](https://travis-ci.org/sharath/go-coinmarketcap)
[![Coverage Status](https://coveralls.io/repos/github/sharath/go-coinmarketcap/badge.svg?branch=master)](https://coveralls.io/github/sharath/go-coinmarketcap?branch=master) 
[![Go Report Card](https://goreportcard.com/badge/github.com/sharath/go-coinmarketcap)](https://goreportcard.com/report/github.com/sharath/go-coinmarketcap)

go-coinmarketcap is a Go client library for accessing the CoinMarketCap Professional API

# Usage

```go
package main

import (
	cmc "github.com/sharath/go-coinmarketcap"
	"fmt"
	"os"
)

func main() {
	// make a new client
	client := cmc.NewClient(os.Getenv("CMC_API_KEY"))

	// fetch top 10 tokens
	listings, _ := client.GetLatestListings(
		cmc.Limit(10),
	)

	// print names and prices of top 10
	for _, tok := range listings {
		token := tok.Name
		price := tok.Quote.USD.Price

		fmt.Println(token, price)
	}
}

```

# Roadmap

Implementation Status:
- [ ] / cryptocurrency
    - [X] / info
    - [X] / map
    - [X] / listings / latest
    - [ ] / market-pairs / latest
    - [ ] / ohlcv / historical
    - [ ] / quotes
        - [ ] / latest
        - [ ] / historical
- [ ] / exchange
    - [ ] / info
    - [ ] / map
    - [ ] / listings / latest
    - [ ] / market-pairs / latest
    - [ ] / quotes
        - [ ] / latest
        - [ ] / historical
- [ ] / global-metrics
    - [ ] / quotes
        - [ ] / latest
        - [ ] / historical
- [ ] / tools
    - [ ] / price-conversion