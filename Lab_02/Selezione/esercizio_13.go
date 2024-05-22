/*
Scrivere un programma che legga da standard input 4 valori a virgola mobile:
	- i primi due valori sono il coefficiente angolare m e il termine noto q di una retta r: y = m*x + q
	- il terzo e il quarto valore sono le coordinate px e py di un punto P(px,py)
Il programma deve determinare se il punto P sta sopra o sotto la retta od appartiene ad essa, e stampare a video il relativo messaggio.
Suggerimento: un punto appartiene ad una retta se sostituendo le sue coordinate nell'equazione della retta l'uguaglianza è verificata. Un punto sta sopra una retta se sostituendo il valore dell'ascissa nell'equazione della retta si ottiene y < py .
*/

package main

import "fmt"

func main() {
	
	var m, q int
	fmt.Print("Inserisci m e q: ")
	fmt.Scan(&m, &q)
	
	var x, y int
	fmt.Print("Inserisci x e y: ")
	fmt.Scan(&x, &y)
	
	if (y == m*x + q) {
		fmt.Println("Il punto appartiene alla retta")
	} else if (y < m*x+q) {
		fmt.Println("Il punto sta sotto la retta")
	} else if (y > m*x+q) {
		fmt.Println("Il punto sta sopra la retta")	
	}

}
