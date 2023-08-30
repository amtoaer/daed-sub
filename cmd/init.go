package cmd

import (
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/amtoaer/daed-sub/cmd/internal"
)

func Init() error {
	config := internal.NewConfig()
	qs := []*survey.Question{
		{
			Name:     "baseUrl",
			Prompt:   &survey.Input{Message: "What is your daed URL(without trailing slash)?"},
			Validate: survey.Required,
		},
		{
			Name:     "username",
			Prompt:   &survey.Input{Message: "What is your username?"},
			Validate: survey.Required,
		},
		{
			Name:     "password",
			Prompt:   &survey.Password{Message: "What is your password?"},
			Validate: survey.Required,
		},
	}
	if err := survey.Ask(qs, config); err != nil {
		return err
	}
	client := internal.NewClient(config.BaseUrl)
	if err := client.Auth(config.Username, config.Password); err != nil {
		return err
	}
	subscriptions := client.GetSubscriptions()
	qs = []*survey.Question{
		{
			Name: "subscriptionId",
			Prompt: &survey.MultiSelect{
				Message: "Which subscription do you want to subscribe?",
				Options: func() []string {
					var options []string
					for _, subscription := range subscriptions {
						options = append(options, fmt.Sprintf("%s: %s", subscription.Id, subscription.Tag))
					}
					return options
				}(),
			},
			Validate: survey.Required,
		},
		{
			Name:     "interval",
			Prompt:   &survey.Input{Message: "What is your interval(in minutes)?"},
			Validate: survey.Required,
		},
	}
	if err := survey.Ask(qs, config); err != nil {
		return err
	}
	for idx, subscription := range config.SubscriptionId {
		config.SubscriptionId[idx] = strings.SplitN(subscription, ":", 2)[0]
	}
	if err := config.Save(); err != nil {
		return err
	}
	return nil
}
