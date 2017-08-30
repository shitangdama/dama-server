# A ChatOps Server Framework for Node.js

### develop
```
npm install -g typescript

npm install

```
### run server

```
npm start
```

### deploy
you should tsc your code with prod mode before deploying.
```
docker build -t <image-name> .

docker run -p <container-port>:<host-port> --name <container-name> <image-name>
```
