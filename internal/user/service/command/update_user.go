package command

import (
	"context"

	"to-work-api/internal/sqlc"
	"to-work-api/internal/user/domain"

	"github.com/jackc/pgx/v5/pgtype"
)

type UpdateUserParams struct {
	UserID       int64
	Name         string
	Email        string
	PasswordHash string
	PromptpayID  string
	PromptpayQr  string
}

type UpdateUserHandler struct {
	queries *sqlc.Queries
}

func NewUpdateUserHandler(queries *sqlc.Queries) *UpdateUserHandler {
	return &UpdateUserHandler{
		queries: queries,
	}
}

func (h *UpdateUserHandler) Handle(ctx context.Context, params UpdateUserParams) (domain.User, error) {
	arg := sqlc.UpdateUserParams{
		UserID:       params.UserID,
		Name:         params.Name,
		Email:        params.Email,
		PasswordHash: params.PasswordHash,
		PromptpayID:  pgtype.Text{String: params.PromptpayID, Valid: params.PromptpayID != ""},
		PromptpayQr:  pgtype.Text{String: params.PromptpayQr, Valid: params.PromptpayQr != ""},
	}

	user, err := h.queries.UpdateUser(ctx, arg)
	if err != nil {
		return domain.User{}, err
	}

	return domain.FromSqlc(user), nil
}
