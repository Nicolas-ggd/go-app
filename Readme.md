### Build swagger

#### Re-build swagger after something change
``` 
  swag init --parseDependency  --parseInternal -g main.go  
```

### Used migrations from:
```
  https://github.com/golang-migrate
```

### Create migration
```
  migrate create -seq -ext .sql -dir internal/migrations create_example_table
```
 - When using Migrate CLI we need to pass to database URL. Let's export it to a variable for convenience: `export POSTGRESQL_URL='postgres://username:password@localhost:5432/database_name?sslmode=disable'`

### Run migration
```
  migrate create -ext sql -dir internal/db/migrations -seq create_example_table
```