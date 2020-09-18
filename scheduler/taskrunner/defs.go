//============================================================
// 描述:
// 作者: Henry
// 日期: 2019-12-04 14:54
// 版权: henry @Since 2018
//
//============================================================
package taskrunner

const (
	READY_TO_DISPATCH = "d"
	READY_TO_EXECUTE  = "e"
	CLOSE             = "c"

	VIDEO_PATH = "./videos/"
)

type controlChan chan string

type dataChan chan interface{}

type fn func(dc dataChan) error
