package Commands

import (
	CommandTypes "GatorRss/Commands/Types"
	"GatorRss/internal/database"
	"context"
	"database/sql"
	"fmt"
	"html"
	"time"
)

const (
	feedUrl = "https://www.wagslane.dev/index.xml"
)

func aggregateHandler(s *CommandTypes.State, cmd CommandTypes.Command) error {
	if len(cmd.Arguments) != 1 {
		return fmt.Errorf("Invalid argument provided. Aggregate expects update frequency argument")
	}

	duration, err := time.ParseDuration(cmd.Arguments[0] + "s")
	if err != nil {
		return err
	}

	fmt.Printf("Collecting Feeds Every %dm%ds", int(duration.Minutes()), int(duration.Seconds())%60)

	scrapeFeeds(s)
	ticker := time.NewTicker(duration)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *CommandTypes.State) error {
	trackedFeeds, err := s.Db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	for _, trackedFeed := range trackedFeeds {

		updateTime := time.Now()
		updateParams := database.UpdateFeedParams{
			ID:          trackedFeed.ID,
			LastFetchAt: sql.NullTime{Time: updateTime},
			UpdatedAt:   updateTime,
		}

		err = s.Db.UpdateFeed(context.Background(), updateParams)
		if err != nil {
			return err
		}

		feed, err := fetchFeed(context.Background(), trackedFeed.Url)
		if err != nil {
			return fmt.Errorf("Could not fetch feed: %w", err)
		}

		title := html.UnescapeString(feed.Channel.Title)
		fmt.Println("", title)

		// description := html.UnescapeString(feed.Channel.Description)
		// fmt.Println("", description)
		for _, item := range feed.Channel.Item {
			title = html.UnescapeString(item.Title)
			fmt.Println("", title)

			// description = html.UnescapeString(item.Description)
			// fmt.Println("", description)
		}
	}
	return nil
}
