package main

import (
	"fmt"
	"math"
)

type Punto struct {
	x float64
	y float64
}

func main() {
	var A, B, C Punto
	fmt.Print("Inserisci i valori dell'ascissa e dell'ordinata del punto A: ")
	fmt.Scan(&A.x, &A.y)
	fmt.Print("Inserisci i valori dell'ascissa e dell'ordinata del punto B: ")
	fmt.Scan(&B.x, &B.y)
	fmt.Print("Inserisci i valori dell'ascissa e dell'ordinata del punto C: ")
	fmt.Scan(&C.x, &C.y)
	fmt.Println()

	latoAB, latoBC, latoCA := Distanza(A, B), Distanza(B, C), Distanza(C, A)
	if latoAB == latoBC && latoBC == latoCA {
		fmt.Println("Il triangolo ABC è equilatero.")
		fmt.Printf("Lunghezza del lato: %v\n", latoAB)
	} else if latoCA == latoBC {
		fmt.Println("Il triangolo ABC è isoscele.")
		fmt.Println("I lati di lunghezza uguale sono AC e BC.")
		fmt.Printf("Lunghezza dei lati AC e BC: %v", latoCA)
	} else if latoAB == latoBC {
		fmt.Println("Il triangolo ABC è isoscele.")
		fmt.Println("I lati di lunghezza uguale sono AB e BC.")
		fmt.Printf("Lunghezza dei lati AB e BC: %v", latoAB)
	} else if latoAB == latoCA {
		fmt.Println("Il triangolo ABC è isoscele.")
		fmt.Println("I lati di lunghezza uguale sono AB e AC.")
		fmt.Printf("Lunghezza dei lati AB e AC: %v", latoAB)
	} else {
		fmt.Println("Il triangolo è escaleno.")
		fmt.Println("Lunghezza del lato AB:", latoAB)
		fmt.Println("Lunghezza del lato BC:", latoBC)
		fmt.Println("Lunghezza del lato CA:", latoCA)
	}
	fmt.Println()
}

func Distanza(p1, p2 Punto) float64 {
	// restituisce un valore float64 pari alla distanza euclidea tra i punti P1 e P2
	return math.Sqrt(math.Pow((p1.x-p2.x), 2) + math.Pow((p1.y-p2.y), 2))
}
