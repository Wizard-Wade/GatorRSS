package CommandHelpers

import (
	CommandTypes "GatorRss/Commands/Types"
	"GatorRss/internal/database"
	"context"
)

func DefaultLogin(handler func(s *CommandTypes.State, cmd CommandTypes.Command, user database.User) error) func(*CommandTypes.State, CommandTypes.Command) error {

	adaptedFunc := func(s *CommandTypes.State, cmd CommandTypes.Command) error {
		uid, err := s.Db.GetUser(context.Background(), s.Config.Current_user_name)
		if err != nil {
			return err
		}
		return handler(s, cmd, uid)
	}

	return adaptedFunc
}
