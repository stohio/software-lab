#!/bin/bash




download() {
	echo "Requesting Local IP"
	#Request a local server IP
	server_IP=$(curl -X GET -s 'http://40.71.25.155:8080/get_ip')
	#server_IP="10.1.26.45"
	
	echo "IP is $server_IP"
	#download and get speed
	echo "Checking Download Speed"
	speed=$(curl -X GET -s "$server_IP:8080/application?id=$2" -o test.tmp -w "%{speed_download}")
	
	echo "Speed: $speed"
	#send speed to remote server
	echo "Sending $1"
	curl -H "Content-Type: application/json" -X POST -d "{\"hostname\": \"$1\",\"server\": \"$server_IP\", \"speed\":$speed}" "http://40.71.25.155:8080/data"
	
}

while :
do
	download "client-1" 3&
	download "client-2" 3&
	download "client-3" 3&
	download "client-4" 3&
	download "client-5" 3&
	sleep 1
	download "client-6" 3&
	sleep 1
	download "client-7" 3

done
