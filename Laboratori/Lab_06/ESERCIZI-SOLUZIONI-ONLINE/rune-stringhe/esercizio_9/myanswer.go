package main

import (
	"fmt"
)

func main() {
	var stringa string

	fmt.Print("Parola in input: ")
	fmt.Scan(&stringa)

	for i := 0; i < len(stringa)-i; i++ {
		fmt.Printf("Sottostringa %d: %s\n", i+1, stringa[i:len(stringa)-i])
	}

}
