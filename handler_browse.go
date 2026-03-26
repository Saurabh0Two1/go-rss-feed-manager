package main

import (
	"context"
	"fmt"
	"gator/m/v2/internal/database"
	"strconv"

	"github.com/google/uuid"
)

func BrowseHandler(s *State, command Command, user database.User) error {

	var limit int64 = 2
	var err error
	if len(command.Args) >= 2 {
		limit, err = strconv.ParseInt(command.Args[1], 10, 32)
		if err != nil {
			return fmt.Errorf("failed to parse limit %v", err)
		}
	}

	var userIDNull uuid.NullUUID
	if len(user.ID) > 0 {
		userIDNull = uuid.NullUUID{
			UUID:  user.ID,
			Valid: true,
		}
	}

	getPostsForUserArgs := database.GetPostsForUserParams{
		UserID: userIDNull,
		Limit:  int32(limit),
	}

	ctx := context.Background()
	posts, err := s.db.GetPostsForUser(ctx, getPostsForUserArgs)

	if err != nil {
		return fmt.Errorf("failed to get posts %v", err)
	}

	for index, post := range posts {
		fmt.Printf("============== %v ================ \n", index+1)
		fmt.Printf("%v", post.Description)
	}

	return nil
}
