package go_coinmarketcap

import (
	"net/url"
	"strconv"
)

// https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis

// Limit: Limit on Response from API
// Default: 100
// Allowed: 1 - 5000
func Limit(limit int) func(values url.Values) string {
	return func(q url.Values) string {
		q.Add("limit", strconv.Itoa(limit))
		return q.Encode()
	}
}

// Start: ID to start results from
// Default: 1
// Allowed: >= 1
func Start(start int) func(values url.Values) string {
	return func(q url.Values) string {
		q.Add("start", strconv.Itoa(start))
		return q.Encode()
	}
}

// ListingStatus: option to filter by
// Default: "active"
// Allowed: "active", "inactive"
func ListingStatus(status string) func(values url.Values) string {
	return func(q url.Values) string {
		q.Add("listing_status", status)
		return q.Encode()
	}
}

// Sort: option to convert currency
// Default: "USD"
// Allowed: "USD", "AUD", "BRL", "CAD", "CHF", "CLP", "CNY", "CZK", "DKK", "EUR", "GBP", "HKD", "HUF",
//          "IDR", "ILS", "INR", "JPY", "KRW", "MXN", "MYR", "NOK", "NZD", "PHP", "PKR", "PLN", "RUB",
//          "SEK", "SGD", "THB", "TRY", "TWD", "ZAR",
func Convert(currency string) func(values url.Values) string {
	return func(q url.Values) string {
		q.Add("convert", currency)
		return q.Encode()
	}
}

// Sort: option to sort listing by
// Default: "market_cap"
// Allowed: "market_cap", "name", "symbol", "c", "price", "circulating_supply", "total_supply",
//          "max_supply", "num_market_pairs", "volume_24h", "percent_change_1h", "percent_change_24h",
//          "percent_change_7d"
func Sort(sort_by string) func(values url.Values) string {
	return func(q url.Values) string {
		q.Add("sort", sort_by)
		return q.Encode()
	}
}

// SortDir: option to sort listing in either ascending or descending order
// Default: "desc"
// Allowed: "asc", "desc"
func SortDir(dir string) func(values url.Values) string {
	return func(q url.Values) string {
		q.Add("sort_dir", dir)
		return q.Encode()
	}
}

// CryptocurrencyType: option to limit by type
// Default: "all"
// Allowed: "all", "coins", "tokens"
func CryptocurrencyType(ctype string) func(values url.Values) string {
	return func(q url.Values) string {
		q.Add("cryptocurrency_type", ctype)
		return q.Encode()
	}
}

// ID: ID to filter results for
// Allowed: >= 1
func ID(id int) func(values url.Values) string {
	return func(q url.Values) string {
		q.Add("start", strconv.Itoa(id))
		return q.Encode()
	}
}

// Symbol: cryptocurrency by symbol
// Examples: "BTC", "ETH", "LTC"
func Symbol(symbol string) func(values url.Values) string {
	return func(q url.Values) string {
		q.Add("symbol", symbol)
		return q.Encode()
	}
}
