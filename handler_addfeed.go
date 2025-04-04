package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/j-tws/go-aggregator/internal/database"
)

func HandlerAddFeed(s *state, cmd cmd) error {
	if len(cmd.args) < 2 {
		return fmt.Errorf("URL and name of feed is required")
	}

	feedName := cmd.args[0]
	feedUrl := cmd.args[1]

	currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("Error finding current user: %w", err)
	}

	params := database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: feedName,
		Url: feedUrl,
		UserID: currentUser.ID,
	}

	feed, err := s.db.CreateFeed(context.Background(), params)

	if err != nil {
		return fmt.Errorf("Error creating feed: %w", err)
	}

	fmt.Println("Successfully created feed:")
	fmt.Println(feed)

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