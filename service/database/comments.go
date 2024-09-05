package database

func (db *appdbimpl) DeleteComment(commentId string, userId string) error {
	_, err := db.c.Exec(`DELETE FROM Comments WHERE commentId=? AND userId=?`, commentId, userId)
	return err
}

func (db *appdbimpl) PutComment(commentId string, userId string, photoId string, content string) error {
	_, err := db.c.Exec(`INSERT INTO Comments (commentId, userId, photoId, content) VALUES (?, ?, ?, ?)`, commentId, userId, photoId, content)
	return err
}

func (db *appdbimpl) DeleteCommentFromPhoto(photoId string) error {
	_, err := db.c.Exec(`DELETE FROM Comments WHERE photoId=?`, photoId)
	return err
}

func (db *appdbimpl) CheckCommentOwnership(commentId string, userId string) error {
	var placeholder string
	err := db.c.QueryRow(`SELECT commentId FROM Comments WHERE commentId=? AND userId=?`, commentId, userId).Scan(&placeholder)
	return err
}

func (db *appdbimpl) RetrieveComments(photoId string) ([]Comment, error) {
	rows, err := db.c.Query(`SELECT
    C.commentId,
    C.content,
    U.username
		FROM
			Comments AS C
		JOIN
			Users AS U ON C.userId = U.userId
		JOIN
			Photos AS P ON C.photoId = P.photoId
		WHERE
			C.photoId =?;`, photoId)

	if err != nil {
		return nil, err
	}

	var commentId string
	var content string
	var username string
	comments := []Comment{}
	for rows.Next() {
		err = rows.Scan(&commentId, &content, &username)
		if err != nil {
			return nil, err
		}
		comment := Comment{content, commentId, username}
		// fmt.Println(image)
		comments = append(comments, comment)
	}

	if rows.Err() != nil {
		return comments, rows.Err()
	}
	return comments, err
}
