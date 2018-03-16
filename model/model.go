package model
type AverageRating struct{
	Set float64 `json:"set"`
}
type NumVotes struct{
	Set int64 `json:"set"`
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
	Set string `json:"set"`
}
type Ratings struct{
    ID string  `json:"id"`
    AverageRating AverageRating  `json:"averageRating"`
	NumVotes NumVotes `json:"numVotes"`
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