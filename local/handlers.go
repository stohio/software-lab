package main

import (
	//"encoding/json"
	//"log"
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"time"
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
	fmt.Printf("ID: %d", n.ID)
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
	w.Header().Set("Content-Disposition", "attachment; filename=\""+name+"\"")
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

func PackageGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	packID, _ := strconv.Atoi(vars["package_id"])
	verID, _ := strconv.Atoi(vars["version_id"])
	pack := network.Stack.Packages[packID-1]
	// Create empty ZIP
	if _, err := os.Stat("packages"); os.IsNotExist(err) {
		os.Mkdir("packages", 0755)
	}
	if _, err := os.Stat("packages/" + vars["package_id"]); os.IsNotExist(err) {
		os.Mkdir("packages/"+vars["package_id"], 0755)
	}

	zipfile, err := os.Create("packages/" + vars["package_id"] + "/" + vars["version_id"] + ".zip")
	if err != nil {
		panic(err)
	}
	defer zipfile.Close()

	packageZip := zip.NewWriter(w)
	defer packageZip.Close()

	filename := pack.Name + ".zip"

	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))

	for _, s := range pack.Softwares {
		header := &zip.FileHeader{
			Name:         s.Name + " " + s.Versions[verID-1].OS + " " + s.Versions[verID-1].Architecture + s.Versions[verID-1].Extension,
			Method:       zip.Store,
			ModifiedTime: uint16(time.Now().UnixNano()),
			ModifiedDate: uint16(time.Now().UnixNano()),
		}
		fw, err := packageZip.CreateHeader(header)
		if err != nil {
			panic(err)
		}

		// Open the file so it can be processed into the zip file
		fname := "software/" + strconv.Itoa(s.ID) + "/" + vars["version_id"] + s.Versions[verID-1].Extension
		fi, err := os.Open(fname)

		if err != nil {
			panic(err)
		}

		if _, err = io.Copy(fw, fi); err != nil {
			panic(err)
		}

		if err = fi.Close(); err != nil {
			panic(err)
		}
	}
}
