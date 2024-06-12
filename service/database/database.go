/*
Package database is the middleware between the app database and the code. All data (de)serialization (save/load) from a
persistent database are handled here. Database specific logic should never escape this package.

To use this package you need to apply migrations to the database if needed/wanted, connect to it (using the database
data source name from config), and then initialize an instance of AppDatabase from the DB connection.

For example, this code adds a parameter in `webapi` executable for the database data source name (add it to the
main.WebAPIConfiguration structure):

	DB struct {
		Filename string `conf:""`
	}

This is an example on how to migrate the DB and connect to it:

	// Start Database
	logger.Println("initializing database support")
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		logger.WithError(err).Error("error opening SQLite DB")
		return fmt.Errorf("opening SQLite: %w", err)
	}
	defer func() {
		logger.Debug("database stopping")
		_ = db.Close()
	}()

Then you can initialize the AppDatabase and pass it to the api package.
*/
package database

import (
	"database/sql"
	"errors"
	"myproject/service/types"
)

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	GetName() (string, error)
	GetUsers() ([]string, error)
	GetUser(id string, username string) (types.User, error)
	InsertUser(user types.User) error
	InsertFollower(userId string, followerId string) error
	RemoveFollow(followerId string, userId string) error
	InsertPost(post types.Post) error
	RemovePost(ownerId string, postId string) error
	BanUser(bannedUserId string, userId string) error
	InsertLike(likeid string, postId string, user types.User) error
	RemoveLike(postId string, likerId string) error
	InsertComment(commentId string, postId string, viewerId string, content string) error
	RemoveComment(postId string, ownerId string, commentId string) error
	UpdateUsername(id string, newUsername string) error
	RevokeBan(profile types.User, viewerId string) error
	GetFeed(user types.User) (*types.Feed, error)
	GetProfile(targetUser types.User, authUserId string) (*types.User, error)
	Ping() error
}

type appdbimpl struct {
	c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}

	// Users table
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS users (
        id TEXT PRIMARY KEY,
    	username TEXT NOT NULL
    );`)
	if err != nil {
		return nil, err
	}
	// follower table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS followers (
        user_id TEXT,
		follower_id TEXT,
		PRIMARY KEY (user_id, follower_id),
		FOREIGN KEY (user_id) REFERENCES users(id),
		FOREIGN KEY (follower_id) REFERENCES users(id)
    );`)
	if err != nil {
		return nil, err
	}
	// posts table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS posts (
        post_id TEXT primary key,
		owner_id TEXT,
		photo TEXT,
		created_at TIMESTAMP, owner_username TEXT,
		FOREIGN KEY (owner_id) REFERENCES users(id)
    );`)
	if err != nil {
		return nil, err
	}
	// comments table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS comments (
        comment_id TEXT PRIMARY KEY,
		user_id TEXT,
		post_id TEXT,
		comment_content TEXT NOT NULL,
		FOREIGN KEY (post_id) REFERENCES posts(post_id) ON DELETE CASCADE,
		Foreign key (user_id) REFERENCES users(id) ON DELETE CASCADE
    );`)
	if err != nil {
		return nil, err
	}

	// bans table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS bans (
        banned_id TEXT,
		user_id TEXT,
		PRIMARY KEY (banned_id, user_id),
		FOREIGN KEY (banned_id) REFERENCES users(id),
		FOREIGN KEY (user_id) REFERENCES users(id)
    );`)
	if err != nil {
		return nil, err
	}
	// likes table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS likes (
        like_id TEXT PRIMARY KEY,
		user_id TEXT,
		post_id TEXT,
		FOREIGN KEY (user_id) REFERENCES users(id)  ON DELETE CASCADE,
		FOREIGN KEY (post_id) REFERENCES posts(post_id) ON DELETE CASCADE
    );`)
	if err != nil {
		return nil, err
	}
	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
