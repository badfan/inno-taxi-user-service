CREATE TABLE IF NOT EXISTS users
(
    id SERIAL PRIMARY KEY,
    user_uuid uuid NOT NULL DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    phone_number VARCHAR(25) UNIQUE NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    user_rating REAL NOT NULL DEFAULT 0.0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- name: CreateUser :one
INSERT INTO users (name, phone_number, email, password)
VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetUsers :many
SELECT * FROM users;

-- name: GetUserByPhoneAndPassword :one
SELECT * FROM users
WHERE phone_number = $1 AND password = $2;

-- name: GetUserIDByPhone :one
SELECT id FROM users
WHERE phone_number = $1;

-- name: GetUserRatingByID :one
SELECT user_rating FROM users
WHERE id = $1;

-- name: UpdateUser :one
UPDATE users
SET name = $1, phone_number = $2, email = $3, password= $4
WHERE id=$5 RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
