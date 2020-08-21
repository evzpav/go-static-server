package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var defaultPort string = "9999"

func fileServerHandler(staticPath string) http.Handler {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("failed to get working directory: %v", err)
	}

	return http.FileServer(http.Dir(fmt.Sprintf("%s/%s/", dir, staticPath)))
}

func staticHandler(staticPath string) func(w http.ResponseWriter, r *http.Request) {
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

func main() {

	staticFolderPath := "frontend/dist"

	mux := http.NewServeMux()
	mux.HandleFunc("/", staticHandler(staticFolderPath))

	log.Printf("Serving %v on http://localhost:%v", staticFolderPath, defaultPort)
	log.Fatal(http.ListenAndServe(":"+defaultPort, mux))

}
