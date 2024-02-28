package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	s1 := leggiSlice()
	s2 := leggiSlice()
	positivi := trovaUguali(s1, s2)
	fmt.Println(positivi)
}

func trovaUguali(s1, s2 []int) (uguali []int) {
	m := make(map[int]bool)
	for _, n := range s1 {
		m[n] = true
	}

	for _, n := range s2 {
		if m[n] == true {
			uguali = append(uguali, n)
		}
	}

	return uguali
}

func leggiSlice() (sl []int) {
	s := bufio.NewReader(os.Stdin)

	fmt.Print("Inserisci una sequenza di interi: ")

	riga, _ := s.ReadString('\n')

	slStr := strings.Split(riga, " ")

	for _, nStr := range slStr {
		nInt, err := strconv.Atoi(nStr)
		if err == nil {
			sl = append(sl, nInt)
		}
	}
	return sl
}
