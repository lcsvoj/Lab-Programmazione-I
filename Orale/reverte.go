package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	str := os.Args[1:]
	strDireta := strings.Join(str, " ")

	strReversa := ""
	for _, c := range strDireta {
		strReversa = string(c) + strReversa
	}
	fmt.Println("La stringa invertita è:", strReversa)
}
