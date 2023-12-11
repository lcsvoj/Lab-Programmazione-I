package main

import "fmt"

func main() {
	var m map[string]int
	if m == nil {
		fmt.Println("'nil' è lo zero-value di una variabile di tipo 'reference type' non inizializzata.")
	}
	//m["mamma"] = 5 /* panic: assignment to entry in nil map */
	m = make(map[string]int)
	for _, s := range []string{"questo", "è", "un", "test"} {
		m[s] = len([]rune(s))
	}

	for k, v := range m {
		fmt.Println(k, "->", v)
	}
	fmt.Println()

	for k := range m {
		fmt.Println(k, "->", m[k])
	}
}
