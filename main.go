package main

import (
	"faking"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/faking", faking.FakingHandle)
	log.Fatal(http.ListenAndServe(":2020", nil))
}
