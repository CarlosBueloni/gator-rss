package main

import (
	"errors"
	"fmt"
)

type command struct {
	Name string
	Args []string
}

type commands struct {
	Commands map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.Commands[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	f, ok := c.Commands[cmd.Name]
	if !ok {
		fmt.Println("command not registered")
		return errors.New("command not registered")
	}
	err := f(s, cmd)
	if err != nil {
		return err
	}
	return nil
}
