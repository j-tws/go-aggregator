package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/j-tws/go-aggregator/internal/config"
	"github.com/j-tws/go-aggregator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main(){
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DbUrl)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	defer db.Close()

	dbQueries := database.New(db)

	programState := state{
		cfg: &cfg,
		db: dbQueries,
	}

	//some commentblabla
	commands := commands{
		list: make(map[string]func(*state, cmd) error),
	}

	// random comment
	commands.register("login", HandlerLogin)
	commands.register("register", HandlerRegister)
	commands.register("reset", HandlerReset)
	commands.register("users", HandlerUsers)
	commands.register("agg", HandlerAgg)
	commands.register("addfeed", middlewareLoggedIn(HandlerAddFeed))
	commands.register("feeds", HandlerFeeds)
	commands.register("follow", middlewareLoggedIn(HandlerFollow))
	commands.register("following", middlewareLoggedIn(HandlerFollowing))
	commands.register("unfollow", middlewareLoggedIn(HandlerUnfollow))
	commands.register("browse", middlewareLoggedIn(HandlerBrowse))

	programArgs := os.Args

	if len(programArgs) < 2 {
		log.Fatal("Error: Command name is required")
	}

	commandName := programArgs[1]
	commandArgs := programArgs[2:]

	cmdErr := commands.run(&programState, cmd{name: commandName, args: commandArgs}) 
	if cmdErr != nil {
		log.Fatal(cmdErr)
	}
}