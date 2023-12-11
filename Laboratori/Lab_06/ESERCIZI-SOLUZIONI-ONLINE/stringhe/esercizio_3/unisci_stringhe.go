package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var stringaCompleta string

	scanner := bufio.NewScanner(os.Stdin)
	
	fmt.Println("inserisci una stringa:")
	
	for scanner.Scan() {
		
		var stringaInput string 
		
		stringaInput = scanner.Text()
		if stringaInput == "" {
			break
		}
		stringaCompleta+= " " + stringaInput
		fmt.Println("inserisci una stringa:")
	}


	fmt.Println("stringa completa:" + stringaCompleta)

}
