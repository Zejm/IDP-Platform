package main

import (
	"fmt"

	"github.com/Zejm/IDP-Platform/internal/handler"
)

func main() {
	fmt.Println("IDP Platform backbone started")

	handler.RunCLI()
}
