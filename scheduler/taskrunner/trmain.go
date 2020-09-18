//============================================================
// 描述:
// 作者: Henry
// 日期: 2019-12-04 14:52
// 版权: henry @Since 2018
//
//============================================================
package taskrunner

import "time"

type Worker struct {
	ticker *time.Ticker
	runner *Runner
}

func NewWorker(interval time.Duration, r *Runner) *Worker {
	return &Worker{
		ticker: time.NewTicker(interval * time.Second),
		runner: r,
	}
}

func (w *Worker) startWorker() {
	for {
		select {
		case <-w.ticker.C:
			go w.runner.StartAll()
		}
	}
}

func Start() {
	r := NewRunner(3, true,
		VideoClearDispatcher, VideoClearExcutor)
	w := NewWorker(3, r)
	go w.startWorker()
}
