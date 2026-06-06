package handler

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Zejm/IDP-Platform/internal/service"
)

func DisplayResult(text string) {
	fmt.Println("<", text)
}

func handleCommand(input string) {
	parts := strings.Fields(input)

	if len(parts) == 0 {
		return
	}

	command := parts[0]
	args := parts[1:]

	switch command {
	case "help":
		DisplayResult("Available commands: help, hello, echo, exit")

	case "add":
		if len(args) < 2 {
			DisplayResult("Arguments not specified. Correct: add -u <username>")
			return
		}

		arg := args[0]
		usrName := strings.TrimSpace(args[1])

		if arg == "-u" && usrName != "" {
			DisplayResult(service.AddUser(usrName))
		}

	default:
		DisplayResult("Unknown command. Try 'help' for assistance")
	}
}

func RunCLI() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")

		input, err := reader.ReadString('\n')
		if err != nil {
			DisplayResult(fmt.Sprintf("Error reading input: %s", err))
			break
		}

		input = strings.TrimSpace(input)

		if input == "exit" {
			DisplayResult("Exiting...")
			break
		}

		handleCommand(input)
	}
}
