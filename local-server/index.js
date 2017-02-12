// grab the packages we need
var express = require('express');
var app = express();
var port = process.env.PORT || 8080;
var fileSystem = require('fs');
var path = require('path');
var http = require('http');
var cors = require('cors');
var request = require('request');
var bodyParser = require('body-parser');

app.use(cors());
app.use(bodyParser.json()); // support json encoded bodies
app.use(bodyParser.urlencoded({ extended: true })); // support encoded bodies

// software list

var files = {"software": [{"clean_name": "Android Studio", "id" : 1, "name": "android-studio-ide-145.3360264-linux.zip"}, {"clean_name": "JRE 1.8", "id":2 ,"name": "jre-8u121-linux-i586.tar.gz"} , {"clean_name": "Postman", "id":3 ,"name":"Postman-linux-x64-4.9.3.tar.gz"} , {"clean_name": "AuthPy", "id":4 ,"name":"authy-authy-python-f085687.zip"} , {"clean_name": "MonoDevelop", "id":5 ,"name":"monodevelop-6.1.2.44-1.flatpak"} , {"clean_name": "SimpleSMS", "id":6,"name":"simpleSMS-master.zip"} , {"clean_name": "Ngrok x64", "id":7 ,"name":"ngrok-stable-linux-amd64.zip"} , {"clean_name": "Java OCR", "id":8 ,"name":"javaocr-20100605.zip"} , {"clean_name": "JDK 1.8", "id":9 ,"name":"jdk-8u121-linux-i586.tar.gz"} ]};
	
console.log(files["software"][0]["name"]);

// placeholder filename
var filename = 'test.txt';

// placeholder for json
//var options = {
	//name: 'Brick Hack',
	//private_ip: '10.2.1.91',
	//software: []
	
//};

// initialize server information with remote
request.post({
	url:'http://software-lab.azurewebsites.net/server',
	json: {
		'name': 'Brick Hack',
		'private_ip':'10.2.1.91',
		'software': [{"id": 1, "name": "android studio", "os": "Windows", "architecture": "x86"}]

	}
}, function(err,httpResponse,body){
	console.log(body);
	if(err){
		console.log('Error in something', err);
		return;
	}
});

console.log("sent posts");
app.get('/test', function(req, res) {
console.log(req.query["id"]);	
});
app.get('/application', function(req, res) {
	var fileName = files["software"][req.query["id"]]["name"];
	var filePath= path.join(__dirname, fileName);
	var stat = fileSystem.statSync(filePath);
	res.writeHead(200, {
		'Content-Type': 'application/octet-stream',
		'Content-Disposition': 'attachment; filename=' + fileName
	});
	var readStream = fileSystem.createReadStream(filePath);
	readStream.pipe(res);

});

// start the server
app.listen(port);
console.log('Server started! At http://localhost:' + port);
