package main

import (
	"fmt"
	"strconv"
)

func main() {

	sl := []int64{1, 2, 4, 5, 8, 15, 32}
	sl1 := converti(sl)
	sl2 := convertiPari(sl)

	fmt.Println(sl1)
	fmt.Println(sl2)
}

func converti(slIn []int64) []string {

	slOut := make([]string, len(slIn))

	for i := 0; i < len(slIn); i++ {
		slOut[i] = strconv.FormatInt(slIn[i], 2)
	}

	return slOut
}

func convertiPari(slIn []int64) []string {

	var slOut []string

	for i := 0; i < len(slIn); i++ {
		if slIn[i]%2 == 0 {
			slOut = append(slOut, strconv.FormatInt(slIn[i], 2))
		}
	}

	return slOut
}
