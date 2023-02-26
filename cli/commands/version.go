package commands

import "fmt"

type Version struct {
	name        string
	description string
	usage       string
}

func (v *Version) Init() error {
	v.name = "version"
	v.description = "Shows the presently installed version"
	v.usage = "kyoto version"

	return nil
}

func (v Version) Execute(args ...string) error {
	fmt.Println("v0.1.0 (Development build)")

	return nil
}

func (v Version) Name() string {
	return v.name
}

func (v Version) Description() string {
	return v.description
}

func (v Version) Usage() string {
	return v.usage
}
