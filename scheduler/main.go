package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"video-server/scheduler/handlers"
	"video-server/scheduler/taskrunner"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.GET("/video-delete-record/:vid-id", handlers.VidDelRecHandler)

	return router
}

func main() {
	go taskrunner.Start()
	r := RegisterHandlers()
	http.ListenAndServe("localhost:9001", r)
}
