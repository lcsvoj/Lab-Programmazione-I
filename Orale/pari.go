package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sl := leggiSlice()
	fmt.Println("Ho trovato", contaPari(sl), "pari nella sequenza.")
}

func leggiSlice() (sl []int) {
	s := bufio.NewScanner(os.Stdin)

	fmt.Print("Inserisci una sequenza di interi: ")
	for s.Scan() {
		riga := s.Text()
		if riga == "" {
			break
		}

		slStr := strings.Split(riga, " ")

		for _, nStr := range slStr {
			nInt, err := strconv.Atoi(nStr)
			if err == nil {
				sl = append(sl, nInt)
			}
		}
	}
	return sl
}

func contaPari(sl []int) (contatore int) {
	for _, n := range sl {
		if n%2 == 0 {
			contatore++
		}
	}
	return contatore
}
