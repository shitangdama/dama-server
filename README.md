# A ChatOps Server Framework for Node.js

### develop
```
npm install -g typescript
npm install -g webpack

npm install

tsc --watch
```
### deploy
you should tsc your code with prod mode before deploying.
```
docker build -t <image-name> .

docker run -d -p <container-port>:<host-port> --name <container-name> <image-name>
```
