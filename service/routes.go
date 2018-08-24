package service

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     String
	handlerFunc http.HandleFunc
}

type Routes []Route

var routes = Routes {

	Route {
		"GetAccount",
		"GET",
		"/accounts/{accountId}",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.Write([]byte("{\"result\" : \"Ok\"}"))
		}
	}
}