package cryptoIndex

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type CryptoJsonData struct {
	Base     string
	Currency string
	Amount   string
}

type CryptoJson struct {
	Data CryptoJsonData
}

var client = &http.Client{Timeout: 10 * time.Second}

func GetGalaPriceInDollars(target CryptoJson) string {
	resp, err := client.Get("https://api.coinbase.com/v2/prices/GALA-USD/spot")
	if err != nil {
		log.Fatal(err)
		return ""
	}
	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(target)
	return target.Data.Amount
}
