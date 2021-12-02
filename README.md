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

ORM logging is disabled by default, but if you want to enable it, define an environment variable.

```
ENABLE_ORM_LOGGING= ./scripts/attach-graphql scripts/test
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
./scripts/attach-graphql scripts/gqlgen
```

### generarte graphql client code

```
./scripts/attach-application-ui npm run codegen
```

### lint frontend code

```
./scripts/attach-application-ui yarn lint
```

### test graphql resolver

```
docker-compose up graphql
# go to http://localhost:8080/
```

## todo

- resolverの複雑度の計算
- resolverのregresion test、golden testとかいうやつをググる

## テスト方針 - 草稿

```
バックエンド：

resolver毎のリグレッションテストは必ず書く
下位のモジュールのテストは重要なところは必ずかく
モジュールと下位モジュールの依存性はコンストラクタ経由で明示しなければならない
     テストコード上は下位モジュールの振る舞いは上位モジュールでテストする必要はない
カバレッジを見ながら抜けている振る舞いがないか検証していく
カバレッジは基本70以上とする

フロントエンド：

typed-vuexでテストできるかわからないので一旦保留
```

