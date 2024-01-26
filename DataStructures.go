package main

import (
	"fmt"
	"strconv"
)

func MultiArrayCreator() {
	table := [5][6]string{}
	for row := 0; row < 5; row++ {
		for col := 0; col < 6; col++ {
			table[row][col] = strconv.Itoa(row) + "," + strconv.Itoa(col)
		}
	}
	fmt.Println(table)
}
