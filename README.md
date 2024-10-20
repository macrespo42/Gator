# Gator 🐊 

gator ? My favorite blog aggregator

## Requirements

In order to make gator work you need:

- postgresql
- go

## Installation

To install on an unix system just type
```bash
go install github.com/macrespo42/Gator@latest
```

## Configuration

You'll need to have a .gatorconfig.json in your $HOME directory location  like
```bash
.gatorconfig.json
```

Example of a valid config file:
```json
{
    "db_url":"postgres://macrespo:@localhost:5432/gator?sslmode=disable", // mandatory
    "current_user_name":"kahya" // optionnal, be set with the cli
}
```

## Commands

You can run a command with 
```bash
Gator command_name arg1 arg2...
```

### There is a list of all availables commands:

- ``register $name_of_user`` | register a new user
- ``login $name_of_user`` | log a registered user
- ``reset`` | ⚠️  DEV command, reset database
- ``users`` | Get a list of registered users
- ``agg $10s`` | Scrape feeds from rss flux at the interval given as argument
- ``feeds`` | Get a list of added feeds
- ``addfeed $title $url`` | add a feed to the database
- ``follow $feed_title`` | follow feed by title
- ``unfollow $feed_title`` | unfollow feed by title
- ``following`` | list of your following feeds
- ``browse $limit=2`` | browse feeds from your following list
