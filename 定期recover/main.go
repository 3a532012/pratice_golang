package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTicker(1 * time.Second)
	go func() {
		for {
			select {
			case <-timer.C:
				go func() {
					defer func() {
						if err := recover(); err != nil {
							fmt.Println(err)
						}
					}()
					proc()
				}()
			}
		}
	}()
	select {}
}

func proc() {
	panic("ok")
}
