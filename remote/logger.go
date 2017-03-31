package main

import (
	"net/http"
	"os"
	"time"

	"github.com/Sirupsen/logrus"
)

var routeLog = logrus.New()
var downloadLog = logrus.New()
var consoleLog = logrus.New()

// InitLogger initilizes the base loggers
func InitLogger() {
	//Create softwarelab folder if it doesnt exist
	if _, err := os.Stat("softwarelab"); os.IsNotExist(err) {
		os.Mkdir("softwarelab", 0755)
	}

	//Create a log folder if it doesnt exist
	if _, err := os.Stat("softwarelab/log"); os.IsNotExist(err) {
		os.Mkdir("softwarelab/log", 0755)

	}

	//Format date/time into string for log folder name
	t := time.Now()
	logFolderName := (t.Format("2006-01-02_15:04:05"))

	//Make folder for current session
	if _, err := os.Stat("softwarelab/log/" + logFolderName); os.IsNotExist(err) {
		os.Mkdir("softwarelab/log/"+logFolderName, 0755)
	}

	//Create Route logger, set output to route.log
	routeLog.Formatter = new(logrus.JSONFormatter)

	if _, err := os.Stat("softwarelab/log/" + logFolderName + "/route.log"); os.IsNotExist(err) {
		file, err := os.Create("softwarelab/log/" + logFolderName + "/route.log")
		if err != nil {
			panic(err)
		}
		routeLog.Out = file
	}

	//Create Download logger, set output to downloads.log
	downloadLog.Formatter = new(logrus.JSONFormatter)

	if _, err := os.Stat("softwarelab/log/" + logFolderName + "/route.log"); os.IsNotExist(err) {
		file, err := os.Create("softwarelab/log/" + logFolderName + "/route.log")
		if err != nil {
			panic(err)
		}
		downloadLog.Out = file
	}

}

// Route Logger logs route info for a route handler function
func RouteLogger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner.ServeHTTP(w, r)

		routeLog.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)

		consoleLog.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)

	})
}
