# Forum service

Manage forum posts.

- [Installation](#installation)
- [Commands](#commands)

## Installation

Set the database up.

```bash
make db-setup
```

## Commands

Run the API:

```bash
make run
```

Run the internal API (used by google cloud internal services):

```bash
make run-internal
```

Run tests:

```bash
make test
```

Connect to the database:

```bash
make db
```

Connect to the test database:

```bash
make db-test
```
