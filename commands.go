package main

import (
    "fmt"
    "gator/internal/config"
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
    if c.m == nil {
        c.m = make(map[string]func(*state, command) error)
    }
    c.m[name] = f
}

func (c *commands) run(s *state, cmd command) error {
    h, ok := c.m[cmd.name]
    if !ok {
        return fmt.Errorf("unknown command: %s", cmd.name)
    }
    return h(s, cmd)
}