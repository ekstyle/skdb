package main

import (
	"github.com/ekstyle/skdb/routes"
	"log"
	"net/http"
)

func main() {
	r := routes.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", r))
}
