package main

/* Pressuposti:
1. divisori propri di un numero naturale sono tutti i suoi divisori, tranne il numero stesso;
2. un numero naturale è perfetto se è uguale alla somma dei suoi divisori propri (per esempio, 6 è perfetto perché 6 = 1+2+3 );
3. valori letti da riga di comando siano specificati nel formato corretto;
4. N > MAX > MIN > 0
*/

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// leggi da riga di comando tre numeri positivi: N, DIVISORIMIN e DIVISORIMAX
	N, _ := strconv.Atoi(os.Args[1])
	DIVISORIMIN, _ := strconv.Atoi(os.Args[2])
	DIVISORIMAX, _ := strconv.Atoi(os.Args[3])
	// fmt.Printf("\nNumeri letti: N = %d, DIVISORIMIN = %d, DIVISORIMAX = %d\n", N, DIVISORIMIN, DIVISORIMAX)

	stampaCoppie(N, DIVISORIMAX, DIVISORIMIN)
}

func DivisoriPropri(n int) (divisori []int) {
	// restituice tutti i divisori propri di n
	// fmt.Printf("\nFinding the divisors of %d:\n", n)
	for i := 1; i <= n; i++ {
		if n%i == 0 && n != i {
			divisori = append(divisori, i)
		}
	}
	// fmt.Printf("The divisors are: %v\n", divisori)
	return divisori
}

func ÈPerfetto(n int) bool {
	// restituice true se il numero è perfetto, false altrimenti
	var somma int
	divisori := DivisoriPropri(n)
	for i := 0; i < len(divisori); i++ {
		somma += divisori[i]
	}
	// fmt.Printf("Is %v a perfect number? %v\n", n, (somma == n))
	return (somma == n)
}

func stampaCoppie(N int, DIVISORIMAX int, DIVISORIMIN int) {
	// stampa tutte le coppie di interi positivi a <= N e b <= N che:
	for a := 2; a <= N; a++ {
		for b := a; b <= N; b++ {
			// fmt.Printf("\n#################### a = %d, b = %d ####################\n", a, b)

			// - almeno uno tra a e b sia un numero perfetto
			if !(ÈPerfetto(a) || ÈPerfetto(b)) {
				// fmt.Printf("\nNone of them is perfect. Moving on to the next pair.\n")
				continue
			}

			// - abbiano in comune un numero di divisori propri >= DIVISORIMIN e <= DIVISORIMAX
			numeroDivisoriComuni := divisoriComuni(a, b)
			if (numeroDivisoriComuni < DIVISORIMIN) || (numeroDivisoriComuni > DIVISORIMAX) {
				// fmt.Printf("Criteria: MIN = %d, MAX = %d\n", DIVISORIMIN, DIVISORIMAX)
				// fmt.Printf("This number doesn't meet the criteria (%d <= %d || %d >= %d). Next pair.\n", numeroDivisoriComuni, DIVISORIMIN, numeroDivisoriComuni, DIVISORIMAX)
				continue
			}

			// fmt.Printf("\nSo this is a valid pair, let's print it:\n")
			fmt.Println(a, b)
		}
	}
}

func divisoriComuni(a int, b int) (numeroDivisoriComuni int) {
	// fmt.Printf("\nOne of them is perfect, we can move on.\nLet's check how many divisors they have in common:\n")
	divisori_a, divisori_b := DivisoriPropri(a), DivisoriPropri(b)
	// fmt.Printf("\nDivisors of a: %v\nDivisors of b: %v\n", divisori_a, divisori_b)
	for i := 0; i < len(divisori_a); i++ {
		for j := 0; j < len(divisori_b); j++ {
			if divisori_a[i] == divisori_b[j] {
				numeroDivisoriComuni++
			}
		}
	}
	// fmt.Printf("\nSo they have %d divisors in common.\n", numeroDivisoriComuni)
	return numeroDivisoriComuni
}
