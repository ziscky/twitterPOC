package twitter

import (
	"io/ioutil"
	"log"
)

//Twitter is a representation of all the details needed by twitter to use the REST API
type Twitter struct {
	ConsumerKey    string
	ConsumerSecret string
	AccessToken    string
	AccessSecret   string
	Username       string
	Debug          bool //to set OAUTH log level, if you want to see the response headers
}

//Tweet tweets a given status, returns the raw twitter api response
func (twitter *Twitter) Tweet(status string) (string, error) {
	var endpoint = "https://api.twitter.com/1.1/statuses/update.json"

	client := new(Client)
	client.ConsumerKey = twitter.ConsumerKey
	client.ConsumerSecret = twitter.ConsumerSecret
	accessToken := &Token{AccessToken: twitter.AccessToken, AccessSecret: twitter.AccessSecret}

	params := map[string]string{
		"status": status,
	}

	response, err1 := client.Request(endpoint, "POST", "", params, accessToken)
	if err1 != nil {
		log.Println("LOG_FATAL", err1)
		return "", err1
	}
	defer response.Body.Close()

	respBody, err2 := ioutil.ReadAll(response.Body)
	if err2 != nil {
		log.Println("LOG_FATAL", err2)
		return "", err2
	}
	return string(respBody), nil
}
