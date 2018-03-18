package reader

import (
	"bufio"
	_ "encoding/csv"
	"encoding/json"
	"fmt"
	_ "io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ashik112/goimdb/httpreq"
	"github.com/ashik112/goimdb/model"
	"github.com/rtt/Go-Solr"
)

/*WriteTitleBasics does ...*/
func WriteTitleBasics(csvData [][]string) {
	start := time.Now()
	titleBasics := make([]model.TitleBasics, len(csvData))
	for index, each := range csvData {
		titleBasics[index].ID = each[0]
		titleBasics[index].TitleType.Set = each[1]
		titleBasics[index].PrimaryTitle.Set = each[2]
		titleBasics[index].OriginalTitle.Set = each[3]
		isAdult, _ := strconv.ParseInt(each[4], 0, 64)
		startYear, _ := strconv.ParseInt(each[5], 0, 64)
		endYear, _ := strconv.ParseInt(each[6], 0, 64)
		runtimeMinutes, _ := strconv.ParseInt(each[7], 0, 64)
		titleBasics[index].IsAdult.Set = isAdult
		titleBasics[index].StartYear.Set = startYear
		titleBasics[index].EndYear.Set = endYear
		titleBasics[index].RunTimeMinutes.Set = runtimeMinutes
		titleBasics[index].Genres.Set = strings.Split(each[8], ",")
	}
	fmt.Println("Processing ", len(csvData), " data took ", time.Since(start))
	titleBasics = titleBasics[1:]
	cnv, _ := json.Marshal(titleBasics)
	err := ioutil.WriteFile("./files/json/title.basics.json", cnv, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

/*WriteRatings does ...*/
func WriteRatings(csvData [][]string) {
	start := time.Now()
	ratings := make([]model.Ratings, len(csvData))
	for index, each := range csvData {
		avgRating, _ := strconv.ParseFloat(each[1], 64)
		votes, _ := strconv.ParseInt(each[2], 0, 64)
		ratings[index].ID = each[0]
		ratings[index].AverageRating.Set = avgRating
		ratings[index].NumVotes.Set = votes
	}
	fmt.Println("Processing ", len(csvData), " data took ", time.Since(start))
	ratings = ratings[1:]
	cnv, _ := json.Marshal(ratings)
	err := ioutil.WriteFile("./files/json/title.ratings.json", cnv, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

/*WriteCrew does ...*/
func WriteCrew(csvData [][]string) {
	start := time.Now()
	data := make([][3]interface{}, len(csvData))
	for index, each := range csvData {
		data[index][0] = each[0]
		data[index][1] = strings.Split(each[1], ",")
		data[index][2] = strings.Split(each[2], ",")
	}
	crew := make([]model.Crew, len(csvData))
	for index, each := range data {
		crew[index].ID = each[0]
		crew[index].Directors.Set = each[1]
		crew[index].Writers.Set = each[2]
	}
	fmt.Println("Processing ", len(csvData), " data took ", time.Since(start))
	crew = crew[1:]
	cnv, _ := json.Marshal(crew)
	err := ioutil.WriteFile("./files/json/title.crew.json", cnv, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

/*WriteNameBasics does ...*/
func WriteNameBasics(csvData [][]string) {
	start := time.Now()
	data := make([][6]interface{}, len(csvData))
	count := 0
	for index, each := range csvData {
		count++
		if count == len(csvData)-2 {
			break
		}
		data[index][0] = each[0]
		data[index][1] = each[1]
		birthYear, _ := strconv.ParseInt(each[2], 0, 64)
		deathYear, _ := strconv.ParseInt(each[3], 0, 64)
		data[index][2] = birthYear
		data[index][3] = deathYear
		data[index][4] = strings.Split(each[4], ",")
		data[index][5] = strings.Split(each[5], ",")
	}
	items := make([]model.NameBasics, len(data))
	for index, each := range data {
		items[index].ID = each[0]
		items[index].PrimaryName.Set = each[1]
		items[index].BirthYear.Set = each[2]
		items[index].DeathYear.Set = each[3]
		items[index].PrimaryProfession.Set = each[4]
		items[index].KnownForTitles.Set = each[5]
	}
	fmt.Println("Processing ", len(csvData), " data took ", time.Since(start))
	items = items[1:]
	cnv, _ := json.Marshal(items)
	err := ioutil.WriteFile("./files/json/name.basics.json", cnv, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

/*WriteTitlePrincipals does ...*/
// func WriteTitlePrincipals(csvData [][]string) {
// 	start := time.Now()
// 	items := make([]model.TitlePrincipals, len(csvData))
// 	for index, each := range csvData {
// 		items[index].ID = each[0]
// 		items[index].PrincipalCast.Set = strings.Split(each[1], ",")
// 	}
// 	items = items[1:]
// 	cnv, _ := json.Marshal(items)
// 	err := ioutil.WriteFile("./files/json/title.principals.json", cnv, 0644)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println("Total Process ", len(csvData), " data took ", time.Since(start))
// }
func Readln(r *bufio.Reader) (string, error) {
	var (
		isPrefix = true
		err      error
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
}

func ShowProgress(done chan bool) {
	stop := false
	for {
		select {
		case <-done:
			stop = true
		default:
			fmt.Print("=")
			time.Sleep(5*time.Second)
		}
		if(stop==true){
			break
		}
	}
}

/*PrintData does..*/
func PrintData(path string) {
	start := time.Now()
	httpreq.DeleteAll("localhost", 8983, "people")
	f, err := os.Open(path)
	if err != nil {
	}
	defer f.Close()
	r := bufio.NewReader(f)
	solrConn, err := solr.Init("localhost", 8983, "people")
	if err != nil {
		fmt.Println(err)
		return
	}
	line, _ := Readln(r)
	line, e := Readln(r)
	done := make(chan bool)
	for e == nil {
		go ShowProgress(done)
		row := strings.Split(line, "\t")
		data := map[string]interface{}{
			"add": []interface{}{
				map[string]interface{}{
					"title":      row[0],
					"ordering":   map[string]interface{}{"set": row[1]},
					"person":     map[string]interface{}{"set": row[2]},
					"category":   map[string]interface{}{"set": row[3]},
					"job":        map[string]interface{}{"set": row[4]},
					"characters": map[string]interface{}{"set": row[5]},
				},
			},
		}
		solrConn.Update(data, false)
		// if err != nil {
		// 	 fmt.Println("error =>", err)
		// } else {
		// 	fmt.Println("resp =>", resp)
		// }
		line, e = Readln(r)
	}
	done <- true
	fmt.Println("Finished for loop")
	fmt.Println("Total Process took", time.Since(start))
}

/*ReadTSV does ...*/
func ReadTSV(directory, target string) {
	PrintData(directory + target)

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
