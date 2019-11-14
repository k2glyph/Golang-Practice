package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("Received jobs: ", j)
			} else {
				fmt.Println("Received all jobs")
				done <- true
				return
			}
		}
	}()
	for j := 1; j <= 2; j++ {
		jobs <- j
		fmt.Println("Sends job: ", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")
	<-done
}
