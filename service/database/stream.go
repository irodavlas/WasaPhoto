package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) GetFullStream(userId string) ([]Photo, error) {

	rows, err := db.c.Query(`
	SELECT image, photoId, userId
	FROM Photos 
	JOIN (
    SELECT followingUserId 
    FROM Followers 
    WHERE userId=?
	) AS a ON Photos.userId = a.followingUserId
	WHERE NOT EXISTS (
    SELECT * 
    FROM Bans 
    WHERE userId = a.followingUserId 
    AND bannedUserId = ?
	)
	ORDER BY date DESC
	`, userId, userId)

	if err != nil {
		return nil, err
	}
	var image []byte
	var photoId string
	var author string
	stream := []Photo{}
	for rows.Next() {
		err = rows.Scan(&image, &photoId, &author)
		if err != nil {
			return nil, err
		}
		comments, err := db.RetrieveComments(photoId)
		if err != nil {
			return nil, err
		}
		likes, err := db.RetrieveLikes(photoId)
		if err != nil {
			return nil, err
		}
		username, err := db.RetrieveUsername(author)
		if err != nil {
			return nil, err
		}
		likeId, err := db.retrieveLikeId(photoId, userId)
		if errors.Is(err, sql.ErrNoRows) {
			likeId = ""
		}

		newPhoto := Photo{image, photoId, username, comments, likes, likeId}
		stream = append(stream, newPhoto)
	}

	if rows.Err() != nil {
		return stream, rows.Err()
	}
	return stream, err
}
