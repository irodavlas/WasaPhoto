package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) GetProfile(userId string, toVisitUserId string) ([]Photo, error) {
	rows, err := db.c.Query("SELECT image, photoId FROM Photos WHERE userId=? AND NOT EXISTS (SELECT * FROM Bans WHERE userId=? AND bannedUserId=?)", toVisitUserId, toVisitUserId, userId)
	if err != nil {
		return nil, err
	}

	var image []byte
	var photoId string
	images := []Photo{}
	for rows.Next() {
		err = rows.Scan(&image, &photoId)
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
		username, err := db.RetrieveUsername(toVisitUserId)
		if err != nil {
			return nil, err
		}
		likeId, err := db.retrieveLikeId(photoId, userId)
		if errors.Is(err, sql.ErrNoRows) {
			likeId = ""
		}

		photo := Photo{image, photoId, username, comments, likes, likeId}
		images = append(images, photo)
	}

	if rows.Err() != nil {
		return images, rows.Err()
	}
	return images, err
}
