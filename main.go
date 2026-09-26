package main

import (
	"GatorRss/Commands"
	"GatorRss/internal/config"
	"GatorRss/internal/database"
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	c, err := config.Read()
	if err != nil {
		log.Fatalln("", err)
	}

	db, err := sql.Open("postgres", c.Url)
	if err != nil {
		log.Fatalf("error connecting to db: %v", err)
	}
	defer db.Close()
	dbQueries := database.New(db)

	main_state := &Commands.State{
		Db:     dbQueries,
		Config: &c,
	}

	commandList := Commands.Commands{
		Handler_Map: make(map[string]func(*Commands.State, Commands.Command) error),
	}

	err = commandList.RegisterCommands()
	if err != nil {
		log.Fatalf("Failed to register commands: %v", err)
	}

	args := os.Args
	if len(args) < 2 {
		log.Fatalln("Invalid argument provided. Expected at least one argument beyond application name")
	}

	cmd := Commands.Command{
		Name:      args[1],
		Arguments: args[2:],
	}
	err = commandList.Run(main_state, cmd)
	if err != nil {
		log.Fatal(err)
	}
}
