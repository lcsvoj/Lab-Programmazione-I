package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	numeri := os.Args[1:]

	for i := 0; i < len(numeri); i++ {

		_, err := strconv.Atoi(numeri[i])
		if err != nil {
			fmt.Printf("Valore in posizione %d non valido (not integer).\n", i)
			os.Exit(0)
		}

		// posizione di indice pari
		if (i%2 == 0) && (i != 0) {
			if numeri[i] <= numeri[i-1] {
				fmt.Printf("Valore in posizione %d non valido (minore del precedente).\n", i)
				os.Exit(0)
			}

			// posizione di indice dispari
			if i%2 != 0 {
				if numeri[i] >= numeri[i-1] {
					fmt.Printf("Valore in posizione %d non valido (maggiore del precedente).\n", i)
					os.Exit(0)
				}
			}
		}
	}
	fmt.Println("Sequenza valida.")
}
