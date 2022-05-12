package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"video_server/scheduler/taskrunner"
)

func RegisterHandlers() *httprouter.Router {
	r := httprouter.New()
	r.GET("/del/:del_vid", DeleteTaskVideo)
	return r
}

func main() {
	go taskrunner.Start()
	r := RegisterHandlers()
	http.ListenAndServe(":8001", r)
}
