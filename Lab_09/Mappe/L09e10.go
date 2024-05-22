package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	sottosequenze := TrovaSottosequenze(LeggiSequenza())
	StampaMappa(sottosequenze)
}

func LeggiSequenza() []string {
	sequenza := os.Args[1:]
	return sequenza
}

func TrovaSottosequenze(sequenza []string) []string {
	sottosequenze := make([]string, 0)
	for i := 0; i < len(sequenza); i++ {
		for j := len(sequenza); j > i; j-- {
			sottosequenza := sequenza[i:j]
			if len(sottosequenza) >= 3 && (sottosequenza[0] == sottosequenza[len(sottosequenza)-1]) {
				sottosequenze = append(sottosequenze, strings.Join(sottosequenza, " "))
			}
		}
	}
	return sottosequenze
}

func OrganizzaSottosequenze(sottosequenze []string) map[string]int {
	risultato := make(map[string]int)
	for _, sottosequenza := range sottosequenze {
		risultato[sottosequenza]++
	}
	return risultato
}

func StampaMappa(sottosequenze []string) {
	sottosequenzeMap := OrganizzaSottosequenze(sottosequenze)
	ordine := OrdinaSottosequenze(sottosequenzeMap)

	for _, sottosequenza := range ordine {
		fmt.Printf("%s -> Occorrenze: %d\n", sottosequenza, sottosequenzeMap[sottosequenza])
	}
}

func OrdinaSottosequenze(m map[string]int) []string {
	ordine := make([]string, 0)
	for sottosequenza, _ := range m {
		ordine = append(ordine, sottosequenza)
	}
	sort.Slice(ordine, func(i, j int) bool {
		return len(ordine[i]) > len(ordine[j])
	})
	return ordine
}
