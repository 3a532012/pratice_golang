package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	ch := make(chan int)
	go func(ch chan int) {
		time.Sleep(time.Second)
		<-ch
	}(ch)
	c := time.Tick(1 * time.Second)
	for range c {
		fmt.Printf("#goroutines: %d\n", runtime.NumGoroutine())
	}
}
