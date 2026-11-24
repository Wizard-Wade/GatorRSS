package Commands

import (
	CommandTypes "GatorRss/Commands/Types"
	"context"
	"fmt"
	"html"
)

const (
	feedUrl = "https://www.wagslane.dev/index.xml"
)

func aggregateHandler(s *CommandTypes.State, cmd CommandTypes.Command) error {
	if len(cmd.Arguments) > 0 {
		return fmt.Errorf("Invalid argument provided. Aggregate accepts 0 arguments")
	}

	feed, err := fetchFeed(context.Background(), feedUrl)
	if err != nil {
		return fmt.Errorf("Could not fetch feed: %w", err)
	}

	title := html.UnescapeString(feed.Channel.Title)
	fmt.Println("", title)

	description := html.UnescapeString(feed.Channel.Description)
	fmt.Println("", description)
	for _, item := range feed.Channel.Item {
		title = html.UnescapeString(item.Title)
		fmt.Println("", title)

		description = html.UnescapeString(item.Description)
		fmt.Println("", description)
	}
	return nil
}
