package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
)

func main() {
	testo := LeggiTesto()
	occorrenze := ContaLettere(testo)
	StampaIstogramma(occorrenze)
}

func LeggiTesto() string {
	scanner := bufio.NewScanner(os.Stdin)
	var result string

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		result += line + "\n"
	}

	return result
}

func ContaLettere(testo string) (occorrenze map[rune]int) {
	occorrenze = make(map[rune]int)
	for _, c := range testo {
		if unicode.IsLetter(c) {
			occorrenze[c]++
		}
	}
	return occorrenze
}

func StampaIstogramma(occorrenze map[rune]int) {

	letterePresenti := make([]rune, 0)
	for lettera, _ := range occorrenze {
		letterePresenti = append(letterePresenti, lettera)
	}

	sort.Slice(letterePresenti, func(i, j int) bool {
		return letterePresenti[i] < letterePresenti[j]
	})

	for _, lettera := range letterePresenti {
		fmt.Printf("%v: %v\n", string(lettera), strings.Repeat("*", occorrenze[lettera]))
	}
}
