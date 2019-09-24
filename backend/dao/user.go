package dao

import "time"

// User 用户表
type User struct {
	Uid       int64 `grom:"primary_key"`
	Openid    string
	Nickname  string
	Gender    int
	Avatar    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
