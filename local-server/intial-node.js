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
app.use(bodyParser.json()); 
app.use(bodyParser.urlencoded({ extended: true })); 

remote_server_ip = 'http://40.71.25.155:8080'
var j = 0;

// software list
var files = {"software": [{"clean_name": "Android Studio", "id" : 1, "name": "android-studio-ide-145.3360264-linux.zip"}, {"clean_name": "JRE 1.8", "id":2 ,"name": "jre-8u121-linux-i586.tar.gz"} , {"clean_name": "Postman", "id":3 ,"name":"Postman-linux-x64-4.9.3.tar.gz"} , {"clean_name": "AuthPy", "id":4 ,"name":"authy-authy-python-f085687.zip"} , {"clean_name": "MonoDevelop", "id":5 ,"name":"monodevelop-6.1.2.44-1.flatpak"} , {"clean_name": "SimpleSMS", "id":6,"name":"simpleSMS-master.zip"} , {"clean_name": "Ngrok x64", "id":7 ,"name":"ngrok-stable-linux-amd64.zip"} , {"clean_name": "Java OCR", "id":8 ,"name":"javaocr-20100605.zip"} , {"clean_name": "JDK 1.8", "id":9 ,"name":"jdk-8u121-linux-i586.tar.gz"}, {"clean_name": "Android Bundle", "id":10 ,"name":"android_bundle.zip"}  ]};
	
// placeholder filename
var filename = 'test.txt';

//initialize server information with remote
//console.log("running");
request.get({
	url: remote_server_ip + '/init_node',
	json: {
		'name': 'node',
		'private_ip':'10.1.26.32'
	}
}, function(err,httpResponse,body){
});

app.get('/application', function(req, res) {
    console.log("got request!!!!!!");
	  console.log(files["software"][req.query["id"]]["name"]);

    request.post({
    	url: remote_server_ip + '/add_downloading_user',
      json: {
        'name': 'Brick Hack',
        'ip': '10.1.26.32'
      }
    }, function(err, httpResponse, body) {
      if(err) {
        console.log('Error in adding count from node');
        return;
      }
    });

	var fileName;

  for (var i = 0; i < files["software"].length; i++) {
    if (files["software"][i]["id"] == req.query["id"]) {
      fileName = files["software"][i]["name"]
      break;
    }
  }
	var filePath= path.join(__dirname, fileName);
	var stat = fileSystem.statSync(filePath);
	res.writeHead(200, {
		'Content-Type': 'application/octet-stream',
		'Content-Disposition': 'attachment; filename=' + fileName
	});
	var readStream = fileSystem.createReadStream(filePath);
	var result = readStream.pipe(res);
  result.on('finish', function(){

    console.log("FINISHED IT ALL");
    j++;
	  console.log(files["software"][req.query["id"]]["name"]);

    request.post({
    	url: remote_server_ip + '/subtract_downloading_user',
      json: {
        'name': 'Brick Hack',
        'ip': '10.1.26.32'
      }
    }, function(err, httpResponse, body) {
      if(err) {
        console.log('Error in substractin count from node');
        return;
      }
    });

  });

  console.log(j, "This is the count i");

});

// start the server
app.listen(port);
console.log('Server started! At http://localhost:' + port);
