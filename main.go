package main

import (
	"fmt"
	"strconv"
)

type point struct {
	x float32
	y float32
	z float32

	func (p point) length() float64 {
		return math.Sqrt(
		(math.Pow(float64(p.x), 2) +
		math.Pow(float64(p.y), 2) +
		math.Pow(float64(p.z), 2)))
		}
}


func constructor(x, y, z float32) *point {
	p := point{x: x, y: y, z: z}
	return &p
	}
type Node struct{
	next Node

}
func main() {
	jultiArrayCreator()
my_map:= make(map[string] int)
my_map["kola"] = 89

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
	for i < 100 {
		fmt.Printf("%v\n", i)
		i = i + 1
	}
	city := "London"

	switch city {
	case "London":
		fmt.Print(city)
	case "Angeles":
		fmt.Print(city)
	default:
		fmt.Printf("%v \n", city)
	}

	greetUser()
	if v, err := doSth(); !err {
		//
	} else {
		fmt.Printf("%v", v)
	}
}

func greetUser() {
	var name string
	fmt.Scanf("%s", &name)
	age, _ := strconv.Atoi(name)
	fmt.Printf("you entered %v", age)
}

func checker() bool {
	num := 6
	condition := num%2 == 0
	return condition
}

func doSth() (int, bool) {
	for pos, char := range "Hello, world!" {
		fmt.Println("\n", pos, char)
	}
	return 20, false
}

func jultiArrayCreator() {
	table := [5][6]string{}
	for row := 0; row < 5; row++ {
		for col := 0; col < 6; col++ {
			table[row][col] = strconv.Itoa(row) + "," + strconv.Itoa(col)
		}
	}
	fmt.Printf("%v \n", table)
}
