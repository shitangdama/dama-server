const socket = require('socket.io')
const _ = require('lodash')
// 这个轮子是可以使用的但是没有考虑redis的问题
const io = socket()

// 公共变量
// 格式是房间名：socketid：状态 默认房间为default
// 状态为1为正常，2为短线从连
// 方便统计
// 也可以使用io.sockets.sockets直接管理
// io.sockets.sockets的格式为socket.id: socket
const rooms = {}


// io.of('/').adapter.clients(['room1', 'room2'], function (err, clients) {
//   console.log(clients); // an array containing socket ids in 'room1' and/or 'room2'
// });
//
// // you can also use
//
// io.in('room3').clients(function (err, clients) {
//   console.log(clients); // an array containing socket ids in 'room3'
// });
// 先加入
// 选择使用加入房间的方式

io.on('connection',onConnection)

function onConnection(socket){
    // 这里存在很多问题
    // 还是按照基本的方式来处理,无法优化
    // 第一步是返回创建成功
    // sending to individual socketid
    // console.log(io.sockets.sockets)
    socket.emit('open', 'connection is ok');
    // 第二部选择房间
    // 加入房间
    socket.on('subscribe', (data) => {
      socket.join(data.room);
      // 对房间进行注册
      if(_.isEmpty(rooms[data.room])){
        rooms[data.room] = {};
        rooms[data.room][socket.id] = 1;
      }else {
        rooms[data.room][socket.id] = 1;
      }

      // 同时把房间名绑定到socket上
      socket.room = data.room
      socket.emit('subscribe', 'room is ok');
    })
    // 离开房间
    socket.on('unsubscribe', (data) => {
      socket.leave(data.room);
      delete rooms[data.room][socket.id];
      if(_.isEmpty(rooms[data.room])){
        delete rooms[data.room];
      }
    })
    // 弹幕发送
    socket.on('danmu', (data) => {
      io.sockets.in(socket.room).emit('danmu', data);
    })

    // 事件短线
    // 从连请求是客户端做，因此对于新请求的session和id是由web获取的session和socket.io-redis
    // 来处理，socket.io-redis中有ad
    socket.on('disconnect', () => {})
    socket.on('disconnecting', () => {})
    // 错误时间
    socket.on('error', () => {})
}

// 返回有多少key
function hasValuew(value) {
    return Object.keys(value).length;
}

io.listen(3000)

// 想要在做一个ws的推送，转移到index1中做
// // 控制用
// const Koa = require('koa');
// const app = new Koa();

// app.use(ctx => {
//   // 使用一个函数返回当前信息
//   ctx.body = 'Hello Koa';
// });
// app.listen(3001);
