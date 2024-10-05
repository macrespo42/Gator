package main

import (
	"fmt"
	"log"
	"os"

	"github.com/macrespo42/Gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal("Error when reading config file")
	}

	s := state{
		Cfg: &cfg,
	}

	cmds := commands{
		Names: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)

	if len(os.Args) < 2 {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	cmd := command{
		Name:      os.Args[1],
		Arguments: os.Args[2:],
	}

	err = cmds.run(&s, cmd)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
}
