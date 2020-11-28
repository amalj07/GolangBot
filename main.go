package main

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"net/url"
)

const (
	consumerKey = "your consumer key"
	consumerSecret = "your consumer secret"
	accessToken = "your access token"
	accessTokenSecret = "your access token secret"
)

func main() {
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)

	stream := api.PublicStreamFilter(url.Values{
		"track": []string{"#golang, #Golang, #Goprogramminglanguage"},
	})

	defer stream.Stop()

	for v := range stream.C {
		fmt.Printf("%s\n", v)
	}
}