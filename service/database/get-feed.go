package database

import (
	"myproject/service/types"
)

// GetName is an example that shows you how to query data
func (db *appdbimpl) GetFeed(user types.User) (*types.Feed, error) {

	var feed types.Feed

	followingList, err := FetchUsers(db, user.Id)
	//error is above on the following list
	if err != nil {
		return nil, err
	}

	for _, profile := range followingList {

		var _user types.User
		_user.Post, err = GetPosts(db, profile, user.Id)

		if err != nil {
			return nil, err
		}

		feed.Feed = append(feed.Feed, _user.Post...)

	}

	return &feed, nil
}
func FetchUsers(db *appdbimpl, profileId string) ([]types.User, error) {
	var users []types.User
	//from your account get YOUR following and grab (id, username) with a query
	//then as for the grab profile take all the Posts from the list obtained above
	//order the list using the time.time functions
	//ez
	query := `
	SELECT users.id, users.username
	FROM followers
	JOIN users ON followers.user_id  = users.id
	WHERE followers.follower_id  = ?;`

	rows, err := db.c.Query(query, profileId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var userID string
		var username string
		if err := rows.Scan(&userID, &username); err != nil {
			return nil, err
		}
		users = append(users, *types.NewUser(userID, username))
	}
	return users, nil
}
