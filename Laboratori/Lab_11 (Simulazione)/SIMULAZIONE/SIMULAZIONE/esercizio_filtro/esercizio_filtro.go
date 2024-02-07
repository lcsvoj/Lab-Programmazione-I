package main

import (
  "fmt"
  "os"
  "strconv"
//  "unicode"
)

func main() {

  // Leggi la stringa con la codifica compressa
  elementi := os.Args[1:]

  // Guarda ogni elemento letto
  var letteraDaStampare string
  for i, el := range elementi {
    // Negli indici pari ci sono lettere da stampare n volte,
    // Dove n viene nel successivo indice (dispari)
    if i == 0 || i % 2 == 0 {
      letteraDaStampare = el
    } else {
      ripetizione, _ := strconv.Atoi(el)
      for j := 0 ; j < ripetizione; j++ {
        fmt.Print(letteraDaStampare)
      }
    }
  }
}
