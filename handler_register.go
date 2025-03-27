package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/j-tws/go-aggregator/internal/database"
)

func HandlerRegister(s *state, cmd cmd) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("Name is required to register.")
	}

	name := cmd.args[0]

	params := database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: name,
	}

	newUser, err := s.db.CreateUser(context.Background(), params)
	if err != nil {
		return fmt.Errorf("Error creating user: %w", err)
	}

	if err := s.cfg.SetUser(newUser.Name); err != nil {
		return fmt.Errorf("Could not set user: %w", err)
	}

	fmt.Printf("User '%s' successfully registered and set as current user!", newUser.Name)
	return nil
}