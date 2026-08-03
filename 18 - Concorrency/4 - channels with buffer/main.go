package main

import "fmt"

func main() {
	channel := make(chan string, 2)

	channel <- "Hello World"
	channel <- "Golang"
	message1 := <-channel
	message2 := <-channel

	fmt.Println(message1)
	fmt.Println(message2)
}
