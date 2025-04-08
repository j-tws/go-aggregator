package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/j-tws/go-aggregator/internal/database"
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
		if err := createPostFromFeed(s, item, nextFeed.ID); err != nil {
			if strings.Contains(err.Error(), `duplicate key value violates unique constraint "posts_url_key"`) {
				fmt.Println("Post already created. Moving on...")
			} else {
				return fmt.Errorf("Error creating post: %w", err)
			}
		}
	}

	log.Printf("Feed %s collected, %v posts found", nextFeed.Name, len(markedFeed.Channel.Item))
	return nil
}

func createPostFromFeed(s *state, post RSSItem, feedID uuid.UUID) error {
	layout := "Mon, 02 Jan 2006 15:04:05 -0700"
	publishedTime, err := time.Parse(layout, post.PubDate)
	if err != nil {
		return fmt.Errorf("Error parsing published time: %w", err)
	}

	postParams := database.CreatePostParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Title: post.Title,
		Url: post.Link,
		Description: sql.NullString{String: post.Description, Valid: true},
		PublishedAt: sql.NullTime{Time: publishedTime, Valid: true},
		FeedID: feedID,
	}

	createdPost, err := s.db.CreatePost(context.Background(), postParams)
	if err != nil {
		return fmt.Errorf("Error creating post: %w", err)
	}

	fmt.Printf("Post successfully created: %s\n", createdPost.Title)
	return nil
}
