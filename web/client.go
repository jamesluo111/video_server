package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

var httpClient *http.Client

func init() {
	httpClient = &http.Client{}

}

func request(b *ApiBody, w http.ResponseWriter, r *http.Request) {
	var resp *http.Response
	var err error
	switch b.Method {
	case http.MethodGet:
		log.Println("jinlaila==============")
		log.Println(b.Url)
		req, _ := http.NewRequest("GET", b.Url, nil)
		req.Header = r.Header
		resp, err = httpClient.Do(req)
		if err != nil {
			log.Printf("err: %v", err)
			return
		}
		normalResponse(w, resp)
	case http.MethodPost:
		req, _ := http.NewRequest("POST", b.Url, bytes.NewBuffer([]byte(b.ReqBody)))
		req.Header = r.Header
		resp, err = httpClient.Do(req)
		if err != nil {
			log.Printf("err: %v", err)
			return
		}
		normalResponse(w, resp)
	case http.MethodDelete:
		req, _ := http.NewRequest("DELETE", b.Url, nil)
		req.Header = r.Header
		resp, err = httpClient.Do(req)
		if err != nil {
			log.Printf("err: %v", err)
			return
		}
		normalResponse(w, resp)
	default:
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "bad api request")
	}
}

func normalResponse(w http.ResponseWriter, res *http.Response) {
	rsp, err := ioutil.ReadAll(res.Body)
	if err != nil {
		re, _ := json.Marshal(IternalError)
		w.WriteHeader(500)
		io.WriteString(w, string(re))
		return
	}
	w.WriteHeader(res.StatusCode)
	io.WriteString(w, string(rsp))
}
