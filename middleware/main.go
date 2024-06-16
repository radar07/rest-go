package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Executing middleware before request phase")
		handler.ServeHTTP(w, r)
		fmt.Println("Executing middleware after response phase")
	})
}

func mainLogic(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Executing main handler...")
	w.Write([]byte("Ok"))
}

func main() {
	mainLogicHandler := http.HandlerFunc(mainLogic)
	http.Handle("/", middleware(mainLogicHandler))

	srvr := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	log.Fatal(srvr.ListenAndServe())
}
