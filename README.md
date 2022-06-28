# User-service

***
The service provides functionality to work with users.

## External requirements

***

    Go 1.18
    Docker
    Docker-compose
    Sqlc
    Goose
    Viper
    Gin
    Zap
    GRPC
    Swagger

## Configuration

The service could be configured by providing environment variables.

| Name        | Meaning                                   | Example   |
|-------------|-------------------------------------------|-----------|
| DBHOST      | Database connection host                  | 127.0.0.1 |
| DBPORT      | Database connection port                  | 5432      |
| DBUSER      | Database connection user                  | postgres  |
| DBNAME      | Database connection name                  | postgres  |
| DBPASSWORD  | Database connection password              | password  |
| SSLMODE     | Database connection sslmode               | disable   |
| APIPORT     | API port                                  | 8000      |
| RPCPORT     | RPC port                                  | 5050      |
| TOKENTTL    | Authorization token lifetime (in minutes) | 2         |
