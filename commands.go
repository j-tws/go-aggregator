package main

type cmd struct {
	name string
	args []string
}

type commands struct {
	list map[string]func(*state, cmd) error
}

func (c *commands) register(name string, f func(*state, cmd) error) {
	c.list[name] = f
}

func (c *commands) run(s *state, cmd cmd) error {
	handlerFunc := c.list[cmd.name]
	err := handlerFunc(s, cmd)
	if err != nil {
		return err
	}

	return nil
}
