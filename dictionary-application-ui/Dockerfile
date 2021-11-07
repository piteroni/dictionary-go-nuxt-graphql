# base
FROM node:14-alpine3.12 AS base

RUN apk update && \
    apk --no-cache add tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
    apk del tzdata

ENV TZ=Asia/Tokyo

WORKDIR /app

COPY package.json /app/package.json
COPY yarn.lock /app/yarn.lock

RUN yarn

# development
FROM base as dev

ENV PS1='\h /\W\ # '

CMD ["yarn", "dev"]

# production
FROM base as builder

COPY . .

RUN yarn build

FROM base AS production

COPY --from=builder /app/.nuxt /app/.nuxt
COPY --from=builder /app/static /app/static

CMD ["yarn", "start"]
