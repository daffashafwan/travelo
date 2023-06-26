package main

import (
	"log"
	"os"
	"runtime/debug"

	appl "travelo/cmd"
)

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Llongfile)

	err := appl.Run(logger)
	if err != nil {
		trace := debug.Stack()
		logger.Fatalf("%s\n%s", err, trace)
	}
}
