package main

import (
	"log"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		log.Fatal("Please provide a migration direction: up or down")
	}

	
}
