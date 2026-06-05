package handler

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func handleCommand(input string) {
	parts := strings.Fields(input)

	if len(parts) == 0 {
		return
	}

	command := parts[0]
	args := parts[1:]

	switch command {
	case "help":
		fmt.Println("")
	case "echo":
		fmt.Println(strings.Join(args, " "))
	default:
		fmt.Println("Unknown command. Try 'help' for assistance")
	}
}

func RunCLI() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			break
		}

		input = strings.TrimSpace(input)

		if input == "exit" {
			break
		}

		handleCommand(input)
	}
}
