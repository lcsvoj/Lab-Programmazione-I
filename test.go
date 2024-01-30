package main

import (
	"fmt"
	"time"
)

func main() {
	// Define a time.Time object, here we use the current time
	d := time.Now()

	// Format the date
	x := d.Format("2006-01-02")

	// Print the formatted date
	fmt.Println(x)
}
