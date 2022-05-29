package resources

import (
	"context"

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
		r.logger.Errorf("error occured while creating user: %s", err.Error())
		return 0, err
	}

	return int(res.ID), err
}

func (r *Resource) GetUserIDByPhone(ctx context.Context, phone string) (int, error) {
	queries := sqlc.New(r.Db)

	res, err := queries.GetUserIDByPhone(ctx, phone)
	if err != nil {
		r.logger.Infof("error occured while getting user's id by phone: %s", err.Error())
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
		r.logger.Errorf("error occured while getting user by phone and password: %s", err.Error())
		return nil, err
	}

	res := sqlcUserConvert(&user)

	return res, err
}

func (r *Resource) GetUserRatingByID(ctx context.Context, id int) (float32, error) {
	queries := sqlc.New(r.Db)

	res, err := queries.GetUserRatingByID(ctx, int32(id))
	if err != nil {
		r.logger.Errorf("error occured while getting user's rating by id: %s", err.Error())
		return 0, err
	}

	return res, err
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
