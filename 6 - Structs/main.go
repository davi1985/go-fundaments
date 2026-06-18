package main

import "fmt"

type user struct {
	name    string
	age     uint8
	address address
}

type address struct {
	street string
	number uint8
}

func main() {
	var user1 user
	user1.name = "Davi"
	user1.age = 40

	address1 := address{"Street 1", 4}
	user1.address = address1
	fmt.Println(user1)

	user2 := user{
		name: "Sarah",
		age:  3,
	}

	fmt.Println(user2)

	user3 := user{name: "Laura"}
	fmt.Println(user3)

	user4 := user{"Joelma", 35, address1}

	fmt.Println(user4)
}
