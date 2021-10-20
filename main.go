package main

import (
	"bufio"
	"fmt"
	"myapp/doctor"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	whatToSay := doctor.Intro()

	fmt.Println(whatToSay)

	// Need to loop. go only has a FOR loop

	for {

		fmt.Print("-> ")
		userInput, _ := reader.ReadString('\n')

		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)

		if userInput == "quit" {
			break
		} else {
			// call function in doctor.go to get a response
			fmt.Println(doctor.Response(userInput))
		}

	}
}
