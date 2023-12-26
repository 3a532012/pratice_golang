package main

import (
	"context"
	"fmt"
)

func main() {

	wokerNumber := 5
	jobNumber := 20
	job := make(chan int, jobNumber)
	result := make(chan int, jobNumber)
	done := make(chan bool)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	initWokrer(ctx, cancel, wokerNumber, job, result, done)
	initJob(jobNumber, job)
	close(job)
	waitJobFinish(wokerNumber, done)
	close(result)
	close(done)
}
func waitJobFinish(wokerNumber int, done <-chan bool) {
	for i := 0; i < wokerNumber; i++ {
		<-done
	}
}
func initJob(jobNumber int, job chan<- int) {
	for i := 0; i < jobNumber; i++ {
		val := i * 10
		job <- val
	}
}

func initWokrer(c context.Context, cancel context.CancelFunc, wokerNum int, job <-chan int, result chan<- int, done chan<- bool) {
	for i := 0; i < wokerNum; i++ {
		go func(wokerID int) {
			for {
				select {
				case jobValue, ok := <-job:
					if !ok {
						done <- true
						return
					}
					if jobValue%7 == 0 {
						fmt.Printf("worker: %d exit \n", wokerID)
						done <- true
						cancel()
						return
					}
					val := jobValue * 2
					result <- val
					fmt.Printf("worker: %d val: %d \n", wokerID, val)
				case <-c.Done():
					done <- true
					return
				}
			}
		}(i)
	}
}
