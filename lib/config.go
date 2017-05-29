package softwarelab

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// SocketAddresses describes the various ip addresses and
// Ports used by the seperate components of software lab
// TODO for now there will be just one NetworkSetup, but
// for the future there could be a development and production
type SocketAddresses struct {
	RemoteServer string `json:"remote_server"`
	RemotePort   string `json:"remote_port"`
	LocalServer  string `json:"local_server"`
	LocalPort    string `json:"local_port"`
}

// LocalServerSettings contains general settings for the local
// server
type LocalServerSettings struct {
	DownloadRetries int `json:"download_retries"`
}

// Settings contains all the settings for server lab
type Settings struct {
	SocketAddresses     SocketAddresses     `json:"socket_addresses"`
	LocalServerSettings LocalServerSettings `json:"local_server_settings"`
}

// GetSettings returns a NetworkSetup struct detailing the
// proper ips and ports to use
func GetSettings() Settings {
	configFile, err := ioutil.ReadFile("../config/config.json")
	if err != nil {
		fmt.Printf("There was a problem opening the config file: %s\n", err)
	}

	var config Settings
	if err := json.Unmarshal(configFile, &config); err != nil {
		fmt.Printf("There was a problem Decoding the config JSON: %s\n", err)
	}

	return config
}

// GetRemoteServer returns the software labs remote server' IP address
func GetRemoteServer() string {
	return GetSettings().SocketAddresses.RemoteServer
}

// GetRemotePort returns the software labs remote server's port number
func GetRemotePort() string {
	return GetSettings().SocketAddresses.RemotePort
}

// GetLocalServer returns the software labs local server's IP address
func GetLocalServer() string {
	return GetSettings().SocketAddresses.LocalServer
}

// GetLocalPort returns the software labs local server's port number
func GetLocalPort() string {
	return GetSettings().SocketAddresses.LocalPort
}

// GetLocalDownloadRetries returns the number of tries the local server should
// retry download its software
func GetLocalDownloadRetries() int {
	return GetSettings().LocalServerSettings.DownloadRetries
}
