# Go Static Server

Go server for SPAs(Single Page Applications) and static files with zero dependencies.

In the main.go define the variable `staticFolderPath` which is the folder of the built files (static files).  
In this example, the folder is `frontend/dist`.
It expects that there is a `index.html` file inside the folder.

```go

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


```


Run the server:
```bash
go run main.go
```
Project will be running on [http://localhost:9999](http://localhost:9999)