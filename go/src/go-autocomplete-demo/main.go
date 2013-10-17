package main

import (
	"flag"
	"github.com/gorilla/mux"
	"go-autocomplete-demo/routes"
	"kylelemons.net/go/daemon"
	"log"
	"net/http"
	"time"
)

var (
	fork     = daemon.ForkPIDFlags("fork", "pid", "/tmp/go-autocomplete-demo.pid")
	logLevel = daemon.LogLevelFlag("log-level")
)

func init() {
	timingFilter := func(route string, fn http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			defer func(start time.Time) {
				if err := recover(); err != nil {
					daemon.Error.Printf("%s %s | %s %s | err: %v", r.Method, r.RequestURI, route, time.Since(start), err)
					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				} else {
					daemon.Info.Printf("%s %s | %s %s", r.Method, r.RequestURI, route, time.Since(start))
				}
			}(time.Now())

			fn(w, r)
		}
	}

	router := mux.NewRouter().StrictSlash(true)
	http.Handle("/", router)

	for _, route := range routes.Routes {
		router.Path(route.Path).
			Methods(route.Methods).
			HandlerFunc(timingFilter(route.Name, route.Handler)).
			Name(route.Name)
	}
}

func main() {
	flag.Parse()
	fork.Fork()
	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatal(err)
		}
	}()
	daemon.Run()
}
