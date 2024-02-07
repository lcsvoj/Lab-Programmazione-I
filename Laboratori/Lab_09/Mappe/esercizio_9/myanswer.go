/*
Legga da riga di comando una sequenza di numeri interi e stampi tutte le sottosequenze che iniziano e finiscono con lo stesso numero.
Ciascuna sottosequenza deve essere stampata su una riga diversa.

Esempi:
	Se la sequenza di input è 1 2 3 -14 2 5 , l'unica sottosequenza è 2 3 -14 2 .
	Se la sequenza di input è 1 2 1 2 3 , abbiamo 2 sottosequenze: 1 2 1 e 2 1 2 .

Si consideri che:
	- se a riga di comando non viene specificata alcuna sequenza, il programma non deve stampare nulla;
	- una sottosequenza può essere contenuta in una sottosequenza più grande;
	- ogni sottosequenza deve comparire una sola volta tra quelle stampate a video;
	- le sottosequenze devono essere stampate in ordine di lunghezza (dalla più corta alla più lunga).
*/

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	sequenza := LeggiSequenza()

	sottosequenze := TrovaSottosequenze(sequenza)

	StampaSottosequenze(sottosequenze)
}

func TrovaSottosequenze(sequenza []int) map[string][]int {

	sottosequenze := make(map[string][]int)

	for i := 0; i < len(sequenza)-1; i++ {
		for j := len(sequenza); j > i; j-- {

			sottosequenza := sequenza[i:j]

			if len(sottosequenza) <= 1 {
				continue
			} else if sottosequenza[0] == sottosequenza[len(sottosequenza)-1] {
				stringa := fmt.Sprintf("%v", sottosequenza)
				sottosequenze[stringa] = sottosequenza
			}
		}
	}
	return sottosequenze
}

func LeggiSequenza() []int {

	elementi := os.Args[1:]

	var sequenza []int
	for _, el := range elementi {
		num, _ := strconv.Atoi(el)
		sequenza = append(sequenza, num)
	}

	return sequenza
}

func StampaSottosequenze(sottosequenze map[string][]int) {

	// Stampi con ordine di lunghezza crescente
	lunghezzaDaStampare := 2

	// Controliamo se sono tutti stampati
	elementiStampati := 0

	// Testa (stampa o non) ogni elemento
	for elementiStampati < len(sottosequenze) {
		for _, el := range sottosequenze {
			if len(el) == lunghezzaDaStampare {
				for _, n := range el {
					fmt.Printf("%d ", n)
				}
				fmt.Println()
				elementiStampati++
			}
		}
		lunghezzaDaStampare++
	}
}
