package main

import "fmt"

func weekDay(number int) string {
	switch number {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-feira"
	case 3:
		return "Terça-feira"
	case 4:
		return "Quarta-feira"
	case 5:
		return "Quinta-feira"
	case 6:
		return "Sexta-feira"
	case 7:
		return "Sábado"
	default:
		return "Dia inválido"
	}
}

func weekDay2(number int) string {
	var weekDay string

	switch {
	case number == 1:
		weekDay = "Domingo"
		fallthrough
	case number == 2:
		weekDay = "Segunda-feira"
	case number == 3:
		weekDay = "Terça-feira"
	case number == 4:
		weekDay = "Quarta-feira"
	case number == 5:
		weekDay = "Quinta-feira"
	case number == 6:
		weekDay = "Sexta-feira"
	case number == 7:
		weekDay = "Sábado"
	default:
		weekDay = "Dia inválido"
	}

	return weekDay
}

func main() {
	fmt.Println(weekDay(3))
	fmt.Println(weekDay(7))
	fmt.Println(weekDay(9))
	fmt.Println("---------------")

	fmt.Println(weekDay2(2))
	fmt.Println(weekDay2(5))
	fmt.Println(weekDay2(200))

	fmt.Println(weekDay2(1))
}
