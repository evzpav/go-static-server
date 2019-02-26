package main

import (
	"flag"
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

	staticFolderPath := flag.String("path", distPath, "folder path")
	route := flag.String("route", "/", "url route")
	port := flag.String("port", "9999", "port")
	flag.Parse()

	fs := http.FileServer(http.Dir(*staticFolderPath))
	http.Handle(*route, http.StripPrefix(*route, fs))
	http.HandleFunc("/*", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	log.Printf("Serving %v on http://localhost:%v%v", *staticFolderPath, *port, *route)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
