package handlers

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"video-server/scheduler/util"
)

func VidDelRecHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid-id")

	if len(vid) == 0 {
		util.SendResponse(w, 400, "video id should not be empty")
		return
	}

	//增加删除记录
	err := AddVideoDeletionRecord(vid)
	if err != nil {
		util.SendResponse(w, 500, "Interbal server error")
		return
	}

	util.SendResponse(w, 200, "del success")
	return
}
