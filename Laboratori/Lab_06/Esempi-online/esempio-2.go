package main

import (
	"fmt"
	"strconv"
)

func main() {
	
	var err error 

	/*
		error is a built-in interface type in Go

		An error variable, if different from 'nil', can describe itself as a string

	*/
	
	var s1 string = "3.1415926535"
	var s2 string = "3"

	
	var v float64
	
	if v, err = strconv.ParseFloat(s1, 32); err != nil {
		fmt.Println(err)
		// return
	}
	fmt.Println("Con bitsize 32:",v)
	if v, err = strconv.ParseFloat(s1, 64); err != nil {
		fmt.Println(err)
		// return
	}
	fmt.Println("Con bitsize 64:",v)
	
	var i int
	
	if i, err = strconv.Atoi(s1); err != nil {
		fmt.Println(err)
		// return
	}
	fmt.Println(i)
	
	if i, err = strconv.Atoi(s2); err != nil {
		fmt.Println(err)
		// return
	}
	fmt.Println(i)
	
	s3 := "Ciaonu" 
	s4 := "Ciaone"
	
	fmt.Println("Ordine alfabetico:")
	if (s3 < s4) {
		fmt.Println(s3)
		fmt.Println(s4)
	} else {
		fmt.Println(s4)
		fmt.Println(s3)
	}
	
}
