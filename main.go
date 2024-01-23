package main

import "fmt"

func main() {
	// var name = "Kola booking"
	// fmt.Printf("Welcome %v to booking app \n", name)

	// var firstname string

	// var age int
	// num := 20
	// num3 := 89
	// final := num3 - num

	// firstname = "Joseph"
	// fmt.Printf("%T\n", age)
	// fmt.Printf("welcome %v, you are %v years old", firstname, age)
	// fmt.Printf("%v", final)

	bookings := [6]string{"Nana", "Kola", "James", "Joseph"}
	// // var anotherBook [7]string

	// // var list []string
	// // list = append(list, "james")
	for i, booking := range bookings {
		if booking == "Joseph" {
			fmt.Println("HI Joseph")
		}
		fmt.Printf("%v\n %v\n ", i, booking)
	}
	i := 20
	for i <100 {
		fmt.Printf("%v", i)
	}
}
