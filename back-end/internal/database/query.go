package database

var (
	CreateTableQuery = `
		CREATE TABLE movies (
        	uid INTEGER PRIMARY KEY AUTOINCREMENT,
        	title VARCHAR(64) NULL,
			director VARCHAR(64) NULL,
         	score INTEGER 0,
        	description VARCHAR(1024) NULL,
			poster VARCHAR(1024) NULL,
			link VARCHAR(1024) NULL
    	);`
	InsertMovieQuery = `
		INSERT INTO 
			movies(title, director, score, description, poster, link)
			values(?, ?, ?, ?, ?, ?);
		`
)
