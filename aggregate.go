package main

import (
	"context"
	"fmt"
)

func scrapeFeed(s *state) error {
	feed, err := s.Db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	_, err = s.Db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		return err
	}

	rss, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return err
	}

	for index := range rss.Channel.Item {
		item := rss.Channel.Item[index]
		fmt.Printf("%s:\n %s\n", item.Title, item.Description)
	}

	return nil
}
