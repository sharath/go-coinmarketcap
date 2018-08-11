package go_coinmarketcap

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type Info map[string]TokenInfo

type CryptocurrencyResponse struct {
	Data   Info   `json:"data"`
	Status Status `json:"status"`
}

type TokenInfo struct {
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

type Status struct {
	Timestamp    time.Time `json:"timestamp"`
	ErrorCode    int       `json:"error_code"`
	ErrorMessage string    `json:"error_message"`
	Elapsed      int       `json:"elapsed"`
	CreditCount  int       `json:"credit_count"`
}

// GetInfoByID is a wrapper function for GetInfo
// Example Usage: GetInfoByID(1, 2, 3)
func (c *Client) GetInfoByID(ids ...int) (Info, error) {
	idsStr := strings.Trim(strings.Replace(fmt.Sprint(ids), " ", ",", -1), "[]")
	return c.GetInfo("id", idsStr)
}

// GetInfoByTicker is a wrapper function for GetInfo
// Example Usage: GetInfoBySymbol("BTC", "ETH", "LTC")
func (c *Client) GetInfoBySymbol(tickers ...string) (Info, error) {
	return c.GetInfo("symbol", strings.Join(tickers, ","))
}

// GetInfo fetches metadeta for specified ids or symbols
// Example usages: GetInfo("id", "1,2,3")
//                 GetInfo("symbols", "BTC,ETH,LTC")
func (c *Client) GetInfo(key, vals string) (Info, error) {
	endpt := "info"
	params := "?" + key + "=" + strings.Replace(vals, " ", "", -1)
	req, err := http.NewRequest("GET", apiURL+cryptocurrency+endpt+params, nil)
	if err != nil {
		return nil, err
	}
	req.Header["X-CMC_PRO_API_KEY"] = []string{c.apiKey}

	resp, err := c.client.Do(req)
	defer resp.Body.Close()

	response := new(CryptocurrencyResponse)
	err = json.NewDecoder(resp.Body).Decode(response)
	if err != nil {
		return nil, err
	}

	if response.Status.ErrorCode >= 400 {
		return nil, errors.New(response.Status.ErrorMessage)
	}
	return response.Data, err
}
