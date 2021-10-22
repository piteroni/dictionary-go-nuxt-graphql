# dictionary go nuxt graphql

## Requirement

- Docker Desktop

## project setup

```
git clone https://github.com/piteroni/dictionary-go-nuxt-graphql.git
cd dictionary-go-nuxt-graphql
docker-compose build
./scripts/construct-node-modules
docker-compose up -d db && sleep 5
./scripts/attach-api go run cmd/migrate/main.go
```

## runs application

```sh
docker-compose up # go to http://localhost:3000/
```

## sub commands

### execute test on graphql container

```
./app/scripts/attach-api scripts/gqlgen
```

### view coverage graphql container

```
./scripts/view-coverage-graphql-container
```

### connect to mysql cli

```
./scripts/connect-db
```

### generarte graphql server code

```
./app/scripts/attach-api scripts/gqlgen
```

### generarte graphql client code

```
./app/scripts/attach-ui npm run codegen
```

## やりのこしたこと

- GraphQLでエラー種別をスキーマで定義すること、[参考リンク](https://www.youtube.com/watch?v=RDNTP66oY2o)
