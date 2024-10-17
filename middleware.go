package main

import (
	"context"
)

func middlewareLoggedIn(handler func(s *state, cmd command) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		_, err := s.Db.GetUser(context.Background(), s.Cfg.CurrentUserName)
		if err != nil {
			return err
		}
		return handler(s, cmd)
	}
}
