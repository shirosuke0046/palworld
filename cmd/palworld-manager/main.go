package main

import (
	"context"
	"fmt"
	"os"
)

var (
	version  string
	revision string
)

func main() {
	if err := rootCommand.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
