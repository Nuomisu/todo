package main

import (
	"log"
	"net/http"
)

func main() {

	router := NewRouter()
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
	log.Fatal(http.ListenAndServe(":8080", router))
}
