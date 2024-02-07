/*
Scrivere un programma che legga da riga di comando due numeri interi n e d .
Il programma deve calcolare e stampare il più piccolo numero ottenibile rimuovendo d cifre
decimali da n .

Si assuma che i valori letti da riga di comando siano nel formato corretto e che il numero di cifre di
n sia maggiore di d.
*/

package main

import (
  "os"
  "fmt"
  "strconv"
)

func main() {
  d, n := os.Args[1], os.Args[2]
  fmt.Println(TagliaNumero(d, n))

}

func TagliaNumero(d, n string) (numeriTagliati []string) {
  // Tratteremo d come una stringa, ma n come un int
  nInt, _ := strconv.Atoi(n)

  for p1 := 0 ; p1 <= len(d) ; p1++ {
    for p2 := len(d) ; p2 > p1 ; p2-- {
      
    }
    stringa := d[:i1]
    numeriTagliati = append(numeriTagliati, d[i:])
  }

  return numeriTagliati
}
