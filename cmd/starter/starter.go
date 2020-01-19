package main

import (
	"flag"
	//"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/micahli/notable-take-home/server"

	"github.com/sirupsen/logrus"
)

var (
	configFileFlag = flag.String("config.file", "config.yml", "Path to the configuration file.")
	debugFlag      = flag.Bool("debug", true, "Show debug information.")
)

func init() {
	// Parse command-line flags
	flag.Parse()

	// Log settings
	if *debugFlag {
		logrus.SetReportCaller(true)
		logrus.SetLevel(logrus.TraceLevel)
	} else {
		logrus.SetReportCaller(false)
		logrus.SetLevel(logrus.InfoLevel)
	}

	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
}

func main() {
	// Create server instance
	instance := server.NewInstance()

	// Interrupt handler
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		logrus.Infof("Received %s signal", <-c)
		instance.Shutdown()
	}()

	// Start server
	instance.Start(*configFileFlag)
}
