package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs<- chan int, result chan<- int) {
	for j:= range jobs {
		fmt.Println("worker:",id, "job", j, "Started")
		fmt.Println("")
		time.Sleep(time.Second)
		fmt.Println("worker:",id, "job", j, "Finished")
		result<-j
	}
}
func main() {
	jobs:=make(chan int,100)
	results:=make(chan int, 100)

	for i:=1;i<=3;i++ {
		go worker(i,jobs, results)
	}
	for i:=1;i<=5;i++ {
		jobs <- i
	}
	close(jobs)
	for i:=1;i<=5;i++ {
		<-results
	}
}