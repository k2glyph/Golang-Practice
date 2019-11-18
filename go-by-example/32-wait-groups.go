package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	fmt.Println("Started jobs of id", id)
	time.Sleep(time.Second)
	fmt.Println("Finished jobs of id", id)
	wg.Done()
}
func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	for i := 1; i <= 5; i++ {
		//wg.Add(1)
		//go worker(i, &wg)
	}

	wg.Wait()
	fmt.Println("Waiting for all threads to finish")

}