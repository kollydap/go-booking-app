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

	i := 20
	for i < 100 {
		fmt.Printf("%v\n", i)
		i = i + 1
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
	var user User
	var userlist [] User
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

	err := json.Unmarshal([]byte(json_string), &user)
	if err != nil {
		panic("ouch")
	}
	err = json.Unmarshal([]byte(json_string), &userlist)
	

}
