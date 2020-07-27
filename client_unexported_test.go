package client

import (
    "time"
    "testing"
)

func TestRegexTimeParse(t *testing.T) {
    raw := []byte("2020-01-03")
    cleaned := RFC_3339.ReplaceAllFunc(raw, apndtz)
    testTime := time.Time{}
    const shortForm = "2013-01-01"
    tme, err := time.Parse(shortForm, "2013-Feb-03")
    if err != nil {
        t.Error("shortform failed")
    }
    if tme.Year() != 2020 {
        t.Error("Failed to extract year")
    }
    err = testTime.UnmarshalJSON(cleaned)
    if err != nil {
        t.Errorf("Cannot unmarshal time %s", cleaned)
    }
}


/*
func TestMapRawResponseGQ(t *testing.T) {
	raw := []byte(`{
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

	rawMap, err := mapRawResponse(raw)
	if err != nil {
		t.Error(err)
	}
	if rawMap["01. symbol"] != "MSFT" {
		t.Error("Failed to map")
	}
}
*/

/*
func TestMapRawResponseTS(t *testing.T) {
	raw := []byte(`{
	       Meta Data": {
	           "1. Information": "Monthly Prices (open, high, low, close) and Volumes",
	           "2. Symbol": "MSFT",
	           "3. Last Refreshed": "2020-01-06",
	           "4. Time Zone": "US/Eastern"
	       },
	       "Monthly Time Series": {
	           "2020-01-06": {
	               "1. open": "158.7800",
	               "2. high": "160.7300",
	               "3. low": "156.5100",
	               "4. close": "159.0300",
	               "5. volume": "62970547"
	           }
	       }
    }`)

	rawMap, err := mapRawResponse(raw)
	if err != nil {
		t.Error(err)
	}
	if rawMap["01. symbol"] != "MSFT" {
	    t.Error("Failed to map")
    }
}
*/
