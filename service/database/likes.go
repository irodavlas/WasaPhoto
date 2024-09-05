package database

func (db *appdbimpl) DeleteLike(likeId string, userId string) error {
	_, err := db.c.Exec(`DELETE FROM Likes WHERE likeId=? AND userId=?`, likeId, userId)
	return err
}

func (db *appdbimpl) PutLike(likeId string, userId string, photoId string) error {
	_, err := db.c.Exec(`INSERT INTO Likes (likeId, userId, photoId) VALUES (?, ?, ?)`, likeId, userId, photoId)
	return err
}

func (db *appdbimpl) DeleteLikeFromPhoto(photoId string) error {
	_, err := db.c.Exec(`DELETE FROM Likes WHERE photoId=?`, photoId)
	return err
}

func (db *appdbimpl) CheckLikeOwnership(likeId string, userId string) error {
	var placeholder string
	err := db.c.QueryRow(`SELECT likeId FROM Likes WHERE likeId=? AND userId=?`, likeId, userId).Scan(&placeholder)
	return err
}

func (db *appdbimpl) RetrieveLikes(photoId string) (int, error) {
	var likeCount int
	err := db.c.QueryRow(`SELECT COUNT(likeId) AS likeCount FROM Likes WHERE photoId=?;`, photoId).Scan(&likeCount)
	return likeCount, err
}

func (db *appdbimpl) retrieveLikeId(photoId string, userId string) (string, error) {
	var placeholder string
	err := db.c.QueryRow(`SELECT likeId FROM Likes WHERE userId=? AND photoId=?`, userId, photoId).Scan(&placeholder)
	return placeholder, err
}
