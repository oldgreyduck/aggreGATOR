package main

import (
    "fmt"
    "os"
    "gator/internal/config"
)

func main() {
    cfg, err := config.Read()
if err != nil {
    // if file missing, start with empty config
    if os.IsNotExist(err) {
        cfg = config.Config{}
    } else {
        fmt.Println(err)
        os.Exit(1)
    }
}
    s := &state{cfg: &cfg}

    cmds := &commands{m: map[string]func(*state, command) error{}}
    cmds.register("login", handlerLogin)

    if len(os.Args) < 2 {
        fmt.Println("a command is required")
        os.Exit(1)
    }
    name := os.Args[1]
    args := os.Args[2:]
    cmd := command{name: name, args: args}

    if err := cmds.run(s, cmd); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
