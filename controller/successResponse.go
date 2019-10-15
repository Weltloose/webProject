package controller

import (
	"fmt"
	"net/http"
)

func SuccessRes(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Submit Success! You've done well.")
}
