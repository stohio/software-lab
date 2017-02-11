// grab the packages we need
var express = require('express');
var app = express();
var port = process.env.PORT || 8080;

var bodyParser = require('body-parser');
app.use(bodyParser.json()); // support json encoded bodies
app.use(bodyParser.urlencoded({ extended: true })); // support encoded bodies


// POST http://localhost:8080/api/users
// parameters sent with 
app.get('/', function(req, res) {
  res.send("Hello there how are oyu");
});


// start the server
app.listen(port);
console.log('Server started! At http://localhost:' + port);
