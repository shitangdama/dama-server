FROM node:latest

USER root

RUN mkdir /app
WORKDIR /app

COPY dist/ /app/
COPY package.json /app/package.json
COPY package-lock.json /app/package-lock.json

RUN npm install --registry=https://registry.npm.taobao.org

EXPOSE 3000

CMD ["node", "./app.js"]
