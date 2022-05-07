package main

import (
	"fmt"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func main() {

	const ConsumerKey = "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	const ConsumerSecret = "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	const AccessToken = "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	const AccessSecret = "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"

	Config := oauth1.NewConfig(ConsumerKey, ConsumerSecret)
	Token := oauth1.NewToken(AccessToken, AccessSecret)

	httpClient := Config.Client(oauth1.NoContext, Token)

	// twitter client
	Client := twitter.NewClient(httpClient)

	// post tweet with media
	Response, _, _ := Client.Statuses.Update("Hello World @hrittikhere yo !", nil)
	fmt.Println(Response, "Posted")
}
