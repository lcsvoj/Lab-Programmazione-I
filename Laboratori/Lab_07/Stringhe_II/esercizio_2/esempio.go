package main

import (
	"fmt"
)

const NumeroLettere = 'Z' - 'A' + 1

func main() {

	var carattere rune = 'Z'

	carattere += 2
	carattere -= 'A'
	carattere %= NumeroLettere
	carattere += 'A'

	fmt.Printf("%c\n", carattere)
}
