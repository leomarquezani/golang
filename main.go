package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"log"
)

func main(){

	router := mux.NewRouter()
	log.Fatal(http.ListenAndServe(":8000", router))
}
