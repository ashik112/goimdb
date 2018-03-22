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

type Titles struct {
	Response struct {
		NumFound int `json:"numFound"`
		Start    int `json:"start"`
		Docs     []struct {
			Tconst         string   `json:"tconst"`
			TitleType      string   `json:"titleType"`
			PrimaryTitle   []string `json:"primaryTitle"`
			OriginalTitle  []string `json:"originalTitle"`
			IsAdult        int      `json:"isAdult"`
			StartYear      string   `json:"startYear"`
			EndYear        string   `json:"endYear"`
			RuntimeMinutes string   `json:"runtimeMinutes"`
			Genres         string   `json:"genres"`
			ID             string   `json:"id"`
			Version        int64    `json:"_version_"`
		} `json:"docs"`
	} `json:"response"`
}

type Ratings struct {
	Response struct {
		NumFound int `json:"numFound"`
		Start    int `json:"start"`
		Docs     []struct {
			Tconst        string  `json:"tconst"`
			AverageRating float64 `json:"averageRating"`
			NumVotes      int     `json:"numVotes"`
			ID            string  `json:"id"`
			Version       int64   `json:"_version_"`
		} `json:"docs"`
	} `json:"response"`
}
type Cast struct {
	Response struct {
		NumFound int `json:"numFound"`
		Start    int `json:"start"`
		Docs     []struct {
			Tconst     string `json:"tconst"`
			Ordering   int    `json:"ordering"`
			Nconst     string `json:"nconst"`
			Category   string `json:"category"`
			Job        string `json:"job"`
			Characters string `json:"characters"`
			ID         string `json:"id"`
			Version    int64  `json:"_version_"`
		} `json:"docs"`
	} `json:"response"`
}
type Crew struct {
	Response struct {
		NumFound int `json:"numFound"`
		Start    int `json:"start"`
		Docs     []struct {
			Tconst    string `json:"tconst"`
			Directors string `json:"directors"`
			Writers   string `json:"writers"`
			ID        string `json:"id"`
			Version   int64  `json:"_version_"`
		} `json:"docs"`
	} `json:"response"`
}

type Person struct {
	Response struct {
		NumFound int `json:"numFound"`
		Start    int `json:"start"`
		Docs     []struct {
			Nconst            string   `json:"nconst"`
			PrimaryName       []string `json:"primaryName"`
			BirthYear         string   `json:"birthYear"`
			DeathYear         string   `json:"deathYear"`
			PrimaryProfession string   `json:"primaryProfession"`
			KnownForTitles    []string `json:"knownForTitles"`
			ID                string   `json:"id"`
			Version           int64    `json:"_version_"`
		} `json:"docs"`
	} `json:"response"`
}