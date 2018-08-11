package go_coinmarketcap

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	c := NewClient("123")
	if c == nil {
		t.FailNow()
	}
}
