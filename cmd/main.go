package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/MatiasE93/project_euw/pkg/httpHandlers"
)

func main() {
	fmt.Println("Starting server on port 8080...")
	router := http.InitHandlers()
	log.Fatal(http.ListenAndServe(":8080", router))
}