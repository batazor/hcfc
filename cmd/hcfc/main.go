package main

import (
	"github.com/batazor/hcfc/cmd/cli"
	"github.com/batazor/hcfc/pkg/logger"
	"log"
)

func init() {
	// Set format log
	log.SetFlags(0)
	log.SetOutput(new(logger.LogWriter))
}

func main() {
	cli.Execute()
}
