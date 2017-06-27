FROM ubuntu:latest
MAINTAINER kbr1990117@gmail.com



RUN apt-get update
RUN apt-get install -y nodejs
RUN apt-get install -y npm
RUN apt-get install -y wget
RUN apt-get install -y curl
RUN npm i -g n
RUN n latest
RUN n use latest
COPY . /src
RUN cd /src

RUN npm install --registry=https://registry.npm.taobao.org


EXPOSE 3000

CMD ["npm", "start"]
