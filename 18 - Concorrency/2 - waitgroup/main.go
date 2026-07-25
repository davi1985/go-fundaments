package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitgroup sync.WaitGroup

	waitgroup.Add(4)

	go func() {
		write("Hello World")
		waitgroup.Done()
	}()

	go func() {
		write("Golang is awesome")
		waitgroup.Done()
	}()

	go func() {
		write("Goroutine 3")
		waitgroup.Done()
	}()

	go func() {
		write("Goroutine 4")
		waitgroup.Done()
	}()

	waitgroup.Wait()
}

func write(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
