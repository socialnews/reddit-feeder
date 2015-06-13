package reddit

import (
	"encoding/json"

	"github.com/socialnews/reddit-feeder/listing"
)

func ListLinks(subreddit string) ([]listing.Listing, error) {
	s, err := getListing(subreddit)
	if err != nil {
		return nil, err
	}

	var p []listing.Listing
	if err := json.Unmarshal([]byte(s), &p); err != nil {
		return nil, err
	}
	return p, nil
}

var getListing = func(subreddit string) (string, error) {
	return "", nil
}
