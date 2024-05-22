/*

Scrivere un programma che legga da standard input un voto v da 0 a 100 e stampi:
- Insufficiente se il voto è inferiore a 60 ( v<60 )
- Sufficiente se il voto è compreso tra 60 e 70 ( v>=60 e v<70 )
- Buono se il voto è compreso tra 70 e 80 ( v>=70 e v<80 )
- Distinto se il voto è comrpeso tra 80 e 90 ( v>=80 e v<90 )
- Ottimo se il voto è compreso tra 90 e 100 ( v>=90 e v<=100 )
- Errore se il voto è negativo o superiore a 100

*/

package main

import "fmt"

func main() {

	var voto int
	fmt.Print("Inserisci il voto: ")
	fmt.Scan(&voto)

	if voto < 60 {
		fmt.Println("Insufficiente")
	} else if voto < 70 {
		fmt.Println("Sufficiente")
	} else if voto < 80 {
		fmt.Println("Buono")
	} else if voto < 90 {
		fmt.Println("Distinto")
	} else if voto <= 100 {
		fmt.Println("Ottimo")
	} else {
		fmt.Println("Errore")
	}

}
