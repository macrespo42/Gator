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

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("feeds", handlerFeeds)
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("follow", middlewareLoggedIn(handlerFollow))
	cmds.register("following", middlewareLoggedIn(handlerFollowing))
	cmds.register("unfollow", middlewareLoggedIn(handleUnfollow))

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
