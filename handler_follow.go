package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/j-tws/go-aggregator/internal/database"
)

func HandlerFollow(s *state, cmd cmd, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Usage: go run . follow <url>")
	}

	url := cmd.args[0]

	feed, err := s.db.GetFeedByUrl(context.Background(), cmd.args[0])
	if err != nil {
		return fmt.Errorf("Error finding feed with url '%s': %w", url, err)
	}

	createFeedFollowParams := database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID: user.ID,
		FeedID: feed.ID,
	}

	createFeedFollowRow, err := s.db.CreateFeedFollow(context.Background(), createFeedFollowParams)
	if err != nil {
		return fmt.Errorf("Error creating feed follow: %w", err)
	}

	fmt.Println("Successfully followed feed!")
	fmt.Printf("Feed name: %s\n", createFeedFollowRow.FeedName)
	fmt.Printf("User name: %s\n", createFeedFollowRow.UserName)

	return nil
}