package query

import (
	"context"

	"to-work-api/internal/sqlc"
	"to-work-api/internal/user/domain"
)

type GetUserParams struct {
	UserID int64
}

type GetUserHandler struct {
	qeries *sqlc.Queries
}

func NewGetUserHandler(queries *sqlc.Queries) *GetUserHandler {
	return &GetUserHandler{
		qeries: queries,
	}
}

func (s *GetUserHandler) Handle(params GetUserParams) (domain.User, error) {
	user, err := s.qeries.GetUser(context.Background(), params.UserID)
	if err != nil {
		return domain.User{}, err
	}
	return domain.FromSqlc(user), nil
}
