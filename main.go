package main

import (
	"encoding/json"
	"fmt"
)

type Contact struct {
	email string
	phone string
	addr  string
}

type Person struct {
	name    string
	age     int
	contact Contact
}

func main() {
	a := new(Peson)
	a.name = "kolawole"
	a.Talk()
	jsonParser()
	i := 20
	for i < 100 {
		fmt.Printf("%v\n", i)
		i = i + 1
	}

	arr := []string{"name", "kola", "osagie"}
	for _, val := range arr {
		println(val)
	}

	var person_map = make(map[int]Person)
	person_map[1] = Person{
		name: "kolawole",
		age:  30,
		contact: Contact{
			email: "kollydap@gmail.com",
			phone: "0590399303",
			addr:  "number 3 shifao kola street",
		},
	}
	person_map[2] = Person{
		name: "kola",
		age:  39,
		contact: Contact{
			email: "kola@gmail.com",
			phone: "0590399303",
			addr:  "number 3 shifao koladapao street",
		},
	}
	new_person := map[int]Person{
		1: Person{
			name: "kola",
			age:  39,
			contact: Contact{
				email: "kola@gmail.com",
				phone: "0590399303",
				addr:  "number 3 shifao koladapao street",
			},
		},
		2: Person{
			name: "kola",
			age:  39,
			contact: Contact{
				email: "kola@gmail.com",
				phone: "0590399303",
				addr:  "number 3 shifao koladapao street",
			},
		},
	}
	for k, v := range new_person {
		fmt.Printf("%v, %v \n", k, v)
	}

}

type User struct {
	Firstname string
	Age       int
}

func jsonParser() {
	// var user User
	var userlist []User
	json_string := `{
		"firstname": "kolawole",
		"age":40
	}`
	json_string = `[
		{
			"firstname": "kolawole",
			"age":40
		},
		{
			"firstname": "kolawole",
			"age":40
		}
	]`

	// err := json.Unmarshal([]byte(json_string), &user)
	// if err != nil {
	// 	panic("ouh")
	// }
	error := json.Unmarshal([]byte(json_string), &userlist)
	if error != nil {
		panic("ouch")
	}
	fmt.Printf("%v", userlist)

}

type VAluestring string

type DigitCounter interface {
	CountOddEven() (int, int)
}

// func (val VAluestring) CountOddEven(int, int) {
// 	odds, evens := 0, 0
// 	for _, digit := range val {
// 		if digit%2 == 0 {
// 			evens++
// 		} else {
// 			odds++
// 		}
// 	}
// 	return odds, evens
// }

type Peson struct {
	name, color string
	age, height int
}

func (person *Peson) Talk() {
	fmt.Println("Hi, my name is", person.name)
}

type Andriod struct {
	Peson
	Model string
}
