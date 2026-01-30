package main

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func FollowingHandler(s *State, command Command) error {

	if len(command.Args) < 1 {
		return fmt.Errorf("usage: %s <FEED_URL>", command.Name)
	}

	currentUserName := s.cfg.CurrentUserName

	fmt.Printf("currentUser: %s", currentUserName)

	ctx := context.Background()
	userRecord, err := s.db.GetUser(ctx, currentUserName)

	if err != nil {
		return fmt.Errorf("failed to get user %v", err)
	}

	var userIDNull uuid.NullUUID
	if len(userRecord.ID) > 0 {
		userIDNull = uuid.NullUUID{
			UUID:  userRecord.ID,
			Valid: true,
		}
	}

	feedFollows, err := s.db.GetFeedFollowsForUser(ctx, userIDNull)

	if err != nil {
		return fmt.Errorf("failed to get feeds %v", err)
	}

	fmt.Println("Here are the feeds you are following:")

	for _, feed := range feedFollows {
		feedRecord, err := s.db.GetFeedById(ctx, feed.FeedID.UUID)
		if err != nil {
			return fmt.Errorf("failed to get feed %v", err)
		}
		fmt.Printf("%s \n", feedRecord.Name)
	}

	return nil
}
