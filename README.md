# snippetbox

## Migrations

create a new set of migration files

```
./migrate create -ext sql -dir migrations -seq session
```

execute migrations

```
./migrate -source file://migrations -database 'mysql://root:SuperSecret!@tcp(localhost:3306)/snippetbox' up 2
```

