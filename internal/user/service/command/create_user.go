package command

import (
	"context"

	"to-work-api/internal/sqlc"
	"to-work-api/internal/user/domain"

	"github.com/jackc/pgx/v5/pgtype"
)

type CreateUserParams struct {
	Name         string
	Email        string
	PasswordHash string
	PromptpayID  string
	PromptpayQr  string
}

type CreateUserHandler struct {
	queries *sqlc.Queries
}

func NewCreateUserHandler(queries *sqlc.Queries) *CreateUserHandler {
	return &CreateUserHandler{
		queries: queries,
	}
}

func (h *CreateUserHandler) Handle(ctx context.Context, params CreateUserParams) (domain.User, error) {
	arg := sqlc.CreateUserParams{
		Name:         params.Name,
		Email:        params.Email,
		PasswordHash: params.PasswordHash,
		PromptpayID:  pgtype.Text{String: params.PromptpayID, Valid: params.PromptpayID != ""},
		PromptpayQr:  pgtype.Text{String: params.PromptpayQr, Valid: params.PromptpayQr != ""},
	}

	user, err := h.queries.CreateUser(ctx, arg)
	if err != nil {
		return domain.User{}, err
	}

	return domain.FromSqlc(user), nil
}
