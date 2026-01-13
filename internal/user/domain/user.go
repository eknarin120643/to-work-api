package domain

import "to-work-api/internal/sqlc"

type User struct {
	UserID int64
	Name   string
	Email  string
}

func FromSqlc(user sqlc.User) User {
	return User{
		UserID: user.UserID,
		Name:   user.Name,
		Email:  user.Email,
	}
}
