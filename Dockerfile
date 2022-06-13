FROM node:14-alpine

RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app

ENV NODE_ENV production

COPY package.json /usr/src/app/
RUN npm install
COPY . /usr/src/app

ENV PORT 8080
EXPOSE 8080

CMD ["node", "./index.js"]
