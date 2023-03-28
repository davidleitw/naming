package main

import (
	"os"

	"github.com/davidleitw/naming/pkg/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
