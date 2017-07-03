var log4js = require('log4js');
var logger = log4js.getLogger();
// logger.level = 'debug';

setInterval(function(){
logger.debug("Some debug messages");
},1000)
