package main

import (
	"fmt"
	"os"
)

const SEPARATORE = ' '

func main() {
	sequenzaInput := os.Args[1:]
	sottosequenze := GeneraSottosequenze(sequenzaInput)
	StampaSottosequenze(sottosequenze)
}

func GeneraSottosequenze(sequenza []string) (sottosequenze map[string]int) {
	sottosequenze = make(map[string]int)
	for i := 0; i < len(sequenza)-1; i++ {
		for j := i + 1; j < len(sequenza); j++ {
			if sequenza[i] == sequenza[j] {
				var sottosequenza string
				for k := i; k < j; k++ {
					sottosequenza += sequenza[k] + string(SEPARATORE)
				}
				sottosequenza += sequenza[j]
				sottosequenze[sottosequenza]++
			}
		}
	}
	return
}

func StampaSottosequenze(sottosequenze map[string]int) {

	lunghezzaMassima := 0
	for k := range sottosequenze {
		lunghezza := numeroDiSeparatori(k) + 1
		if lunghezzaMassima < lunghezza {
			lunghezzaMassima = lunghezza
		}
	}

	for lunghezza := 2; lunghezza <= lunghezzaMassima; lunghezza++ {
		for k := range sottosequenze {
			if lunghezza == numeroDiSeparatori(k)+1 {
				fmt.Println(k)
			}
		}
	}

}

func numeroDiSeparatori(sequenza string) int {
	contatore := 0
	for _, v := range sequenza {
		if v == SEPARATORE {
			contatore++
		}
	}
	return contatore
}
