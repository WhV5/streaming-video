//============================================================
// 描述:
// 作者: Henry
// 日期: 2019-12-04 11:19
// 版权: Henry @Since 2018
//
//============================================================
package main

import "log"

type ConnLimiter struct {
	concurrentConn int
	bucket         chan int
}

func NewConnLimiter(cc int) *ConnLimiter {
	return &ConnLimiter{
		concurrentConn: cc,
		bucket:         make(chan int, cc),
	}
}

func (cl *ConnLimiter) GetConn() bool {
	if len(cl.bucket) >= cl.concurrentConn {
		log.Printf("Reached the rate limitation.")
		return false
	}

	cl.bucket <- 1
	return true
}

func (cl *ConnLimiter) ReleaseConn() {
	c := <-cl.bucket
	log.Printf("New connction coming:%d", c)
}
