package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func HandlerAgg(s *state, cmd cmd) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Usage: go run . agg <time>")
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return fmt.Errorf("Error parsing time with '%s': %w", cmd.args[0], err)
	}
	fmt.Printf("Collecting feeds every %v\n", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <- ticker.C {
		err := scrapFeeds(s)
		if err != nil {
			return fmt.Errorf("error scraping feed: %w", err)
		}
	}
}

func scrapFeeds(s *state) error {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("Error fetching next feed: %w", err)
	}

	if err := s.db.MarkFeedFetched(context.Background(), nextFeed.ID); err != nil {
		return fmt.Errorf("Error marking feed as fetched: %w", err)
	}

	markedFeed, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return fmt.Errorf("Error fetching feed via url '%s': %w", nextFeed.Url, err)
	}

	for _, item := range markedFeed.Channel.Item {
		fmt.Printf("Found post: %s\n", item.Title)
	}

	log.Printf("Feed %s collected, %v posts found", nextFeed.Name, len(markedFeed.Channel.Item))
	return nil
}
