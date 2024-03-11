# Install Go-App Application

This project provides a best practice way to build application using Go.

## Prerequisites

- Migration installed on your system. You can find installation instructions at https://github.com/golang-migrate.


## After installing the application:
**Note:** After installing the application, you need to download all necessary dependency which application requires:

- Run the following command: `go mod download or go mod tidy`

## Create migration

- After successfully installed migration, let's create migration using command: `migrate create -seq -ext .sql -dir internal/migrations create_example_table`
- If there were no errors, we should have two files available under internal/migrations folder:

  - 000001_create_example_table.down.sql
  - 000001_example_users_table.up.sql

## Run migration

- To run migration, let's use command: `migrate -database ${POSTGRESQL_URL} -path db/migrations up`

- When using Migrate CLI we need to pass to database URL. Let's export it to a variable for convenience: `export POSTGRESQL_URL='postgres://username:password@localhost:5432/database_name?sslmode=disable'`

- During migration running if error occurred which tells, `Dirty database version <version>. Fix and force version`. In that case, run command: `migrate -path migrations/ -database postgres://test:test@localhost/dummy?sslmode=disable force <version>`

## Drop migration

**Note:** Please note that before drop migration you need to export convenience variable.

- Drop all migration using command: `migrate -database ${POSTGRESQL_URL} -path db/migrations down`

## Environment Variables

The following environment variables are required to run the application:

| Variable         | Description                                      | Default Value |
|------------------|-------------------------------------------------|---------------|
| `DB_PORT`          | Port of the database server                       | (Set as an environment variable)          |
| `DB_HOST`          | Hostname or IP address of the database server   | (Set as an environment variable)    |
| `DB_USER`          | Username for the database user                   | (Set as an environment variable) |
| `DB_PASSWORD`     | Password for the database user                  | (Set as an environment variable) |
| `DB_NAME`          | Name of the database                             | (Set as an environment variable)  |

**Note:** Do not store the actual password in this file. Instead, set the password as an environment variable before running the container using methods specific to your operating system. For example, on Linux or macOS, you can use the `export` command:

## Updating Swagger Documentation

**Note:** If you're making changes to the code that affect the Swagger documentation, you can regenerate it using the following command.

- Run the following command: `swag init --parseDependency  --parseInternal -g main.go`
