package main

import (
	//"encoding/json"
        "log"
	"fmt"
	"net/http"
	"io"
	//"io/ioutil"
        "os"
	//swl "github.com/stohio/software-lab/lib"
	"github.com/gorilla/mux"
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
        fmt.Printf("Enabled: ", n.Added)
}

// Endpoint to retrieve software from node by softwareID and versionID
func SoftwareGet(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    filename := "software/" + vars["software_id"] + "/" + vars["version_id"] + ".txt";
    fmt.Println(filename)

    // If the file doesnt exist
    if _, err := os.Stat(filename); os.IsNotExist(err) {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(404)
        return
    } else {
        AddClient()
        //I need to increment the counter
        w.Header().Set("Content-Type", "application/octet-stream")
        w.Header().Set("Content-Disposition", "attachment; filename='dong'")
        file, err := os.Open(filename)
        if err != nil {
            log.Fatal(err)
        }
        defer file.Close()
        // Copy sends the file to the client
        n, err := io.Copy(w, file)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println(n, "bytes send")
        RemoveClient()
    }
}
