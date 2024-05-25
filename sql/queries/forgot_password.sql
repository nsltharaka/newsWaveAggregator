-- name: CreatePasswordResetCase :one
INSERT INTO
    forgot_password (case_number, opened, user_id)
VALUES ($1, $2, $3) RETURNING *;

-- name: GetCaseForUser :one
SELECT * FROM forgot_password WHERE user_id = $1;