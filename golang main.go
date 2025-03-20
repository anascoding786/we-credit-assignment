package main

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

// Initialize logger
var log = logrus.New()

func init() {
	// Set output to file
	file, err := os.OpenFile("/var/log/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.Out = file
	} else {
		log.Out = os.Stdout
	}

	// Set log format and level
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetLevel(logrus.DebugLevel)
}

func main() {
	for {
		log.Info("This is an INFO log")
		log.Warn("This is a WARN log")
		log.Debug("This is a DEBUG log")
		log.Error("This is an ERROR log")

		time.Sleep(5 * time.Second)
	}
}
