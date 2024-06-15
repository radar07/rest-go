package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

type CustomServeMux struct{}

func (p *CustomServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		giveRandom(w, r)
		return
	}
	http.NotFound(w, r)
}

func giveRandom(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Random number is %f", rand.Float64())
}

func main() {
	// mux := &CustomServeMux{}

	newMux := http.NewServeMux()

	newMux.HandleFunc("/randomFloat", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "random float is %f", rand.Float64())
	})

	newMux.HandleFunc("/randomInt", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "random int is %d", rand.Int())
	})
	http.ListenAndServe(":8080", newMux)
}
