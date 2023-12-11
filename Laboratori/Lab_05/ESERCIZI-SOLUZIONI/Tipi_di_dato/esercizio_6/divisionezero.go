package main

import (
	"fmt"
	"reflect"
)

func main() {
	var numero float64 = 0
	fmt.Println("Quando 0 è", reflect.TypeOf(numero), ":")

	fmt.Println("0/0 =", 0/numero)

	numero = 1 / numero
	fmt.Println("1/0 = ", numero)

	numero = 1 / numero
	fmt.Println("1/+Inf =", numero)

	numero = (-1) / numero
	fmt.Println("-(1/0) =", numero)

	numero = numero - numero
	fmt.Println("-Inf + +Inf =", numero)

	numero = numero + numero
	fmt.Println("-Inf + -Inf =", numero)
}
