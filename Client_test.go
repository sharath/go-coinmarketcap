package go_coinmarketcap

import (
	"fmt"
	"os"
	"testing"
)

func TestNewClient(t *testing.T) {
	c := NewClient("123")
	if c == nil {
		t.FailNow()
	}
}

func TestClient_GetMetadata(t *testing.T) {
	c := NewClient(os.Getenv("CMC_API_KEY"))
	if c == nil {
		fmt.Println("Warning: CMC_API_KEY not set.")
		t.SkipNow()
	}
	{
		data, err := c.GetMetadata("symbol", "ETH,BTC,LTC")
		if err != nil || data == nil {
			fmt.Println(data, err)
			t.FailNow()
		}
	}
	{
		data, err := c.GetMetadata("id", "1,2,3")
		if err != nil || data == nil {
			fmt.Println(data, err)
			t.FailNow()
		}
	}
}

func TestClient_GetMetadataByID(t *testing.T) {
	c := NewClient(os.Getenv("CMC_API_KEY"))
	if c == nil {
		fmt.Println("Warning: CMC_API_KEY not set.")
		t.SkipNow()

	}
	{
		data, err := c.GetMetadataByID(1, 2, 3)
		if err != nil || data == nil {
			fmt.Println(data, err)
			t.FailNow()
		}
	}
}

func TestClient_GetMetadataBySymbol(t *testing.T) {
	c := NewClient(os.Getenv("CMC_API_KEY"))
	if c == nil {
		fmt.Println("Warning: CMC_API_KEY not set.")
		t.SkipNow()
	}
	{
		data, err := c.GetMetadataBySymbol("BTC", "ETH")
		if err != nil || data == nil {
			fmt.Println(data, err)
			t.FailNow()
		}
	}
}
