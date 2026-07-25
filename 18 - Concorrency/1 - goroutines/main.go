package main

import (
	"fmt"
	"time"
)

// concurrency is the composition of independently executing processes,
// while parallelism is the simultaneous execution of (possibly related) computations.

func main() {
	go write("Hello World")
	write("Golang is awesome")
}

func write(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
