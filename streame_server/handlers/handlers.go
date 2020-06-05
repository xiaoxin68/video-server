package handlers

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
	"video-server/streame_server/config"
	"video-server/streame_server/util"
)

//读取并传输视频的二进制流
func StreamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid-id")
	vl := config.VIDEO_DIR + vid

	//打开文件
	video, err := os.Open(vl)
	if err != nil {
		log.Printf("Error when try to open file: %v", err)
		util.SendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
		return
	}

	//设置请求头信息
	w.Header().Set("Content-Type", "video/mp4")
	//转化成二进制流，传输到http
	http.ServeContent(w, r, "", time.Now(), video)
	//由于文件打开了，故必须关闭
	defer video.Close()
}

//上传
func UploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//获取文件内容
	r.Body = http.MaxBytesReader(w, r.Body, config.MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(config.MAX_UPLOAD_SIZE); err != nil {
		util.SendErrorResponse(w, http.StatusBadRequest, "File is too big")
		return
	}

	//转化为file
	file, _, err := r.FormFile("file")
	if err != nil {
		log.Printf("Error when try to get file: %v", err)
		util.SendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
		return
	}

	//读取文件
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Read file error: %v", err)
		util.SendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
	}

	//写文件
	fn := p.ByName("vid-id")
	err = ioutil.WriteFile(config.VIDEO_DIR+fn, data, 0666)
	if err != nil {
		log.Printf("Write file error: %v", err)
		util.SendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
		return
	}

	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "Uploaded successfully")
}
