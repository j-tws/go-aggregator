package main

import (
	"context"
	"fmt"

	"github.com/j-tws/go-aggregator/internal/database"
)

func HandlerUnfollow(s *state, cmd cmd, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Usage: go run . unfollow <feed url>")
	}

	feedToDelete, err := s.db.GetFeedByUrl(context.Background(), cmd.args[0])
	if err != nil {
		return fmt.Errorf("Error fetching feed via url name: %w", err)
	}

	feedFollowDeleteErr := s.db.DeleteFeedFollows(context.Background(), database.DeleteFeedFollowsParams{
		UserID: user.ID,
		FeedID: feedToDelete.ID,
	})

	if feedFollowDeleteErr != nil {
		return fmt.Errorf("Error deleting feed follow: %w", err)
	}

	fmt.Printf("%s unfollowed successfully!\n", feedToDelete.Name)

	return nil
}