package main

import (
	"gator/internal/config"
	"errors"
)

type state struct {
    cfg *config.Config
}

type command struct {
    name string
    args []string
}

type commands struct {
    m map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
    // TODO
}

func (c *commands) run(s *state, cmd command) error {
    // TODO: look up handler, return error if not found
    return errors.New("unregistered command")
}