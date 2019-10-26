package model

import (
	"fmt"
	"os"

	"github.com/Weltloose/webProject/data"
	"github.com/Weltloose/webProject/lib"
	"github.com/google/uuid"
)

func Register(name, password string) bool {
	return data.Register(name, password)
}

func Login(name, password string) bool {
	return data.Login(name, password)
}

func PublichHomework(name, passwd string) string {

	tuid := data.PublichHomework(name, passwd)
	if tuid == "invalid" {
		return ""
	}
	err := os.Mkdir(lib.GetHomeworkPath(tuid), os.ModePerm)
	if err != nil {
		fmt.Println("not create dir, ", err)
	}
	return tuid
}

func GetHomework(name, passwd, uid string) (string, bool) {
	ok := data.GetHomework(name, passwd, uid)
	if !ok {
		return "", false
	}
	dirName := lib.GetHomeworkPath(uid)
	zippedName := lib.GetZippedFilePath(uuid.New().String())
	lib.Zipit(dirName, zippedName)
	return zippedName, true
}
