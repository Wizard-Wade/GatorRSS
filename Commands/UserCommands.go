package Commands

import (
	CommandTypes "GatorRss/Commands/Types"
	"GatorRss/internal/database"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func usersHandler(s *CommandTypes.State, cmd CommandTypes.Command) error {
	if len(cmd.Arguments) > 0 {
		return fmt.Errorf("Invalid argument provided. Reset accepts 0 arguments")
	}

	users, err := s.Db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	for _, user := range users {
		fmt.Print("", user)
		if user == s.Config.Current_user_name {
			fmt.Print(" (current)")
		}
		fmt.Print("\n")
	}
	return nil
}

func registerHandler(s *CommandTypes.State, cmd CommandTypes.Command) error {
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

func handlerLogin(s *CommandTypes.State, cmd CommandTypes.Command) error {
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
