package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Dimensione non sufficiente")
	} else {
		stringaAsterischi := ""

		for i := 0; i < n; i++ {
			stringaAsterischi += "*"
			fmt.Println(stringaAsterischi)
		}
	}
}
