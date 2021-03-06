package main

import (
	"log"

	cmd "github.com/xielizyh/goprj-tour/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute error: %v", err)
	}
}
