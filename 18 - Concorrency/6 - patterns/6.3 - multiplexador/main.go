package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel := multi(write("Hello World"), write("Golang"))

	for range 10 {
		fmt.Println(<-channel)
	}
}

func multi(channel1, channel2 <-chan string) <-chan string {
	channelOut := make(chan string)

	go func() {
		for {
			select {
			case message := <-channel1:
				channelOut <- message
			case message := <-channel2:
				channelOut <- message
			}
		}
	}()

	return channelOut
}

func write(text string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("Value: %s", text)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return channel
}
