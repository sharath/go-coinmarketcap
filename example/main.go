package main

import (
	"fmt"
	cmc "github.com/sharath/go-coinmarketcap"
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
