// grab the packages we need
var express = require('express');
var bodyParser = require('body-parser');
var app = express();
var path = require('path');

app.use(bodyParser.json()); // support json encoded bodies
app.use(bodyParser.urlencoded({
	  extended: true
}));

//app.use(bodyParser.urlencoded({ extended: true })); // support encoded bodies
var port = process.env.PORT || 8080;

var name;
var private_ip;
var software;

//app.use(express.static(path.join(__dirname, 'public')));
app.get('/', function(req,res) {
  res.send("Hello world");
});

var node_list = [];
node_list.push({'ip':'10.2.0.252', 'active_downloaders':0});

app.get('/apps', function(req, res) {
	var application_info = {"software": [{"clean_name": "Android Studio", "id" : 1, "name": "android-studio-ide-145.3360264-linux.zip"}, {"clean_name": "JRE 1.8", "id":2 ,"name": "jre-8u121-linux-i586.tar.gz"} , {"clean_name": "Postman", "id":3 ,"name":"Postman-linux-x64-4.9.3.tar.gz"} , {"clean_name": "AuthPy", "id":4 ,"name":"authy-authy-python-f085687.zip"} , {"clean_name": "MonoDevelop", "id":5 ,"name":"monodevelop-6.1.2.44-1.flatpak"} , {"clean_name": "SimpleSMS", "id":6,"name":"simpleSMS-master.zip"} , {"clean_name": "Ngrok x64", "id":7 ,"name":"ngrok-stable-linux-amd64.zip"} , {"clean_name": "Java OCR", "id":8 ,"name":"javaocr-20100605.zip"} , {"clean_name": "JDK 1.8", "id":9 ,"name":"jdk-8u121-linux-i586.tar.gz"}, {"clean_name": "Android Bundle", "id":10, "name":"android_bundle.zip"} ]};
	res.send(application_info);	
});


app.get('/init_node', function(req, res) {
	// sends back what the root node is when a new node joins the cluster
	// this lets the new node download all current files from the original root node
	res.send(node_list[0]['ip']);
});
app.post('/register_node', function(req, res) {
	node_list.push({'ip':req.body['ip'], 'active_downloaders':0});
	console.log(node_list);
	res.send('success');

});


// start the server
app.listen(8080);
console.log('Server started! At http://localhost:' + port);
