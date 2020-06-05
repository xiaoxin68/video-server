package handlers

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"net/http"
	"video-server/api/model"
	"video-server/api/session"
	"video-server/api/utils"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//读取r.Body中的信息并转化成对应的结构体
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &model.UserCredential{}

	if err := json.Unmarshal(res, ubody); err != nil {
		utils.SendErrorResponse(w, utils.ErrorRequestBodyParseFailed)
		return
	}

	//添加用户
	if err := AddUserCredetial(ubody.Username, ubody.Pwd); err != nil {
		utils.SendErrorResponse(w, utils.ErrorDBError)
		return
	}

	//添加session信息并插入到数据库中
	id := session.GenerateNewSessionId(ubody.Username)
	su := &model.SignedUp{Success: true, SessionId: id}
	if resp, err := json.Marshal(su); err != nil {
		utils.SendErrorResponse(w, utils.ErrorInternalFaults)
		return
	} else {
		utils.SendNormalResponse(w, string(resp), 201)
	}
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uname := p.ByName("user_name")
	io.WriteString(w, uname)
}
