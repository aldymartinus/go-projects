package main

import (
	"fmt"
)

func main() {
	var pyramidHeight int

	fmt.Print("How tall your pyramid need to be? ")
	fmt.Scan(&pyramidHeight)

	fmt.Print("\n")

	for i := 0; i < pyramidHeight; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print("* ")
		}
		fmt.Print("\n")
	}
}
