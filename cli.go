package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
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

func handlerReset(s *state, cmd command) error {
	err := s.Db.DeleteAllUsers(context.Background())
	if err != nil {
		return err
	}

	fmt.Printf("Command %s executed with success.\n", cmd.Name)
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Arguments) == 0 {
		return fmt.Errorf("the register handler expects a single argument the username")
	}

	userParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Arguments[0],
	}

	usr, err := s.Db.CreateUser(context.Background(), userParams)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	s.Cfg.SetUser(usr.Name)
	fmt.Printf("%v\n", usr)

	return nil
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Arguments) == 0 {
		return fmt.Errorf("the login handler expects a single argument the username")
	}

	_, err := s.Db.GetUser(context.Background(), cmd.Arguments[0])
	if err != nil {
		return fmt.Errorf("can't login the user does not exist")
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
