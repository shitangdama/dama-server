const WebSocket = require('ws');

for(let i=1; i<400;i++){
  const ws = new WebSocket('ws://122.228.212.134:6767');
  ws.on('open', function open() {
    console.log("aaaaaaaaaaaaaa"+i)
    setInterval(function(){
        ws.send('i am ' + i);
        console.log("aaa"+i)
    },i*1000)
  });

  ws.on('message', function incoming(data) {
    console.log(data);
  });
}
