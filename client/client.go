package client

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"time"
)

type ApiFuncType string

type TimeSeriesInterval string

type OutputSize string

var API_KEY string = "UW5RD2YT3PAYVX40"
var LEADING_NUMBERS *regexp.Regexp = regexp.MustCompile("\\d+\\. ")
var RFC_3339 *regexp.Regexp = regexp.MustCompile("(\\d{4}-\\d{2}-\\d{2})")

const (
	ALPHAVANTAGE_URL string             = "https://www.alphavantage.co"
	COMPACT          OutputSize         = "compact"
	FULL             OutputSize         = "full"
	TS_01            TimeSeriesInterval = "1min"
	TS_05            TimeSeriesInterval = "5min"
	TS_15            TimeSeriesInterval = "15min"
	TS_30            TimeSeriesInterval = "30min"
	TS_60            TimeSeriesInterval = "60min"
	QUOTE            ApiFuncType        = "GLOBAL_QUOTE"
	INTRADAY         ApiFuncType        = "TIME_SERIES_INTRADAY"
	DAILY            ApiFuncType        = "TIME_SERIES_DAILY"
	WEEKLY           ApiFuncType        = "TIME_SERIES_WEEKLY"
	MONTHLY          ApiFuncType        = "TIME_SERIES_MONTHLY"
)

type Quote struct {
	Symbol           string    `json:"symbol"`
	Open             float64   `json:"open"`
	High             float64   `json:"high"`
	Low              float64   `json:"low"`
	Price            float64   `json:"price"`
	Volume           uint64    `json:"volume"`
	LatestTradingDay time.Time `json:"latest trading day"`
	PreviousClose    float64   `json:"previous close"`
	Change           float64   `json:"change"`
	ChangePercent    float64   `json:"change percent"`
}

type GlobalQuoteResponse struct {
	GlobalQuote Quote `json:"Global Quote"`
}

func GetUrl(symbol string, functype ApiFuncType, options ...url.Values) (*url.URL, error) {
	var s string
	for _, v := range options {
		s += "&" + v.Encode()
	}
	v, err := url.ParseQuery(s)
	if err != nil {
		log.Fatal(err)
	}
	v.Set("function", string(functype))
	v.Set("symbol", symbol)
	v.Set("apikey", API_KEY)
	return url.Parse(ALPHAVANTAGE_URL + "/query?" + v.Encode())
}

func MarshalQuote(b []byte) *Quote {
	quote := new(GlobalQuoteResponse)
	b = LEADING_NUMBERS.ReplaceAllLiteral(b, []byte(""))
	b = RFC_3339.ReplaceAll(b, []byte("${1}T00:00:00Z"))
	s := string(b)
	print(s)
	err := json.Unmarshal(b, quote)
	if err != nil {
		panic("failed to marshal quote")
	}
	return &quote.GlobalQuote
}

func Get(symbol string, functype ApiFuncType, options ...url.Values) []byte {
	url, err := GetUrl(symbol, functype, options...)
	if err != nil {
		log.Fatal(err)
	}
	response, err := http.Get(url.String())
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	if b, err := ioutil.ReadAll(response.Body); err != nil {
		log.Fatal(err)
	} else {
		return b
	}
	panic("We fucked up")
}

func RFQ(symbol string) (*Quote, error) {
	return nil, nil
}
