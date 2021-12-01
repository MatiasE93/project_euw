package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/MatiasE93/project_euw/cmd/cryptoIndex"
)

var view *template.Template
var port string

type pageData struct {
	Crypto string
	Valor  string
	Moneda string
}

func init() {
	view = template.Must(template.ParseGlob("templates/*.gohtml"))
	port = os.Getenv("PORT")
}

func main() {
	http.HandleFunc("/", indexPage)
	log.Print("Listening on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func indexPage(w http.ResponseWriter, r *http.Request) {
	var galaPrice cryptoIndex.CryptoJson
	galaPrice = cryptoIndex.GetGalaPriceInDollars(galaPrice)

	pd := pageData{
		Crypto: galaPrice.Data.Base,
		Valor:  galaPrice.Data.Amount,
		Moneda: galaPrice.Data.Currency,
	}

	err := view.ExecuteTemplate(w, "index.gohtml", pd)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
