package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arg := os.Args[1]
	feet, err := strconv.ParseFloat(arg, 63)
	if err != nil {
		fmt.Printf("ERROR: %q is not a number. %s\n", arg, err)
		return
	}
	meters := feet * 0.3048

	fmt.Printf("%g feet is %g meters.\n", feet, meters)
}
