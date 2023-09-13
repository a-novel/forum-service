# Forum service

Manage forum posts.

## Prerequisites

- Download [Go](https://go.dev/doc/install)
- Install [Mockery](https://vektra.github.io/mockery/latest/installation/)
- Clone [go-framework](https://github.com/a-novel/go-framework)
    - From the framework, run `docker compose up -d`

## Installation

Create a env file.

```bash
touch .envrc
```
```bash
printf 'export POSTGRES_URL="postgres://forum@localhost:5432/agora_forum?sslmode=disable"
export POSTGRES_URL_TEST="postgres://test@localhost:5432/agora_forum_test?sslmode=disable"
' > .envrc
```
```bash
direnv allow .
```

Set the database up.
```bash
make db-setup
```

## Commands

### Run the API

```bash
make run
```
```bash
curl http://localhost:2041/ping
# Or curl http://localhost:2041/healthcheck
```

### Run the internal API

```bash
make run-internal
```
```bash
curl http://localhost:20041/ping
# Or curl http://localhost:20041/healthcheck
```

### Run tests

```bash
make test
```

### Update mocks

```bash
mockery
```

### Open a postgres console

```bash
make db
# Or make db-test
```
