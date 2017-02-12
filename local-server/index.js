// grab the packages we need
var express = require('express');
var app = express();
var port = process.env.PORT || 8080;
var fileSystem = require('fs');
var path = require('path');
var http = require('http');

var bodyParser = require('body-parser');

app.use(cors());
app.use(bodyParser.json()); // support json encoded bodies
app.use(bodyParser.urlencoded({ extended: true })); // support encoded bodies

// placeholder filename
var filename = 'test.txt';

var options = {
	name: 'Brick Hack',
	private_ip: '10.2.1.91',
	software: []
	
};

// initialize server information with remote
app.get('/application', function(req, res) {
	var filePath= path.join(__dirname, filename);
	var stat = fileSystem.statSync(filePath);
	res.writeHead(200, {
		'Content-Type': 'application/octet-stream',
		'Content-Disposition': 'attachment; filename=' + filename
	});
	var readStream = fileSystem.createReadStream(filePath);
	readStream.pipe(res);

});

// start the server
app.listen(port);
console.log('Server started! At http://localhost:' + port);
