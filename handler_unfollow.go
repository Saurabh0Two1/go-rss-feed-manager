package main

import (
	"context"
	"fmt"
	"gator/m/v2/internal/database"

	"github.com/google/uuid"
)

func UnfollowHandler(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	ctx := context.Background()
	url := cmd.Args[1]

	feedData, err := s.db.GetFeedByUrl(ctx, url)

	if err != nil {
		return fmt.Errorf("failed to get feed: %s", err)
	}

	var userIDNull uuid.NullUUID

	if len(user.ID) > 0 {
		userIDNull = uuid.NullUUID{
			UUID:  user.ID,
			Valid: true,
		}
	}

	var feedIDNull uuid.NullUUID

	if len(feedData.ID) > 0 {
		feedIDNull = uuid.NullUUID{
			UUID:  feedData.ID,
			Valid: true,
		}
	}

	var unfollowFeedParams database.UnfollowFeedParams

	unfollowFeedParams.UserID = userIDNull
	unfollowFeedParams.FeedID = feedIDNull

	_, err = s.db.UnfollowFeed(ctx, unfollowFeedParams)

	if err != nil {
		return fmt.Errorf("failed to unfollow: %s", err)
	}

	fmt.Printf("%s unfollowed successfully!\n", feedData.Name)

	return nil
}
