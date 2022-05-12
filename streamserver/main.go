package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type middleware struct {
	r     *httprouter.Router
	limit *ConnLimiter
}

func NewMiddleware(r *httprouter.Router, limit int) http.Handler {
	m := middleware{}
	connLimit := NewConnLimiter(limit)
	m.r = r
	m.limit = connLimit
	return m
}

func (midd middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !midd.limit.GetConn() {
		sentErrorResponse(w, http.StatusTooManyRequests, "too many token")
		return
	}
	midd.r.ServeHTTP(w, r)
	defer midd.limit.ReleaseConn()
}

func RegisterHandlers() *httprouter.Router {
	r := httprouter.New()
	r.GET("/video/:vid-id", steamHandler)
	r.POST("/video/:vid-id", uploadHandler)
	r.GET("/testpage", testPageHandler)
	return r
}

func main() {
	r := RegisterHandlers()
	mh := NewMiddleware(r, 2)
	http.ListenAndServe(":8002", mh)
}
