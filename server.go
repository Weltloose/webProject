package main

import (
	"net/http"

	"github.com/weltloose/webProject/controller"
)

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/upload", controller.UploadHomework)
	server.ListenAndServe()
}
