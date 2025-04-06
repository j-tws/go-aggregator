package main

import (
	"context"
	"fmt"

	"github.com/j-tws/go-aggregator/internal/database"
)

func middlewareLoggedIn(handler func(*state, cmd, database.User) error) func(*state, cmd) error {
	handlerFunc := func(s *state, cmd cmd) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return fmt.Errorf("Error finding current user: %w", err)
		}

		return handler(s, cmd, user)
	}

	return handlerFunc
}