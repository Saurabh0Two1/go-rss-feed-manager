package main

import (
	"context"
	"gator/m/v2/internal/database"
)

func MiddlewareLoggedIn(handler func(s *State, cmd Command, user database.User) error) func(*State, Command) error {

	return func(s *State, cmd Command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)

		if err != nil {
			return err
		}
		return handler(s, cmd, user)
	}
}
