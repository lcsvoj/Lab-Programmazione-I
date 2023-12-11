package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)



func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("inserisci una stringa:")
	
	for scanner.Scan() {
		
		var stringaInput string 
		
		stringaInput = scanner.Text()
		if stringaInput == "" {
			break
		}
		fmt.Println("stringa in maiuscolo:",strings.ToUpper(stringaInput))
		fmt.Println("inserisci una stringa:")
	}
	
	fmt.Println("ciao")
}
