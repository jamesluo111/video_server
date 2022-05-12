package main

import (
	"io"
	"net/http"
)

func sentErrorResponse(w http.ResponseWriter, httpSc int, errMsg string) {
	w.WriteHeader(httpSc)
	io.WriteString(w, errMsg)
}
