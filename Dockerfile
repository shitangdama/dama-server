FROM node:latest

COPY . /src
WORKDIR /src
RUN npm i

EXPOSE 3000

CMD ["node", "./index.js"]
