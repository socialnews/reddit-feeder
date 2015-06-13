package reddit

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/socialnews/reddit-feeder/listing"
)

func ListLinks(subreddit string) (*listing.Listing, error) {
	s, err := getListing(subreddit)
	if err != nil {
		return nil, err
	}

	var l listing.Listing
	if err := json.Unmarshal([]byte(s), &l); err != nil {
		log.Print(err.Error())
		return nil, err
	}
	return &l, nil
}

var getListing = func(subreddit string) (string, error) {
	return http.Get("http://reddit.com/r/" + subreddit + ".json"), nil
}
