#!/bin/bash

#Request a local server IP
#server_IP= curl -X GET -s 'software-lab.azurewebsites.net/server'
server_IP="10.1.26.45"

echo 'curl -x GET "$server_IP:8080/application?id=1" -o test.tmp -w "%{speed_download}"'


