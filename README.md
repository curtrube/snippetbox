# snippetbox

## Allow connections from your Docker network
If you want to allow the web user to connect from any host, you can create an account with a wildcard host:

```sql
CREATE USER 'web'@'%' IDENTIFIED BY 'pass';

GRANT SELECT, INSERT, UPDATE, DELETE
ON snippetbox.*
TO 'web'@'%';
```

## DB Migrations

Database migrations are managed through [migrate](https://github.com/golang-migrate/migrate).


create a new set of migration files

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
- [ ] export OTEL (traces, logs, etc)
- [ ] DB migrations
- [ ] Handle graceful server shutdown (SIGINT/SIGTERM)
- [ ] Migrate to PostgreSQL
