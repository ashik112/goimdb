package model

type Files struct {
	Title   string
	Ratings string
	People  string
	Persons string
	Crew    string
	Episode string
}
type Solr struct {
	Hostname string
	Port     int
	Core     string
}

type Imdb struct {
	URL string
}
