package commands

import (
	"fmt"

	"github.com/bjvanbemmel/kyoto/cli"
)

type Help struct {
	name        string
	description string
	usage       string
}

func (h *Help) Init() error {
	h.name = "help"
	h.description = "Shows an overview of all available commands"
	h.usage = "kyoto help"

	return nil
}

func (h Help) Execute(...string) error {
	fmt.Println("\nUsage: kyoto [COMMAND] [OPTIONS] [ARGUMENTS]")
	fmt.Printf("\nA lightweight static site generator.\n\n")

	fmt.Println("Commands:")
	for _, arg := range cli.Commands.Commands {
		fmt.Printf("    %-32s%s\n", arg.Usage(), arg.Description())
	}

	fmt.Println("\nFor more information, navigate to https://github.com/bjvanbemmel/kyoto")

	return nil
}

func (h Help) Name() string {
	return h.name
}

func (h Help) Description() string {
	return h.description
}

func (h Help) Usage() string {
	return h.usage
}
