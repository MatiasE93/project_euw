package cryptoIndex

import (
	"net/http"
	"io/ioutil"
	"log"
)

func getGalaPrice() {
	resp, err := http.Get("https://api.coinbase.com/v2/prices/BTC-USD/spot")
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	sb := string(body)
	log.Print(sb)
}