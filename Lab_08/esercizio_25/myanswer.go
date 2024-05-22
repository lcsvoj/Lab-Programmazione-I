package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	n, mazzo := LeggiNumero(), GeneraMazzo()
	EstraiCarte(mazzo, n)
}

func LeggiNumero() (n int) {
	// legge da riga di comando un numero intero e ne restituisce il valore;
	n, _ = strconv.Atoi(os.Args[1])
	if !(n < 10 && n > 0) {
		fmt.Println("Si chiede un numero tra 0 e 10")
	}
	return n
}

func GeneraMazzo() string {
	// restituisce un valore string che rappresenta un mazzo di carte formato dalle sole carte di cuori
	var mazzo string
	for i := '\U0001F0B1'; i <= '\U0001F0BA'; i++ {
		mazzo += string(i) + " "
	}
	fmt.Println("Il mazzo originale:", mazzo)
	return mazzo
}

func EstraiCarta(mazzo string) (cartaEstratta rune, mazzoResiduo string) {
	// riceve in input un valore string nel parametro mazzo e restituisce un valore di tipo rune nella
	// variabile cartaEstratta ed un valore di tipo string nella variabile mazzoResiduo , che
	// rappresentano rispettivamente la carta estratta in modo casuale dal mazzo di carte rappresentato da
	// mazzo ed il mazzo residuo di carte dopo l'estrazione

	runeMazzo := []rune(mazzo)
	var doveEstrarre int

	for {
		doveEstrarre = rand.Intn(len(runeMazzo))
		if runeMazzo[doveEstrarre] != ' ' {
			break
		}
	}

	cartaEstratta = runeMazzo[doveEstrarre]
	mazzoResiduo = string(runeMazzo[:doveEstrarre]) + string(runeMazzo[doveEstrarre+2:])

	return cartaEstratta, mazzoResiduo
}

func EstraiCarte(mazzo string, n int) {
	// riceve in input un valore string nel parametro mazzo ed un valore int nel parametro n e simula l'estrazione casuale (ed in sequenza) di
	// n carte dal mazzo di carte rappresentato da mazzo , stampando a video le carte estratte e quelle
	// rimaste nel mazzo dopo ogni estrazione; la funzione deve utilizzare la funzione EstraiCarta()
	for i := 0; i < n; i++ {
		cartaEstratta, mazzoResiduo := EstraiCarta(mazzo)
		fmt.Printf("Estratta la carta %c - Carte rimaste nel mazzo: %v\n", cartaEstratta, mazzoResiduo)
		mazzo = mazzoResiduo
	}
}
