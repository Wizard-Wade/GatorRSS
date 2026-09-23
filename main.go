package main

import (
	"GatorRss/Commands"
	"GatorRss/internal/config"
	"log"
	"os"
)

func main() {
	c, err := config.Read()
	if err != nil {
		log.Fatalln("", err)
	}

	main_state := &Commands.State{
		Config: &c,
	}

	commandList := Commands.Commands{
		Handler_Map: make(map[string]func(*Commands.State, Commands.Command) error),
	}

	commandList.Register("Login", Commands.HandlerLogin)

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
