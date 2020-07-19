//============================================================
// 描述:
// 作者: Henry
// 日期: 2019-12-04 14:52
// 版权: henry @Since 2018
//
//============================================================
package main

import (
	"io"
	"net/http"
)

func sendResponse(w http.ResponseWriter, sc int, resp string) {
	w.WriteHeader(sc)
	io.WriteString(w, resp)
}
