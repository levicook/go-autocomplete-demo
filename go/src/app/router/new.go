package router

import (
	"app/routes"
	"github.com/gorilla/mux"
	"net/http"
)

func New() (router *mux.Router) {
	router = mux.NewRouter().StrictSlash(true)

	for _, route := range routes.Routes {
		handlerFunc := timingFilter(route.Name, route.Handler)

		router.Path(route.Path).
			Methods(route.Methods).
			HandlerFunc(handlerFunc).
			Name(route.Name)
	}

	router.NotFoundHandler = timingFilter("not-found",
		http.NotFoundHandler().ServeHTTP)

	return
}
