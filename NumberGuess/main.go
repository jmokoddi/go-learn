package main

import (
	"bufio"
	"fmt"
	"os"
)

// outside main function makes it available to everything
// never changes and DONT have to be used in program.
const prompt = "and press ENTER when ready."

func main() {
	var firstNumber int
	var secondNumber int
	var subtraction int
	var answer int
	reader := bufio.NewReader(os.Stdin)

	firstNumber = 2
	secondNumber = 5
	subtraction = 7

	// display a welcome/instructions
	fmt.Println("Guess the Number Game")
	fmt.Println("---------------------")
	fmt.Println("")

	fmt.Println("Think of a number between 1 and 10", prompt)
	reader.ReadString('\n')

	// take them through the games
	fmt.Println("Multiply your number by", firstNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Now multiply your number by", secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Divide the result by the number you originally thought of", prompt)
	reader.ReadString('\n')

	fmt.Println("Now subtract", subtraction, prompt)

	// give them the answer
	answer = firstNumber*secondNumber - subtraction
	fmt.Println("The answer is", answer)
}
