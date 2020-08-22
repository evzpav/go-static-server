package statichandler

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func fileServerHandler(staticPath string) http.Handler {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("failed to get working directory: %v", err)
	}

	return http.FileServer(http.Dir(fmt.Sprintf("%s/%s/", dir, staticPath)))
}

// StaticHandler is the handler to serve spa static files
func StaticHandler(staticPath string) func(w http.ResponseWriter, r *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		path, err := filepath.Abs(r.URL.Path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		path = filepath.Join(staticPath, path)

		_, err = os.Stat(path)
		if os.IsNotExist(err) {
			http.ServeFile(w, r, filepath.Join(staticPath, "index.html"))
			return
		} else if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fileServerHandler(staticPath).ServeHTTP(w, r)
	})
}
