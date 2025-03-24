package main

import (
	"fmt"
	"log"

	"github.com/j-tws/go-aggregator/internal/config"
)

func main(){
	c, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	fmt.Printf("Read config: %+v\n", c)

	err = c.SetUser("dovahkiin")

	c, err = config.Read()

	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config again: %+v\n", c)
}