package main

import (
	"github.com/MartinRusso28/matrixfmt/pkg/server"
)

func main() {
	server := server.GetMainEngine()

	err := server.Run(":8080")

	if err != nil {
		panic("Error running the server")
	}
}
