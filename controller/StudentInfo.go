package controller

import "strconv"

// checkout studentid
func Checkout_studentid(id string) bool {

	if val, err := strconv.ParseInt(id, 10, 64); err == nil {
		if val >= 10000000 && val < 100000000 {
			return true
		}
	}
	return false
}
