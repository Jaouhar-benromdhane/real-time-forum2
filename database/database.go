package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDb() {
	var err error
	DB, err = sql.Open("sqlite3", "./db/forum.db")
	if err != nil {
		log.Fatal(err)
	}

	CreateUserTable := `
	CREATE TABLE IF NOT EXISTS users (
		id TEXT PRIMARY KEY,
		nickname TEXT UNIQUE NOT NULL,
		age INTEGER NOT NULL,
		gender TEXT NOT NULL,
		first_name TEXT NOT NULL,
		last_name TEXT NOT NULL,
		email TEXT UNIQUE NOT NULL,
		password BLOB NOT NULL
	);`
	CreatePostTable := `
	CREATE TABLE IF NOT EXISTS posts (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		content TEXT NOT NULL,
		category TEXT NOT NULL,
		user_id TEXT NOT NULL,
		created_at DATE DEFAULT (datetime('now', 'localtime')),
		FOREIGN KEY(user_id) REFERENCES users(id),
		FOREIGN KEY(category) REFERENCES categories(name)
	);`

	CreateSessionTable := `
	CREATE TABLE IF NOT EXISTS sessions (
		id TEXT PRIMARY KEY,
		user_id TEXT NOT NULL UNIQUE,
		expiration DATE NOT NULL,
		FOREIGN KEY(user_id) REFERENCES users(id)
	);`

	CreateCategoryTable := `
	CREATE TABLE IF NOT EXISTS categories (
		name TEXT PRIMARY KEY,
		description TEXT NOT NULL
	);`

	CreateCommentTable := `
	CREATE TABLE IF NOT EXISTS comments (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		content TEXT NOT NULL,
		post_id INTEGER NOT NULL,
		user_id TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY(post_id) REFERENCES posts(id),
		FOREIGN KEY(user_id) REFERENCES users(id)
	);`

	CreatMessageTable := `
	CREATE TABLE IF NOT EXISTS messages (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		sender_id TEXT NOT NULL,
		receiver_id TEXT NOT NULL,
		content TEXT NOT NULL,
		created_at DATE DEFAULT (datetime('now', 'localtime'))
);`

	_, err = DB.Exec(CreateUserTable)
	if err != nil {
		log.Fatal(err)
	}
	_, err = DB.Exec(CreatePostTable)
	if err != nil {
		log.Fatal(err)
	}
	_, err = DB.Exec(CreateSessionTable)
	if err != nil {
		log.Fatal(err)
	}
	_, err = DB.Exec(CreateCategoryTable)
	if err != nil {
		log.Fatal(err)
	}
	_, err = DB.Exec(CreateCommentTable)
	if err != nil {
		log.Fatal(err)
	}
	_, err = DB.Exec(CreatMessageTable)
	if err != nil {
		log.Fatal(err)
	}

}
