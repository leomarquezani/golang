package service

import(
	"github.com/gorrila/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			handler(route.HandlerFunc)
	}
	return router
	)
}
