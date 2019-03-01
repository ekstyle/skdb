package main

import (
	"github.com/ekstyle/skdb/routes"
	"github.com/ekstyle/skdb/store"
	"log"
	"net/http"
)

func main() {
	store.DemoSave()
	r := routes.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", r))
}
