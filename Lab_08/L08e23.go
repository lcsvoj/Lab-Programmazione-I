package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {

	// 1) Legga da riga di comando un numero intero soglia ;
	input := os.Args[1]
	soglia, _ := strconv.Atoi(input)
	// 2) Generi in modo casuale una sequenza di lunghezza arbitraria di numeri interi da 0 a 100.
	// Il processo di generazione si interrompe quando viene generato un numero inferiore a soglia .
	sequenza := Genera(soglia)

	// 3) Stampi a video tutti i numeri generati.
	fmt.Println("Valori generati:\t", sequenza)

	// 4) Stampi a video tutti i numeri generati superiori a soglia .
	fmt.Println("Valori sopra soglia:\t", sequenza[:len(sequenza)-1])

}

// Ritorna una sequenza di lunghezza arbitraria di numeri interi, generata in base alle specifiche di cui al punto 2
func Genera(soglia int) []int {
	sequenza := make([]int, 0)
	for {
		n := rand.Intn(101)
		sequenza = append(sequenza, n)
		if n < soglia {
			break
		}
	}
	return sequenza
}
