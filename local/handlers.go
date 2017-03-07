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

func SoftwareGet(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    //softwareID, err := strconv.Atoi(vars["software_id"])
    //versionID, err := strconv.Atoi(vars["version_id"])

    filename := "software/" + vars["software_id"] + "/" + vars["version_id"];

    if _, err := os.Stat(filename); os.IsNotExist(err) {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(404)
        return
    } else {
        file, err := os.Open(filename)
        if err != nil {
            log.Fatal(err)
        }
        defer file.Close()
        n, err := io.Copy(w, file)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println(n, "bytes send")
    }
}
