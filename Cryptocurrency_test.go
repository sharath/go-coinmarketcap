package go_coinmarketcap

import (
	"fmt"
	"os"
	"testing"
)

func checkKey(t *testing.T) *Client {
	c := NewClient(os.Getenv("CMC_API_KEY"))
	if c == nil {
		fmt.Println("Warning: CMC_API_KEY not set.")
		fmt.Println("This test doesn't mean anything basically.")
	}
	return c
}

func TestClient_GetInfo(t *testing.T) {
	c := checkKey(t)
	if c == nil {
		return
	}
	{
		data, err := c.GetInfo("symbol", "ETH,BTC,LTC")
		if err != nil || data == nil {
			t.FailNow()
		}
	}
	{
		data, err := c.GetInfo("id", "1,2,3")
		if err != nil || data == nil {
			t.FailNow()
		}
	}
}

func TestClient_GetInfoByID(t *testing.T) {
	c := checkKey(t)
	if c == nil {
		return
	}
	{
		data, err := c.GetInfoByID(1, 2, 3)
		if err != nil || data == nil {
			t.FailNow()
		}
	}
}

func TestClient_GetInfoBySymbol(t *testing.T) {
	c := checkKey(t)
	if c == nil {
		return
	}
	{
		data, err := c.GetInfoBySymbol("BTC", "ETH")
		if err != nil || data == nil {
			t.FailNow()
		}
	}
}

func TestClient_GetIDMapFor(t *testing.T) {
	c := checkKey(t)
	if c == nil {
		return
	}
	{
		data, err := c.GetIDMapFor("BTC", "ETH")
		if err != nil || data == nil {
			fmt.Println(err)
			t.FailNow()
		}
	}
}

func TestClient_GetIDMapWhere(t *testing.T) {
	c := checkKey(t)
	if c == nil {
		return
	}
	{
		data, err := c.GetIDMapWhere()
		if err != nil || data == nil {
			fmt.Println(err)
			t.FailNow()
		}
	}
}
