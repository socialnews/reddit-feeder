package link

type Link struct {
	Kind string   `json:"kind"`
	Data LinkData `json:"data"`
}

type LinkData struct {
	Id          string `json:"id"` // Prepend with "t3_", or use Name
	Name        string `json:"name"`
	NumComments int64  `json:"num_comments"`
	SelfText    string `json:"selftext"`
	SubredditId string `json:"subreddit_id"`
	Ups         int64  `json"ups"`
	Url         string `json:"url"`
}
