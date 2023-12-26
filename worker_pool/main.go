package main

import (
	"fmt"
	"time"
)

func main() {
	worker := 5
	jobNumber := 10
	job := make(chan int, 10)
	result := make(chan int, 10)
	done := make(chan bool)
	initWorkPool(worker, job, result, done)
	initJob(jobNumber, job)
	close(job)
	waitJobFinish(worker, done)
	close(result)
	close(done)
	printAllResult(result)

}
func printAllResult(results chan int) {
	for result := range results {
		fmt.Println(result)
	}
}
func waitJobFinish(workerNumber int, done <-chan bool) {
	for i := 0; i < workerNumber; i++ {
		<-done
	}
}
func initWorkPool(workerNumber int, jobs <-chan int, result chan<- int, done chan<- bool) {
	for i := 0; i < workerNumber; i++ {
		go func(worker int) {
			for job := range jobs {
				val := job * 2
				result <- val
				fmt.Printf("woker: %d value: %d \n", worker, val)
				time.Sleep(2 * time.Second)
			}
			done <- true
		}(i)
	}
}

func initJob(n int, job chan<- int) {
	for i := 0; i < n; i++ {
		val := i * 10
		job <- val
	}
}
