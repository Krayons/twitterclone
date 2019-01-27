package timeline

import (
	"fmt"

	"github.com/Krayons/twitterclone/reader"
)

// Timeline of a user
type Timeline struct {
	User           string
	ReleventTweets []string
}

// BuildTimeline returns users time lines
func BuildTimeline(users map[string]reader.Followers, tweets []reader.Tweet) string {
	var output string
	for user, followers := range users {
		output += fmt.Sprintln(user)
		for _, tweet := range tweets {
			if followers.Contains(tweet.User) {
				output += fmt.Sprintf("\t@%s: %s\n", tweet.User, tweet.Msg)
			}
		}
	}
	return output
}
