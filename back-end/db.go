package main

import (
	"log"

	"github.com/amirhnajafiz/hollyworld/back-end/internal/database"
)

func main() {
	// creating a database connection
	d, err := database.NewConnection()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Created")

	// making the migration
	err = database.Migrate(d)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Migrated")

	// adding initialized data
	err = database.Seed(d)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Seeded")
}
