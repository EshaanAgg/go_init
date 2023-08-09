package api

import (
	"encoding/json"
	"eshaanagg/go/cryptomasters/data"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// We can store the format specifier in the constant itself
const apiUrl = "https://cex.io/api/ticker/%s/USD"

// We typically return pointers to response types as they can be nil when there is an error
func GetRate(currency string) (*data.Rate, error) {
	// Make the API request
	res, err := http.Get(fmt.Sprintf(apiUrl, strings.ToUpper(currency)))
	
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code recieved: %v", res.StatusCode)
	}

	// Convert the HTTP response to bytes
	bodyBtyes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// Parse the JSON and store the parsed data
	var response CEXResponse
	err = json.Unmarshal(bodyBtyes, &response)
	if err != nil {
		return nil, err
	}

	crypoRate := data.Rate{Currency: currency, Price: response.Bid}

	return &crypoRate, nil
}