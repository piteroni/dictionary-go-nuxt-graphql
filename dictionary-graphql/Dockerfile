# base
FROM golang:1.16-alpine AS base

RUN apk upgrade --update && \
    apk --no-cache add git gcc libc-dev

# development
FROM base AS dev

ENV TZ=Asia/Tokyo \
    PS1='\h /\W\ # '

RUN cd /tmp && \
    go get -u github.com/cosmtrek/air && \
    go get -u github.com/rakyll/gotest

WORKDIR /app

COPY go.mod go.sum ./

RUN \
    # for test coverage
    go get -u gopkg.in/yaml.v2@v2.4.0 && \
    go mod download

CMD ["bash"]

# production
FROM base AS builder

WORKDIR /build

COPY . .

RUN go mod download && \
    GOOS=linux go build -ldflags="-s -w" -trimpath -o main .

FROM alpine:3.14.2 AS production

ENV TZ=Asia/Tokyo

WORKDIR /app

COPY --from=builder /build/main ./

CMD ["/app/main"]
