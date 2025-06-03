package api

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/kaipov24/cryptomsters/datatypes"
)

const apiUrl = "https://cex.io/api/ticker/%s/USD"

func GetRate(currencey string) (*datatypes.Rate, error) {
	upperCurrency := strings.ToUpper(currencey)
	res, err := http.Get(fmt.Sprintf(apiUrl, upperCurrency))
	if err != nil {
		return nil, err
	}
	if res.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		json := string(bodyBytes)
		fmt.Println(json)
	} else {
		return nil, fmt.Errorf("status code received: %v", res.StatusCode)
	}
	rate := datatypes.Rate{Currency: currencey, Price: 20}
	return &rate, nil
}
