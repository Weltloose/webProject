package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Weltloose/webProject/grpcForRedis"
	"github.com/Weltloose/webProject/model"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	if model.Register(r.FormValue("name"), r.FormValue("passwd")) == true {
		fmt.Fprintln(w, "register success")
	} else {
		fmt.Fprintln(w, "not registered")
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	passwd := r.FormValue("passwd")
	if model.Login(name, passwd) == true {
		val := grpcForRedis.RPCSetAuthInfo(name, passwd, 600*1e9)
		http.SetCookie(w, &http.Cookie{
			Name:   "Authed",
			Value:  val,
			MaxAge: 600,
		})
		fmt.Fprintln(w, "login success")
	} else {
		fmt.Fprintln(w, "not logined")
	}
}

func PublichHomeworkHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	passwd := r.FormValue("passwd")
	c, err := r.Cookie("Authed")
	if err == nil {
		tname, tpasswd := grpcForRedis.RPCGetAuth(c.Value)
		if tname != "" && tpasswd != "" {
			name = tname
			passwd = tpasswd
		}
	}
	tstring := model.PublichHomework(name, passwd)
	if tstring == "" {
		fmt.Fprintln(w, "wrong username or passwd")
		return
	}
	fmt.Println(tstring)
	fmt.Fprintln(w, "The uid is: ", tstring)
}

func GetHomeworkHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	passwd := r.FormValue("passwd")
	c, err := r.Cookie("Authed")
	if err == nil {
		tname, tpasswd := grpcForRedis.RPCGetAuth(c.Value)
		if tname != "" && tpasswd != "" {
			name = tname
			passwd = tpasswd
		}
	}
	zippedFileName, ok := model.GetHomework(name, passwd, r.FormValue("uid"))
	if !ok {
		fmt.Fprintln(w, "No such homework exist")
	} else {
		data, err := ioutil.ReadFile(zippedFileName)
		if err != nil {
			fmt.Fprintln(w, "no zipped file")
		}
		fmt.Fprintln(w, string(data))
	}

}
