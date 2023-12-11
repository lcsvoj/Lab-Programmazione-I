package main

import (
	"fmt"
)

const Dimensione = 10

func main() {

	var a []int

	for i := 0; i < Dimensione; i++ {
		a = append([]int{i + 1}, a...)
	}
	/*
		primo ciclo: i == 0, a == []
		a = [1]

		secondo ciclo: i == 1, a == [1]
		a = [2 1]

		terzo ciclo: i == 2, a == [2 1]
		a = [3 2 1]

		... 	*/

	fmt.Println(a) // Prints [10 9 8 7 6 5 4 3 2 1]

	b := make([]int, len(a)) // b = [0 0 0 0 0 0 0 0 0 0]
	copy(b, a)
	fmt.Println(b) // Prints [10 9 8 7 6 5 4 3 2 1]

	c := make([]int, Dimensione)
	copy(c, a[Dimensione/2:])
	fmt.Println(c) // Prints [5 4 3 2 1 0 0 0 0 0]

}
