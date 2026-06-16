package main

func main() {
	var name string = "Davi" // explicit declaration
	surname := "Silva"       // implicit declaration (type inference)
	println("Hello, " + name + " " + surname)

	var (
		age  int    = 30
		city string = "New York"
	)

	println("Age:", age)
	println("City:", city)

	var1, var2 := "Go", "Programming" // multiple variable declaration with type inference
	println("Language:", var1)
	println("Topic:", var2)

	const const1 = "This is a constant" // constant declaration

	const PI float64 = 3.14159 // constant declaration with type inference
	println("Constant:", const1)
	println("PI:", PI)

	var1, var2 = var2, var1 // swapping values of var1 and var2
	println("Swapped var1:", var1)
	println("Swapped var2:", var2)
}
