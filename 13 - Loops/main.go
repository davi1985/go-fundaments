package main

import (
	"fmt"
	"time"
)

func main() {
	// i := 0

	// for i < 10 {
	// 	i++
	// 	fmt.Println("Incrementing j... ", i)
	// 	time.Sleep(time.Second)
	// }
	// fmt.Println(i)

	// for j := 0; j < 10; j += 2 {
	// 	fmt.Println("Incrementing j... ", j)
	// 	time.Sleep(time.Second)
	// }

	names := []string{"Davi", "Laura", "Sarah", "Joelma"}

	for key, value := range names {
		fmt.Println(key, value)
	}

	for _, value := range names {
		fmt.Println(value)
	}

	for key, value := range "Davi Silva" {
		fmt.Println(key, string(value))
	}

	user := map[string]string{
		"name":    "Davi",
		"surname": "Silva",
	}

	for key, value := range user {
		fmt.Println(key, value)
	}

	for {
		fmt.Println("Infinity for...")
		time.Sleep(time.Second)
	}
}
