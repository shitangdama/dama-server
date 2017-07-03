FROM node:latest

COPY . /src
RUN cd /src
RUN npm i -g yarn
RUN yarn install

EXPOSE 3000

CMD ["node", "./src/index.js"]
