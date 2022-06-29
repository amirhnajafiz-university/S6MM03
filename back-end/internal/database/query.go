package database

var (
	CreateTableQuery = `
		CREATE TABLE movies (
        	uid INTEGER PRIMARY KEY AUTOINCREMENT,
        	title VARCHAR(64) NULL,
			director VARCHAR(64) NULL,
         	score INTEGER,
        	description VARCHAR(1024) NULL,
			type VARCHAR(8) NULL,
			link VARCHAR(1024) NULL
    	);`
	InsertMovieQuery = `
		INSERT INTO 
			movies(title, director, score, description, type, link)
			values(?, ?, ?, ?, ?, ?);
		`
)
