package router

import (
	"net/http"
)

func ServeFile(urlPath, filePath string) {
	http.HandleFunc(urlPath,
		timingFilter("serve-file", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, filePath)
		}))
}
