const log4js = require('log4js');
const WebSocket = require('ws');
var logger = log4js.getLogger();

const = new WebSocket.Server({ port: 3000 });


//
// 
wss.on('connection', function connection(ws) {
  ws.on('message', function incoming(message) {
    console.log('received: %s', message);
      console.log(111111111111)
      console.log("111111111111")
      console.log("asdasdasd")
      console.log("测试")
      logger.debug("Some debug messages");
      ws.send('something');
  });
});
