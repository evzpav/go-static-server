package main

import (
	"go-static-server/statichandler"
	"log"
	"net/http"
)

var defaultPort string = "9999"

func main() {

	staticFolderPath := "frontend/dist"

	mux := http.NewServeMux()

	mux.HandleFunc("/", statichandler.StaticHandler(staticFolderPath))

	log.Printf("Serving %v on http://localhost:%v", staticFolderPath, defaultPort)
	
	log.Fatal(http.ListenAndServe(":"+defaultPort, mux))

}
