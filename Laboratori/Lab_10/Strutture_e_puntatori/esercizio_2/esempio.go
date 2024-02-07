package main

import "fmt"

func main() {
	var a, b = 15, 20
	f(&a, &b)
	fmt.Println(a, b)
}
func f(x, y *int) {
	*x, *y = *y, *x
}
