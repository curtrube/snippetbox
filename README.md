# snippetbox

MYSQL_HOST=mysql.internal
MYSQL_PORT=3306
MYSQL_DATABASE=users_db
MYSQL_USERNAME=users_service
MYSQL_PASSWORD=${MYSQL_PASSWORD}


## DB Migrations

Database migrations are managed through [migrate](https://github.com/golang-migrate/migrate).

To create a new set of migration files

```
./migrate create -ext sql -dir migrations -seq create_snippet_table
```

execute migrations

```
./migrate -source file://db/migrations -database 'mysql://root:SuperSecret!@tcp(localhost:3306)/snippetbox' up
```


## Enhancements

- [ ] Read ENV variables (db connString, port, etc)
- [ ] CI/CD
    - [ ] lint - golangci-lint
    - [ ] test - gotest
    - [ ] build - build the image
    - [ ] release - https://github.com/cycjimmy/semantic-release-action
- [ ] export OTEL (traces, logs, etc)
- [X] DB migrations
- [ ] Handle graceful server shutdown (SIGINT/SIGTERM)
- [ ] Migrate to PostgreSQL
