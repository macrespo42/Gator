package main

import (
	"context"
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

	return nil
}
