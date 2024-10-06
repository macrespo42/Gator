package main

import (
	"fmt"

	"github.com/macrespo42/Gator/internal/config"
	"github.com/macrespo42/Gator/internal/database"
)

type state struct {
	Cfg *config.Config
	Db  *database.Queries
}

type command struct {
	Name      string
	Arguments []string
}

type commands struct {
	Names map[string]func(*state, command) error
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
