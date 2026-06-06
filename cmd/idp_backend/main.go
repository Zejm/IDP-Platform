package main

import (
	"github.com/Zejm/IDP-Platform/cmd/network"
	"github.com/Zejm/IDP-Platform/internal/handler"
)

func main() {
	go network.RunNetwork()
	handler.RunCLI()
}
