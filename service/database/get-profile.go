package database

import (
	"errors"
	"myproject/service/types"
)

// GetName is an example that shows you how to query data
func (db *appdbimpl) GetProfile(profile types.User, viewerId string) (*types.User, error) {
	var err error
	modality, err := cleanBanFeed(db, profile.Id, viewerId)
	if modality == "0" {
		return nil, errors.New("Attempt to visualize a profile either banned or banned by")
	}

	profile.Follower, err = getFollower(db, profile)
	if err != nil {
		return nil, err
	}
	profile.Following, err = getFollowing(db, profile)
	if err != nil {
		return nil, err
	}
	profile.Post, err = GetPosts(db, profile, viewerId)
	if err != nil {

		return nil, err
	}
	profile.Npost = len(profile.Post)
	//random return
	return &profile, nil
}
func getFollowing(db *appdbimpl, profile types.User) ([]types.Follower, error) {
	var following []types.Follower
	query := `
		SELECT
			users.id AS linked_id,
			users.username AS linked_username
		FROM
			followers
		JOIN
			users ON followers.user_id = users.id
		WHERE
			followers.follower_id = ?`
	rows, err := db.c.Query(query, profile.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user types.Follower
		err := rows.Scan(&user.Id, &user.Username)
		if err != nil {
			return nil, err
		}
		following = append(following, user)
	}
	return following, err
}
func getFollower(db *appdbimpl, profile types.User) ([]types.Follower, error) {
	var followers []types.Follower
	query := `
		SELECT
			users.id AS linked_id,
			users.username AS linked_username
		FROM
			followers
		JOIN
			users ON followers.follower_id = users.id
		WHERE
			followers.user_id = ?`

	// Execute the query
	rows, err := db.c.Query(query, profile.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var user types.Follower
		err := rows.Scan(&user.Id, &user.Username)
		if err != nil {
			return nil, err
		}
		followers = append(followers, user)
	}
	return followers, nil
}
func GetPosts(db *appdbimpl, profile types.User, viewerId string) ([]types.Post, error) {
	var posts []types.Post

	rows, err := db.c.Query("SELECT post_id, owner_username, photo, created_at FROM posts WHERE owner_id = ? ORDER BY created_at DESC", profile.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var post types.Post
		if err := rows.Scan(&post.PostId, &post.OwnerUsername, &post.Photo, &post.Time); err != nil {
			return nil, err
		}
		post.OwnerId = profile.Id
		//grab likes and commments from relative table for each post
		post.Likes, err = getLikes(db, post.PostId, viewerId)
		if err != nil {
			return nil, err
		}
		post.Comments, err = getComments(db, post.PostId, viewerId)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
func getComments(db *appdbimpl, postId string, viewerId string) ([]types.Comment, error) {
	var comments []types.Comment
	query := `
		SELECT users.username, comments.comment_content, comments.comment_id
		FROM users
		JOIN comments ON users.id = comments.user_id
		WHERE comments.post_id = ? 
		AND users.username NOT IN ( 
			SELECT users.username 
			FROM users 
			JOIN bans ON users.id = bans.banned_id 
			WHERE bans.user_id = ?
		)
		EXCEPT 
		SELECT users.username, NULL AS comment_content, NULL AS comment_id
		FROM users 
		JOIN bans ON users.id = bans.user_id 
		WHERE bans.banned_id = ?;`
	rows, err := db.c.Query(query, postId, viewerId, viewerId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var comment types.Comment
		if err := rows.Scan(&comment.Username, &comment.Message, &comment.CommentId); err != nil {
			return nil, err
		}
		comments = append(comments, comment)

	}
	return comments, nil
}
func getLikes(db *appdbimpl, postId string, viewerId string) ([]string, error) {
	var likes []string
	query := `
	SELECT users.username
	FROM users
	JOIN likes ON users.id = likes.user_id
	WHERE likes.post_id = ?
	and users.username not in ( 
	select users.username 
	from users 
	JOIN bans ON users.id = bans.banned_id 
	where bans.user_id = ?
	) 
	EXCEPT 
	select users.username 
	from users 
	JOIN bans ON users.id = bans.user_id 
	where bans.banned_id = ?
	`
	rows, err := db.c.Query(query, postId, viewerId, viewerId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var username string
		if err := rows.Scan(&username); err != nil {
			return nil, err
		}
		likes = append(likes, username)

	}

	return likes, nil
}

func cleanBanFeed(db *appdbimpl, profile string, viewerId string) (string, error) {
	query := `
		SELECT
			CASE
				WHEN EXISTS (SELECT 1 FROM bans WHERE banned_id = ? AND user_id = ?) THEN '0'
				WHEN EXISTS (SELECT 1 FROM bans WHERE banned_id = ? AND user_id = ?) THEN '0'
				ELSE ''
			END AS ban_status
	`

	var banStatus string
	err := db.c.QueryRow(query, viewerId, profile, profile, viewerId).Scan(&banStatus)
	if err != nil {
		return "", err
	}

	return banStatus, nil

}
