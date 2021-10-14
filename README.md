# dictionary go nuxt graphql

## Requirement

- Docker Desktop

## project setup

```
git clone https://github.com/piteroni/dictionary-go-nuxt-graphql.git
cd dictionary-go-nuxt-graphql/app
docker-compose build
./scripts/construct-node-modules
./scripts/up-db -d && sleep 5
./scripts/attach-api go run cmd/migrate/main.go
```

## runs application

```sh
./app/scripts/up # go to http://localhost:3000/
```

## sub commands

### generarte graphql server

```
./app/scripts/attach-api scripts/gqlgen
```

### generarte graphql client

```
./app/scripts/attach-ui npm run codegen
```
