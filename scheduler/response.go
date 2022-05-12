package main

import (
	"io"
	"net/http"
)

func ResponseInternal(w http.ResponseWriter, httpSc int, errMsg string) {
	w.WriteHeader(httpSc)
	io.WriteString(w, errMsg)
}
