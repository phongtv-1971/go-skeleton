# go-skeleton
Skeleton for api app

## Config
```shell
cp config.yaml.example config.yaml
docker-compose up -d
```

## View available make command
```shell
make help
```

### Running test
```shell
make setup-test
make test
```

### Start server
Available mode is development, test, or production. default mode is development.
Server running on 0.0.0.0:8080
```shell
make create-database APP_ENV=mode
make migrate-up APP_ENV=mode
make server APP_ENV=mode
```
Ex: `make server APP_ENV=production`

### Generate migrate file
```shell
make generate-migration name=file_name
```
Ex: `make generate-migration name=create_table_users`

### Generate sqlc code
```shell
make sqlc
```

### Build a binary file and execute
```shell
make build
APP_ENV=production build/go-skeleton
```

### Tool
- sqlc v1.8.0
- go-migrate v4.14.1
- swagger source html, css, js
- mockgen v1.6.0
- golangci-lint v1.36.0

### Available Endpoint example
```shell
GET    /api_doc
GET    /health_check
POST   /api/v1/users
GET    /api/v1/users/:id
GET    /api/v1/users
```
