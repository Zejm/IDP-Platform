package handler

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Zejm/IDP-Platform/internal/service"
)

func handleCommand(input string) {
	parts := strings.Fields(input)

	if len(parts) == 0 {
		return
	}

	command := parts[0]
	args := parts[1:]

	//mt.Println("command:", command)
	fmt.Println("hey")

	switch command {
	case "help":
		fmt.Println("Available commands: help, hello, echo, exit")
	case "echo":
		fmt.Println(strings.Join(args, " "))
	case "hello":
		service.SayHello()
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
