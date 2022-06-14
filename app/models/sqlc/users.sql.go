// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: users.sql

package sqlc

import (
	"context"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (name, phone_number, email, password)
VALUES ($1, $2, $3, $4) RETURNING id, user_uuid, name, phone_number, email, password, user_rating, created_at, updated_at
`

type CreateUserParams struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Name,
		arg.PhoneNumber,
		arg.Email,
		arg.Password,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.UserUuid,
		&i.Name,
		&i.PhoneNumber,
		&i.Email,
		&i.Password,
		&i.UserRating,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getUserByPhoneAndPassword = `-- name: GetUserByPhoneAndPassword :one
SELECT id, user_uuid, name, phone_number, email, password, user_rating, created_at, updated_at FROM users
WHERE phone_number = $1 AND password = $2
`

type GetUserByPhoneAndPasswordParams struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

func (q *Queries) GetUserByPhoneAndPassword(ctx context.Context, arg GetUserByPhoneAndPasswordParams) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByPhoneAndPassword, arg.PhoneNumber, arg.Password)
	var i User
	err := row.Scan(
		&i.ID,
		&i.UserUuid,
		&i.Name,
		&i.PhoneNumber,
		&i.Email,
		&i.Password,
		&i.UserRating,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserIDByPhone = `-- name: GetUserIDByPhone :one
SELECT id FROM users
WHERE phone_number = $1
`

func (q *Queries) GetUserIDByPhone(ctx context.Context, phoneNumber string) (int32, error) {
	row := q.db.QueryRowContext(ctx, getUserIDByPhone, phoneNumber)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const getUserRatingByID = `-- name: GetUserRatingByID :one
SELECT user_rating FROM users
WHERE id = $1
`

func (q *Queries) GetUserRatingByID(ctx context.Context, id int32) (float32, error) {
	row := q.db.QueryRowContext(ctx, getUserRatingByID, id)
	var user_rating float32
	err := row.Scan(&user_rating)
	return user_rating, err
}

const getUserUUIDAndRatingByID = `-- name: GetUserUUIDAndRatingByID :one
SELECT user_uuid, user_rating FROM users
WHERE id = $1
`

type GetUserUUIDAndRatingByIDRow struct {
	UserUuid   uuid.UUID `json:"user_uuid"`
	UserRating float32   `json:"user_rating"`
}

func (q *Queries) GetUserUUIDAndRatingByID(ctx context.Context, id int32) (GetUserUUIDAndRatingByIDRow, error) {
	row := q.db.QueryRowContext(ctx, getUserUUIDAndRatingByID, id)
	var i GetUserUUIDAndRatingByIDRow
	err := row.Scan(&i.UserUuid, &i.UserRating)
	return i, err
}

const getUserUUIDByID = `-- name: GetUserUUIDByID :one
SELECT user_uuid FROM users
WHERE id = $1
`

func (q *Queries) GetUserUUIDByID(ctx context.Context, id int32) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, getUserUUIDByID, id)
	var user_uuid uuid.UUID
	err := row.Scan(&user_uuid)
	return user_uuid, err
}

const getUsers = `-- name: GetUsers :many
SELECT id, user_uuid, name, phone_number, email, password, user_rating, created_at, updated_at FROM users
`

func (q *Queries) GetUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, getUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.UserUuid,
			&i.Name,
			&i.PhoneNumber,
			&i.Email,
			&i.Password,
			&i.UserRating,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET name = $1, phone_number = $2, email = $3, password= $4
WHERE id=$5 RETURNING id, user_uuid, name, phone_number, email, password, user_rating, created_at, updated_at
`

type UpdateUserParams struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	ID          int32  `json:"id"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.Name,
		arg.PhoneNumber,
		arg.Email,
		arg.Password,
		arg.ID,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.UserUuid,
		&i.Name,
		&i.PhoneNumber,
		&i.Email,
		&i.Password,
		&i.UserRating,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
