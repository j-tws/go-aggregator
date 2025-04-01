package main

import (
	"context"
	"fmt"
)

func HandlerFeeds(s *state, cmd cmd) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("Error fetching all feeds: %w", err)
	}

	for _, feed := range feeds {
		fmt.Println("----------")
		fmt.Printf("Feed name: %s\n", feed.Name)
		fmt.Printf("Feed URL: %s\n", feed.Url)
		fmt.Printf("Feed User Name: %s\n", feed.Name_2)
		fmt.Println("----------")
	}

	return nil
}