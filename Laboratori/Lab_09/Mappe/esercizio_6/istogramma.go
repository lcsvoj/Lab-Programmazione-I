package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {

	StampaOccorrenze(Occorrenze(LeggiTesto()))

}

func LeggiTesto() string {
	scanner := bufio.NewScanner(os.Stdin)
	testo := ""
	for scanner.Scan() {
		testo += scanner.Text()
	}
	return testo
}

func Occorrenze(s string) map[rune]int {
	occorrenze := map[rune]int{}
	for _, l := range s {
		if unicode.IsLetter(l) {
			occorrenze[l]++
		}
	}
	return occorrenze
}

func StampaOccorrenze(occorrenze map[rune]int) {

	chiavi := []rune{}
	for k := range occorrenze {
		chiavi = append(chiavi, k)
	}

	SortRunes(chiavi)

	fmt.Println("Occorrenze:")
	for _, k := range chiavi {
		fmt.Printf("%c: %s\n", k, strings.Repeat("*", occorrenze[k]))
	}
}

func SortRunes(a []rune) {
	for i := 0; i < len(a)-1; i++ {
		indiceMin := i
		for j := i + 1; j < len(a); j++ {
			if a[indiceMin] > a[j] {
				indiceMin = j
			}
		}
		a[i], a[indiceMin] = a[indiceMin], a[i]
	}
}
