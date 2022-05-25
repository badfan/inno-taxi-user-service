package resources

import (
	"context"

	"github.com/badfan/inno-taxi-user-service/app/models/sqlc"
)

func (r *Resource) CreateUser(user sqlc.User) (int, error) {
	ctx := context.Background()

	queries := sqlc.New(r.Db)

	res, err := queries.CreateUser(ctx, sqlc.CreateUserParams{
		Name:        user.Name,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		Password:    user.Password,
	})
	if err != nil {
		r.logger.Errorf("error occured while creating user: %s", err.Error())
	}

	return int(res.ID), err
}

func (r *Resource) GetUserIDByPhone(phone string) (int, error) {
	ctx := context.Background()

	queries := sqlc.New(r.Db)

	res, err := queries.GetUserIDByPhone(ctx, phone)
	if err != nil {
		r.logger.Infof("error occured while getting user's id by phone: %s", err.Error())
	}

	return int(res), err
}

func (r *Resource) GetUserByPhoneAndPassword(phone, password string) (sqlc.User, error) {
	ctx := context.Background()

	queries := sqlc.New(r.Db)

	res, err := queries.GetUserByPhoneAndPassword(ctx, sqlc.GetUserByPhoneAndPasswordParams{
		PhoneNumber: phone,
		Password:    password,
	})
	if err != nil {
		r.logger.Errorf("error occured while getting user by phone and password: %s", err.Error())
	}

	return res, err
}

func (r *Resource) GetUserRatingByID(id int) (float32, error) {
	ctx := context.Background()

	queries := sqlc.New(r.Db)

	res, err := queries.GetUserRatingByID(ctx, int32(id))
	if err != nil {
		r.logger.Errorf("error occured while getting user's rating by id: %s", err.Error())
	}

	return res, err
}
