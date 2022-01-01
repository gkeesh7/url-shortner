package cron

import (
	"context"
	"github.com/robfig/cron"
	"log"
	"strings"
	"url-shortner/logic"
)

var (
	crn         *cron.Cron
	crnShutdown = make(chan struct{})
	ctx         = context.Background()
)

func init() {
	crn = cron.New()
}

func Start() error {
	log.Print("Starting Cron.....")

	//TODO: Make this frequency configurable
	cronFrequency := "0 0/5 * * * ?"

	expressions := GetExpression(cronFrequency)

	for _, expression := range expressions {
		err := crn.AddFunc(expression, deleteExpiredURLs)
		if err != nil {
			log.Printf("Couldn't add deletion of expired URLs into cron with Expression %v", err.Error())
			return err
		}
	}

	crn.Start()
	return nil
}

func GetExpression(cronStateFrequency string) []string {
	return strings.Split(cronStateFrequency, "$")
}

func deleteExpiredURLs() {
	//TODO take a distributed lock on a common key before executing the logic
	logic.DeleteExpiredURLs(ctx)
}

func Stop() <-chan struct{} {
	go func() {
		log.Print("Stopping Cron...")
		crn.Stop()
		close(crnShutdown)
	}()
	return crnShutdown
}
