package main

import (
	"context"
	"fmt"
)

func HandlerLogin(s *state, cmd cmd) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("Username is required for login")
	}

	name := cmd.args[0]
	user, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		return fmt.Errorf("No such user in database. Please register.")
	}

	if err := s.cfg.SetUser(user.Name); err != nil {
		return err
	}

	fmt.Printf("User has been set as %s.", cmd.args[0])
	return nil
}
