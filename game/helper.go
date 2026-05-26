package game

import (
	"bufio"
	"fmt"
)

func readInput(scanner *bufio.Scanner) (string, bool) {
	if !scanner.Scan() {
		fmt.Println("\nGoodbye")
		return "", true
	}

	text := scanner.Text()

	if text == "exit" {
		fmt.Println("Goodbye")
		return "", true
	}

	return text, false

}
