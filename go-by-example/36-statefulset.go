package main

import (
	"fmt"
	"math/rand"

	"sync/atomic"
	"time"
)
type readOp struct {
	 key int
	 resp chan int
}
type writeOp struct {
	 key int
	 val int
	 resp chan bool
}
func main() {
	var readOps uint64
	var writeOps uint64

	var reads=make(chan readOp)
	var writes=make(chan writeOp)
	var state = make(map[int]int)
	go func() {
		for{
			select {
			case read:=<-reads:
				read.resp<-state[read.key]
			case write:=<-writes:
				state[write.key]=write.val
				write.resp<-true
			}
			}
		}()
	for w:=0;w<100;w++ {
		go func() {
			for{
				write:=writeOp{
					key:  rand.Intn(1000),
					val:  rand.Intn(1000),
					resp: make(chan bool),
				}
				writes<-write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	for r:=0;r<100;r++ {
		go func() {
			for {
				read:=readOp{
					key:  rand.Intn(5),
					resp: make(chan int),
				}
				reads<-read
				<-reads
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	time.Sleep(time.Second*4)
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)
	fmt.Println("State:", state)
	}




//for r:=0;r<100;r++ {
	//	go func() {
	//		total:=0
	//		for {
	//
	//		}
	//	}
	//}
//}