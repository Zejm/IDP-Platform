package main

import (
	"fmt"

	"github.com/Zejm/IDP-Platform/cmd/network"
	"github.com/Zejm/IDP-Platform/internal/handler"
)

func main() {
	fmt.Println("IDP Platform backbone started")

	go network.RunNetwork()
	handler.RunCLI()
}
