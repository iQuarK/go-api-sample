package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/iquark/go-api-sample/handlers"
)

func main() {
	router := mux.NewRouter()
	sub := router.PathPrefix("/api/v1").Subrouter()
	sub.Methods("GET").Path("/teams/{id:[0-9]+}").HandlerFunc(handlers.GetTeamsHandler)
	sub.Methods("GET").Path("/members/{id:[0-9]+}").HandlerFunc(handlers.GetMembersHandler)

	fmt.Println("Starting server...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
