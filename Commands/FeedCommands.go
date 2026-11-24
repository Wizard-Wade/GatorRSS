package Commands

import (
	CommandTypes "GatorRss/Commands/Types"
	"GatorRss/internal/database"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func addFeedHandler(s *CommandTypes.State, cmd CommandTypes.Command, user database.User) error {
	if len(cmd.Arguments) != 2 {
		return fmt.Errorf("Invalid argument provided. Add Feed accepts 2 arguments NAME & URL")
	}

	params := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Arguments[0],
		Url:       cmd.Arguments[1],
		UserID:    user.ID,
	}

	feed, err := s.Db.CreateFeed(context.Background(), params)
	if err != nil {
		return err
	}

	_, err = addFollow(s, feed.Url, user)
	if err != nil {
		return err
	}

	fmt.Println("added the url")
	fmt.Println(feed)

	return nil
}

func listFeedsHandler(s *CommandTypes.State, cmd CommandTypes.Command) error {
	if len(cmd.Arguments) != 0 {
		return fmt.Errorf("Invalid argument provided. List Feeds accepts NO arguments.")
	}

	feedRows, err := s.Db.GetFeeds(context.Background())

	if err != nil {
		return err
	}

	for _, row := range feedRows {
		userName, err := s.Db.GetUserByID(context.Background(), row.UserID)
		if err != nil {
			return err
		}

		fmt.Println("", row.Name, row.Url, userName.Name)
	}

	return nil
}
