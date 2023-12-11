package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var ultimo string

	scanner := bufio.NewScanner(os.Stdin)
	
	fmt.Println("inserisci un cognome:")
	
	for scanner.Scan() {
		
		var stringaInput string 
		
		stringaInput = scanner.Text()
		if stringaInput == "" {
			break
		}
		
		if ultimo == ""  || stringaInput > ultimo {
			ultimo = stringaInput
		} 
		
		fmt.Println("inserisci un cognome:")
	}

	fmt.Println("ultimo cognome:", ultimo)
	
}
