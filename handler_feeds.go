package main

import (
	"context"
	"fmt"
)

func FeedsHandler(s *State, command Command) error {
	ctx := context.Background()
	feeds, err := s.db.GetFeeds(ctx)
	if err != nil {
		return fmt.Errorf("error fetching feeds %v", err)
	}

	for _, feed := range feeds {
		fmt.Printf("%s \n", feed.Name)
		fmt.Printf("%s \n", feed.Url)
		user, err := s.db.GetUserById(ctx, feed.UserID.UUID)
		if err != nil {
			return fmt.Errorf("error fetching user with given userID %v", err)
		}

		fmt.Printf("%s \n", user.Name)
	}

	return nil
}
