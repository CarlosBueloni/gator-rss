package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/carlosbueloni/gator-rss/internal/config"
	"github.com/carlosbueloni/gator-rss/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	dbQueries := database.New(db)

	programState := &state{
		db:  dbQueries,
		cfg: &cfg,
	}
	cmds := commands{
		Commands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)

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
