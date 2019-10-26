package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/Weltloose/webProject/lib"
)

// Upload file
func UploadHomework(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024)
	fileHeader := r.MultipartForm.File["homework"][0]
	file, err := fileHeader.Open()
	if err != nil {
		fmt.Fprintln(w, "not take file")
	}
	StudentId := r.FormValue("StudentId")
	uuid := r.FormValue("uuid")
	if !CheckoutStudentid(StudentId) {
		fmt.Fprintln(w, "Invalid studentid, try again!")
		return
	}
	if err != nil {
		fmt.Fprintln(w, "No file detected!")
		return
	} else {

		data, err := ioutil.ReadAll(file)
		if err == nil {
			// fmt.Fprintln(w, string(data)) // 实操，这里只能用println才能正确的返回，使用printf会导致传回去的文件出错
			err = ioutil.WriteFile(string(lib.GetHomeworkPath(uuid)+"/"+fileHeader.Filename), data, 0644)
			// fmt.Println(string("./homeworks/" + fileHeader.Filename))
			if err != nil {
				fmt.Fprintln(w, "File not savedable!")
				return
			}
		}

	}
	fmt.Fprintf(w, "Submit Success! You've done well.")
}

// checkout studentid
func CheckoutStudentid(id string) bool {

	if val, err := strconv.ParseInt(id, 10, 64); err == nil {
		if val >= 10000000 && val < 100000000 {
			return true
		}
	}
	return false
}
