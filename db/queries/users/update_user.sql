-- name: UpdateUser :one
UPDATE "public"."users"
  set "name" = $2,
  "email" = $3,
  "password_hash" = $4,
  "promptpay_id" = $5,
  "promptpay_qr" = $6
WHERE "user_id" = $1
RETURNING *;
