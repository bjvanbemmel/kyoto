package cli

import (
	"errors"
	"fmt"
)

var Commands commands = commands{
	Commands: make(map[string]Command),
}

type Command interface {
	Init() error
	Execute(...string) error
	Name() string
	Description() string
	Usage() string
}

type commands struct {
	Commands map[string]Command
}

func (h *commands) Append(arg Command) error {
	h.Commands[arg.Name()] = arg

	return nil
}

func (h commands) Get(key string) (Command, error) {
	arg := h.Commands[key]

	if arg == nil {
		return nil, errors.New(fmt.Sprintf("kyoto: '%s' is not an available command.", key))
	}

	return arg, nil
}
