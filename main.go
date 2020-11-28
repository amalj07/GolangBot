package main

import (
	"github.com/ChimeraCoder/anaconda"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"net/url"
	"os"
)

type logger struct {
	*logrus.Logger
}

func main() {

	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("Failed to load .env file")
	}

	var (
		consumerKey = os.Getenv("CONSUMER_KEY")
		consumerSecret = os.Getenv("CONSUMER_SECRET")
		accessToken = os.Getenv("ACCESS_TOKEN")
		accessTokenSecret = os.Getenv("ACCESS_TOKEN_SECRET")
	)

	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)

	log := &logger{logrus.New()}
	api.SetLogger(log)

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

		if t.RetweetedStatus != nil {
			continue
		}

		_, err := api.Retweet(t.Id, false)
		if err != nil{
			logrus.Errorf("Failed to Retweet %d: %v", t.Id, err)
		}

		logrus.Infof("Retweeted: %d", t.Id)
	}
}

func (log *logger) Critical(args ...interface{}) { log.Error(args...) }
func (log *logger) Criticalf(format string, args ...interface{}) { log.Errorf(format, args...) }
func (log *logger) Notice(args ...interface{}) { log.Info(args...) }
func (log *logger) Noticef(format string, args ...interface{}) { log.Infof(format, args...) }