//============================================================
// 描述:
// 作者: Henry
// 日期: 2019-12-04 11:11
// 版权: Henry @Since 2018
//
//============================================================
package main

import (
	"io"
	"net/http"
)

func sendErrorResponse(w http.ResponseWriter, sc int, errMsg string) {
	w.WriteHeader(sc)
	io.WriteString(w, errMsg)
}
