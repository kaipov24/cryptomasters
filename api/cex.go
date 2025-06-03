package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/kaipov24/cryptomsters/datatypes"
)

const apiUrl = "https://cex.io/api/ticker/%s/USD"

func GetRate(currencey string) (*datatypes.Rate, error) {
	if len(currencey) != 3 {
		return nil, fmt.Errorf("3 characters required")
	}
	upperCurrency := strings.ToUpper(currencey)
	res, err := http.Get(fmt.Sprintf(apiUrl, upperCurrency))
	if err != nil {
		return nil, err
	}
	var response Response
	if res.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(bodyBytes, &response)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("status code received: %v", res.StatusCode)
	}
	rate := datatypes.Rate{Currency: currencey, Price: response.Bid}
	return &rate, nil
}
