package main

import "time"

func main() {
	chan1, chan2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			chan1 <- "from chan1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			chan2 <- "from chan2"
		}
	}()

	for {
		select {
		case m1 := <-chan1:
			println(m1)
		case m2 := <-chan2:
			println(m2)
		}
	}
}
