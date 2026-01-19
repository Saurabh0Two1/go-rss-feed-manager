package main

import (
	"context"
	"fmt"
)

func UsersHandler(s *State, command Command) error {

	ctx := context.Background()
	users, err := s.db.GetUsers(ctx)

	if err != nil {
		return fmt.Errorf("Failed to get users due to: %w", err)
	}

	currentUser := s.cfg.CurrentUserName

	for _, user := range users {
		if user == currentUser {
			fmt.Printf("%s (current) \n", user)
		} else {
			fmt.Println(user)
		}
	}

	return nil
}
