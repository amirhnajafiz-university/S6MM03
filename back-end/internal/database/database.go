package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

const (
	DbFile = "./sqlite.db"
)

// NewConnection opens a connection to our sqlite file
func NewConnection() (*sql.DB, error) {
	return sql.Open("sqlite3", DbFile)
}

// Migrate our database file
func Migrate(conn *sql.DB) error {
	_, err := conn.Exec(CreateTableQuery)

	return err
}

// Seed data into our database
func Seed(conn *sql.DB) error {
	type Movie struct {
		Title       string
		Director    string
		Score       int
		Description string
		Type        string
		Link        string
	}

	s, err := conn.Prepare(InsertMovieQuery)
	if err != nil {
		return err
	}

	movies := []Movie{
		{
			Title:    "Doctor Strange in the Multiverse of Madness",
			Director: "Sam Raimi",
			Score:    7,
			Description: "Doctor Strange teams up with a " +
				"mysterious teenage girl from his dreams who " +
				"can travel across multiverses, to battle multiple threats, " +
				"including other-universe versions of himself, which threaten " +
				"to wipe out millions across the multiverse. " +
				"They seek help from Wanda the Scarlet Witch, Wong and others.",
			Type: "DASH1",
		},
		{
			Title:    "Obi-Wan Kenobi",
			Director: "Deborah Chow",
			Score:    8,
			Description: "Jedi Master Obi-Wan Kenobi " +
				"has to save young Leia after she is kidnapped, " +
				"all the while being pursued by Imperial Inquisitors " +
				"and his former Padawan, now known as Darth Vader.",
			Type: "DASH1",
		},
		{
			Title:    "Top Gun: Maverick",
			Director: "Bob Standford",
			Score:    9,
			Description: "After more than thirty years of service" +
				" as one of the Navy's top aviators, Pete" +
				" Mitchell is where he belongs, pushing" +
				" the envelope as a courageous test pilot" +
				" and dodging the advancement in rank that would ground",
			Type: "DASH2",
			Link: "https://amir-farshid.arvanvod.com/YnGp7wpPVQ/D6OvYMvJ7x/h_,144_200,240_400,360_800,480_1500,k.mp4.list/manifest.mpd",
		},
		{
			Title:    "Stranger Things",
			Director: "Brad Richards",
			Score:    9,
			Description: "When a young boy disappears, his mother, " +
				"a police chief and his friends must " +
				"confront terrifying supernatural forces in order to get him back.",
			Type: "DASH2",
			Link: "https://amir-farshid.arvanvod.com/YnGp7wpPVQ/La61Pn75Y9/h_,144_200,240_400,360_800,480_1483,k.mp4.list/manifest.mpd",
		},
		{
			Title:    "The Boys",
			Director: "Eric Kripke",
			Score:    8,
			Description: "A group of vigilantes set out to take" +
				" down corrupt superheroes who abuse their superpowers.",
			Type: "HLS",
			Link: "https://amir-farshid.arvanvod.com/YnGp7wpPVQ/vL3zBRzqmP/h_,144_200,240_400,k.mp4.list/master.m3u8",
		},
		{
			Title:    "The Black Phone",
			Director: "Scott Derrickson",
			Score:    7,
			Description: "After being abducted by a child killer and locked in a soundproof basement, " +
				"a 13-year-old boy starts receiving calls on " +
				"a disconnected phone from the killer's previous victims.",
			Type: "HLS",
			Link: "https://amir-farshid.arvanvod.com/YnGp7wpPVQ/KBbv76jDAG/h_,144_200,240_400,360_800,480_1500,k.mp4.list/master.m3u8",
		},
	}

	for _, movie := range movies {
		_, err = s.Exec(
			movie.Title,
			movie.Director,
			movie.Score,
			movie.Description,
			movie.Type,
			movie.Link,
		)
		if err != nil {
			return err
		}
	}

	return nil
}
