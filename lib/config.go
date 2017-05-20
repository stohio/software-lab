package softwarelab

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// NetworkSetup describes the various ip addresses and
// Ports used by the seperate components of software lab
// TODO for now there will be just one NetworkSetup, but
// for the future there could be a development and production
type NetworkSetup struct {
	WebsiteServer string `json:"website_server"`
	WebsitePort   string `json:"website_port"`
	RemoteServer  string `json:"remote_server"`
	RemotePort    string `json:"remote_port"`
	LocalServer   string `json:"local_server"`
	LocalPort     string `json:"local_port"`
}

// GetNetworkSettings returns a NetworkSetup struct detailing the
// proper ips and ports to use
func GetNetworkSettings() NetworkSetup {
	configFile, err := ioutil.ReadFile("../config/config.json")
	if err != nil {
		fmt.Printf("There was a problem opening the config file: %s\n", err)
	}

	var config NetworkSetup
	if err := json.Unmarshal(configFile, &config); err != nil {
		fmt.Printf("There was a problem Decoding the config JSON: %s\n", err)
	}

	return config
}

// GetWebsiteServer returns the software labs website's IP address
func GetWebsiteServer() string {
	return GetNetworkSettings().WebsiteServer
}

// GetWebsitePort returns the software labs website's port number
func GetWebsitePort() string {
	return GetNetworkSettings().WebsitePort
}

// GetRemoteServer returns the software labs remote server' IP address
func GetRemoteServer() string {
	return GetNetworkSettings().RemoteServer
}

// GetRemotePort returns the software labs remote server's port number
func GetRemotePort() string {
	return GetNetworkSettings().RemotePort
}

// GetLocalServer returns the software labs local server's IP address
func GetLocalServer() string {
	return GetNetworkSettings().LocalServer
}

// GetLocalPort returns the software labs local server's port number
func GetLocalPort() string {
	return GetNetworkSettings().LocalPort
}
