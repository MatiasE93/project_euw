package cryptoIndex

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type CryptoJson struct {
	Data CryptoJsonData
}

type CryptoJsonData struct {
	Base     string
	Currency string
	Amount   string
}

var client = &http.Client{Timeout: 10 * time.Second}

func GetGalaPriceInDollars(target CryptoJson) {
	resp, err := client.Get("https://api.coinbase.com/v2/prices/GALA-USD/spot")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(target)

	base := target.Data.Base
	if base != "" {
		log.Println(base)
	} else {
		log.Println("Base esta vacio")
	}

	currency := target.Data.Currency
	if currency != "" {
		log.Println(currency)
	} else {
		log.Println("Currency esta vacio")
	}

	amount := target.Data.Amount
	if amount != "" {
		log.Println(amount)
	} else {
		log.Println("Amount esta vacio")
	}
}
