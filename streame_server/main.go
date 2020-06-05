package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"video-server/streame_server/handlers"
	"video-server/streame_server/token_bucket"
	"video-server/streame_server/util"
)

type middleWareHandler struct {
	r *httprouter.Router
	l *token_bucket.ConnLimiter
}

func NewMiddleWareHandler(r *httprouter.Router, cc int) http.Handler {
	m := middleWareHandler{}
	m.r = r
	m.l = token_bucket.NewConnLimiter(cc)
	return m
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.GET("/videos/:vid-id", handlers.StreamHandler)

	router.POST("/upload/:vid-id", handlers.UploadHandler)

	return router
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !m.l.GetConn() {
		util.SendErrorResponse(w, http.StatusTooManyRequests, "Too many requests")
		return
	}

	m.r.ServeHTTP(w, r)
	defer m.l.ReleaseConn()
}

func main() {
	r := RegisterHandlers()
	mh := NewMiddleWareHandler(r, 2)
	http.ListenAndServe("localhost:9000", mh)
}
