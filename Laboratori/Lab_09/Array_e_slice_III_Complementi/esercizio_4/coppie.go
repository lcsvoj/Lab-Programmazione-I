package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {

	sliceBidimensionale := CreaSliceBidimensionale(LeggiNumero())
	InizializzaSliceBidimensionale(sliceBidimensionale)

	coppie := Coppie(sliceBidimensionale)
	for _, coppia := range coppie {
		fmt.Println(coppia)
	}
}

func LeggiNumero() (n int) {
	n, _ = strconv.Atoi(os.Args[1])
	return
}

func CreaSliceBidimensionale(dimensione int) (risultato [][]int) {
	risultato = make([][]int, dimensione)
	for i := 0; i < dimensione; i++ {
		risultato[i] = make([]int, dimensione)
	}
	return
}

func InizializzaSliceBidimensionale(sliceBidimensionale [][]int) {
	rand.Seed(int64(time.Now().Nanosecond()))
	for i := range sliceBidimensionale {
		for j := range sliceBidimensionale[i] {
			sliceBidimensionale[i][j] = rand.Intn(2)
		}
	}
}

func Coppie(sliceBidimensionale [][]int) (coppie [][]int) {
	for i, riga := range sliceBidimensionale {
		for j, valore := range riga {
			if valore == 1 {
				coppie = append(coppie, []int{i, j})
			}
		}
	}
	return
}
