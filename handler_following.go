package main

import (
	"context"
	"fmt"

	"github.com/j-tws/go-aggregator/internal/database"
)

func HandlerFollowing(s *state, cmd cmd, user database.User) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("Usage: go run . following")
	}
	
	userFeeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("Error getting user feeds: %w", err)
	}

	fmt.Printf("Current feeds for user '%s':\n", user.Name)
	for i, userFeed := range userFeeds {
		fmt.Printf("%d, Feed Name: %s\n", i, userFeed.Name)
	}

	return nil
}