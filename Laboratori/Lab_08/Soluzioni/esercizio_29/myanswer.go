package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	numeri := os.Args[1:]
	fmt.Println(numeri)
	fmt.Printf("len(numeri) = %v\n", len(numeri))

	for i := 0; i < len(numeri); i++ {

		numeroAttuale, err := strconv.Atoi(numeri[i])
		fmt.Printf("\ni = %v, numeroAttuale = %v\n", i, numeroAttuale)
		if err != nil {
			fmt.Printf("Valore in posizione %d non valido (not integer).\n", i)
			os.Exit(0)
		}
		fmt.Printf("Elemento %v considerato valido come intero.\n", i)

		var numeroPrecedente int
		if i != 0 {
			numeroPrecedente, _ = strconv.Atoi(numeri[i-1])
		}

		// posizione di indice pari
		if (i%2 == 0) && (i != 0) {
			fmt.Printf("Posizione %v considerata pari\n", i)
			fmt.Printf("Comparando %v com %v\n", numeroAttuale, numeroPrecedente)
			if numeroAttuale <= numeroPrecedente {
				fmt.Printf("Valore in posizione %d non valido (minore del precedente).\n", i)
				os.Exit(0)
			}
			fmt.Printf("%v > %v --> OK \n", numeroAttuale, numeroPrecedente)
		}

		// posizione di indice dispari
		if i%2 != 0 {
			fmt.Printf("Posizione %v considerata dispari\n", i)
			fmt.Printf("Comparando %v com %v, \n", numeroAttuale, numeroPrecedente)
			if numeroAttuale >= numeroPrecedente {
				fmt.Printf("Valore in posizione %d non valido (maggiore del precedente).\n", i)
				os.Exit(0)
			}
			fmt.Printf("%v < %v --> OK \n", numeroAttuale, numeroPrecedente)

		}
	}
	fmt.Println("Sequenza valida.")
}
