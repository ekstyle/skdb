package controllers

import "net/http"

//test api
func YourHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API!\n"))
}
