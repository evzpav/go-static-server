package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	dir, _ := os.Getwd()
	fs := http.FileServer(http.Dir(dir + "/dist"))
	http.Handle("/", fs)
	http.HandleFunc("/*", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	log.Println("Serving " + dir)
	http.ListenAndServe(":80", nil)
}
