package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	sufficienti, insufficienti := FiltraVoti(LeggiNumeri())
	fmt.Printf("Voti sufficienti: %d\nVoti insufficienti: %d\n", sufficienti, insufficienti)

}

func LeggiNumeri() (numeri []int) {
	// restituisce un valore numeri di tipo []int in cui sono
	// memorizzati i valori numerici interi specificati a riga di comando
	numeri = make([]int, len(os.Args)-1)
	for i, c := range os.Args[1:] {
		numeri[i], _ = strconv.Atoi(c)
	}
	return numeri
}

func FiltraVoti(voti []int) (sufficienti, insufficienti []int) {
	/* riceve in input un valore []int nel parametro voti e restituisce due valori di tipo []int ,
	sufficienti e insufficienti , in cui sono memorizzati rispettivamente i voti sufficienti e
	insufficienti presenti in voti */
	for _, c := range voti {
		if c >= 60 {
			sufficienti = append(sufficienti, c)
		} else {
			insufficienti = append(insufficienti, c)
		}
	}
	return sufficienti, insufficienti
}
