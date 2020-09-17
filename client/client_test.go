package client_test

import (
	"github.com/John-Nagle-Iv/stocks/client"
	"net/url"
	"testing"
)

var RAW_RESPONSE []byte = []byte(`{
	"Global Quote": {
		"01. symbol": "MSFT",
		"02. open": "158.3200",
		"03. high": "159.9450",
		"04. low": "158.0600",
		"05. price": "158.6200",
		"06. volume": "21121681",
		"07. latest trading day": "2020-01-03",
		"08. previous close": "160.6200",
		"09. change": "-2.0000",
		"10. change percent": "-1.2452%"
	}
}`)

func TestRegexTimeParse(t *testing.T) {
	raw := []byte("2020-01-03")

	f := client.RFC_3339.FindIndex(raw)
	if f[0] != 0 {
		t.Error("Failed to find regex date patern")
		return
	}
	sm := client.RFC_3339.FindSubmatch(raw)[0]
	if string(sm) != "2020-01-03" {
		t.Error("Failed to find submatch date patern")
		return
	}
	l := string(client.RFC_3339.ReplaceAll(raw, []byte("${1}T00:00:00Z")))
	if l != "2020-01-03T00:00:00Z" {
		t.Error("Failed to format time as expected.")
		return
	}
}

func TestGetUrl(t *testing.T) {
	originalKey := client.API_KEY
	defer func() {
		client.API_KEY = originalKey
	}()
	client.API_KEY = "demo"
	expectedUrl, err := url.Parse("https://www.alphavantage.co/query?apikey=demo&function=GLOBAL_QUOTE&symbol=MSFT")
	if err != nil {
		t.Error(err)
	}
	actualUrl, err := client.GetUrl("MSFT", client.QUOTE)
	if err != nil {
		t.Error(err)
	}
	if actualUrl.Query().Encode() != expectedUrl.Query().Encode() {
		t.Errorf("\nExpected: %q,\n  Actual: %q", expectedUrl.String(), actualUrl.String())
	}
	if actualUrl.Hostname() != expectedUrl.Hostname() {
		t.Errorf("\nExpected: %q,\n  Actual: %q", expectedUrl.String(), actualUrl.String())
	}
}

func TestMarshalQuote(t *testing.T) {
	gq := client.MarshalQuote(RAW_RESPONSE)
	if gq.Symbol != "MSFT" {
		t.Error("Failed to marshal quote")
	}
}

/*
func TestGet(t *testing.T) {
	b := client.Get("MSFT", client.QUOTE)
	if b == nil {
		t.Error("bytes nil")
	}
	q := new(client.GlobalQuoteResponse)
	err := json.Unmarshal(b, q)
	if err != nil {
		t.Error(err)
	}
	if q.GlobalQuote.Symbol != "MSFT" {
		t.Error("Incorrect GlobalQuote Symbol")
	}
}
*/
