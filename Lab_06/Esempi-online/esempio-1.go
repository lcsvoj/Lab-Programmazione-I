package main

import "fmt"

func main() {
	
	var x rune = '\u0061'	// rappresentazione in bit Unicode/UTF-8; 4 cifre esadecimali (2 byte)
	var y rune = 97			// Unicode value/code (valore/codice numerico in base 10)
	var z rune = 'a'
	
	fmt.Println(x,y,z)
	fmt.Println()
	fmt.Println(string(x),string(y),string(z))
	
	fmt.Println()
	
	var beta_1 rune = '\u03B2'
	var beta_2 rune = 946
	fmt.Println(string(beta_1),string(beta_2))
	
}