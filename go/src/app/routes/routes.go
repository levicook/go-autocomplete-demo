package routes

import (
	"app/pages/hello"
	"net/http"
)

type (
	routes []route
	route  struct {
		Name    string
		Methods string
		Path    string
		Handler func(http.ResponseWriter, *http.Request)
	}
)

var Routes = routes{
	{
		Handler: hello.Hello,
		Methods: "GET",
		Name:    "hello",
		Path:    "/hello",
	},
}
