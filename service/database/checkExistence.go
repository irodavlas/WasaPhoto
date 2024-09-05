package database

func (db *appdbimpl) CheckUserExistence(token string) error {
	var placeholder string
	err := db.c.QueryRow("SELECT userId FROM Users WHERE userId=?", token).Scan(&placeholder)
	return err
}

func (db *appdbimpl) CheckPhotoExistence(photoId string) error {
	var placeholder string
	err := db.c.QueryRow("SELECT userId FROM Photos WHERE photoId=?", photoId).Scan(&placeholder)
	return err
}

func (db *appdbimpl) CheckLikeExistence(likeId string) error {
	var placeholder string
	err := db.c.QueryRow("SELECT likeId FROM Likes WHERE likeId=?", likeId).Scan(&placeholder)
	return err
}

func (db *appdbimpl) CheckLikeExistenceFromPhoto(photoId string, userId string) error {
	var placeholder string
	err := db.c.QueryRow("SELECT likeId FROM Likes WHERE photoId=? AND userId=?", photoId, userId).Scan(&placeholder)
	return err
}

func (db *appdbimpl) CheckCommentExistence(commentId string) error {
	var placeholder string
	err := db.c.QueryRow("SELECT commentId FROM Comments WHERE commentId=?", commentId).Scan(&placeholder)
	return err
}
