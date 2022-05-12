package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"video_server/scheduler/dbops"
)

func DeleteTaskVideo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("del_vid")
	err := dbops.TaskAddVideoDelete(vid)
	if err != nil {
		ResponseInternal(w, http.StatusInternalServerError, "db error")
		return
	}
	ResponseInternal(w, http.StatusOK, "successfully!")
	return

}
