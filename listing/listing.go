package listing

import "github.com/socialnews/reddit-feeder/link"

type Listing struct {
	Kind string `json:"kind"`
	Data Data   `json:"data"`
}

type Data struct {
	Modhash  string      `json:"modhash"`
	Children []link.Link `json:"children"`
	After    string      `json:"after"`
	Before   string      `json:"before"`
}
