package main

import (
	"fmt"
	"os"

	"github.com/bjvanbemmel/kyoto/cli"
	"github.com/bjvanbemmel/kyoto/cli/commands"
)

func main() {
	var init cli.Command = &commands.Init{}
	var help cli.Command = &commands.Help{}
	var version cli.Command = &commands.Version{}
	var serve cli.Command = &commands.Serve{}

	init.Init()
	cli.Commands.Append(init)

	help.Init()
	cli.Commands.Append(help)

	version.Init()
	cli.Commands.Append(version)

	serve.Init()
	cli.Commands.Append(serve)

	if len(os.Args) < 2 {
		help.Execute()

		return
	}

	arg, err := cli.Commands.Get(os.Args[1])

	if err != nil {
		fmt.Println(err)
		fmt.Printf("Try '%s' to list all commands.\n", help.Usage())

		os.Exit(1)
	}

	err = arg.Execute(os.Args[2:]...)

	if err != nil {
		panic(err)
	}
}
