package main

import "fmt"

func main() {

	var decimale, base int
	fmt.Print("Inserisci un numero decimale intero non negativo: ")
	fmt.Scan(&decimale)
	fmt.Print("In che base lo vuoi rappresentato (max: 9)? ")
	fmt.Scan(&base)

	fmt.Printf("Numero convertito: %d, %X\n", decimale, decimale)

	var risultato, cifra int
	var moltiplicatore int = 1
	for decimale > 0 {
		cifra = (decimale % base)
		risultato += (cifra * moltiplicatore)
		decimale /= base
                moltiplicatore *= 10
	}

	fmt.Printf("Numero convertito: %d\n", risultato)
}
