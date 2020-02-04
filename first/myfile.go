package main

import "fmt"
import (
	output "fmt"
	"runtime"
)

func profile() {
	fmt.Print("Loading your ")
	output.Println("profile . . ")
	output.Println("CPU COUNT:", runtime.NumCPU()*2)
}
