package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/Krayons/twitterclone/reader/text"
	"github.com/Krayons/twitterclone/timeline"
)

func main() {
	userFile := flag.String("users", "./users.txt", "file location of the users file")
	tweetsFile := flag.String("tweets", "./tweets.txt", "file location of the tweets file")

	flag.Parse()
	userText, err := text.ReadFile(*userFile)
	if err != nil {
		log.Fatal(err.Error())
	}
	tweetText, err := text.ReadFile(*tweetsFile)
	if err != nil {
		log.Fatal(err.Error())
	}

	tweets, err := text.ToTweets(tweetText)
	if err != nil {
		log.Fatal(err.Error())
	}
	users, err := text.ToFollowers(userText)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(timeline.BuildTimeline(users, tweets))
}
