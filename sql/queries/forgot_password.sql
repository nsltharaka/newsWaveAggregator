-- name: CreatePasswordResetCase :one
INSERT INTO
    forgot_password (case_number, code, user_id)
VALUES ($1, $2, $3)
RETURNING
    *;

-- name: GetCaseForUser :one
SELECT * FROM forgot_password WHERE user_id = $1;

-- name: GetCaseByNumber :one
SELECT * FROM forgot_password WHERE case_number = $1;

-- name: UpdateCode :exec
UPDATE forgot_password SET code = $1 WHERE case_number = $2;

-- name: DeleteCase :exec
DELETE FROM forgot_password WHERE case_number = $1;