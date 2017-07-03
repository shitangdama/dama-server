FROM node:latest

COPY . /src
RUN cd /src
RUN npm i

EXPOSE 3000

CMD ["node", "./src/index.js"]
