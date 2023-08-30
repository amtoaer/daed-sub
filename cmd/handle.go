package cmd

import (
	"fmt"

	"github.com/amtoaer/daed-sub/cmd/internal"
)

func Handle() error {
	config := internal.NewConfig()
	if err := config.Load(); err != nil {
		return err
	}
	client := internal.NewClient(config.BaseUrl)
	if err := client.Auth(config.Username, config.Password); err != nil {
		return err
	}
	failedSubscriptionIds := client.RefreshSubscriptions(config.SubscriptionId)
	if len(failedSubscriptionIds) > 0 {
		fmt.Printf("Failed to refresh subscriptions: %v\n", failedSubscriptionIds)
	} else {
		fmt.Println("Refresh subscriptions successfully!")
	}
	return client.Reload()
}
