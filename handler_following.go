package main

import (
	"context"
	"fmt"
	"gator/m/v2/internal/database"

	"github.com/google/uuid"
)

func FollowingHandler(s *State, command Command, user database.User) error {

	if len(command.Args) < 1 {
		return fmt.Errorf("usage: %s <FEED_URL>", command.Name)
	}

	ctx := context.Background()

	var userIDNull uuid.NullUUID
	if len(user.ID) > 0 {
		userIDNull = uuid.NullUUID{
			UUID:  user.ID,
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
