package main

import (
	"database/sql"
	"fmt"
	"gator/m/v2/internal/config"
	"gator/m/v2/internal/database"
	"log"
	"os"
	"syscall"

	_ "github.com/lib/pq"
)

type State struct {
	cfg *config.Config
	db  *database.Queries
}

func main() {

	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DbUrl)
	if err != nil {
		log.Fatalf("error opening connection to DB: %v", err)
	}

	defer db.Close()
	dbQueries := database.New(db)

	programState := &State{
		cfg: &cfg,
		db:  dbQueries,
	}

	cmds := Commands{
		registeredCommands: make(map[string]func(s *State, command Command) error),
	}

	cmds.Register("login", loginHandler)
	cmds.Register("register", registerHandler)
	cmds.Register("reset", ResetHandler)
	cmds.Register("users", UsersHandler)
	cmds.Register("agg", AggregatorHandler)
	cmds.Register("addfeed", AddFeedHandler)
	cmds.Register("feeds", FeedsHandler)

	if len(os.Args) < 2 {
		fmt.Print("Not enough arguments provided")
		syscall.Exit(1)
		return
	}
	commandName := os.Args[1]

	err = cmds.Run(programState, Command{
		Name: commandName,
		Args: os.Args[1:],
	})

	if err != nil {
		log.Fatal(err)
		syscall.Exit(1)
	}
}
