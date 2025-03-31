package main

import (
	"context"
	"fmt"
)

func HandlerAgg(s *state, cmd cmd) error {
	rssFeed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("Error fetching feed: %w", err)
	}

	fmt.Printf("Feed: %+v\n", rssFeed)

	return nil
}
