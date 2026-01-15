package query

import (
	"context"

	"to-work-api/internal/sqlc"
	"to-work-api/internal/user/domain"
)

// create function GetUsers
type GetUsersParams struct {
	Offset int64
	Limit  int64
}

type GetUsersHandler struct {
	qeries *sqlc.Queries
}

func NewGetUsersHandler(queries *sqlc.Queries) *GetUsersHandler {
	return &GetUsersHandler{
		qeries: queries,
	}
}

func (s *GetUsersHandler) Handle(params GetUsersParams) ([]domain.User, error) {
	usersSqlc, err := s.qeries.GetUsers(context.Background())
	if err != nil {
		return nil, err
	}

	var users []domain.User
	for _, user := range usersSqlc {
		users = append(users, domain.FromSqlc(user))
	}

	return users, nil
}
