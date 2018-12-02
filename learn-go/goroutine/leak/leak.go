// +build OMIT

package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func query() int {
	n := rand.Intn(100)
	time.Sleep(time.Duration(n) * time.Millisecond)
	return n
}
func queryAll() int {
	ch := make(chan int)
	go func() { ch <- query() }()
	go func() { ch <- query() }()
	go func() { ch <- query() }()
	return <-ch
}

func main() {
	for i := 0; i < 4; i++ {
		// 每次查询有 3 个值进入 goroutine 但是只有一个接受
		queryAll()
		fmt.Printf("#goroutines: %d\n", runtime.NumGoroutine())
	}
}

// #goroutines: 3
// #goroutines: 5
// #goroutines: 7
// #goroutines: 9
