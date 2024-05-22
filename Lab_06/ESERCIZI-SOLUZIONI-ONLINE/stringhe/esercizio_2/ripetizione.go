package main

import "fmt"

func main() {

	var n int
	var stringaInput string

	fmt.Scan(&n, &stringaInput)

	for i := 0; i < n; i++ {
		fmt.Print(stringaInput)
		if i != n-1 {
			fmt.Print("-")
		}
	}

	fmt.Println()
}
