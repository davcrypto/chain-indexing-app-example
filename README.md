# chain-indexing-app-example

Example implementation for [cryto.com indexing](https://github.com/crypto-com/chain-indexing)

## How to Run

### 1 Prerequisite

- Crypto.com Chain full node
- Postgres Database

### 2 Configuration file

A sample configuration is available under `config/config.sample.toml`.

Copy it, update configuration based on your setup and rename it as `config/config.toml`.

Note: Postgres database password is not available in `config.toml` nor command option. You must provide it as environment variable `DB_PASSWORD` on start.

#### Reminder On Connecting Mainnet

There is a rate limiter on our public nodes. If you hit the rate limit, you may want to run your own nodes.
### 3 Postgres Database

You can have your Postgres setup locally or remotely.

**REMINDER**: I would suggest using our `docker-compose` script to start the DB instance. If you install through `homebrew`, its default setting will need to be adjusted in order to match the indexing server's configuration. 

#### Run Postgres with Docker

**WARNING**: The docker files available under `docker/` is intended only for development and testing purpose only. Never use them in production setup.

For a local test run. A docker-compose file with Postgres database and PgAdmin console is available via running:

```bash
docker-compose --file docker/docker-compose.development.yml up -d
```

This will start the following docker instances on your local network when you use default credentials:
| Docker image | Port | Username | Password | Other Config | Mounted volume |
| --- | --- | --- | --- | --- | --- |
| Postgres | 5432 | postgres | postgres | Database Name = postgres; SSL = true | ./pgdata-dev |
| PgAdmin | 8080 | pgadmin@localhost | pgadmin | N/A | N/A |

### 4 Execute Database Migration

for DB_PASSWORD, never use common word such as "postgres" , "admin", "password", choose at least 16 characters including number, special character, capital letters even in testing environment.
if you don't use strong password, pgmigrate will stop further processing

#### Docker

```bash
docker run -it \
    --env DB_USERNAME=postgres \
    --env DB_PASSWORD=your_postgresql_password \
    --env DB_HOST=host.docker.internal \
    --env DB_PORT=5432 \
    --env DB_NAME=postgres \
    --env DB_SCHEMA=public \
    chain-indexing-app /app/migrate -- -verbose up
```

#### Manual Build

```bash
# In your first run, you need to install the dependency `migrate`
./pgmigrate.sh --install-dependency
# Then you should have `migrate` under your `$PATH`
which migrate

# Run the below command to start the migrate
./pgmigrate.sh -- -verbose up
```

### 5 Run the Service

#### Docker

```bash
docker run \
    -v `pwd`/config:/app/config --read-only \
    -p 28857:28857 \
    --env DB_PASSWORD=your_postgresql_password \
    chain-indexing-app /app/chain-indexing-app
```

#### Manual build

```bash
env DB_PASSWORD=your_postgresql_password ./chain-indexing-app
```
