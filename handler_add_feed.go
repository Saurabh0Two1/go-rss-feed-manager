package main

import (
	"context"
	"fmt"
	"gator/m/v2/internal/database"
	"syscall"
	"time"

	"github.com/google/uuid"
)

func AddFeedHandler(s *State, cmd Command) error {
	if len(cmd.Args) != 3 {
		err := fmt.Errorf("usage: %s <name>", cmd.Name)
		fmt.Printf("%v", err)
		return err
	}

	ctx := context.Background()
	name := cmd.Args[1]
	url := cmd.Args[2]
	currentUser := s.cfg.CurrentUserName

	existingUser, err := s.db.GetUser(ctx, currentUser)

	if err != nil {
		fmt.Printf("error in getting user:-  %v \n", err)
		syscall.Exit(1)
		return nil
	}

	var userIDNull uuid.NullUUID

	if len(existingUser.ID) > 0 {
		userIDNull = uuid.NullUUID{
			UUID:  existingUser.ID,
			Valid: true,
		}
	}

	feedParams := database.CreateFeedParams{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), UserID: userIDNull, Name: name, Url: url}

	feed, err := s.db.CreateFeed(ctx, feedParams)

	if err != nil {
		fmt.Println("error in creating feed")
	}

	var feedIDNull uuid.NullUUID

	if len(feed.ID) > 0 {
		feedIDNull = uuid.NullUUID{
			UUID:  feed.ID,
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

	_, err = s.db.CreateFeedFollows(ctx, feedFollowsParams)

	if err != nil {
		return fmt.Errorf("failed to create feed follow")
	}

	fmt.Println(feed)

	return nil
}
