package main

import "fmt"
import "strconv"
import "os"
import "math"

func main() {

	numeri := LeggiNumeri()
	minimoDispari := MinimoDispari(numeri)
	StampaPariMaggioriDi(numeri, minimoDispari)

}

func LeggiNumeri() (numeri []int) {
	for _, v := range os.Args[1:] {
		if n, err := strconv.Atoi(v); err == nil {
			numeri = append(numeri, n)
		}
	}
	return
}

func MinimoDispari(numeri []int) (minimoDispari int) {
	minimoDispari = int(math.MaxInt64)
	for _, v := range numeri {
		if v%2 != 0 && v < minimoDispari {
			minimoDispari = v
		}
	}
	return
}

func StampaPariMaggioriDi(numeri []int, minimo int) {
	for _, v := range numeri {
		if v%2 == 0 && v > minimo {
			fmt.Print(v, " ")
		}
	}
	fmt.Println()
}
