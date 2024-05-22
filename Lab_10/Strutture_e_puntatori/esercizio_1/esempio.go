package main

import "fmt"

func main() {
	var a, b, c int
	var ptr *int
	a, b, c = 10, 15, 20
	ptr = &a
	*ptr += 10
	a += 10
	ptr = &b
	*ptr += 10
	*ptr = c
	*ptr += 10
	fmt.Println(a, b, c)
}
