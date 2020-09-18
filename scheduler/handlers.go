//============================================================
// 描述:
// 作者: Henry
// 日期: 2019-12-04 14:52
// 版权: henry @Since 2018
//
//============================================================
package main

import (
	"github.com/StreamingVideo/scheduler/dbops"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func vidDelRecHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid-id")

	if len(vid) == 0 {
		sendResponse(w, 400, "video id should not be empty.")
		return
	}
	err := dbops.AddVideoDeletionRecord(vid)
	if err != nil {
		sendResponse(w, 500, "Internal server")
		return
	}

	sendResponse(w, 200, "")
	return
}
