package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/macrespo42/Gator/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Arguments) < 2 {
		return fmt.Errorf("Please provie the name and url of the feed as arguments")
	}

	user, err := s.Db.GetUser(context.Background(), s.Cfg.CurrentUserName)
	if err != nil {
		return err
	}

	feedParams := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		name:      cmd.Arguments[0],
		Url:       cmd.Arguments[1],
		UserID:    user,
	}

	feed, err := s.Db.CreateFeed(context.Background(), feedParams)
	if err != nil {
		return err
	}

	fmt.Printf("%v\n", feed)
	return nil
}

func handlerUsers(s *state, _ command) error {
	users, err := s.Db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	for index := range users {
		if s.Cfg.CurrentUserName == users[index].Name {
			fmt.Printf("* %s (current)\n", users[index].Name)
		} else {
			fmt.Printf("* %s\n", users[index].Name)
		}
	}
	return nil
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
