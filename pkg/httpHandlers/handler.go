package httpHandlers

import "github.com/gorilla/mux"

func InitHandlers() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", welcomeHandler()).Methods("GET")
	return router
}