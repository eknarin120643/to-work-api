-- name: DeleteUser :exec
DELETE FROM "public"."users"
WHERE "user_id" = $1;
