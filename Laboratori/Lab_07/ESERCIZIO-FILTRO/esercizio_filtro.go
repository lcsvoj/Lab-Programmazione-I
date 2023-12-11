package main

import "fmt"

func main() {
	var n string
	fmt.Scan(&n)

	var ultimaCifra string
	fmt.Print("Output: ")
	for i := 0; i < len(n); i++ {
		if i == 0 || string(n[i]) != ultimaCifra {
			fmt.Print(string(n[i]))
		}
		ultimaCifra = string(n[i])
	}
	fmt.Println()
}
