package main

import "fmt"

func main() {
	var a, b int = 10, 20
	var ptr *int
	*ptr = 50
	ptr = &a
	*ptr += b
	fmt.Println(a, b)
}
