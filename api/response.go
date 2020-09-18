//============================================================
// 描述:
// 作者: Henry
// 日期: 2019-12-02 15:47
// 版权: Henry @Since 2018
//
//============================================================
package main

import (
	"encoding/json"
	"github.com/StreamingVideo/api/defs"
	"io"
	"net/http"
)

func sendErrorResponse(w http.ResponseWriter, errResp defs.ErrResponse) {
	w.WriteHeader(errResp.HttpSC)

	resStr, _ := json.Marshal(&errResp.Error)
	io.WriteString(w, string(resStr))
}

func sendNormalResponse(w http.ResponseWriter, resp string, sc int) {
	w.WriteHeader(sc)
	io.WriteString(w, resp)
}
