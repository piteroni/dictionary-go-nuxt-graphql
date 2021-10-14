FROM golang:1.16-alpine

RUN apk upgrade --update && \
    apk --no-cache add git gcc libc-dev

ENV TZ=Asia/Tokyo

ENV PS1='\h /\W\ # '

RUN cd /tmp && \
    go get -u github.com/cosmtrek/air

WORKDIR /app

COPY go.mod go.sum ./

RUN \
    # for testing
    go get -u gorm.io/driver/sqlite && \
    # for test coverage
    go get -u gopkg.in/yaml.v2@v2.3.0 && \
    go mod download

CMD ["air", "-c", ".air.toml"]
