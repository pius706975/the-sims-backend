package main

import (
	"github.com/pius706975/the-sims-backend/cmd"
	"log"
	"os"
)

func main() {
	err := cmd.Run(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
}