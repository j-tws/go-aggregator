package main

import (
	"context"
	"fmt"
)

func HandlerUsers(s *state, cmd cmd) error {
	usersName, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Error fetching all users: %w", err)
	}

	for _, name := range usersName {
		if name == s.cfg.CurrentUserName {
			fmt.Printf("* %s (current)\n", name)
			} else {
			fmt.Printf("* %s\n", name)
		}
	}

	return nil
}
