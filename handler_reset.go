package main

import (
	"context"
	"fmt"
)

func HandlerReset(s *state, cmd cmd) error {
	err := s.db.ResetUsers(context.Background())

	if err != nil {
		return fmt.Errorf("Error resetting users table: %w", err)
	}

	fmt.Println("Users table successfully resetted.")

	return nil
}