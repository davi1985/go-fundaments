package main

import "fmt"

func main() {

	user := map[string]string{
		"name":    "Davi",
		"surname": "Silva",
	}

	fmt.Println(user["name"])

	user2 := map[string]map[string]string{
		"user": {
			"name":    "Davi",
			"surname": "Silva",
		},
	}

	fmt.Println(user2)
	delete(user2["user"], "name")
	fmt.Println(user2)
}
