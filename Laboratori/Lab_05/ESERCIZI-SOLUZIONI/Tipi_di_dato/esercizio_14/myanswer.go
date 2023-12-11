/* Scrivere un programma che legga da standard input sei valori reali rappresentando i punti nel piano cartesiano che corrispondono alle vertici di un triangolo.
Una volta terminata la fase di lettura, il programma deve stampare a video (come mostrato nell'Esempio di
esecuzione), se il triangolo ABC definito dai segmenti/lati AB , BC e AC è equilatero, iscoscele o scaleno. */

package main

import (
	"fmt"
	"math"
)

const EPSILON = 1.0e-9

func main() {

	var xA, xB, xC, yA, yB, yC float64

	fmt.Print("\nInserisci i valori dell'ascissa e dell'ordinata del punto A: ")
	fmt.Scan(&xA, &yA)
	fmt.Print("Inserisci i valori dell'ascissa e dell'ordinata del punto B: ")
	fmt.Scan(&xB, &yB)
	fmt.Print("Inserisci i valori dell'ascissa e dell'ordinata del punto C: ")
	fmt.Scan(&xC, &yC)

	// La lunghezza di ciascun lato di un triangolo è pari alla distanza euclidea tra gli estremi del lato.
	latoAB := Distanza(xA, yA, xB, yB)
	latoBC := Distanza(xB, yB, xC, yC)
	latoAC := Distanza(xC, yC, xA, yA)

	var classificaTriangolo string
	if ÈXUgualeAY(latoAB, latoBC) && ÈXUgualeAY(latoBC, latoAC) {
		// Un triangolo è equilatero se ha tutti e tre i lati di lunghezza uguale.
		classificaTriangolo = "equilatero"
	} else if ÈXUgualeAY(latoAB, latoBC) || ÈXUgualeAY(latoAB, latoAC) || ÈXUgualeAY(latoAC, latoBC) {
		// Un triangolo è isoscele se ha solo due lati di lunghezza uguale.
		classificaTriangolo = "isocele"
	} else {
		// Un triangolo è scaleno se ha tutti e tre i lati di lunghezza diversa.
		classificaTriangolo = "scaleno"
	}

	fmt.Printf("\nLe lunghezze dei lati sono: AB = %.2f, BC = %.2f, AC = %.2f\n", latoAB, latoBC, latoAC)
	fmt.Printf("Quindi il triangolo ABC è %s.\n\n", classificaTriangolo)

}

func Distanza(x1, y1, x2, y2 float64) float64 {
	return (math.Pow(math.Pow(x1-x2, 2)+math.Pow(y1-y2, 2), 0.5))
}

/* La funzione `func ÈXUgualeAY(x, y float64) bool` riceve in input due
valore reali nei parametri `x` e `y` e restituisce `true` se `x` è uguale
di `y` a meno di una costante EPSILON */
/* ÈXUgualeAY(5.0,4.999) -> false */
/* ÈXUgualeAY(5.0,4.9999999999) -> true */
func ÈXUgualeAY(x, y float64) bool {
	return math.Abs(x-y) <= EPSILON
}
