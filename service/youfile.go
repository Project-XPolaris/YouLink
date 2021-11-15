package service

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
)

var DefaultYouFileClient = NewYouFileClient()

type YouFileClient struct {
	client *resty.Client
}

func NewYouFileClient() *YouFileClient {
	return &YouFileClient{
		client: resty.New(),
	}
}

func (c *YouFileClient) Init(baseUrl string) {
	c.client.SetBaseURL(baseUrl)
}

func (c *YouFileClient) MoveFile(inputs []*Variable, function *Function) {
	resp, err := c.client.R().
		SetBody(map[string]interface{}{
			"callbackId": function.Id,
			"inputs":     inputs,
		}).
		Post("/youlink/movefile")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
	return
}
