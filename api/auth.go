//============================================================
// 描述:
// 作者: Henry
// 日期: 2019-12-03 17:58
// 版权: Henry @Since 2018
//
//============================================================
package main

import (
	"github.com/StreamingVideo/api/defs"
	"github.com/StreamingVideo/api/session"
	"net/http"
)

var HEADER_FIELD_SESSION = "X-Session-Id"
var HEADER_FIELD_UNAME = "x-User-Name"

func validateUserSession(r *http.Request) bool {
	sid := r.Header.Get(HEADER_FIELD_SESSION)
	if len(sid) == 0 {
		return false
	}
	uName, ok := session.IsSessionExpired(sid)
	if ok {
		return false
	}
	r.Header.Add(HEADER_FIELD_UNAME, uName)
	return true
}

func ValidateUser(w http.ResponseWriter, r *http.Request) bool {
	uName := r.Header.Get(HEADER_FIELD_UNAME)
	if len(uName) == 0 {
		sendErrorResponse(w, defs.ErrorNotAuthUser)
		return false
	}
	return true
}
