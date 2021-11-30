package main

import (
	"log"
	"html/template"
	"net/http"
	"os"
	"cryptoIndexPkg"
)

var view *template.Template
var port string

type pageData struct {
	Title string
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
	pd := pageData {
		Title: "Pagina Indice",
	}
	cryptoIndex.getGalaPrice()

	err := view.ExecuteTemplate(w, "index.gohtml", pd)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}