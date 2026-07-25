package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	channel := make(chan string)

	go write("Hello World", channel)

	fmt.Println("After goroutine")

	// for {
	// 	message, isOpen := <-channel
	// 	if !isOpen {
	// 		fmt.Println("channel is closed")
	// 		break
	// 	}

	// 	fmt.Println("channel message received:", message)
	// }

	for message := range channel {
		fmt.Println("channel message received:", message)
	}
}

func write(text string, channel chan string) {
	for i := range 5 {
		var upperMessage = strings.ToUpper(text) + " - " + strconv.Itoa(i)
		channel <- upperMessage
		time.Sleep(time.Second)
	}

	close(channel)
}
