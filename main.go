package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/macrespo42/Gator/internal/config"
	"github.com/macrespo42/Gator/internal/database"
)

func registerCommands(cmds commands) error {
	err := cmds.register("login", handlerLogin)
	if err != nil {
		return err
	}

	err = cmds.register("register", handlerRegister)
	if err != nil {
		return err
	}

	err = cmds.register("reset", handlerReset)
	if err != nil {
		return err
	}

	err = cmds.register("users", handlerUsers)
	if err != nil {
		return err
	}

	err = cmds.register("agg", handlerAgg)
	if err != nil {
		return err
	}

	err = cmds.register("feeds", handlerFeeds)
	if err != nil {
		return err
	}

	err = cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	if err != nil {
		return err
	}

	err = cmds.register("follow", middlewareLoggedIn(handlerFollow))
	if err != nil {
		return err
	}
	err = cmds.register("following", middlewareLoggedIn(handlerFollowing))
	if err != nil {
		return err
	}
	err = cmds.register("unfollow", middlewareLoggedIn(handleUnfollow))
	if err != nil {
		return err
	}
	err = cmds.register("browse", middlewareLoggedIn(handlerBrowse))
	if err != nil {
		return err
	}
	return nil
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal("Error when reading config file")
	}

	s := state{
		Cfg: &cfg,
	}

	db, err := sql.Open("postgres", s.Cfg.Db_URL)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	dbQueries := database.New(db)
	s.Db = dbQueries

	cmds := commands{
		Names: make(map[string]func(*state, command) error),
	}

	err = registerCommands(cmds)
	if err != nil {
		fmt.Printf("error when registering command")
		os.Exit(1)
	}
	if len(os.Args) < 2 {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	cmd := command{
		Name:      os.Args[1],
		Arguments: os.Args[2:],
	}

	err = cmds.run(&s, cmd)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
}
