package Commands

import (
	"GatorRss/internal/config"
	"fmt"
	"strings"
)

type State struct {
	Config *config.Config
}

type Command struct {
	Name      string
	Arguments []string
}

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Arguments) == 0 {
		return fmt.Errorf("No arguments included. Expected login name.")
	}

	err := s.Config.SetUser(cmd.Arguments[0])
	if err != nil {
		return err
	}

	fmt.Println("User set: ", cmd.Arguments[0])
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

func (c *Commands) Register(name string, function func(s *State, cmd Command) error) error {
	c.Handler_Map[strings.ToLower(name)] = function
	return nil
}
