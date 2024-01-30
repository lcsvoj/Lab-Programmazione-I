package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(InvertiStringa(s))
}

func InvertiStringa(s string) string {
	sl := []rune(s)
	if len(sl) > 1 {
		return string(sl[len(sl)-1:]) + InvertiStringa(string(sl[:len(sl)-1]))
	}
	return s
}
