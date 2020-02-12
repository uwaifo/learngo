package main

import (
	"fmt"
	"os"
)

const (
	usgae       = "Usage: [username] [password]"
	errUser     = "Access deied for %s \n"
	errPassword = "Invalid password for %s \n"
	granted     = "Access granted to  %s.\n"
)

func main() {
	args := os.Args
	username, password := "james", "bond"

	if len(args) != 3 {
		//if len(argUsername) < 0 || len(argPassword) < 0 {
		fmt.Println(usgae)
		return

	} else if args[1] != username {
		fmt.Printf(errUser, args[1])

	} else if args[2] != password {
		fmt.Printf(errPassword, args[2])

	} else {
		fmt.Printf(granted, args[1])

	}

}
