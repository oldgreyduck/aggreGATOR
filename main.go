package main

import (
	"fmt"
	"gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("read error:", err)
		return
	}

	if err := (&cfg).SetUser("yourname"); err != nil {
		fmt.Println("set user error:", err)
		return
	}

	cfg2, err := config.Read()
	if err != nil {
		fmt.Println("read2 error:", err)
		return
	}

	fmt.Println(cfg2)
}
