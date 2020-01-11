package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"url-shortner/cron"
	"url-shortner/server"
	"url-shortner/utils/concurrencyutils"
)

const (
	shutdownTimeout = 10 * time.Second
	logTag          = "main"
)

func main() {
	server.Start()

	err := cron.Start()
	if err != nil {
		log.Printf("Error initialising cron... %v", err.Error())
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-interrupt

	concurrencyutils.WaitChannels(server.Stop(), cron.Stop())
}
