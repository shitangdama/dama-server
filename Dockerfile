FROM node:latest

COPY . /src
RUN cd /src

EXPOSE 3000

CMD ["node", "./src/index.js"]
