package timeline

import (
	"testing"

	"github.com/Krayons/twitterclone/reader"
)

func TestPrintTimeline(t *testing.T) {
	usersToFollowers := make(map[string]reader.Followers)
	usersToFollowers["Kyle"] = reader.Followers{"Kyle", "Test"}
	usersToFollowers["Test"] = reader.Followers{"Test"}
	tweets := []reader.Tweet{reader.Tweet{User: "Kyle", Msg: "Hello!"}, reader.Tweet{User: "Test", Msg: "Test Tweet"}}
	output := BuildTimeline(usersToFollowers, tweets)
	expectedOutput := `Kyle
	@Kyle: Hello!
	@Test: Test Tweet
Test
	@Test: Test Tweet
`
	if output != expectedOutput {
		t.Log("output not as expected")
		t.Fail()
	}
}
