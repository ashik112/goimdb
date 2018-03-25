package reader

import (
	"log"
	"strings"
	"encoding/csv"
	_ "encoding/json"
	"fmt"
	"io"
	_ "io/ioutil"
	"os"
	"strconv"
	_ "strings"
	_ "time"

	"github.com/ashik112/goimdb/gosolr"
	_ "github.com/ashik112/goimdb/model"
)

/*WriteTitleBasics does ...*/
func WriteTitleBasics(csvData [][]string) {
	// start:=time.Now()
	// titleBasics:=make([]model.TitleBasics,len(csvData))
	// for index, each := range csvData {
	// 	titleBasics[index].ID= each[0]
	// 	titleBasics[index].TitleType.Set= each[1]
	// 	titleBasics[index].PrimaryTitle.Set= each[2]
	// 	titleBasics[index].OriginalTitle.Set= each[3]
	// 	isAdult,_:=strconv.ParseInt(each[4],0,64)
	// 	startYear,_:=strconv.ParseInt(each[5],0,64)
	// 	endYear,_:=strconv.ParseInt(each[6],0,64)
	// 	runtimeMinutes,_:=strconv.ParseInt(each[7],0,64)
	// 	titleBasics[index].IsAdult.Set= isAdult
	// 	titleBasics[index].StartYear.Set= startYear
	// 	titleBasics[index].EndYear.Set= endYear
	// 	titleBasics[index].RunTimeMinutes.Set= runtimeMinutes
	// 	titleBasics[index].Genres.Set= strings.Split(each[8],",")
	// }
	// fmt.Println("Processing ", len(csvData)," data took ",time.Since(start))
	// titleBasics=titleBasics[1:]
	// cnv,_:=json.Marshal(titleBasics)
	// err := ioutil.WriteFile("./files/json/title.basics.json", cnv, 0644)
	// if err!=nil{
	// 	fmt.Println(err)
	// }
}

/*WriteRatings does ...*/
func WriteRatings(csvData [][]string) {
	// start:=time.Now()
	// ratings:=make([]model.Ratings,len(csvData))
	// for index, each := range csvData {
	// 	avgRating,_:=strconv.ParseFloat(each[1],64)
	// 	votes,_:=strconv.ParseInt(each[2],0,64)
	// 	ratings[index].ID= each[0]
	// 	ratings[index].AverageRating.Set= avgRating
	// 	ratings[index].NumVotes.Set=votes
	// }
	// fmt.Println("Processing ", len(csvData)," data took ",time.Since(start))
	// ratings=ratings[1:]
	// cnv,_:=json.Marshal(ratings)
	// err := ioutil.WriteFile("./files/json/title.ratings.json", cnv, 0644)
	// if err!=nil{
	// 	fmt.Println(err)
	// }
}

/*WriteCrew does ...*/
func WriteCrew(csvData [][]string) {
	// 	start:=time.Now()
	// 	data:=make([][3]interface{},len(csvData))
	// 	for index, each := range csvData {
	// 	   data[index][0]=each[0]
	// 	   data[index][1]=strings.Split(each[1],",")
	// 	   data[index][2]=strings.Split(each[2],",")
	//    }
	// 	crew:=make([]model.Crew,len(csvData))
	// 	for index, each := range data {
	// 		crew[index].ID= each[0]
	// 		crew[index].Directors.Set= each[1]
	// 		crew[index].Writers.Set=each[2]
	// 	}
	// 	fmt.Println("Processing ", len(csvData)," data took ",time.Since(start))
	// 	crew=crew[1:]
	// 	cnv,_:=json.Marshal(crew)
	// 	err := ioutil.WriteFile("./files/json/title.crew.json", cnv, 0644)
	// 	if err!=nil{
	// 		fmt.Println(err)
	// 	}
}

/*WriteNameBasics does ...*/
func WriteNameBasics(csvData [][]string) {
	// start:=time.Now()
	// data:=make([][6]interface{},len(csvData))
	// count:=0
	// for index, each := range csvData {
	// 	count++
	// 	if count==len(csvData)-2{
	// 		break
	// 	}
	//     data[index][0]=each[0]
	//     data[index][1]=each[1]
	//     birthYear,_:=strconv.ParseInt(each[2],0,64)
	//     deathYear,_:=strconv.ParseInt(each[3],0,64)
	//     data[index][2]=birthYear
	//     data[index][3]=deathYear
	//     data[index][4]=strings.Split(each[4],",")
	//     data[index][5]=strings.Split(each[5],",")
	// }
	// items:=make([]model.NameBasics,len(data))
	// for index, each := range data {
	// 	items[index].ID= each[0]
	// 	items[index].PrimaryName.Set= each[1]
	// 	items[index].BirthYear.Set=each[2]
	// 	items[index].DeathYear.Set=each[3]
	// 	items[index].PrimaryProfession.Set=each[4]
	// 	items[index].KnownForTitles.Set=each[5]
	// }
	// fmt.Println("Processing ", len(csvData)," data took ",time.Since(start))
	// items=items[1:]
	// cnv,_:=json.Marshal(items)
	// err := ioutil.WriteFile("./files/json/name.basics.json", cnv, 0644)
	// if err!=nil{
	// 	fmt.Println(err)
	// }
}

/*WriteTitlePrincipals does ...*/
func WriteTitlePrincipals(csvData [][]string) {
	// start:=time.Now()
	// items:=make([]model.TitlePrincipals,len(csvData))
	// for index, each := range csvData {
	// 	items[index].ID= each[0]
	// 	items[index].PrincipalCast.Set= strings.Split(each[1],",")
	// }
	// items=items[1:]
	// cnv,_:=json.Marshal(items)
	// err := ioutil.WriteFile("./files/json/title.principals.json", cnv, 0644)
	// if err!=nil{
	// 	fmt.Println(err)
	// }
	// fmt.Println("Total Process ", len(csvData)," data took ",time.Since(start))
}

/*InsertTitleRatings does..*/
func InsertTitleRatings(path string) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	var data []interface{}
	csvr := csv.NewReader(f)
	csvr.Comma = '\t'
	csvr.LazyQuotes = true
	csvr.FieldsPerRecord = -1
	headers, err := csvr.Read()
	if err == nil {
		fmt.Println(headers)
	}
	count := 0
	for {
		count++
		row, err := csvr.Read()
		if err == io.EOF {
			if data != nil {
				gosolr.Update(data)
			}
			break
		}
		fmt.Println(row)

		averageRating, err := strconv.ParseFloat(row[1], 64)
		if err != nil {
			averageRating = 0
		}
		numVotes, err := strconv.Atoi(row[2])
		if err != nil {
			numVotes = 0
		}
		item := map[string]interface{}{
			"id": row[0],
			"averageRating": map[string]interface{}{
				"set": averageRating,
			},
			"numVotes": map[string]interface{}{
				"set": numVotes,
			},
		}
		data = append(data, item)
		// fmt.Println(row[0], " ", row[1], " ", row[2], " ", row[3], " ", row[4], " ", startYear, " ", endYear, " ", runtimeMinutes, " ", row[8])
		if count == 1000 {
			if data != nil {
				gosolr.Update(data)
				data = nil
			}
			count = 0
		}
	}
}

/*InsertTitleBasics does..*/
func InsertTitleBasics(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	// var data []interface{}
	csvr := csv.NewReader(f)
	csvr.Comma = '\t'
	// csvr.FieldsPerRecord=9
	// csvr.TrimLeadingSpace=true
	csvr.LazyQuotes = true
	
	headers, err := csvr.Read()
	if err == nil {
		fmt.Println(headers)
	} else{
		log.Fatalln(err)
	}
	count := 0
	for i:=0;i<1064855;i++{
		count++
		row, err := csvr.Read()
		if err!=nil{
			log.Fatalln(err)
		}
		if err == io.EOF {
			// if data != nil {
			// 	gosolr.Update(data)
			// }
			break
		}
		
		// fmt.Println(count+1, " : ", row)

		// startYear, err := strconv.Atoi(row[5])
		// if err != nil {
		// 	startYear = 0
		// }
		// endYear, err := strconv.Atoi(row[6])
		// if err != nil {
		// 	endYear = 0
		// }
		// runtimeMinutes, err := strconv.Atoi(row[7])
		// if err != nil {
		// 	runtimeMinutes = 0
		// }
		// item := []interface{}{
		// 	map[string]interface{}{
		// 		"id": row[0],
		// 		"titleType": map[string]interface{}{
		// 			"set": row[1],
		// 		},
		// 		"primaryTitle": map[string]interface{}{
		// 			"set": row[2],
		// 		},
		// 		"originalTitle": map[string]interface{}{
		// 			"set": row[3],
		// 		},
		// 		"isAdult": map[string]interface{}{
		// 			"set": row[4],
		// 		},
		// 		"startYear": map[string]interface{}{
		// 			"set": startYear,
		// 		},
		// 		"endYear": map[string]interface{}{
		// 			"set": endYear,
		// 		},
		// 		"runtimeMinutes": map[string]interface{}{
		// 			"set": runtimeMinutes,
		// 		},
		// 		"genres": map[string]interface{}{
		// 			"set": row[8],
		// 		},
		// 	},
		// }
		// if row[1] == "tt1300854" {
			// fmt.Printf("%T ",row[0])
			fmt.Println(count,": ",row)
		if strings.EqualFold("tt1300854",row[0]){
			break
		}
		// }

		// gosolr.Update(item)
		// data = append(data, item)
		// fmt.Println(row[0], " ", row[1], " ", row[2], " ", row[3], " ", row[4], " ", startYear, " ", endYear, " ", runtimeMinutes, " ", row[8])
		// if count == 1000 {
		// 	if data != nil {
		// 		gosolr.Update(data)
		// 		data = nil
		// 	}
		// 	count = 0
		// }
	}
}

/*ReadTSV does ...*/
func ReadTSV(directory, target string) {
	InsertTitleBasics(directory + target)
	// InsertTitleRatings(directory + target)
	// start:=time.Now()
	// csvFile, err := os.Open(directory+target)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer csvFile.Close()

	// reader := csv.NewReader(csvFile)
	// reader.Comma = '\t'
	// reader.LazyQuotes = true
	// reader.FieldsPerRecord = -1

	// fmt.Println("Reading Data from file...")
	// csvData, err := reader.ReadAll()
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// elasped:=time.Since(start)

	// switch target {
	// case "name.basics.tsv":
	// 	WriteNameBasics(csvData)
	// case "title.ratings.tsv":
	// 	WriteRatings(csvData)
	// case "title.crew.tsv":
	// 	WriteCrew(csvData)
	// case "title.basics.tsv":
	// 	WriteTitleBasics(csvData)
	// case "title.principals.tsv":
	// 	WriteTitlePrincipals(csvData)
	// default:
	// 	panic("Unexpected error: couldn't locate file")
	// }
	// fmt.Println("Reading data took ",elasped)
	// fmt.Println("Data length: ",len(csvData))
}
