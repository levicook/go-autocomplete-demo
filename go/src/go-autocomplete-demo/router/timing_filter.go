package router

import (
	"kylelemons.net/go/daemon"
	"net/http"
	"time"
)

func timingFilter(route string, fn http.HandlerFunc) http.HandlerFunc {
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
