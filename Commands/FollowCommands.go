package Commands

import (
	CommandTypes "GatorRss/Commands/Types"
	"GatorRss/internal/database"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func followHandler(s *CommandTypes.State, cmd CommandTypes.Command, user database.User) error {
	if len(cmd.Arguments) != 1 {
		return fmt.Errorf("Invalid argument provided. List Feeds accepts one URL argument.")
	}

	followRow, err := addFollow(s, cmd.Arguments[0], user)
	if err != nil {
		return err
	}

	fmt.Printf("%s followed %s\n", followRow.UserName, followRow.FeedName)

	return nil
}

func addFollow(s *CommandTypes.State, url string, user database.User) (*database.CreateFeedFollowRow, error) {
	feed, err := s.Db.GetFeedByURL(context.Background(), url)

	if err != nil {
		return nil, err
	}

	feedfollow := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	followRow, err := s.Db.CreateFeedFollow(context.Background(), feedfollow)
	if err != nil {
		return nil, err
	}
	return &followRow, nil
}

func followingHandler(s *CommandTypes.State, cmd CommandTypes.Command, user database.User) error {
	if len(cmd.Arguments) != 0 {
		return fmt.Errorf("Invalid argument provided. List Feeds accepts one NO argument.")
	}

	feeds, err := s.Db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	for i, feed := range feeds {
		fmt.Printf("\t%d: %s\n", i, feed.FeedName)
	}

	return nil
}
