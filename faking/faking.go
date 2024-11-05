package faking

import (
	"fmt"
	"main/puking"
	"net/http"
)

func FakingHandle(response http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response, "faking men")
	puking.Damn()
	// demn.silentDenmt
}

func FakingPOSTHandle(response http.ResponseWriter, request *http.Request) {
	fmt.Println("POST GES")
}
