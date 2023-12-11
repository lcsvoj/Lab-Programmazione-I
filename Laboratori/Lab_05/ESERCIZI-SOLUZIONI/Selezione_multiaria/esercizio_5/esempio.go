package main

import (
	"fmt"
)

func main() {

	var a1, b1 float64 = 5, 2

	var c1 complex128
		
	c1 = complex(a1, b1) 					// La funzione complex restituisce un valore complex128 con parte reale 
											// uguale a 5 e parte immaginaria uguale a 2i
	fmt.Println(c1)      					// Questa istruzine stampa a video "(5+2i)"

	var a2, b2 float32 = 6, 3
	
	c2 := complex(a2, b2) 					// La funzione complex restituisce un valore complex64 con parte reale 
											// uguale a 6 e parte immaginaria uguale a 3i
	fmt.Println(c2)       					// Questa istruzine stampa a video "(6+3i)"
	
	var c3 complex64 = 7 + 3i 				// Altro modo per inizializzare un variabile di tipo complex64, valido 
											// anche per inizializzare un variabile di tipo complex128
	
	parteReale := real(c3) 					// La funzione real(c3) restituisce il valore float32 della parte reale di c3
	fmt.Println(parteReale) 				// Questa istruzine stampa a video "7"; 	
	parteImmaginaria := imag(c1)			// La funzione imag(c1) restituisce il valore float64 della parte immaginaria di c1
	fmt.Println(parteImmaginaria) 			// Questa istruzine stampa a video "2";
	
	fmt.Println(c1+complex128(c2)+complex128(c3)) 	// Questa istruzine stampa a video "(18+8i)"
}
