package app

import (
	"errors"

	"github.com/spf13/viper"
)

type APIConfig struct {
	APIPort  string
	TokenTTL string
}

type DBConfig struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
	SSLMode    string
}

func NewAPIConfig() (*APIConfig, error) {
	res := &APIConfig{}
	viper.AutomaticEnv()

	res.APIPort = viper.Get("APIPORT").(string)
	if res.APIPort == "" {
		return nil, errors.New("env var APIPORT is empty")
	}

	res.TokenTTL = viper.Get("TOKENTTL").(string)
	if res.TokenTTL == "" {
		return nil, errors.New("env var TOKENTTL is empty")
	}

	return res, nil
}

func NewDBConfig() (*DBConfig, error) {
	res := &DBConfig{}
	viper.AutomaticEnv()

	res.DBUser = viper.Get("DBUSER").(string)
	if res.DBUser == "" {
		return nil, errors.New("env var DBUSER is empty")
	}

	res.DBPassword = viper.Get("DBPASSWORD").(string)
	if res.DBPassword == "" {
		return nil, errors.New("env var DBPASSWORD is empty")
	}

	res.DBHost = viper.Get("DBHOST").(string)
	if res.DBHost == "" {
		return nil, errors.New("env var DBHOST is empty")
	}

	res.DBPort = viper.Get("DBPORT").(string)
	if res.DBPort == "" {
		return nil, errors.New("env var DBPORT is empty")
	}

	res.DBName = viper.Get("DBNAME").(string)
	if res.DBName == "" {
		return nil, errors.New("env var DBNAME is empty")
	}

	res.SSLMode = viper.Get("SSLMODE").(string)
	if res.SSLMode == "" {
		return nil, errors.New("env var SSLMODE is empty")
	}

	return res, nil
}
