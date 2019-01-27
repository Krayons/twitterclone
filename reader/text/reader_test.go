package text

import (
	"testing"
)

func TestReadFile_success(t *testing.T) {
	lines, err := ReadFile("./users_test.txt")
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	if len(lines) != 3 {
		t.Log("Incorrect number of lines returned")
		t.Fail()
	}
	if lines[0] != "Ward follows Alan" {
		t.Log("Incorrect text found")
		t.Fail()
	}
}

func TestToFollowers_success(t *testing.T) {
	lines, _ := ReadFile("./users_test.txt")
	m, err := ToFollowers(lines)
	if err != nil {
		t.Log("Failed to create follows map")
		t.Fail()
	}
	if len(m) != 3 {
		t.Log("Expected 3 members in the followers file")
		t.Fail()
	}
	if len(m["Kyle"]) != 3 {
		t.Log("Kyle expected to be following 2 users plus himself")
		t.Fail()
	}
}

func TestToFollowers_incorrectFormatting(t *testing.T) {
	lines := []string{"kyle follows "}
	_, err := ToFollowers(lines)
	if err == nil {
		t.Log("Expected error but got none")
		t.Fail()
	}
	if err.Error() != "Failed to parse: Users formatted expected as 'username follows username, username" {
		t.Log("Expected failed to parse but got different error")
		t.Fail()
	}
}

func TestToFollowers_incorrectFormattingTwo(t *testing.T) {
	lines := []string{"kyle welsh follows kyle"}
	_, err := ToFollowers(lines)
	if err == nil {
		t.Log("Expected error but got none")
		t.Fail()
	}
	if err.Error() != "Failed to parse: Users formatted expected as 'username follows username, username" {
		t.Log("Expected failed to parse but got different error")
		t.Fail()
	}
}

func TestToTweets_success(t *testing.T) {
	lines, _ := ReadFile("./tweets_test.txt")
	m, err := ToTweets(lines)
	if err != nil {
		t.Log("Successful conversion expected")
		t.Fail()
	}
	if len(m) != 4 {
		t.Log("expected 4 tweets")
		t.Fail()
	}
}

func TestToTweets_TooLongTweet(t *testing.T) {
	lines, _ := ReadFile("./tweets_test_toolong.txt")
	_, err := ToTweets(lines)
	if err == nil {
		t.Log("Expected tweet too long exception")
		t.Fail()
	}
	if err.Error() != "Tweets can be a maximum of 140 characters" {
		t.Log("expected too long of a tweet but got a different error")
		t.Fail()
	}
}

func TestToTweets_BadTweetFormat(t *testing.T) {
	line := []string{"kyle test follows test"}
	_, err := ToTweets(line)
	if err == nil {
		t.Log("Expected poorly formatted input")
		t.Fail()
	}
	if err.Error() != "Failed to parse tweets, expected format 'user> tweet'" {
		t.Log("Expected poorly formatted error but got a different error")
		t.Fail()
	}
}

func TestReadFile_noFileFound(t *testing.T) {
	lines, err := ReadFile("./missing_file.txt")
	if lines != nil || err == nil {
		t.Log("expected error")
		t.Fail()
	}
}
