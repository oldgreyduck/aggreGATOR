package main

import "fmt"


func handlerLogin(s *state, cmd command) error {
    if len(cmd.args) == 0 {
        return fmt.Errorf("username is required")
    }
    if len(cmd.args) > 1 {
        return fmt.Errorf("only one argument (username) is allowed")
    }

    if s == nil || s.cfg == nil {
        return fmt.Errorf("state not initialized")
    }

    username := cmd.args[0]

    if err := s.cfg.SetUser(username); err != nil {
        return err
    }

    fmt.Printf("logged in as %s\n", username)
    return nil
}