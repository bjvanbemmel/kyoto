package commands

type Serve struct {
    name string
    description string
    usage string
}

func (s *Serve) Init() error {
    s.name = "serve"
    s.description = "Starts an HTTP webserver to serve the generated website"
    s.usage = "kyoto serve"

    return nil
}

func (s Serve) Execute(args ...string) error {
    return nil
}

func (s Serve) Name() string {
    return s.name
}

func (s Serve) Description() string {
    return s.description
}

func (s Serve) Usage() string {
    return s.usage
}
