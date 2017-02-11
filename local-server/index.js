// grab the packages we need
var express = require('express');
var app = express();
var port = process.env.PORT || 8080;

var bodyParser = require('body-parser');
app.use(bodyParser.json()); // support json encoded bodies
app.use(bodyParser.urlencoded({ extended: true })); // support encoded bodies


// POST http://localhost:8080/api/users
// parameters sent with 
app.post('/api/users', function(req, res) {
	var user_id = req.headers.id;
	var token = req.headers.token;
	var geo = req.headers.geo
	console.log(req.headers.id);
	res.send(user_id + ' ' + token + ' ' + geo);
});

// start the server
app.listen(port);
console.log('Server started! At http://localhost:' + port);
