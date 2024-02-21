package main

import (
	"github.com/charmbracelet/log"
	"github.com/osean-man/pinner/cmd"
)

func main() {
	log.SetTimeFormat("")
	cmd.Execute()
}
