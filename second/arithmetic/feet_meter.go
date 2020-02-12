package main

import (
	"fmt"
	"os"
	"strconv"
)

//FeetMeter . .. .
func FeetMeter() {
	argument := os.Args[1]
	feet, _ := strconv.ParseFloat(argument, 63)
	meters := feet * 0.3048

	fmt.Printf("%g feet is %g meters.\n", feet, meters)

}
