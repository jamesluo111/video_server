package main

import "log"

type ConnLimiter struct {
	concurrentConn int
	bucket         chan int
}

func NewConnLimiter(c int) *ConnLimiter {
	return &ConnLimiter{
		concurrentConn: c,
		bucket:         make(chan int, c),
	}
}

func (limiter *ConnLimiter) GetConn() bool {
	if len(limiter.bucket) >= limiter.concurrentConn {
		log.Println("Reached the rate limiter")
		return false
	}
	limiter.bucket <- 1
	return true
}

func (limiter *ConnLimiter) ReleaseConn() {
	c := <-limiter.bucket
	log.Println("new conn coming", c)
}
