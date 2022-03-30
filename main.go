package main

import (
	"github.com/alganbr/kedai-authsvc/cmd"
	"os"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
