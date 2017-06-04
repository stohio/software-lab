package main

import (
	//"encoding/json"
	//"log"
	"archive/zip"
	"encoding/json"
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
	filename := "softwarelab/software/" + vars["software_id"] + "/" + vars["version_id"] + version.Extension
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

	fileStat, err := file.Stat()
	if err != nil {
		panic(err)
	}
	fileSize := int(fileStat.Size())
	w.Header().Set("Content-Length", strconv.Itoa(fileSize))
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

// PackageGet is an endpoint to retrieve packages from node by packageID and versionID
func PackageGet(w http.ResponseWriter, r *http.Request) {

	//Unpack mux http params
	vars := mux.Vars(r)
	packID, _ := strconv.Atoi(vars["package_id"])
	verID, _ := strconv.Atoi(vars["version_id"])

	//Locate the package being used, if it doesn't exist, return a 404.
	pack := RepoFindPackage(packID)

	if pack == nil {
		paramError := swl.ParamError{
			Error: "Could not find package",
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(404)
		if err := json.NewEncoder(w).Encode(paramError); err != nil {
			panic(err)
		}
		return
	}

	//Create a zip writer that writes the the http writer
	packageZip := zip.NewWriter(w)
	packageSize := 0
	defer packageZip.Close()

	//Get each software in the package and add it to the zip
	for _, s := range pack.Softwares {
		version := RepoFindSoftwareVersion(s.ID, verID)
		header := &zip.FileHeader{
			Name:         s.Name + " " + version.OS + " " + version.Architecture + version.Extension,
			Method:       zip.Store,
			ModifiedTime: uint16(time.Now().UnixNano()),
			ModifiedDate: uint16(time.Now().UnixNano()),
		}

		zipWriter, err := packageZip.CreateHeader(header)
		if err != nil {
			panic(err)
		}

		// Open the file so it can be processed into the zip file
		softwareFilename := "softwarelab/software/" + strconv.Itoa(s.ID) + "/" + vars["version_id"] + version.Extension

		softwareFile, err := os.Open(softwareFilename)
		if err != nil {
			panic(err)
		}

		//Process file into zip
		n, err := io.Copy(zipWriter, softwareFile)

		if err != nil {
			panic(err)
		}
		packageSize += int(n)

		if err = softwareFile.Close(); err != nil {
			panic(err)
		}
	}

	//Set response header info
	filename := pack.Name + ".zip"
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))
	w.Header().Set("Content-Length", strconv.Itoa(packageSize))
}
