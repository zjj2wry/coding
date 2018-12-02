package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("do somethin one")
		wg.Done()
	}()

	go func() {
		fmt.Println("do somethin two")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("done")
}
