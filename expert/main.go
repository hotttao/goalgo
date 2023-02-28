package main

import (
	"fmt"
	"sync"
	"sync/atomic"

	"golang.org/x/sync/semaphore"
	"golang.org/x/sync/singleflight"
)

type T struct {
	n int
	s string
}

func (T) M1() {}
func (T) M2() {}

type NonEmptyInterface interface {
	M1()
	M2()
}

func main() {
	var t = T{
		n: 17,
		s: "hello, interface",
	}
	// var ei interface{}
	// ei = t

	var i NonEmptyInterface
	i = t
	// fmt.Println(ei)
	fmt.Println(i)
	_ = sync.Map{}
	_ = sync.Pool{}
	_ = atomic.Value{}

	atomic.AddInt32()
	_ = semaphore.NewWeighted(10)
	_ = singleflight.Group{}
}
