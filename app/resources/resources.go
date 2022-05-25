package resources

import (
	"database/sql"
	"fmt"
	"github.com/badfan/inno-taxi-user-service/app/models/sqlc"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/pressly/goose/v3"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type IResource interface {
	CreateUser(user sqlc.User) (int, error)
	GetUserIDByPhone(phone string) (int, error)
	GetUserByPhoneAndPassword(phone, password string) (sqlc.User, error)
	GetUserRatingByID(id int) (float32, error)
}

type Resource struct {
	Db     *sql.DB
	logger *zap.SugaredLogger
}

func NewResource(logger *zap.SugaredLogger) (*Resource, error) {
	viper.AutomaticEnv()

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		viper.Get("DBHOST"), viper.Get("DBPORT"), viper.Get("DBUSER"), viper.Get("DBPASSWORD"),
		viper.Get("DBNAME"), viper.Get("SSLMODE"))

	db, err := goose.OpenDBWithDriver("pgx", connStr)
	if err != nil {
		return nil, err
	}

	logger.Info("Migration start")

	err = goose.Up(db, "./migrations/")
	if err != nil {
		return nil, err
	}

	logger.Info("Migration ended")

	return &Resource{Db: db, logger: logger}, nil
}
