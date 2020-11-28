package main

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"
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

	result, _ := api.GetSearch("#Golang", nil)
	for _, tweet := range result.Statuses{
		fmt.Println(tweet.Text)
	}
}