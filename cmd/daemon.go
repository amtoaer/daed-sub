package cmd

import (
	"log"
	"time"

	"github.com/amtoaer/daed-sub/cmd/internal"
)

func Daemon() error {
	config := internal.NewConfig()
	if err := config.Load(); err != nil {
		return err
	}
	client := internal.NewClient(config.BaseUrl)
	for ticker := time.NewTicker(time.Duration(config.Interval) * time.Minute); ; <-ticker.C {
		if err := client.Auth(config.Username, config.Password); err != nil {
			return err
		}
		failedSubscriptionIds := client.RefreshSubscriptions(config.SubscriptionId)
		if len(failedSubscriptionIds) > 0 {
			log.Printf("Failed to refresh subscriptions: %v\n", failedSubscriptionIds)
		} else {
			log.Println("Refresh subscriptions successfully!")
		}
		if err := client.Reload(); err != nil {
			log.Printf("Failed to reload: %v\n", err)
		} else {
			log.Println("Reload successfully!")
		}
	}
}
