package main

import "fmt"

func invertSignal(n int) int {
	return n * -1
}

func invertSignalWithPointer(n *int) {
	*n = *n * -1
}

func main() {
	n := 20
	numberInverted := invertSignal(n)

	fmt.Println(numberInverted)
	fmt.Println(n)

	newNumber := 40
	fmt.Println(newNumber)
	invertSignalWithPointer(&newNumber)
	fmt.Println(newNumber)
}
