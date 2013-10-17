package main

import (
	"flag"
	"go-autocomplete-demo/router"
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

		//	serveDir("/app-js/", "./app-js/")
		//	serveDir("/vendor/", "./vendor/")

		//	if steal.Env == "development" {
		//		serveDir("/can/", "./can/")
		//		serveDir("/funcunit/", "./funcunit/")
		//		serveDir("/steal/", "./steal/")
		//	}

		//	serveFile("/steal.js", "./vendor/steal/steal.js")
		//	serveFile("/steal.production.js", "./vendor/steal/steal.production.js")
		//	serveFile("/stealconfig.js", "./stealconfig.js")

		http.Handle("/", router.New())

		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatal(err)
		}
	}()

	daemon.Run()
}
