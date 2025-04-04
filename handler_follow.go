package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/j-tws/go-aggregator/internal/database"
)

func HandlerFollow(s *state, cmd cmd) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Usage: go run . follow <url>")
	}

	url := cmd.args[0]

	feed, err := s.db.GetFeedByUrl(context.Background(), cmd.args[0])
	if err != nil {
		return fmt.Errorf("Error finding feed with url '%s': %w", url, err)
	}
	
	currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("Error finding username '%s': %w", s.cfg.CurrentUserName, err)
	}

	createFeedFollowParams := database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID: currentUser.ID,
		FeedID: feed.ID,
	}

	createFeedFollowRow, err := s.db.CreateFeedFollow(context.Background(), createFeedFollowParams)
	if err != nil {
		return fmt.Errorf("Error creating feed follow: %w", err)
	}

	fmt.Println("Successfully created feed follow!")
	fmt.Printf("Feed name: %s", createFeedFollowRow.FeedName)
	fmt.Printf("User name: %s", createFeedFollowRow.UserName)

	return nil
}