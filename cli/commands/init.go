package commands

import (
	"errors"
	"fmt"

	"github.com/bjvanbemmel/kyoto/project"
)

type Init struct {
	name        string
	description string
	usage       string
}

func (i *Init) Init() error {
	i.name = "init"
	i.description = "Initializes a new Kyoto project"
	i.usage = "kyoto init <project title>"

	return nil
}

func (i Init) Execute(args ...string) error {
	if len(args) < 1 {
		return errors.New("No trailing argument(s) or value(s)!")
	}

	if err := project.Generate(args[0]); err != nil {
		return err
	}

	fmt.Printf("Successfully created '%s'. We hope you build something lovely!\n", args[0])
	fmt.Printf("\nRun 'cd %s' to enter the new directory.\n", args[0])

	return nil
}

func (i Init) Name() string {
	return i.name
}

func (i Init) Description() string {
	return i.description
}

func (i Init) Usage() string {
	return i.usage
}
