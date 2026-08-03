package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel := write(8, 5)

	for message := range channel {
		fmt.Println(message)
	}
}

/*
write starts a concurrent generator that emits random passwords of a given
length over a read-only string channel at 500ms intervals.

It limits generation to the specified quantity and automatically closes
the channel upon completion to allow safe iteration via range loops.

Example:

	ch := write(8, 3)
	for password := range ch {
	    fmt.Println(password)
	}
*/
func write(width int, quantity int) <-chan string {
	channel := make(chan string)

	go func() {
		defer close(channel)

		for range quantity {
			channel <- randomPassword(width)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return channel
}

func randomPassword(n int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	chars := make([]byte, n)

	for i := range chars {
		chars[i] = charset[rand.Intn(len(charset))]
	}

	return string(chars)
}
