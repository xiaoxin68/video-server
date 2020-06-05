package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func SendErrorResponse(w http.ResponseWriter, errPesp ErrResponse) {
	w.WriteHeader(errPesp.HttpSC)
	resStr, _ := json.Marshal(&errPesp.Error)
	io.WriteString(w, string(resStr))
}

func SendNormalResponse(w http.ResponseWriter, resp string, sc int) {
	w.WriteHeader(sc)
	io.WriteString(w, resp)
}
