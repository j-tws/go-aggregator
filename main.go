package main

import (
	"log"
	"os"

	"github.com/j-tws/go-aggregator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main(){
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	programState := state{cfg: &cfg}

	commands := commands{
		list: make(map[string]func(*state, cmd) error),
	}

	commands.register("login", HandlerLogin)

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