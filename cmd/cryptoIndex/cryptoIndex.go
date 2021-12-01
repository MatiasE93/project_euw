package cryptoIndex

import (
	"encoding/json"
	"io/ioutil"
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

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Body: " + string(body))

	bodyJsonified, err := json.Marshal(string(body))
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(bodyJsonified, &target)
	log.Println("Body Jsonificado: " + string(bodyJsonified))

	base := target.Data.Base
	if base != "" {
		log.Println(base)
	} else {
		log.Println("Base esta vacio")
	}
}
