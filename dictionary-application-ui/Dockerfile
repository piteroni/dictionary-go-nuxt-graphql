# base
FROM node:14-alpine3.12 AS base

RUN apk update && \
    apk --no-cache add tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
    apk del tzdata

ENV TZ=Asia/Tokyo

WORKDIR /app

ARG HTTP_ENDPOINT
ARG BROWSER_HTTP_ENDPOINT

ENV HTTP_ENDPOINT=${HTTP_ENDPOINT}
ENV BROWSER_HTTP_ENDPOINT=${BROWSER_HTTP_ENDPOINT}

# development
FROM base as dev

ENV PS1='\h /\W\ # '

CMD ["bash"]

# production
FROM base as builder

COPY . .

COPY package.json /app/package.json
COPY yarn.lock /app/yarn.lock

RUN yarn && yarn build

FROM base AS production

COPY /static /app/static

COPY --from=builder /app/.nuxt /app/.nuxt

CMD ["yarn", "start"]
