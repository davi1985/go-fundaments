package main

import "fmt"

type user struct {
	name string
	age  uint8
}

func (u user) save() {
	fmt.Printf("Saving user data - %s\n", u.name)
}

func (u user) eighth() bool {
	return u.age > 18
}

func (u *user) birthday() {
	u.age++
}

func main() {
	user1 := user{"User 1", 30}
	fmt.Println(user1)
	user1.save()

	user2 := user{"Davi", 30}
	fmt.Println(user2)
	user2.save()
	fmt.Println(user2.eighth())
	user2.birthday()
	fmt.Println(user2)
}
