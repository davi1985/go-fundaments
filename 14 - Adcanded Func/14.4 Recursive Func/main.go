package main

import "fmt"

func fibo(pos uint) uint {
	if pos <= 1 {
		return pos
	}

	return fibo(pos-2) + fibo(pos-1)
}

func main() {

	pos := uint(50)

	for i := uint(0); i < pos; i++ {
		fmt.Println(i, fibo(i))
	}

	// fmt.Println(fibo(pos))
}
