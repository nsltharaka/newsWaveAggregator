-- name: CreateUser :one
insert into users(
        created_at,
        updated_at,
        username,
        email,
        password
    )
values ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetUserByApiKey :one
SELECT * FROM users WHERE api_key = $1;

-- name: GetUserByEmail :one
select *
from users
where email = $1;