# Database Migration using Go Migrate 

## Installation
### Windows
Using Scoop
```bash
$ scoop install migrate
```

### Linux
```bash
$ curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey| apt-key add -
$ echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/migrate.list
$ apt-get update
$ apt-get install -y migrate
```

### Mac
```bash
$ brew install golang-migrate
```

## init create init database file
```bash
$ migrate create -ext sql -dir database/migration/ -seq init_db
```

## Create another 

```bash 
$ migrate create -ext sql -dir database/migration/ -seq data_seeder
```

## Run Migration Up

```bash
migrate -path database/migration/ -database "postgresql://username:secretkey@localhost:5432/database_name?sslmode=disable" -verbose up
```

## Rollback Migration

```bash
$ migrate -path database/migration/ -database "postgresql://username:secretkey@localhost:5432/database_name?sslmode=disable" -verbose down
```