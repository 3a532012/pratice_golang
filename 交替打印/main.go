package main

import (
	"fmt"
)

func main() {
	done := make(chan struct{})
	letter := make(chan struct{})
	number := make(chan struct{})
	go func() {
		j := 'A'
		for {
			select {
			case <-letter:
				if j >= 'Z' {
					done <- struct{}{}
					return
				} else {
					fmt.Print(string(j))
					j++
					fmt.Print(string(j))
					j++
					number <- struct{}{}
				}
			}
		}
	}()
	go func() {
		i := 1
		for {
			select {
			case <-number:
				fmt.Printf("%d", i)
				i++
				fmt.Printf("%d", i)
				i++
				letter <- struct{}{}
			}
		}
	}()

	number <- struct{}{}
	<-done
}
