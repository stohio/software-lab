// grab the packages we need
var express = require('express');
var bodyParser = require('body-parser');
var app = express();
var path = require('path');

app.use(bodyParser.json()); // support json encoded bodies
//app.use(bodyParser.urlencoded({ extended: true })); // support encoded bodies
var port = process.env.PORT || 8080;

var object;
var name;
var private_ip;
var software;

app.use(express.static(path.join(__dirname, 'public')));

app.get('/server', function(req, res) {
  res.send(object);
});

app.post('/server', function(req, res) {
  object = req.body;
  console.log(req.body);

  name = object.name;
  private_ip = object.private_ip;
  software = object.software;

  res.send(req.body);
});

// start the server
app.listen(8080);
console.log('Server started! At http://localhost:' + port);
