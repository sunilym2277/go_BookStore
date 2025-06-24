package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sunilym2277/go_BookStore/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("Localhost:9010", r))

}
