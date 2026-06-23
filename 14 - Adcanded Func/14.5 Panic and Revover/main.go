package main

import "fmt"

func recoverExec() {
	fmt.Println("Running recoverExec()")

	if r := recover(); r != nil {
		fmt.Println("Try recover")
	}
}

func studentIsApproved(n1, n2 float32) bool {
	defer recoverExec()
	average := (n1 + n2) / 2

	if average > 6 {
		return true
	} else if average < 6 {
		return false
	}

	panic("Average equals 6")
}

func main() {
	fmt.Println(studentIsApproved(6, 6))
	fmt.Println("App finished")
}
