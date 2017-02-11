// grab the packages we need
var express = require('express');
var app = express();
var port = process.env.PORT || 8080;
var fileSystem = require('fs');
var path = require('path');

var bodyParser = require('body-parser');
app.use(bodyParser.json()); // support json encoded bodies
app.use(bodyParser.urlencoded({ extended: true })); // support encoded bodies



app.get('/application', function(req, res) {
	var filePath= path.join(__dirname, 'test.txt');
	var stat = fileSystem.statSync(filePath);
	res.writeHead(200, {'Content-Type': 'application/octet-stream'});
	var readStream = fileSystem.createReadStream(filePath);
	readStream.pipe(res);

});

// start the server
app.listen(port);
console.log('Server started! At http://localhost:' + port);
