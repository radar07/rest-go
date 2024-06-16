package main

import (
	jsonparse "encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
)

type Args struct {
	ID string
}

type Book struct {
	ID     string `"json:string,omitempty"`
	Name   string `"json:name,omitempty"`
	Author string `"json:author,omitempty"`
}

type JSONServer struct{}

func (t *JSONServer) GiveBookDetail(r *http.Request, args *Args, reply *Book) error {
	var books []Book

	raw, err := ioutil.ReadFile("./books.json")
	if err != nil {
		panic(err)
	}

	marshallErr := jsonparse.Unmarshal(raw, &books)
	if marshallErr != nil {
		panic(err)
	}

	for _, book := range books {
		if book.ID == args.ID {
			*reply = book
			break
		}
	}

	return nil
}

func main() {
	s := rpc.NewServer()
	s.RegisterCodec(json.NewCodec(), "application/json")

	s.RegisterService(new(JSONServer), "")
	r := mux.NewRouter()
	r.Handle("/rpc", s)
	http.ListenAndServe(":1234", r)
}
