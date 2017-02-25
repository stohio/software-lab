// grab the packages we need
var express = require('express');
var bodyParser = require('body-parser');
var app = express();
var path = require('path');
var PubNub = require('pubnub')

app.use(bodyParser.json()); // support json encoded bodies
app.use(bodyParser.urlencoded({
	  extended: true
}));

//app.use(bodyParser.urlencoded({ extended: true })); // support encoded bodies
var port = process.env.PORT || 8080;

var name;
var private_ip;
var software;

app.use(express.static(path.join(__dirname, 'public')));

var node_list = [];
node_list.push({'ip':'10.1.26.32', 'active_downloaders':0});
//node_list.push({'ip':'1.252', 'active_downloaders':2});
//node_list.push({'ip':'1.999', 'active_downloaders':5});

app.get('/get_cluster', function(req, res) {
	res.send(node_list);
});
app.get('/apps', function(req, res) {
	var application_info = {"software": [{"clean_name": "Android Studio", "id" : 1, "name": "android-studio-ide-145.3360264-linux.zip"}, {"clean_name": "JRE 1.8", "id":2 ,"name": "jre-8u121-linux-i586.tar.gz"} , {"clean_name": "Postman", "id":3 ,"name":"Postman-linux-x64-4.9.3.tar.gz"} , {"clean_name": "AuthPy", "id":4 ,"name":"authy-authy-python-f085687.zip"} , {"clean_name": "MonoDevelop", "id":5 ,"name":"monodevelop-6.1.2.44-1.flatpak"} , {"clean_name": "SimpleSMS", "id":6,"name":"simpleSMS-master.zip"} , {"clean_name": "Ngrok x64", "id":7 ,"name":"ngrok-stable-linux-amd64.zip"} , {"clean_name": "Java OCR", "id":8 ,"name":"javaocr-20100605.zip"} , {"clean_name": "JDK 1.8", "id":9 ,"name":"jdk-8u121-linux-i586.tar.gz"}, {"clean_name": "Android Bundle", "id":10, "name":"android_bundle.zip"} ]};
	res.send(application_info);	
});


app.get('/init_node', function(req, res) {
	console.log("init node");
	console.log(req.private_ip);
	// sends back what the root node is when a new node joins the cluster
	// this lets the new node download all current files from the original root node
	res.send(node_list[0]['ip']);
});

app.post('/register_node', function(req, res) {
	node_list.push({'ip':req.body['ip'], 'active_downloaders':0});
	console.log(node_list);
	res.send('success');

});

app.get('/get_ip', function(req, res) {
	// find ip with least amount of active downloaders
	var min = node_list[0]["active_downloaders"];
	var min_ip = node_list[0]["ip"];
	for (var i = 1, len = node_list.length; i < len; i++) {
		if(node_list[i]["active_downloaders"] < min){
			min = node_list[i]["active_downloaders"];
			min_ip = node_list[i]["ip"];
		}
	  }
	console.log(min, min_ip);
	res.send(min_ip);
	
});

app.post('/subtract_downloading_user', function(req, res) {
	var arrayFound = node_list.filter(function(node_list) {
		    return node_list.ip == req.body['ip'];
	});
	arrayFound[0]['active_downloaders']--;
	console.log(node_list);
	res.send('success');
});

app.post('/add_downloading_user', function(req, res) {
	var arrayFound = node_list.filter(function(node_list) {
		    return node_list.ip == req.body['ip'];
	});
	arrayFound[0]['active_downloaders']++;
	console.log(node_list);
	res.send('success');
});

app.post('/data', function(req, res) {
	
	console.log("hostname: " + req.body.hostname);
	console.log("server: " + req.body.server);
	console.log("speed: " + req.body.speed);
	
	pubnub = new PubNub({
		publishKey : 'pub-c-ee515da9-2add-4288-8f22-69129e66fb8d',
		subscribeKey : 'sub-c-37a10fcc-f112-11e6-99a6-02ee2ddab7fe'
	});
	var publishConfig = {
		channel : "clientdata",
		message: req.body
	}
	pubnub.publish(publishConfig, function(status, response) {
		console.log(status, response);
	});

	res.send('success');
});

// start the server
app.listen(8080);
console.log('Server started! At http://localhost:' + port);
