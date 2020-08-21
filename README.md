# Go Static Server

Go server for SPAs(Single Page Applications) and static files with zero dependencies.

In the main.go define the variable `staticFolderPath` which is the folder of the built files (static files).  
In this example, the folder is `frontend/dist`.
It expects that there is a `index.html` file inside the folder.

Run the server:
```bash
go run main.go
```
Project will be running on [http:/localhost:9999](http:/localhost:9999)