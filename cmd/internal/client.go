package internal

import (
	"context"

	"github.com/carlmjohnson/requests"
)

type Client struct {
	BaseUrl string
	Token   string
}

type Token struct {
	Token string `json:"token"`
}

type Subscriptions struct {
	Subscriptions []Subscription `json:"subscriptions"`
}

type TokenWrapper struct {
	Data Token `json:"data"`
}

type SubscriptionsWrapper struct {
	Data Subscriptions `json:"data"`
}

type Subscription struct {
	Id   string `json:"id"`
	Link string `json:"link"`
	Tag  string `json:"tag"`
}

func NewClient(baseUrl string) *Client {
	return &Client{
		BaseUrl: baseUrl,
	}
}

func (c *Client) Auth(username, password string) error {
	resp := TokenWrapper{}
	err := requests.URL(c.BaseUrl + "/graphql").ContentType("application/json").BodyJSON(map[string]any{
		"query": "query Token($username: String!, $password: String!) {\n  token(username: $username, password: $password)\n}",
		"variables": map[string]string{
			"username": username,
			"password": password,
		},
		"operationName": "Token",
	}).ToJSON(&resp).Fetch(context.TODO())
	if err != nil {
		return err
	}
	c.Token = resp.Data.Token
	return nil
}

func (c *Client) GetSubscriptions() []Subscription {
	resp := SubscriptionsWrapper{}
	requests.URL(c.BaseUrl+"/graphql").Header("Authorization", "Bearer "+c.Token).ContentType("application/json").BodyJSON(map[string]any{
		"query":         "query Subscriptions {\n  subscriptions {\n    id\n    tag\n    link\n  }\n}",
		"operationName": "Subscriptions",
	}).ToJSON(&resp).Fetch(context.TODO())
	return resp.Data.Subscriptions
}

func (c *Client) RefreshSubscriptions(subscriptions []string) []string {
	failed := []string{}
	for _, subscription := range subscriptions {
		err := requests.URL(c.BaseUrl+"/graphql").Header("Authorization", "Bearer "+c.Token).ContentType("application/json").BodyJSON(map[string]any{
			"query": "mutation UpdateSubscription($id: ID!) {\n  updateSubscription(id: $id) {\n    id\n  }\n}",
			"variables": map[string]string{
				"id": subscription,
			},
			"operationName": "UpdateSubscription",
		}).Fetch(context.TODO())
		if err != nil {
			failed = append(failed, subscription)
		}
	}
	return failed
}

func (c *Client) Reload() error {
	return requests.URL(c.BaseUrl+"/graphql").Header("Authorization", "Bearer "+c.Token).ContentType("application/json").BodyJSON(map[string]any{
		"query": "mutation Run($dry: Boolean!) {\n  run(dry: $dry)\n}",
		"variables": map[string]bool{
			"dry": false,
		},
		"operationName": "Run",
	}).Fetch(context.TODO())
}
