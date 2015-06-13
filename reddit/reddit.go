package reddit

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/socialnews/reddit-feeder/listing"
)

func ListLinks(subreddit string) (*listing.Listing, error) {
	data, err := getListing(subreddit)
	if err != nil {
		return nil, err
	}

	var l listing.Listing
	if err := json.Unmarshal(data, &l); err != nil {
		log.Print(err.Error())
		return nil, err
	}
	return &l, nil
}

var getListing = func(subreddit string) ([]byte, error) {
	response, err := http.Get("http://reddit.com/r/" + subreddit + ".json")
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}
	body, err := ioutil.ReadAll(io.LimitReader(response.Body, 1049576))
	if err != nil {
		return nil, err
	}
	return body, nil
}
