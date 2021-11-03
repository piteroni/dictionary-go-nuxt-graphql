FROM golang:1.16-alpine AS base

RUN apk upgrade --update && \
    apk --no-cache add git gcc libc-dev

FROM base AS builder

WORKDIR /build

COPY . .

RUN go mod download && go build -o main .

FROM alpine:3.14.2 AS production

ENV TZ=Asia/Tokyo

WORKDIR /app

COPY --from=builder /build/main ./

ENTRYPOINT [ "/app/main" ]

FROM base AS dev

ENV TZ=Asia/Tokyo

ENV PS1='\h /\W\ # '

RUN cd /tmp && \
    go get -u github.com/cosmtrek/air && \
    go get -u github.com/rakyll/gotest

WORKDIR /app

COPY go.mod go.sum ./

RUN \
    # for testing
    go get -u gorm.io/driver/sqlite && \
    # for test coverage
    go get -u gopkg.in/yaml.v2@v2.3.0 && \
    go mod download

CMD ["bash"]
