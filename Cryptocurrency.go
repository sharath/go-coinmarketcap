package go_coinmarketcap

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
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
	Status CMCStatus          `json:"status"`
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
	IsActive            bool      `json:"is_active"`
	FirstHistoricalData time.Time `json:"first_historical_data"`
	LastHistoricalData  time.Time `json:"last_historical_data"`
}
type CryptocurrencyMap []CryptocurrencyMapData
type CryptocurrencyMapResponse struct {
	Data   CryptocurrencyMap `json:"data"`
	Status CMCStatus         `json:"status"`
}

// GetIDMapFor fetches CMC ID Maps for specified parameters
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

func (c *Client) GetIDMapWhere(status string, start, limit int) (CryptocurrencyMap, error) {
	endpt := "map"
	req, err := http.NewRequest("GET", apiURL+cryptocurrency+endpt, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("listing_status", status)
	q.Add("start", strconv.Itoa(start))
	q.Add("limit", strconv.Itoa(limit))
	req.URL.RawQuery = q.Encode()

	req.Header["X-CMC_PRO_API_KEY"] = []string{c.apiKey}

	resp, err := c.client.Do(req)
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
