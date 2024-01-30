package main

import (
	"fmt"
)

func main() {

	insiemeResiduo := ""

	fmt.Scan(&insiemeResiduo)

	permutazioni := Permutazioni(insiemeResiduo)
	fmt.Printf("%d permutazioni: %v\n", len(permutazioni), permutazioni)

}

func Permutazioni(insiemeResiduo string) (permutazioni []string) {
	if len(insiemeResiduo) == 0 {
		return []string{""}
	} else {
		for i, v := range insiemeResiduo {
			sottopermutazioni := Permutazioni(insiemeResiduo[:i] + insiemeResiduo[i+1:])
			for _, permutazione := range sottopermutazioni {
				permutazioni = append(permutazioni, string(v)+permutazione)
			}
		}
		return
	}
}
