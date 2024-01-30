package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	dividendo, divisore := LeggiValori()
	quoziente, resto := Dividi(dividendo, divisore)
	fmt.Printf("Quoziente = %d, Resto = %d\n", quoziente, resto)

}

func LeggiValori() (a, b int) {
	a, _ = strconv.Atoi(os.Args[1])
	b, _ = strconv.Atoi(os.Args[2])
	return
}

func Dividi(dividendo, divisore int) (quoziente, resto int) {
	if dividendo < divisore {
		return 0, dividendo
	} else {
		quoziente, resto = Dividi(dividendo-divisore, divisore)
		return quoziente + 1, resto
	}
}
