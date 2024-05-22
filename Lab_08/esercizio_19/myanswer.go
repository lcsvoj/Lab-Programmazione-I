package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	n, _ := strconv.Atoi(os.Args[1])
	fmt.Println(Fattoriali(n))
}

func Fattoriali(n int) (f []int) {
	// riceve in input un valore int nel parametro n e restituisce il valore f di
	// tipo []int in cui in f[0] è memorizzato il fattoriale di 1 , in f[1] è
	// memorizzato il fattoriale di 2 , ..., in f[n-1] è memorizzato il fattoriale di n

	f = make([]int, n+1)
	for i := 0; i < len(f); i++ {
		f[len(f)-1-i] = FattorialeDiUnNumero(n - i)
	}
	return f[1:]
}

func FattorialeDiUnNumero(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * FattorialeDiUnNumero(n-1)
}
