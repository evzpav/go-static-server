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
	log.Println("Serving " + distPath)
	log.Fatal(http.ListenAndServe(":80", nil))
}
