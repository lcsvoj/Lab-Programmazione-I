/*
Legga da riga di comando una sequenza di valori che rappresentano caratteri appartenenti all'alfabeto inglese (e quindi ASCII);
Stampi a video tutte le sottosequenze di caratteri presenti in s che:
	i. iniziano e finiscono con lo stesso carattere;
	ii. sono formate da almeno 3 caratteri.
Ciascuna sottosequenza deve essere stampata un'unica volta, riportando il relativo numero di occorrenze della sottosequenza nella sequenza input.
Le sottosequenze devono essere stampate in ordine di lunghezza (dalla più lunga alla più corta).
Se non esistono sottosequenze che soddisfano le condizioni 1 e 2, il programma non deve stampare nulla.
Si noti che una sottosequenza può essere contenuta in un'altra sottosequenza più grande.
Si assuma che la sequenza di valori specificata a riga di comando sia nel formato corretto e includa almeno 3 caratteri.
*/

package main

import (
	"fmt"
	"os"
	"strings"
)

var lunghezzaMassima int = 0

type sottosequenza struct {
	elementi  []string
	ocorrenze int
}

func main() {

	sequenza := LeggiSequenza()

	sottosequenze := TrovaSottosequenze(sequenza)

	StampaSottosequenze(sottosequenze)
}

func TrovaSottosequenze(sequenza []string) map[string]sottosequenza {

	sottosequenze := make(map[string]sottosequenza)

	for i := 0; i < len(sequenza)-1; i++ {
		for j := len(sequenza); j > i; j-- {

			s := sottosequenza{
				elementi:  make([]string, 0),
				ocorrenze: 1,
			}

			s.elementi = sequenza[i:j]

			if len(s.elementi) <= 2 {
				continue
			} else if s.elementi[0] == s.elementi[len(s.elementi)-1] {
				stringa := fmt.Sprintf("%v", s.elementi)

				_, exists := sottosequenze[stringa]
				if exists {
					s.ocorrenze += 1
				}

				sottosequenze[stringa] = s

				// Tiene atualizzata la lunghezza della maggiore sottosequenza
				if len(s.elementi) > lunghezzaMassima {
					lunghezzaMassima = len(s.elementi)
				}
			}
		}
	}
	return sottosequenze
}

func LeggiSequenza() []string {

	elementi := os.Args[1:]

	return elementi
}

func StampaSottosequenze(sottosequenze map[string]sottosequenza) {

	// Stampi con ordine di lunghezza crescente
	lunghezzaDaStampare := lunghezzaMassima

	// Controliamo se sono tutti stampati
	elementiStampati := 0

	// Testa (stampa o non) ogni elemento
	for elementiStampati < len(sottosequenze) {
		for _, el := range sottosequenze {
			if len(el.elementi) == lunghezzaDaStampare {
				fmt.Printf("%s -> Occorrenze: %d", strings.Join(el.elementi, " "), el.ocorrenze)
				fmt.Println()
				elementiStampati++
			}
		}
		lunghezzaDaStampare--
	}
}
