package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func PageHandler(filePath string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if data, err := ioutil.ReadFile(filePath); err == nil {
			fmt.Fprint(w, string(data))
		}
	})
}
