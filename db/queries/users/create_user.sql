-- name: CreateUser :one
INSERT INTO "public"."users" (
  "name", "email", "password_hash", "promptpay_id", "promptpay_qr"
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;
