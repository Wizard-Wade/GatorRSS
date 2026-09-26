package Commands

import (
	"GatorRss/internal/config"
	"GatorRss/internal/database"
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

type State struct {
	Db     *database.Queries
	Config *config.Config
}

type Command struct {
	Name      string
	Arguments []string
}

func handlerLogin(s *State, cmd Command) error {
	if len(cmd.Arguments) == 0 {
		return fmt.Errorf("No arguments included. Expected login name.")
	}

	_, usrerr := s.Db.GetUser(context.Background(), cmd.Arguments[0])
	if usrerr != nil {
		return fmt.Errorf("Login Name not found.")
	}

	err := s.Config.SetUser(cmd.Arguments[0])
	if err != nil {
		return err
	}

	fmt.Println("User set: ", cmd.Arguments[0])
	return nil
}

func registerHandler(s *State, cmd Command) error {
	if len(cmd.Arguments) == 0 {
		return fmt.Errorf("No arguments included. Expected login name.")
	}

	params := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Arguments[0],
	}

	_, writeError := s.Db.CreateUser(context.Background(), params)
	if writeError != nil {
		return writeError
	}

	handlerLogin(s, cmd)

	fmt.Printf("Created user %v. ID: %v, Created: %v, Updated: %v\n", params.Name, params.ID, params.CreatedAt, params.UpdatedAt)
	return nil
}

type Commands struct {
	Handler_Map map[string]func(*State, Command) error
}

func (c *Commands) Run(s *State, cmd Command) error {
	function, found := c.Handler_Map[strings.ToLower(cmd.Name)]
	if !found {
		return fmt.Errorf("function %s does not exist", cmd.Name)
	}
	return function(s, cmd)
}

func (c *Commands) register(name string, function func(s *State, cmd Command) error) error {
	c.Handler_Map[strings.ToLower(name)] = function
	return nil
}

func (c *Commands) RegisterCommands() error {
	err := c.register("Login", handlerLogin)
	if err != nil {
		return err
	}

	err = c.register("Register", registerHandler)
	if err != nil {
		return err
	}

	return nil
}
