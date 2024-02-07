package main

import "fmt"

func main() {
	b := [3]rune{'a', 'b', 'c'}

	for _, v := range b {
		fmt.Printf("%c ", v)
	}

}
