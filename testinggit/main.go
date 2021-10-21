package main

import (
	"fmt"
	"strconv"
)

func main() {
	// one way - declare then assign (two steps)
	firstNumber := 2

	fmt.Println("This thing is neat " + strconv.Itoa(firstNumber))
}
