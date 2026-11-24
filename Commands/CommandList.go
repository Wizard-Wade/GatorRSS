package Commands

import (
	CommandHelpers "GatorRss/Commands/Helpers"
	CommandTypes "GatorRss/Commands/Types"
	"context"
	"fmt"
	"strings"
)

func resetHandler(s *CommandTypes.State, cmd CommandTypes.Command) error {
	if len(cmd.Arguments) > 0 {
		return fmt.Errorf("Invalid argument provided. Reset accepts 0 arguments")
	}

	resetErr := s.Db.ClearUsers(context.Background())
	if resetErr != nil {
		return resetErr
	}

	fmt.Println("User Table Reset")
	return nil
}

type Commands struct {
	Handler_Map map[string]func(*CommandTypes.State, CommandTypes.Command) error
}

func (c *Commands) Run(s *CommandTypes.State, cmd CommandTypes.Command) error {
	function, found := c.Handler_Map[strings.ToLower(cmd.Name)]
	if !found {
		return fmt.Errorf("function %s does not exist", cmd.Name)
	}
	return function(s, cmd)
}

func (c *Commands) register(name string, function func(s *CommandTypes.State, cmd CommandTypes.Command) error) error {
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

	err = c.register("Reset", resetHandler)
	if err != nil {
		return err
	}

	err = c.register("Users", usersHandler)
	if err != nil {
		return err
	}

	err = c.register("agg", aggregateHandler)
	if err != nil {
		return err
	}

	err = c.register("addfeed", CommandHelpers.DefaultLogin(addFeedHandler))
	if err != nil {
		return err
	}

	err = c.register("feeds", listFeedsHandler)
	if err != nil {
		return err
	}

	err = c.register("follow", CommandHelpers.DefaultLogin(followHandler))
	if err != nil {
		return err
	}

	err = c.register("following", CommandHelpers.DefaultLogin(followingHandler))
	if err != nil {
		return err
	}

	return nil
}
