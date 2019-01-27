package reader

// Tweet Contains the structure of a tweet
type Tweet struct {
	User string
	Msg  string
}

// Followers list of usernames
type Followers []string

// Contains returns true if a username is in a followers list
func (f *Followers) Contains(username string) bool {
	for _, s := range *f {
		if s == username {
			return true
		}
	}
	return false
}
