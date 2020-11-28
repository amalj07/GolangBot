package main

import (
	"github.com/ChimeraCoder/anaconda"
	"github.com/sirupsen/logrus"
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
		"track": []string{"#testgolangbotproject"},
	})

	defer stream.Stop()

	for v := range stream.C {
		t, ok := v.(anaconda.Tweet)

		if !ok {
			logrus.Warningf("Unexpeted value: %T", v)
			continue
		}

		_, err := api.Retweet(t.Id, false)
		if err != nil{
			logrus.Errorf("Failed to Retweet %d: %v", t.Id, err)
		}

		logrus.Infof("Retweeted: %d", t.Id)
	}
}