package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	LeggiEStampa(scanner)
}

func InvertiStringa(s string) string {
	sl := []rune(s)
	if len(sl) > 1 {
		return string(sl[len(sl)-1:]) + InvertiStringa(string(sl[:len(sl)-1]))
	}
	return s
}

func LeggiEStampa(scanner *bufio.Scanner) {

	if scanner.Scan() {
		s := scanner.Text()
		LeggiEStampa(scanner)
		fmt.Println(InvertiStringa(s))
	}

}
