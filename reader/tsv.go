package reader

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
	"strconv"
	"encoding/json"
	_"encoding/xml"
	"io/ioutil"
	"github.com/ashik112/goimdb/model"
)

/*WriteTitleBasics does ...*/
func WriteTitleBasics(csvData [][]string){
	start:=time.Now()

	fmt.Println(len(csvData))
	titleBasics:=make([]model.TitleBasics,len(csvData))

	for index, each := range csvData {
		titleBasics[index].ID= each[0]
		titleBasics[index].TitleType.Set= each[1]
		titleBasics[index].PrimaryTitle.Set= each[2]
		titleBasics[index].OriginalTitle.Set= each[3]
		isAdult,_:=strconv.ParseInt(each[4],0,64)
		startYear,_:=strconv.ParseInt(each[5],0,64)
		endYear,_:=strconv.ParseInt(each[6],0,64)
		runtimeMinutes,_:=strconv.ParseInt(each[7],0,64)
		titleBasics[index].IsAdult.Set= isAdult
		titleBasics[index].StartYear.Set= startYear
		titleBasics[index].EndYear.Set= endYear
		titleBasics[index].RunTimeMinutes.Set= runtimeMinutes
		titleBasics[index].Genres.Set= each[8]
	}
	fmt.Println("Processing ", len(csvData)," data took ",time.Since(start))
	titleBasics=titleBasics[1:]
	cnv,_:=json.Marshal(titleBasics)
	err := ioutil.WriteFile("./files/decompressed/title.basics.json", cnv, 0644)
	if err!=nil{
		fmt.Println(err)
	}
}
/*WriteRatings does ...*/
func WriteRatings(csvData [][]string){
	start:=time.Now()

	//var ratings []Ratings
	ratings:=make([]model.Ratings,len(csvData))
	//count:=0
	for index, each := range csvData {
		
		avgRating,_:=strconv.ParseFloat(each[1],64)
		votes,_:=strconv.ParseInt(each[2],0,64)
		ratings[index].ID= each[0]
		ratings[index].AverageRating.Set= avgRating		
		ratings[index].NumVotes.Set=votes
	}
	fmt.Println("Processing ", len(csvData)," data took ",time.Since(start))
	ratings=ratings[1:]
	cnv,_:=json.Marshal(ratings)
	err := ioutil.WriteFile("./files/decompressed/title.ratings.json", cnv, 0644)
	if err!=nil{
		fmt.Println(err)
	}
}

/*ReadTSV does ...*/
func ReadTSV(target string) {
	start:=time.Now()
	csvFile, err := os.Open(target)

	if err != nil {
		fmt.Println(err)
	}

	defer csvFile.Close()

	reader := csv.NewReader(csvFile)

	reader.Comma = '\t' // Use tab-delimited instead of comma <---- here!
	reader.LazyQuotes = true
	 reader.FieldsPerRecord = -1

	fmt.Println("Reading Data from file...")
	csvData, err := reader.ReadAll()	
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	elasped:=time.Since(start)

	//WriteRatings(csvData)
	WriteTitleBasics(csvData)

	fmt.Println("Reading data took ",elasped)
	fmt.Println("Data length: ",len(csvData))
	
}
