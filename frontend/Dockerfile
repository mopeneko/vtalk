FROM node:12-alpine

WORKDIR /usr/local/src/app

COPY package.json yarn.lock ./
RUN yarn install

COPY . .
CMD [ "yarn", "dev" ]

