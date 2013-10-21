package main

import (
	"app/router"
	"flag"
	"kylelemons.net/go/daemon"
	"log"
	"net/http"
)

var (
	fork     = daemon.ForkPIDFlags("fork", "pid", "go-autocomplete-demo.pid")
	logLevel = daemon.LogLevelFlag("log-level")
)

func main() {
	flag.Parse()
	fork.Fork()

	go func() {
		http.Handle("/", router.New())

		router.ServeDir("/js/", "js/pkg/")

		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatal(err)
		}
	}()

	daemon.Run()
}
