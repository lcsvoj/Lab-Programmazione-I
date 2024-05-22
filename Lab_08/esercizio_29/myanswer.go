package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	numeri := os.Args[1:]
	fmt.Println(numeri)

	for i := 0; i < len(numeri); i++ {

		numeroAttuale, err := strconv.Atoi(numeri[i])
		if err != nil {
			fmt.Printf("Valore in posizione %d non valido.\n", i)
			os.Exit(0)
		}

		var numeroPrecedente int
		if i != 0 {
			numeroPrecedente, _ = strconv.Atoi(numeri[i-1])
		}

		// posizione di indice pari
		if (i%2 == 0) && (i != 0) {
			if numeroAttuale <= numeroPrecedente {
				fmt.Printf("Valore in posizione %d non valido.\n", i)
				os.Exit(0)
			}
		}

		// posizione di indice dispari
		if i%2 != 0 {
			if numeroAttuale >= numeroPrecedente {
				fmt.Printf("Valore in posizione %d non valido.\n", i)
				os.Exit(0)
			}

		}
	}
	fmt.Println("Sequenza valida.")
}
