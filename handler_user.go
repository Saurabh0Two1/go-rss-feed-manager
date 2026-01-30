package main

import (
	"context"
	"fmt"
	"syscall"
)

func LoginHandler(s *State, cmd Command) error {
	if len(cmd.Args) != 2 {
		err := fmt.Errorf("usage: %s <name>", cmd.Name)
		fmt.Printf("%v", err)
		syscall.Exit(1)
		return err
	}

	ctx := context.Background()
	userName := cmd.Args[1]

	existingUser, err := s.db.GetUser(ctx, userName)

	if err != nil {
		fmt.Printf("error in getting user:-  %v \n", err)
		syscall.Exit(1)
		return nil
	}

	err = s.cfg.SetUser(existingUser.Name)

	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Print("User switched successfully!")

	return nil
}
