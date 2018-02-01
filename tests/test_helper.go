package tests

import (
	"net/http"
)

//CheckPageResponse checks if a page that should respond is found correctly
func CheckPageResponse(url string) bool {
	response, err := http.Get(url)
	if err != nil {
		return false
	}
	if response == nil {
		return false
	}
	if response.Status == "404 Not Found" {
		return false
	}
	return true
}

//CheckNoPageResponse checks if a page that does not exist responds with a 404 Error
func CheckNoPageResponse(url string) bool {
	response, err := http.Get(url)
	if err != nil {
		return true
	}
	if response == nil {
		return true
	}
	if response.Status == "404 Not Found" {
		return true
	}
	return false
}
