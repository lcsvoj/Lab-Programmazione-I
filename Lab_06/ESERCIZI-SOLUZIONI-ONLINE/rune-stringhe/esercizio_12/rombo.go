package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Dimensione non sufficiente")
	} else {
		dimensione := 2*n + 1
		stringaSpazi := ""
		stringaAsterischi := ""

		for i := 0; i < dimensione; i++ {
			stringaSpazi += " "
			stringaAsterischi += "*"
		}

		for i := 0; i < dimensione/2; i++ {
			fmt.Print(stringaSpazi[:dimensione/2-i], stringaAsterischi[:2*i+1], "\n")
		}
		fmt.Println(stringaAsterischi)
		for i := dimensione/2 - 1; i >= 0; i-- {
			fmt.Print(stringaSpazi[:dimensione/2-i], stringaAsterischi[:2*i+1], "\n")
		}
	}

}
