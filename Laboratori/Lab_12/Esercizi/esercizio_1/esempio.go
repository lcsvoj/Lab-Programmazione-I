package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("Funzione 'main' - Inizio\n")
	Conteggio(3)
	fmt.Printf("Funzione 'main' - Fine\n")
}

func Conteggio(n uint) {
	rientro := strings.Repeat("\t", int(3-n))
	fmt.Printf("%sFunzione 'Conteggio' - Inizio esecuzione per n = %d\n", rientro, n)
	if n == 0 {
		fmt.Printf("%sPartenza!\n", rientro)
	} else {
		fmt.Printf("%s%d\n", rientro, n)
		Conteggio(n - 1)
	}
	fmt.Printf("%sFunzione 'Conteggio' - Fine esecuzione per n = %d\n", rientro, n)
}
