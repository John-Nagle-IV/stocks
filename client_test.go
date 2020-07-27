package client_test

import (
    "fmt"
	"encoding/json"
	"github.com/stocks/market/alphavantage/client"
	"net/url"
	"testing"
)

func TestGetUrl(t *testing.T) {
	originalKey := client.API_KEY
	defer func() {
		client.API_KEY = originalKey
	}()
	client.API_KEY = "demo"
	expectedUrl, err := url.Parse("https://www.alphavantage.co/query?function=GLOBAL_alphavantage.QUOTE&symbol=MSFT&apikey=demo")
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

func TestUnmarshalQuote(t *testing.T) {
        raw := []byte(`{
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
        }`)
        var q client.Quote
        l := client.LEADING_NUMBERS.ReplaceAllLiteral(raw, []byte(""))
        l = client.RFC_3339.ReplaceAll(l, []byte("$1T00:00:00Z"))
        fmt.Println(string(l))
        json.Unmarshal(l, &q)
        if q.Symbol != "MSFT" || q.Volume != 21121681 {
            t.Error("Quote not unmarshaled")
        }
}

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
