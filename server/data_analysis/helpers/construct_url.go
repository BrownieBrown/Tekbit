package helpers

import (
	"fmt"
	"net/url"
)

func ConstructAlphaVantageURL(params map[string]string) string {
	baseURL := "https://www.alphavantage.co/query"
	urlParams := url.Values{}
	for key, value := range params {
		urlParams.Add(key, value)
	}

	return fmt.Sprintf("%s?%s", baseURL, urlParams.Encode())
}
