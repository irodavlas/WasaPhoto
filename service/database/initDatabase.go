package database

import (
	"database/sql"
)

func initDatabase(db *sql.DB) error {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS Users(
		userId VARCHAR(16) NOT NULL PRIMARY KEY,
		username VARCHAR(16) NOT NULL,
		UNIQUE(userId)
		UNIQUE(username)
	);
	`)

	if err != nil {
		return err
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS Bans(
		userId VARCHAR(16) NOT NULL,
		bannedUserId VARCHAR(16) NOT NULL
	);
	`)

	if err != nil {
		return err
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS Followers(
		userId VARCHAR(16) NOT NULL,
		followingUserId VARCHAR(16) NOT NULL
	);
	`)

	if err != nil {
		return err
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS Photos(
		photoId VARCHAR(16) NOT NULL PRIMARY KEY,
		userId VARCHAR(16) NOT NULL,
		image TEXT,
		date TEXT NOT NULL,
		UNIQUE(photoId)
	);
	`)

	if err != nil {
		return err
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS Likes(
		likeId VARCHAR(16) NOT NULL PRIMARY KEY,
		userId VARCHAR(16) NOT NULL,
		photoId VARCHAR(16) NOT NULL,
		UNIQUE(likeId)
	);
	`)

	if err != nil {
		return err
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS Comments(
		commentId VARCHAR(16) NOT NULL PRIMARY KEY,
		userId VARCHAR(16) NOT NULL,
		photoId VARCHAR(16) NOT NULL,
		content TEXT NOT NULL,
		UNIQUE(commentId)
	);
	`)

	if err != nil {
		return err
	}
	return nil
}
