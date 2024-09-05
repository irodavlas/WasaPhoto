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
	"fmt"
)

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	SetName(userId string, username string) error
	CheckUserExistence(token string) error
	CheckPhotoExistence(photoId string) error
	RetrieveId(username string) (string, error)
	RegisterUser(username string, userId string) error
	Ping() error
	CheckFollow(userId string, followingUserId string) error
	FollowUser(userId string, followingUserId string) error
	DeleteFollow(userId string, toUnfollowUserId string) error
	CheckBan(toCheckUserId string, userId string) error
	BanUser(userId string, toBanUserId string) error
	DeleteBan(userId string, toUnbanUserId string) error
	PostPhoto(photoId string, userId string, image []byte, date string) error
	DeleteLike(likeId string, userId string) error
	PutLike(likeId string, userId string, photoId string) error
	CheckLikeExistence(likeId string) error
	DeleteComment(commentId string, userId string) error
	DeleteCommentFromPhoto(photoId string) error
	PutComment(commentId string, userId string, photoId string, content string) error
	DeleteLikeFromPhoto(photoId string) error
	DeletePhoto(photoId string, userId string) error
	CheckLikeExistenceFromPhoto(photoId string, userId string) error
	GetProfile(userId string, toVisitUserId string) ([]Photo, error)
	CheckCommentExistence(commentId string) error
	CheckLikeOwnership(likeId string, userId string) error
	CheckPhotoOwnership(photoId string, userId string) error
	CheckCommentOwnership(commentId string, userId string) error
	GetFullStream(userId string) ([]Photo, error)
	RetrieveComments(photoId string) ([]Comment, error)
	RetrieveLikes(photoId string) (int, error)
	RetrieveUsername(userId string) (string, error)
	retrieveLikeId(photoId string, userId string) (string, error)
	RetrieveFollowers(userId string) (int, error)
	RetrieveFollowing(userId string) (int, error)
	// RetrieveImage(photoId string) (string, error)
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

	// Check if table exists. If not, the database is empty, and we need to create the structure
	err := initDatabase(db)

	if err != nil {
		return nil, fmt.Errorf("error creating tables: %w", err)
	}

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
