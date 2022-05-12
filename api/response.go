package main

import (
	"encoding/json"
	"io"
	"net/http"
	"video_server/api/defs"
)

func sendErrorResponse(w http.ResponseWriter, errRes defs.ErrorResponse) {
	w.WriteHeader(errRes.HttpSC)
	resStr, _ := json.Marshal(&errRes.Error)
	io.WriteString(w, string(resStr))
}

func sendNormalResponse(w http.ResponseWriter, httpSc int, resStr string) {
	w.WriteHeader(httpSc)
	io.WriteString(w, resStr)
}
