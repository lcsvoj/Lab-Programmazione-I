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
	DIVISORIMAX, _ := strconv.Atoi(os.Args[2])
	DIVISORIMIN, _ := strconv.Atoi(os.Args[3])

	stampaCoppie(N, DIVISORIMAX, DIVISORIMIN)
}

func DivisoriPropri(n int) (divisori []int) {
	// restituice tutti i divisori propri di n
	for i := 1; i < n; i++ {
		if n%i == 0 {
			divisori = append(divisori, i)
		}
	}
	return divisori
}

func ÈPerfetto(n int) bool {
	// restituice true se il numero è perfetto, false altrimenti
	var somma int
	divisori := DivisoriPropri(n)
	for i := 0; i < len(divisori); i++ {
		somma += divisori[i]
	}
	return (somma == n)
}

func stampaCoppie(N int, DIVISORIMAX int, DIVISORIMIN int) {
	// stampa tutte le coppie di interi positivi a <= N e b <= N che:
	for a := 0; a <= N; a++ {
		for b := 0; b <= N; b++ {

			// - almeno uno tra a e b sia un numero perfetto
			if !(ÈPerfetto(a) || ÈPerfetto(b)) {
				continue
			}

			// - abbiano in comune un numero di divisori propri >= DIVISORIMIN e <= DIVISORIMAX
			var numeroDivisoriInComune int
			divisori_a, divisori_b := DivisoriPropri(a), DivisoriPropri(b)
			for i := 0; i < len(divisori_a); i++ {
				for j := 0; j < len(divisori_b); j++ {
					if divisori_a[i] == divisori_b[j] {
						numeroDivisoriInComune++
					}
				}
			}
			if (numeroDivisoriInComune <= DIVISORIMIN) || (numeroDivisoriInComune >= DIVISORIMAX) {
				continue
			}

			fmt.Println(a, b)
		}
	}
}
