package main

import "fmt"

func HandlerLogin(s *state, cmd cmd) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("Username is required for login")
	}

	err := s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return err
	}

	fmt.Printf("User has been set as %s.", cmd.args[0])
	return nil
}
