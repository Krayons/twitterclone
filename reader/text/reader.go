package text

import (
	"bufio"
	"errors"
	"os"
	"regexp"
	"strings"

	"github.com/Krayons/twitterclone/reader"
)

// ToFollowers : Returns users to follow map. It is assumed you have to follow your self
func ToFollowers(lines []string) (map[string]reader.Followers, error) {
	userFollowerMap := make(map[string]reader.Followers)
	var re = regexp.MustCompile(`(?m)\w+`)
	for _, line := range lines {
		words := re.FindAllString(line, -1)
		if words == nil || len(words) < 3 {
			return nil, errors.New("Failed to parse: Users formatted expected as 'username follows username, username")
		}
		if strings.ToUpper(words[1]) != "FOLLOWS" {
			return nil, errors.New("Failed to parse: Users formatted expected as 'username follows username, username")
		}
		user := words[0]
		var followers reader.Followers
		for _, follower := range words[2:] {
			followers = append(followers, follower)
		}
		followers = append(followers, user) // You have to follow yourself
		userFollowerMap[user] = followers

	}
	return userFollowerMap, nil
}

// ToTweets : Returns a map that maps from user name to tweets
func ToTweets(lines []string) ([]reader.Tweet, error) {
	var tweets []reader.Tweet
	for _, tweet := range lines {
		posOfGuillemet := strings.Index(tweet, "> ")
		if posOfGuillemet < 1 {
			return nil, errors.New("Failed to parse tweets, expected format 'user> tweet'")
		}
		textComponent := tweet[posOfGuillemet+2:]
		if len(textComponent) > 140 {
			return nil, errors.New("Tweets can be a maximum of 140 characters")
		}
		user := tweet[:posOfGuillemet]
		tweets = append(tweets, reader.Tweet{User: user, Msg: textComponent})
	}
	return tweets, nil
}

// ReadFile read in a text file to an array of strings
func ReadFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	var lines []string
	defer file.Close()

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		strings.Trim(text, " ")
		// Ignore empty lines
		if len(text) != 0 {
			lines = append(lines, text)
		}
	}

	return lines, scanner.Err()
}
