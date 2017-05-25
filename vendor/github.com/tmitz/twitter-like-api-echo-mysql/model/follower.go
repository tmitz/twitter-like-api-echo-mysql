package model

type Follower struct {
	FollowerID int `json:"follower_id" db:"follower_id"`
	FollowedID int `json:"followed_id" db:"followed_id"`
}
