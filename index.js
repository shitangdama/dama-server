const log4js = require('log4js');
const WebSocket = require('ws');
var logger = log4js.getLogger();

const wss = new WebSocket.Server({ port: 3000 });

wss.on('connection', function connection(ws) {
  ws.on('message', function incoming(message) {
    console.log('received: %s', message);
    logger.debug("Some debug messages");
  });
});
