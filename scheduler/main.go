//============================================================
// 描述:
// 作者: Henry
// 日期: 2019-12-04 14:52
// 版权: henry @Since 2018
//
//============================================================
package main

import (
	"github.com/StreamingVideo/scheduler/taskrunner"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.GET("/video-delete-record/:vid-id", vidDelRecHandler)

	return router
}

func main() {
	//c := make(chan int)
	go taskrunner.Start()
	r := RegisterHandlers()
	http.ListenAndServe(":9001", r)
	//<- c
}
