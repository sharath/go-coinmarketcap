package go_coinmarketcap

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// cryptocurrency/info
type CryptocurrencyInfoData struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Symbol   string   `json:"symbol"`
	Category string   `json:"category"`
	Slug     string   `json:"slug"`
	Logo     string   `json:"logo"`
	Tags     []string `json:"tags"`
	Urls     struct {
		Website      []string `json:"website"`
		Explorer     []string `json:"explorer"`
		SourceCode   []string `json:"source_code"`
		MessageBoard []string `json:"message_board"`
		Chat         []string `json:"chat"`
		Announcement []string `json:"announcement"`
		Reddit       []string `json:"reddit"`
		Twitter      []string `json:"twitter"`
	} `json:"urls"`
}
type CryptocurrencyInfo map[string]CryptocurrencyInfoData
type CryptocurrencyInfoResponse struct {
	Data   CryptocurrencyInfo `json:"data"`
	Status ResponseStatus     `json:"status"`
}

// GetInfoByID is a wrapper function for GetInfo
// Example Usage: GetInfoByID(1, 2, 3)
func (c *Client) GetInfoByID(ids ...int) (CryptocurrencyInfo, error) {
	idsStr := strings.Trim(strings.Replace(fmt.Sprint(ids), " ", ",", -1), "[]")
	return c.GetInfo("id", idsStr)
}

// GetInfoByTicker is a wrapper function for GetInfo
// Example Usage: GetInfoBySymbol("BTC", "ETH", "LTC")
func (c *Client) GetInfoBySymbol(tickers ...string) (CryptocurrencyInfo, error) {
	return c.GetInfo("symbol", strings.Join(tickers, ","))
}

// GetInfo fetches metadeta for specified ids or symbols
// Example usages: GetInfo("id", "1,2,3")
//                 GetInfo("symbols", "BTC,ETH,LTC")
func (c *Client) GetInfo(key, vals string) (CryptocurrencyInfo, error) {
	endpt := "info"
	req, err := http.NewRequest("GET", apiURL+cryptocurrency+endpt, nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add(key, vals)
	req.URL.RawQuery = q.Encode()

	req.Header["X-CMC_PRO_API_KEY"] = []string{c.apiKey}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	response := new(CryptocurrencyInfoResponse)
	err = json.NewDecoder(resp.Body).Decode(response)
	if err != nil {
		return nil, err
	}

	if response.Status.ErrorCode >= 400 {
		return nil, errors.New(response.Status.ErrorMessage)
	}
	return response.Data, err
}

// cryptocurrency/map
type CryptocurrencyMapData struct {
	ID                  int       `json:"id"`
	Name                string    `json:"name"`
	Symbol              string    `json:"symbol"`
	Slug                string    `json:"slug"`
	IsActive            int       `json:"is_active"` // Note: the documentation says its a bool, but response isn't
	FirstHistoricalData time.Time `json:"first_historical_data"`
	LastHistoricalData  time.Time `json:"last_historical_data"`
}
type CryptocurrencyMap []CryptocurrencyMapData
type CryptocurrencyMapResponse struct {
	Data   CryptocurrencyMap `json:"data"`
	Status ResponseStatus    `json:"status"`
}

// GetIDMapFor fetches CMC ID Maps for specified symbols
func (c *Client) GetIDMapFor(symbols ...string) (CryptocurrencyMap, error) {
	endpt := "map"
	req, err := http.NewRequest("GET", apiURL+cryptocurrency+endpt, nil)
	if err != nil {
		return nil, err
	}

	params := strings.Join(symbols, ",")
	q := req.URL.Query()
	q.Add("symbol", params)
	req.URL.RawQuery = q.Encode()

	req.Header["X-CMC_PRO_API_KEY"] = []string{c.apiKey}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	response := new(CryptocurrencyMapResponse)
	err = json.NewDecoder(resp.Body).Decode(response)
	if err != nil {
		return nil, err
	}

	if response.Status.ErrorCode >= 400 {
		return nil, errors.New(response.Status.ErrorMessage)
	}
	return response.Data, err
}

// GetIDMapWhere fetches CMC ID Maps for specified parameters
// Allowed values for status: "active", "inactive"
// start >= 1
// limit 1 - 5000
func (c *Client) GetIDMapWhere(opts ...func(values url.Values) string) (CryptocurrencyMap, error) {
	endpt := "map"
	req, err := http.NewRequest("GET", apiURL+cryptocurrency+endpt, nil)
	if err != nil {
		return nil, err
	}
	for _, opt := range opts {
		req.URL.RawQuery = opt(req.URL.Query())
	}
	req.Header["X-CMC_PRO_API_KEY"] = []string{c.apiKey}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	response := new(CryptocurrencyMapResponse)
	err = json.NewDecoder(resp.Body).Decode(response)
	if err != nil {
		return nil, err
	}

	if response.Status.ErrorCode >= 400 {
		return nil, errors.New(response.Status.ErrorMessage)
	}
	return response.Data, err
}

// listings/latest
type CryptocurrencyListingsLatestPrice struct {
	Price            float64   `json:"price"`
	Volume24H        float64   `json:"volume_24h"`
	PercentChange1H  float64   `json:"percent_change_1h"`
	PercentChange24H float64   `json:"percent_change_24h"`
	PercentChange7D  float64   `json:"percent_change_7d"`
	MarketCap        float64   `json:"market_cap"`
	LastUpdated      time.Time `json:"last_updated"`
}
type CryptocurrencyListingsLatestData struct {
	ID                int       `json:"id"`
	Name              string    `json:"name"`
	Symbol            string    `json:"symbol"`
	Slug              string    `json:"slug"`
	CmcRank           int       `json:"cmc_rank"`
	NumMarketPairs    int       `json:"num_market_pairs"`
	CirculatingSupply float64   `json:"circulating_supply"`
	TotalSupply       float64   `json:"total_supply"`
	MaxSupply         float64   `json:"max_supply"`
	LastUpdated       time.Time `json:"last_updated"`
	DateAdded         time.Time `json:"date_added"`
	Quote             struct {
		USD CryptocurrencyListingsLatestPrice `json:"USD"`
		BTC CryptocurrencyListingsLatestPrice `json:"BTC"`
	} `json:"quote"`
}
type CryptocurrencyListingsLatest []CryptocurrencyListingsLatestData
type CryptocurrencyListingsLatestResponse struct {
	Data   CryptocurrencyListingsLatest `json:"data"`
	Status ResponseStatus               `json:"status"`
}

// GetIDMapFor fetches CMC ID Maps for specified symbols
func (c *Client) GetLatestListings(opts ...func(values url.Values) string) (CryptocurrencyListingsLatest, error) {
	endpt := "listings/latest"
	req, err := http.NewRequest("GET", apiURL+cryptocurrency+endpt, nil)
	if err != nil {
		return nil, err
	}
	for _, opt := range opts {
		req.URL.RawQuery = opt(req.URL.Query())
	}
	req.Header["X-CMC_PRO_API_KEY"] = []string{c.apiKey}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	response := new(CryptocurrencyListingsLatestResponse)
	err = json.NewDecoder(resp.Body).Decode(response)
	if err != nil {
		return nil, err
	}

	if response.Status.ErrorCode >= 400 {
		return nil, errors.New(response.Status.ErrorMessage)
	}
	return response.Data, err
}

// / market-pairs / latest
type CryptocurrencyMarketPairsLatest struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Symbol         string `json:"symbol"`
	NumMarketPairs int    `json:"num_market_pairs"`
	MarketPairs    []struct {
		Exchange struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Slug string `json:"slug"`
		} `json:"exchange"`
		MarketPair     string `json:"market_pair"`
		MarketPairBase struct {
			CurrencyID     int    `json:"currency_id"`
			CurrencySymbol string `json:"currency_symbol"`
			CurrencyType   string `json:"currency_type"`
		} `json:"market_pair_base"`
		MarketPairQuote struct {
			CurrencyID     int    `json:"currency_id"`
			CurrencySymbol string `json:"currency_symbol"`
			CurrencyType   string `json:"currency_type"`
		} `json:"market_pair_quote"`
		Quote struct {
			ExchangeReported struct {
				Price          float64   `json:"price"`
				Volume24HBase  float64   `json:"volume_24h_base"`
				Volume24HQuote float64   `json:"volume_24h_quote"`
				LastUpdated    time.Time `json:"last_updated"`
			} `json:"exchange_reported"`
			USD struct {
				Price       float64   `json:"price"`
				Volume24H   float64   `json:"volume_24h"`
				LastUpdated time.Time `json:"last_updated"`
			} `json:"USD"`
		} `json:"quote"`
	} `json:"market_pairs"`
}
type CryptocurrencyMarketPairsLatestResponse struct {
	Data   *CryptocurrencyMarketPairsLatest `json:"data"`
	Status ResponseStatus                   `json:"status"`
}

// GetLatestMarketPairsBy fetches CMC Market pairs for specified symbols
func (c *Client) GetLatestMarketPairsBy(opts ...func(values url.Values) string) (*CryptocurrencyMarketPairsLatest, error) {
	endpt := "/market-pairs/latest"
	req, err := http.NewRequest("GET", apiURL+cryptocurrency+endpt, nil)
	if err != nil {
		return nil, err
	}
	for _, opt := range opts {
		req.URL.RawQuery = opt(req.URL.Query())
	}
	req.Header["X-CMC_PRO_API_KEY"] = []string{c.apiKey}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	response := new(CryptocurrencyMarketPairsLatestResponse)
	err = json.NewDecoder(resp.Body).Decode(response)
	if err != nil {
		return nil, err
	}

	if response.Status.ErrorCode >= 400 {
		return nil, errors.New(response.Status.ErrorMessage)
	}
	return response.Data, err
}
