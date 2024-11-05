package main

import (
	"log"
	"main/faking"
	"net/http"
)

func fakingMiddleware(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "POST":
		faking.FakingPOSTHandle(response, request)
	default:
		faking.FakingHandle(response, request)
	}
}

func main() {
	http.HandleFunc("/faking", fakingMiddleware)
	log.Fatal(http.ListenAndServe(":2020", nil))
}
