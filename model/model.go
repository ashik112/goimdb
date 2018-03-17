package model
type AverageRating struct{
	Set float64 `json:"zset"`
}
type NumVotes struct{
	Set int64 `json:"set"`
}
type Ratings struct{
    ID string  `json:"id"`
    AverageRating AverageRating  `json:"averageRating"`
	NumVotes NumVotes `json:"numVotes"`
}
type TitleType struct{
	Set string `json:"set"`
}
type PrimaryTitle struct{
	Set string `json:"set"`
}
type OriginalTitle struct{
	Set string `json:"set"`
}
type IsAdult struct{
	Set int64 `json:"set"`
}
type StartYear struct{
	Set int64 `json:"set"`
}
type EndYear struct{
	Set int64 `json:"set"`
}
type RunTimeMinutes struct{
	Set int64 `json:"set"`
}
type Genres struct{
	Set interface{} `json:"set"`
}
/*TitleBasics does..*/
type TitleBasics struct{
    ID string  `json:"id"`
	TitleType TitleType `json:"titleType"`
	PrimaryTitle PrimaryTitle `json:"primaryTitle"`
	OriginalTitle OriginalTitle `json:"originalTitle"`
	IsAdult IsAdult `json:"isAdult"`
	StartYear StartYear `json:"startYear"`
	EndYear EndYear `json:"endYear"`
	RunTimeMinutes RunTimeMinutes `json:"runtimeMinutes"`
	Genres Genres `json:"genres"`
}
type Crew struct{
    ID interface{}  `json:"id"`
    Directors Directors  `json:"directors"`
	Writers Writers `json:"writers"`
}
type Directors struct{
	Set interface{} `json:"set"`
}
type Writers struct{
	Set interface{} `json:"set"`
}
type TitleEpisode struct{
	ID string `json:"id"`
    Parent Parent  `json:"parentID"`
	SeasonNumber SeasonNumber `json:"seasonNumber"`
	EpisodeNumber EpisodeNumber `json:"episodeNumber"`
}
type Parent struct{
	Set string `json:"set"`
}
type SeasonNumber struct{
	Set string `json:"set"`
}
type EpisodeNumber struct{
	Set string `json:"set"`
}

type TitlePrincipals struct{
	ID interface{}  `json:"id"`
    PrincipalCast PrincipalCast  `json:"principalCast"`
}
type PrincipalCast struct{
	Set interface{} `json:"set"`
}

type NameBasics struct{
	ID interface{}  `json:"id"`
    PrimaryName PrimaryName  `json:"primaryName"`
	BirthYear BirthYear `json:"birthYear"`
	DeathYear DeathYear `json:"deathYear"`
	PrimaryProfession PrimaryProfession `json:"primaryProfession"`
	KnownForTitles KnownForTitles `json:"knownForTitles"`
}
type PrimaryName struct{
	Set interface{} `json:"set"`
}
type BirthYear struct{
	Set interface{} `json:"set"`
}
type DeathYear struct{
	Set interface{} `json:"set"`
}
type PrimaryProfession struct{
	Set interface{} `json:"set"`
}
type KnownForTitles struct{
	Set interface{} `json:"set"`
}