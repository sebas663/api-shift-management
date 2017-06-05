package main

import (
	"log"
	"net/http"
)

func main() {

	router := APIRouter()

	log.Fatal(http.ListenAndServe(":4200", router))
}
