package main

import (
	"log"
	"html/template"
	"net/http"
	"os"
)

var view *template.Template

type pageData struct {
	Title string
	Firstname string
}

func init() {
	view = template.Must(template.ParseGlob("template/*.gohtml"))
}

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", indexPage)
	log.Print("Listening on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func indexPage(w http.ResponseWriter, r *http.Request) {
	pd := pageData {Title:"Pagina Indice"}

	err := view.ExecuteTemplate(w, "index.gohtml", pd)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}