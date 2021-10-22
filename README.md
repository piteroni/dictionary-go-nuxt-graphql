# dictionary go nuxt graphql

## Requirement

- Docker Desktop

## project setup

```
git clone https://github.com/piteroni/dictionary-go-nuxt-graphql.git
cd dictionary-go-nuxt-graphql
docker-compose build
./scripts/construct-node-modules
docker-compose up -d db
./scripts/attach-graphql go run cmd/migrate/main.go
```

## runs application

```sh
docker-compose up
```

**port binding**

container|url
--|--
application-ui|https://localhost:3000
graphql|https://localhost:8080
mysql|https://localhost:3306

## sub commands

### execute test on graphql container

```
./scripts/attach-graphql scripts/test
```

### view coverage graphql container

```
./scripts/view-coverage-graphql-container
```

### connect to mysql cli

```
./scripts/connect-db
```

### refresh database records

```
./scripts/attach-graphql go run cmd/drop/main.go
./scripts/attach-graphql go run cmd/migrate/main.go
```

### generarte graphql server code

```
./scripts/attach-application-graphql scripts/gqlgen
```

### generarte graphql client code

```
./scripts/attach-application-ui npm run codegen
```

## todo

- GraphQLでエラー種別をスキーマで定義すること、[参考リンク](https://www.youtube.com/watch?v=RDNTP66oY2o)
