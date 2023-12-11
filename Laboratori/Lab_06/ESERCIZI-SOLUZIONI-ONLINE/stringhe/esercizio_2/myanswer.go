package main

import "fmt"

func main() {
	var (
		n int
		s string
	)
	fmt.Scan(&n, &s)

	for i := 0; i < n; i++ {
		fmt.Print(s)
		if i != n-1 {
			fmt.Print("-")
		}
	}
	fmt.Println()
}
