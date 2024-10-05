package main

import (
	"fmt"

	"github.com/macrespo42/Gator/internal/config"
)

type state struct {
	Cfg *config.Config
}

type command struct {
	Name      string
	Arguments []string
}

type commands struct {
	Names map[string]func(*state, command) error
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Arguments) == 0 {
		return fmt.Errorf("the login handler expects a single argument the username")
	}

	s.Cfg.SetUser(cmd.Arguments[0])

	fmt.Println("the user has been set")
	return nil
}

func (c *commands) register(name string, f func(*state, command) error) error {
	_, ok := c.Names[name]
	if ok {
		return fmt.Errorf("a command with this name already exist")
	}

	c.Names[name] = f
	return nil
}

func (c *commands) run(s *state, cmd command) error {
	to_run, ok := c.Names[cmd.Name]

	if !ok {
		return fmt.Errorf("command not found")
	}

	err := to_run(s, cmd)
	if err != nil {
		return err
	}
	return nil
}
