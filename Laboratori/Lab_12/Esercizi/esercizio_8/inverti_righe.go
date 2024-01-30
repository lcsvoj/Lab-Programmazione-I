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

func LeggiEStampa(scanner *bufio.Scanner) {

	if scanner.Scan() {
		s := scanner.Text()
		LeggiEStampa(scanner)
		fmt.Println(s)
	}

}
