package main

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/macrespo42/Gator/internal/database"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", "gator")

	c := http.Client{Timeout: time.Duration(3) * time.Second}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var rss RSSFeed

	err = xml.Unmarshal(body, &rss)
	if err != nil {
		return nil, err
	}

	rss.Channel.Title = html.EscapeString(rss.Channel.Title)
	rss.Channel.Description = html.EscapeString(rss.Channel.Description)

	for index := range rss.Channel.Item {
		rss.Channel.Item[index].Title = html.EscapeString(rss.Channel.Item[index].Title)
		rss.Channel.Item[index].Description = html.EscapeString(rss.Channel.Item[index].Description)
	}

	return &rss, nil
}

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

		layout := "Mon, 02 Jan 2006 15:04:05 -0700"
		value := item.PubDate

		publicationDate, err := time.Parse(layout, value)
		if err != nil {
			log.Printf("unexpected date format....")
			continue
		}

		createPostParams := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       item.Title,
			Url:         item.Link,
			Description: item.Description,
			PublishedAt: publicationDate,
			FeedID:      feed.ID,
		}
		post, err := s.Db.CreatePost(context.Background(), createPostParams)
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			}
			log.Println(err)
			continue
		}
		log.Printf("Post: %s, registered with success\n", post.Title)
	}
	log.Printf("Feed: %s, collected %v posts founds\n", feed.Name, len(rss.Channel.Item))
	return nil
}
