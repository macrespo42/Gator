package main

import (
	"fmt"
	"log"

	"github.com/macrespo42/Gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal("Error when reading config file")
	}

	cfg.SetUser("macrespo")

	cfg, err = config.Read()
	if err != nil {
		log.Fatal("Error when reading config file")
	}

	fmt.Printf("CFG: %+v\n", cfg)
}
