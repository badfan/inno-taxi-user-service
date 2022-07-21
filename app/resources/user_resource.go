package resources

import (
	"context"

	"github.com/google/uuid"

	"github.com/badfan/inno-taxi-user-service/app/models"
	"github.com/badfan/inno-taxi-user-service/app/models/sqlc"
)

func (r *Resource) CreateUser(ctx context.Context, user *models.User) (int, error) {
	queries := sqlc.New(r.Db)

	res, err := queries.CreateUser(ctx, sqlc.CreateUserParams{
		Name:        user.Name,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		Password:    user.Password,
	})
	if err != nil {
		return 0, err
	}

	return int(res.ID), err
}

func (r *Resource) GetUserIDByPhone(ctx context.Context, phone string) (int, error) {
	queries := sqlc.New(r.Db)

	res, err := queries.GetUserIDByPhone(ctx, phone)
	if err != nil {
		return 0, err
	}

	return int(res), err
}

func (r *Resource) GetUserByPhoneAndPassword(ctx context.Context, phone, password string) (*models.User, error) {
	queries := sqlc.New(r.Db)

	user, err := queries.GetUserByPhoneAndPassword(ctx, sqlc.GetUserByPhoneAndPasswordParams{
		PhoneNumber: phone,
		Password:    password,
	})
	if err != nil {
		return nil, err
	}

	res := sqlcUserConvert(&user)

	return res, err
}

func (r *Resource) GetUserRatingByID(ctx context.Context, id int) (float32, error) {
	queries := sqlc.New(r.Db)

	res, err := queries.GetUserRatingByID(ctx, int32(id))
	if err != nil {
		return 0, err
	}

	return res, err
}

func (r *Resource) GetUserUUIDByID(ctx context.Context, id int) (uuid.UUID, error) {
	queries := sqlc.New(r.Db)

	res, err := queries.GetUserUUIDByID(ctx, int32(id))
	if err != nil {
		return uuid.UUID{}, err
	}

	return res, nil
}

func (r *Resource) GetUserUUIDAndRatingByID(ctx context.Context, id int) (uuid.UUID, float32, error) {
	queries := sqlc.New(r.Db)

	res, err := queries.GetUserUUIDAndRatingByID(ctx, int32(id))
	if err != nil {
		return uuid.UUID{}, 0, err
	}

	return res.UserUuid, res.UserRating, nil
}

func sqlcUserConvert(source *sqlc.User) *models.User {
	res := &models.User{
		ID:          source.ID,
		UserUuid:    source.UserUuid,
		Name:        source.Name,
		PhoneNumber: source.PhoneNumber,
		Email:       source.Email,
		Password:    source.Password,
		UserRating:  source.UserRating,
		CreatedAt:   source.CreatedAt,
		UpdatedAt:   source.UpdatedAt,
	}

	return res
}
