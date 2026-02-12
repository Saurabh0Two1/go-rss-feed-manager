package main

import (
	"context"
	"fmt"
	"gator/m/v2/internal/database"
	"time"

	"github.com/google/uuid"
)

func RegisterHandler(s *State, cmd Command) error {
	fmt.Printf("%v", cmd)
	if len(cmd.Args) != 2 {
		err := fmt.Errorf("usage: %s <name>", cmd.Name)
		fmt.Printf("%v", err)
		return err
	}

	ctx := context.Background()
	userName := cmd.Args[1]

	existingUser, err := s.db.GetUser(ctx, userName)

	if err != nil {
		fmt.Println("error in getting user")
	}

	if len(existingUser.Name) > 0 {
		err = s.cfg.SetUser(existingUser.Name)

		if err != nil {
			return fmt.Errorf("couldn't set current user: %w \n", err)
		}
		return nil
	}

	userParams := database.CreateUserParams{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: userName}

	user, err := s.db.CreateUser(ctx, userParams)

	if err != nil {
		fmt.Println("error in creating user")
	}

	err = s.cfg.SetUser(user.Name)

	if err != nil {
		return fmt.Errorf("couldn't set current user: %w \n", err)
	}

	fmt.Printf("Created a new user: %v at %v \n", user.Name, user.CreatedAt.Format(time.RFC1123))

	return nil
}
