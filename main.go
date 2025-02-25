package main

import (
	"fmt"
	"log"

	"github.com/carlosbueloni/gator-rss/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config file: %v", err)
	}
	fmt.Printf("Read config: %v\n", cfg)

	err = cfg.SetUser("Steven Universe")
	if err != nil {
		log.Fatalf("error setting usename: %v", err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config file: %v", err)
	}
	fmt.Printf("Read config again: %v\n", cfg)

}
