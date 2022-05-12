package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"video_server/api/session"
)

type middleWareHandler struct {
	r *httprouter.Router
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := middleWareHandler{}
	m.r = r
	return m
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//checkoutSession
	validateUserSession(r)
	m.r.ServeHTTP(w, r)
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.POST("/createuser", CreateUser)
	router.POST("/user/:user_name", Login)
	router.GET("/user/:user_name", GetUserInfo)
	router.POST("/user/:user_name/videos", AddNewVideos)
	router.GET("/user/:user_name/videos", ListAllVideos)
	router.DELETE("/user/:user_name/videos/:vid-id", DeleteVideos)
	router.POST("/videos/:vid-id/comments", PostComments)
	router.GET("/videos/:vid-id/comments", ShowComments)
	return router
}

func Prepare() {
	session.LoadSessionsFromDB()
}

func main() {
	Prepare()
	r := RegisterHandlers()
	mh := NewMiddleWareHandler(r)
	http.ListenAndServe(":8000", mh)
}
