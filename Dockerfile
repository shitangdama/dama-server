FROM node:latest

USER root

RUN mkdir /app
WORKDIR /app

COPY src/ /app/src
COPY index.js /app/index.js
COPY package.json /app/package.json
COPY package-lock.json /app/package-lock.json
COPY tsconfig.json /app/tsconfig.json

RUN npm install --registry=https://registry.npm.taobao.org

EXPOSE 3000

CMD ["npm", "start"]
