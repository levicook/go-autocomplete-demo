package router

import (
	"net/http"
	"path"
	"strings"
)

func ServeDir(urlPath, fromDir string) {
	http.HandleFunc(urlPath,
		timingFilter("serve-dir", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, path.Join(fromDir, strings.Split(r.URL.Path, urlPath)[1]))
		}))
}
