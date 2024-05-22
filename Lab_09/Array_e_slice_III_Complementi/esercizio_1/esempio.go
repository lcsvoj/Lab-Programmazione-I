package main

import (
	"fmt"
)

func main() {

	var a [3][2]int = [3][2]int{{1, 2}, {10, 20}, {100, 200}}
	//a := [3][2]int{{1, 2}, {10, 20}, {100, 200}}
	//a := [...][2]int{{1, 2}, {10, 20}, {100, 200}}

	fmt.Println("a:", a)
	/* a è un array bi-dimensionale di tipo [3][2]int con lunghezza pari a 3 */
	fmt.Printf("Tipo di a: %T\n", a)

	fmt.Println()

	for i, r := range a {
		fmt.Printf("a[%d]: %v\n", i, r)
	}
	for i := 0; i < len(a); i++ {
		/* a[i], con 0 <= i < len(a), è di tipo [2]int */
		fmt.Printf("Tipo di a[%d]: %T\n", i, a[i])
	}

	fmt.Println()

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			fmt.Printf("%d ", a[i][j])
		}
		fmt.Println()
	}
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			/* a[i][j], con 0 <= i < len(a) e con 0 <= j < len(a[i]), è di tipo int */
			fmt.Printf("a[%d][%d] è il valore alla riga %d e colonna %d dell'array bi-dimensionale a: %d\n", i, j, i, j, a[i][j])
		}
	}

}
