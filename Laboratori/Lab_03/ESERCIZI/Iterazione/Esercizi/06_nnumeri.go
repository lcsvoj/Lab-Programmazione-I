/*
Scrivere un programma che, dopo aver letto da standard input un numero intero n , chiede
all'utente di inserire n numeri interi (sempre da standard input).

Dopo aver letto gli n numeri interi, il programma deve stampare:
- la somma degli n numeri letti;
- il minimo tra i numeri letti;
- il massimo tra i numeri letti;
- il numero di interi letti strettamente positivi (maggiori di 0), strettamente negativi (minori di 0), e

nulli.
Suggerimento: per leggere n numeri, potete utilizzare un ciclo for che per n volte utilizza
fmt.Scan
*/

package main

import (
	"fmt"
)

func main() {

	var numeroNumeri, numeroInserito int
	var somma, minimo, massimo, positivi, negativi, nulli int

	// Leggiamo il numero di numeri che saranno inseriti
	fmt.Scan(&numeroNumeri)

	fmt.Println("Inserisci", numeroNumeri, "numeri:")
	// Leggiamo, in loop, i numeri e facciamo cosa dobbiamo fare per ognuno
	for i := 0; i < numeroNumeri; i++ {
		fmt.Scan(&numeroInserito)

		// Aggiungiamo il numero inserito alla somma
		somma += numeroInserito

		// Troviamo il segno del numero inserito, aggiungiamolo alla string giusta
		if numeroInserito < 0 {
			negativi += 1
		} else if numeroInserito == 0 {
			nulli += 1
		} else if numeroInserito > 0 {
			positivi += 1
		}

		// Troviamo se il numero deve sostituire il massimo o minimo attuale
		if numeroInserito > massimo {
			massimo = numeroInserito
		}

		if numeroInserito < minimo {
			minimo = numeroInserito
		}
	}

	// Stampiamo i risultati
	fmt.Println("Somma =", somma)
	fmt.Println("Valore minimo =", minimo)
	fmt.Println("Valore massimo =", massimo)
	fmt.Println("Interi > 0 =", positivi)
	fmt.Println("Interi < 0 =", negativi)
	fmt.Println("Interi = 0 =", nulli)
}
