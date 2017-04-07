package main

import (
	//"encoding/json"
	//"log"
	"fmt"
	"io"
	"net/http"
	//"io/ioutil"
	"os"
	//swl "github.com/stohio/software-lab/lib"
	"strconv"

	"github.com/gorilla/mux"
	swl "github.com/stohio/software-lab/lib"
)

// Test endpoint also gets information about the Node
func Test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Node is up")
	n := GetNode()
	fmt.Printf("Id: %d", n.Id)
	fmt.Printf("Name: %s", *n.Name)
	fmt.Printf("IP: %s", *n.IP)
	fmt.Printf("Enabled: %t", n.Enabled)
	fmt.Printf("Clients: %d", n.Clients)
	fmt.Printf("Added: %v", n.Added)
}

// SoftwareGet is an endpoint to retrieve software from node by softwareID and versionID
func SoftwareGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	softID, _ := strconv.Atoi(vars["software_id"])
	verID, _ := strconv.Atoi(vars["version_id"])
	software := network.Stack.Softwares[softID-1]
	version := software.Versions[verID-1]
	filename := "software/" + vars["software_id"] + "/" + vars["version_id"] + version.Extension
	fmt.Println("Sending ", filename)
	swl.DownloadLog.Info("Software Request " + vars["software_id"] + " Version " + vars["version_id"])
	swl.ConsoleLog.Info("Software Request " + vars["software_id"] + " Version " + vars["version_id"])

	// If the file doesnt exist
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		return
	}
	//name := s.Name
	name := software.Name + " " + version.OS + " " + version.Architecture + version.Extension
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename='"+name+"'")
	file, err := os.Open(filename)
	if err != nil {

		panic(err)
	}
	defer file.Close()
	// Copy sends the file to the client
	AddClient()
	n, err := io.Copy(w, file)
	if err != nil {
		swl.DownloadLog.Info("Cancelled Request " + vars["software_id"] + " Version " + vars["version_id"])
		swl.ConsoleLog.Info("Cancelled Request " + vars["software_id"] + " Version " + vars["version_id"])
		//panic(err)
	}
	fmt.Println(n, "bytes sent")
	RemoveClient()

}
