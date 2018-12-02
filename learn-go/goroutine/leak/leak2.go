// +build OMIT

package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var ch chan int
	if false {
		ch = make(chan int, 1)
		ch <- 1
	}
	// goroutine 没有初始化，只有接受者没有发送方
	go func(ch chan int) {
		<-ch
	}(ch)
	c := time.Tick(1 * time.Second)
	for range c {
		fmt.Printf("#goroutines: %d\n", runtime.NumGoroutine())
	}
}
