package main

import (
	"context"
	"fmt"
)

func HandlerFollowing(s *state, cmd cmd) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("Usage: go run . following")
	}
	
	currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("Error finding current user: %w", err)
	}

	userFeeds, err := s.db.GetFeedFollowsForUser(context.Background(), currentUser.ID)
	if err != nil {
		return fmt.Errorf("Error getting user feeds: %w", err)
	}

	fmt.Printf("Current feeds for user '%s':\n", currentUser.Name)
	for i, userFeed := range userFeeds {
		fmt.Printf("%d, Feed Name: %s\n", i, userFeed.Name)
	}

	return nil
}