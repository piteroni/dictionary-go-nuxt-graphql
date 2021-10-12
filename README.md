# dictionary go nuxt graphql

## Requirement

- Docker Desktop

## project setup

```
git clone https://github.com/piteroni/dictionary-go-nuxt-graphql.git
cd dictionary-go-nuxt-graphql/app
docker-compose build
./scripts/construct-node-modules
./scripts/attach-api go run cmd/migrate/main.go
```

## runs application

```sh
basename $(pwd) # => dictionary-go-nuxt-graphql
./app/scripts/up # go to http://localhost:3000/
```
