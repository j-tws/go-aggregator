package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/j-tws/go-aggregator/internal/database"
)

func HandlerBrowse(s *state, cmd cmd, user database.User) error {
	var limit int
	if len(cmd.args) == 0 {
		limit = 2
	} else if len(cmd.args) == 1 {
		parsedLimit, err := strconv.Atoi(cmd.args[0])
		if err != nil {
			return fmt.Errorf("Error parsing limit: %w", err)
		}

		limit = parsedLimit
	} else {
		return fmt.Errorf("Usage: go run . browse <optional_limit>")
	}

	getPostsParams := database.GetPostsParams{
		UserID: user.ID,
		Limit: int32(limit),
	}

	posts, err := s.db.GetPosts(context.Background(), getPostsParams)
	if err != nil {
		return fmt.Errorf("Error fetching posts: %w", err)
	}

	fmt.Println("Here are all your saved posts:")
	for _, post := range posts {
		fmt.Printf("Title: %s\n", post.Title)
		fmt.Printf("URL: %s\n", post.Url)
		fmt.Printf("Description: %s\n", post.Description.String)
		fmt.Printf("Published at: %s\n", post.PublishedAt.Time)
		fmt.Println("------------------------------")
	}

	return nil
}