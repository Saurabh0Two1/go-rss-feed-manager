package main

import (
	"context"
	"fmt"
	"gator/m/v2/internal/database"
	"time"

	"github.com/google/uuid"
)

func FollowHandler(s *State, command Command) error {

	if len(command.Args) < 2 {
		return fmt.Errorf("usage: %s <FEED_URL>", command.Name)
	}

	url := command.Args[1]
	currentUserName := s.cfg.CurrentUserName
	ctx := context.Background()

	userRecord, err := s.db.GetUser(ctx, currentUserName)

	if err != nil {
		return fmt.Errorf("failed to get current user %v", err)
	}

	var userIDNull uuid.NullUUID
	if len(userRecord.ID) > 0 {
		userIDNull = uuid.NullUUID{
			UUID:  userRecord.ID,
			Valid: true,
		}
	}

	existingFeed, err := s.db.GetFeedByUrl(ctx, url)

	if err != nil || len(userRecord.ID) == 0 {
		feedParams := database.CreateFeedParams{
			ID:        uuid.New(),
			Url:       url,
			UserID:    userIDNull,
			Name:      url,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		existingFeed, err = s.db.CreateFeed(ctx, feedParams)

		if err != nil {
			return fmt.Errorf("failed to create feed %v", err)
		}
	}

	var feedIDNull uuid.NullUUID

	if len(existingFeed.ID) > 0 {
		feedIDNull = uuid.NullUUID{
			UUID:  existingFeed.ID,
			Valid: true,
		}
	}

	feedFollowsParams := database.CreateFeedFollowsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    userIDNull,
		FeedID:    feedIDNull,
	}

	feedFollow, err := s.db.CreateFeedFollows(ctx, feedFollowsParams)

	if err != nil {
		return fmt.Errorf("failed to create feedFollow %v", err)
	}

	fmt.Printf("Successfully following the feed - feed %s for current user %s", feedFollow.FeedName, feedFollow.UserName)

	return nil
}
