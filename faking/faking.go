package faking

import (
	"fmt"
	"net/http"
	"puking"
)

func FakingHandle(response http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response, "faking men")
	puking.Damn()
	// demn.silentDenmt
}
