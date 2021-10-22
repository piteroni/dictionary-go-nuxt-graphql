FROM node:14-alpine3.12

RUN apk update && \
    apk --no-cache add tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
    apk del tzdata

ENV TZ=Asia/Tokyo

WORKDIR /app
    
ENV PS1='\h /\W\ # '

COPY package.json /app/package.json
COPY yarn.lock /app/yarn.lock

RUN yarn

CMD ["yarn", "dev"]
