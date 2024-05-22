package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	n := LeggiNumero()
	mazzo := GeneraMazzo()
	EstraiCarte(mazzo, n)
}

func LeggiNumero() int {
	input := os.Args[1]
	n, _ := strconv.Atoi(input)
	return n
}

// GeneraMazzo restituisce un (string) mazzo di carte formato dalle sole carte ♥
func GeneraMazzo() string {
	var mazzo string
	for card := '🂱'; card <= '🂺'; card++ {
		mazzo += string(card)
	}
	return mazzo
}

// EstraiCarta restituisce (rune) la carta estratta in modo casuale dal mazzo di
// carte ed il (string) mazzo residuo di carte dopo l'estrazione
func EstraiCarta(mazzo string) (cartaEstratta rune, mazzoResiduo string) {
	mazzoRune := []rune(mazzo)
	n := rand.Intn(len(mazzoRune))
	cartaEstratta = mazzoRune[n]
	mazzoResiduo = string(mazzoRune[:n]) + string(mazzoRune[n+1:])
	return cartaEstratta, mazzoResiduo
}

// EstraiCarte simula l'estrazione casuale (ed in sequenza) di n carte dal mazzo di
// carte rappresentato da mazzo , stampando a video le carte estratte e quelle
// rimaste nel mazzo dopo ogni estrazione
func EstraiCarte(mazzo string, n int) {
	for i := 0; i < n; i++ {
		cartaEstratta, mazzoResiduo := EstraiCarta(mazzo)
		fmt.Printf("Estratta la carta %v - Carte rimaste nel mazzo: %v\n", string(cartaEstratta), mazzoResiduo)
		mazzo = mazzoResiduo
	}
}
