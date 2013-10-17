package router

import (
	"net/http"
)

func ServeFile(urlPath, filePath string) {
	timingFilter(filePath, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filePath)
	})
}
