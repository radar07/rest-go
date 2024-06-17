package main

import (
	"log"

	"github.com/radar07/rest-go/postgres/models"
)

func main() {
	_, err := models.InitDB()
	if err != nil {
		log.Println(err)
	}
}
