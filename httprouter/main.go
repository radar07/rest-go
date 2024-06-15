package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func helloWorld(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	router := httprouter.New()

	router.GET("/hello", helloWorld)
	log.Fatal(http.ListenAndServe(":8080", router))
}
