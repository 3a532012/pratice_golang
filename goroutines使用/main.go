// 写代码实现两个 goroutine，其中一个产生随机数并写入到 go channel 中，另外一个从 channel 中读取数字并打印到标准输出。最终输出五个随机数。
package main

import (
	"fmt"
)

func main() {
	done := make(chan struct{})
	number := make(chan int, 100)
	go func() {
		for i := 0; i < 100; i++ {
			number <- i
		}
		close(number)
	}()
	go func() {
		for n := range number {
			fmt.Println(n)
		}
		done <- struct{}{}
	}()

	<-done
}
