package main

import (
	"net/http"

	"github.com/Weltloose/webProject/controller"
)

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/loginPage", controller.PageHandler("pages/login.html"))
	http.HandleFunc("/uploadPage", controller.PageHandler("pages/upload.html"))
	http.HandleFunc("/registerPage", controller.PageHandler("pages/register.html"))
	http.HandleFunc("/publishHomeworkPage", controller.PageHandler("pages/publishHomework.html"))
	http.HandleFunc("/getHomeworkPage", controller.PageHandler("pages/getHomework.html"))
	http.HandleFunc("/upload", controller.UploadHomework)
	http.HandleFunc("/login", controller.LoginHandler)
	http.HandleFunc("/register", controller.RegisterHandler)
	http.HandleFunc("/publishHomework", controller.PublichHomeworkHandler)
	http.HandleFunc("/getHomework", controller.GetHomeworkHandler)

	server.ListenAndServe()
}
