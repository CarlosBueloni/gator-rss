package main

import (
	"github.com/carlosbueloni/gator-rss/internal/config"
	"log"
	"os"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	programState := &state{
		cfg: &cfg,
	}
	cmds := commands{
		Commands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	newCommand := command{
		Name: cmdName,
		Args: cmdArgs,
	}
	err = cmds.run(programState, newCommand)
	if err != nil {
		log.Fatal(err)
	}

}
