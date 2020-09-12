package main

import (
	"log"

	"github.com/rapde/rap/cmd"
	"github.com/rapde/rap/lib/action"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("[RAP] ")

	action.EnsureEnv()
	cmd.Execute()
}
