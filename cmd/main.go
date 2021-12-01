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
	Title     string
	Firstname string
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
	galaPrice := *new(cryptoIndex.CryptoJson)

	pd := pageData{
		Title: "Pagina Indice",
	}
	price := cryptoIndex.GetGalaPriceInDollars(galaPrice)
	log.Println(price)

	err := view.ExecuteTemplate(w, "index.gohtml", pd)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
