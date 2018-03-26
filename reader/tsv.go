package reader

import (
	"bufio"
	_ "encoding/csv"
	_ "encoding/json"
	"fmt"
	_ "io"
	_ "io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	_ "strings"
	"time"
	_ "time"

	"github.com/ashik112/goimdb/gosolr"
	_ "github.com/ashik112/goimdb/model"
)

/*InsertTitleRatings does..*/
func InsertTitleRatings(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), "\t")
		averageRating, err := strconv.ParseFloat(row[1], 64)
		if err != nil {
			averageRating = 0
		}
		numVotes, err := strconv.Atoi(row[2])
		if err != nil {
			numVotes = 0
		}
		item := []interface{}{
			map[string]interface{}{
				"id": row[0],
				"averageRating": map[string]interface{}{
					"set": averageRating,
				},
				"numVotes": map[string]interface{}{
					"set": numVotes,
				},
			},
		}
		gosolr.Update(item, "localhost", 8983, "imdb")
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

/*InsertTitleBasics does..*/
func InsertTitleBasics(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), "\t")
		startYear, err := strconv.Atoi(row[5])
		if err != nil {
			startYear = 0
		}
		endYear, err := strconv.Atoi(row[6])
		if err != nil {
			endYear = 0
		}
		runtimeMinutes, err := strconv.Atoi(row[7])
		if err != nil {
			runtimeMinutes = 0
		}
		item := []interface{}{
			map[string]interface{}{
				"id": row[0],
				"titleType": map[string]interface{}{
					"set": row[1],
				},
				"primaryTitle": map[string]interface{}{
					"set": row[2],
				},
				"originalTitle": map[string]interface{}{
					"set": row[3],
				},
				"isAdult": map[string]interface{}{
					"set": row[4],
				},
				"startYear": map[string]interface{}{
					"set": startYear,
				},
				"endYear": map[string]interface{}{
					"set": endYear,
				},
				"runtimeMinutes": map[string]interface{}{
					"set": runtimeMinutes,
				},
				"genres": map[string]interface{}{
					"set": row[8],
				},
			},
		}
		gosolr.Update(item, "localhost", 8983, "imdb")
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

/*InsertTitleRatings does..*/
func InsertTitlePrincipals(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), "\t")
		ordering, err := strconv.Atoi(row[1])
		if err != nil {
			ordering = 0
		}
		item := []interface{}{
			map[string]interface{}{
				"tconst":     row[0],
				"ordering":   ordering,
				"nconst":     row[2],
				"category":   row[3],
				"characters": row[5],
			},
		}
		gosolr.Update(item, "localhost", 8983, "cast")
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

/*InsertNameBasics does..*/
func InsertNameBasics(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), "\t")
		birthYear, err := strconv.Atoi(row[2])
		if err != nil {
			birthYear = 0
		}
		deathYear, err := strconv.Atoi(row[3])
		if err != nil {
			deathYear = 0
		}
		item := []interface{}{
			map[string]interface{}{
				"id": row[0],
				"primaryName": map[string]interface{}{
					"set": row[1],
				},
				"birthYear": map[string]interface{}{
					"set": birthYear,
				},
				"deathYear": map[string]interface{}{
					"set": deathYear,
				},
				"primaryProfession": map[string]interface{}{
					"set": row[4],
				},
				"knownForTitles": map[string]interface{}{
					"set": row[5],
				},
			},
		}
		gosolr.Update(item, "localhost", 8983, "person")
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

/*ReadTSV does ...*/
func ReadTSV(directory, target string) {
	start := time.Now()
	// InsertTitleBasics(directory + target)
	// InsertTitleRatings(directory + target)
	// InsertTitlePrincipals(directory + target)
	InsertNameBasics(directory + target)
	fmt.Println("... took ", time.Since(start))
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
