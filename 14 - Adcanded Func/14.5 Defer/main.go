package main

import "fmt"

func fn1() {
	fmt.Println("Fn 1")
}

func fn2() {
	fmt.Println("Fn 2")
}

func studentIsApproved(n1, n2 float32) bool {
	defer fmt.Println("Average calculated")

	average := (n1 + n2) / 2

	if average >= 6 {
		return true
	}

	return false
}

func main() {
	defer fn1()
	fn2()

	fmt.Println(studentIsApproved(7, 8))
}
