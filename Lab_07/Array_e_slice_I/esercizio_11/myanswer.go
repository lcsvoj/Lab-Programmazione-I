package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	numeroGiocatori := numeroGiocatori()
	giocatori := nomiGiocatori(numeroGiocatori)
	numeroTurni := numeroTurni()

	fmt.Println()
	vincitori := make([]string, numeroTurni)
	for i := 0; i < numeroTurni; i++ {
		var risultatoVincente int
		fmt.Printf("=== Turno %d ===\n", i+1)
		for _, nomeGiocatore := range giocatori {
			dado1, dado2 := LanciaDueDadi()
			fmt.Printf("Giocatore %s, lanci di valore: %d e %d\n", nomeGiocatore, dado1, dado2)
			if dado1+dado2 > risultatoVincente {
				risultatoVincente = dado1 + dado2
				vincitori[i] = nomeGiocatore
			}
		}
		fmt.Printf("Turno %d, vincitore: %s\n\n", i+1, vincitori[i])
	}

	fmt.Println("\nVittorie:")
	for i := 0; i < numeroTurni; i++ {
		fmt.Printf("Vincitore turno %d: %s\n", i+1, vincitori[i])
	}
}

func numeroGiocatori() int {
	fmt.Print("Inserisci il numero di giocatori: ")
	var n int
	fmt.Scan(&n)
	return n
}

func numeroTurni() int {
	fmt.Print("Inserisci il numero di turni: ")
	var n int
	fmt.Scan(&n)
	return n
}

func nomiGiocatori(n int) []string {
	fmt.Printf("Inserisci i nomi dei %d giocatori:\n", n)

	scanner := bufio.NewScanner(os.Stdin)
	giocatori := make([]string, n)

	var contatore int = 0
	for scanner.Scan() {
		line := scanner.Text()

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Errore nella lettura dell'input:", err)
		}

		giocatori[contatore] = line

		contatore++
		if contatore == n {
			break
		}
	}
	return giocatori
}

func LanciaDueDadi() (int, int) {
	return rand.Intn(6) + 1, rand.Intn(6) + 1
}
