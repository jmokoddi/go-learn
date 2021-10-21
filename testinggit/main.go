package main

import (
	"fmt"
	"strconv"
)

func main() {
	// one way - declare then assign (two steps)
	firstNumber := 2

	fmt.Println("This thing is neat " + strconv.Itoa(firstNumber))

	fmt.Println("This is my new feature")

	fmt.Println(("This is the second new feature!"))
}
