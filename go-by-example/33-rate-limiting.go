package main

import (
	"fmt"
	"time"
)
/*
	Rate limiting is an important mechanism for controlling resource utilization and maintaining quality of service.
	Go elegantly supports rate limiting with goroutines, channels, and tickers.
*/
func main() {
	//requests:=make(chan int, 5)
	//for i:=1;i<=5;i++ {
	//	requests<-i
	//}
	//close(requests)
	////limiter:=time.Tick(200*time.Millisecond)
	//for req:=range requests {
	//	//<-limiter
	//	//time.Sleep(time.Millisecond*200)
	//	fmt.Println("request", req, time.Now())
	//}
	burstyLimiter:=make(chan time.Time,3)
	//
	for i:=0;i<3;i++ {
		burstyLimiter<-time.Now()
	}
	go func() {
		for t:=range time.Tick(time.Second*2) {
			burstyLimiter<-t
		}
	}()
	burstyRequests:=make(chan int ,6)
	for i:=1;i<=5;i++ {
		burstyRequests<-i
	}
	close(burstyRequests)
	for req:= range burstyRequests {
		<-burstyLimiter
		fmt.Println("requests", req, time.Now())
	}
}