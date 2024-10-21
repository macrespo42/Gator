package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/macrespo42/Gator/internal/database"
)

func handleUnfollow(s *state, cmd command) error {
	if len(cmd.Arguments) < 1 {
		return fmt.Errorf("please provide the feed to unfollow as argument")
	}

	unfollowArg := database.DeleteFeedFollowParams{
		Name: s.Cfg.CurrentUserName,
		Url:  cmd.Arguments[0],
	}

	_, err := s.Db.DeleteFeedFollow(context.Background(), unfollowArg)
	if err != nil {
		return err
	}

	return nil
}

func handlerFollowing(s *state, _ command) error {
	followedFeeds, err := s.Db.GetFeedFollowForUser(context.Background(), s.Cfg.CurrentUserName)
	if err != nil {
		return err
	}

	for index := range followedFeeds {
		fmt.Println(followedFeeds[index].FeedName)
	}

	return nil
}

func handlerFollow(s *state, cmd command) error {
	if len(cmd.Arguments) < 1 {
		return fmt.Errorf("please provide an url as argument")
	}

	feed, err := s.Db.GetFeedByUrl(context.Background(), cmd.Arguments[0])
	if err != nil {
		return err
	}

	user, err := s.Db.GetUser(context.Background(), s.Cfg.CurrentUserName)
	if err != nil {
		return err
	}

	feedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	feedFollow, err := s.Db.CreateFeedFollow(context.Background(), feedFollowParams)
	if err != nil {
		return err
	}

	fmt.Printf("%s now following %s ", feedFollow.UserName, feedFollow.FeedName)
	return nil
}

func handlerFeeds(s *state, _ command) error {
	feeds, err := s.Db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	for index := range feeds {
		feed := feeds[index]
		user, err := s.Db.GetUserNameById(context.Background(), feed.UserID)
		if err != nil {
			return err
		}
		fmt.Printf("%s: %s created by %s\n", feed.Name, feed.Url, user)
	}
	return nil
}

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Arguments) < 2 {
		return fmt.Errorf("please provie the name and url of the feed as arguments")
	}

	user, err := s.Db.GetUser(context.Background(), s.Cfg.CurrentUserName)
	if err != nil {
		return err
	}

	feedParams := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Arguments[0],
		Url:       cmd.Arguments[1],
		UserID:    user.ID,
	}

	feed, err := s.Db.CreateFeed(context.Background(), feedParams)
	if err != nil {
		return err
	}

	feedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	_, err = s.Db.CreateFeedFollow(context.Background(), feedFollowParams)
	if err != nil {
		return err
	}

	fmt.Printf("%v\n", feed)
	return nil
}

func handlerUsers(s *state, _ command) error {
	users, err := s.Db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	for index := range users {
		if s.Cfg.CurrentUserName == users[index].Name {
			fmt.Printf("* %s (current)\n", users[index].Name)
		} else {
			fmt.Printf("* %s\n", users[index].Name)
		}
	}
	return nil
}

func handlerReset(s *state, cmd command) error {
	err := s.Db.DeleteAllUsers(context.Background())
	if err != nil {
		return err
	}

	fmt.Printf("Command %s executed with success.\n", cmd.Name)
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Arguments) == 0 {
		return fmt.Errorf("the register handler expects a single argument the username")
	}

	userParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Arguments[0],
	}

	usr, err := s.Db.CreateUser(context.Background(), userParams)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	s.Cfg.SetUser(usr.Name)
	fmt.Printf("%v\n", usr)

	return nil
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Arguments) == 0 {
		return fmt.Errorf("the login handler expects a single argument the username")
	}

	_, err := s.Db.GetUser(context.Background(), cmd.Arguments[0])
	if err != nil {
		return fmt.Errorf("can't login the user does not exist")
	}

	s.Cfg.SetUser(cmd.Arguments[0])

	fmt.Println("the user has been set")
	return nil
}

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Arguments) < 1 {
		return fmt.Errorf("please provide a time ticker as argument")
	}

	time_between_reqs, err := time.ParseDuration(cmd.Arguments[0])
	if err != nil {
		return err
	}

	fmt.Printf("Collection feeds every %s\n", time_between_reqs.String())
	ticker := time.NewTicker(time_between_reqs)
	for ; ; <-ticker.C {
		err := scrapeFeed(s)
		if err != nil {
			return err
		}
	}
}

func handlerBrowse(s *state, cmd command) error {
	limit := 2

	if len(cmd.Arguments) == 1 {
		i, err := strconv.Atoi(cmd.Arguments[0])
		if err != nil {
			return err
		}
		limit = i
	}

	user, err := s.Db.GetUser(context.Background(), s.Cfg.CurrentUserName)
	if err != nil {
		return err
	}

	params := database.GetPostForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	}

	posts, err := s.Db.GetPostForUser(context.Background(), params)
	if err != nil {
		return err
	}
	for _, post := range posts {
		fmt.Printf("--- %s ---\n", post.Title)
		fmt.Printf("    %v\n", post.Description)
		fmt.Printf("Link: %s\n", post.Url)
		fmt.Println("=====================================")
	}
	return nil
}
