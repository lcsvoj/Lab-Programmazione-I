package main

import (
	"fmt"
	"strconv"
)

func main() {

	var somma int

	for {
		var stringaInput string
		fmt.Println("inserisci un intero:")
		fmt.Scan(&stringaInput)

		val, err := strconv.Atoi(stringaInput); 

		if err != nil {
			break
		}
		
		somma+=val
	}

	fmt.Println("somma:", somma)
}
