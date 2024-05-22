package main

import (
	"fmt"
	"math"
)

func main() {

	var a, b, c float64

	fmt.Scan(&a, &b, &c)

	discriminante := b*b - 4*a*c

	switch {
	case discriminante > 0:
		x1 := (-b + math.Sqrt(discriminante)) / (2 * a)
		x2 := (-b - math.Sqrt(discriminante)) / (2 * a)
		fmt.Println("Due radici reali", x1, x2)
	case discriminante == 0:
		x := -b / (2 * a)
		fmt.Println("Due radici reali coincidenti", x, x)
	case discriminante < 0:
		x1 := complex(-b, math.Sqrt(-discriminante)) / complex(2*a, 0)
		// equivalente a x1 := complex(-b/(2*a), math.Sqrt(-discriminante)/(2*a))
		x2 := complex(-b, -math.Sqrt(-discriminante)) / complex(2*a, 0)
		// equivalente a x2 := complex(-b/(2*a), -math.Sqrt(-discriminante)/(2*a))
		fmt.Println("Due radici complesse", x1, x2)
	}

}
