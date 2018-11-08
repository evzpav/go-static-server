package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	distPath := dir + "/dist"
	fs := http.FileServer(http.Dir(distPath))
	http.Handle("/", fs)
	http.HandleFunc("/*", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	port := "9999"
	log.Printf("Serving %s on http://localhost:%s", distPath, port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
