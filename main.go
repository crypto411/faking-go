package main

import (
	"log"
	"main/faking"
	"net/http"
)

func main() {
	http.HandleFunc("/faking", faking.FakingHandle)
	log.Fatal(http.ListenAndServe(":2020", nil))
}
