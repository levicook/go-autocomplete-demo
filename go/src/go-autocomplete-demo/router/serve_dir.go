package router

import (
	"net/http"
)

func ServeDir(urlPath, dirPath string) {
	// TODO wrap with timingFilter
	http.Handle(urlPath, http.StripPrefix(urlPath, http.FileServer(http.Dir(dirPath))))
}
