package main

import "fmt"

func main() {

	var a, b float64

	fmt.Scan(&a, &b)

	if a == 0 {
		if b == 0 {
			fmt.Println("Indeterminata")
		} else {
			fmt.Println("Impossibile")
		}
	} else {
		fmt.Println("La radice è", -b/a)
	}

	/*
		if a == 0 && b == 0{
			fmt.Println("Indeterminata")
		} else if a == 0 {
			fmt.Println("Impossibile")
		} else {
			fmt.Println("La radice è", -b/a)
		}
	*/

}
