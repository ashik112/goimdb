package gosolr

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	_ "net/url"
	"strconv"
	_ "strings"

	"github.com/ashik112/goimdb/model"
)

var SolrConfig = model.Solr{"localhost", 8983, "imdb", "cast"}

func DeleteAll(hostname string, port int, core string) {
	url := "http://" + hostname + ":" + strconv.Itoa(port) + "/solr/" + core + "/update?commit=true"
	fmt.Println("URL:>", url)
	var xmlStr = []byte(`<delete><query>*:*</query></delete>`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(xmlStr))
	req.Header.Set("Content-Type", "text/xml")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

func CreateSolrFields(hostname string, port int, core string, path string, done chan bool) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Creating fields from...")
	url := "http://" + hostname + ":" + strconv.Itoa(port) + "/solr/" + core + "/schema"
	fmt.Println("URL:>", url)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	done <- true
}

func Update(data []interface{}, hostname string, port int, core string) {
	url := "http://" + hostname + ":" + strconv.Itoa(port) + "/solr/" + core + "/update?autoCommit=true&commitWithin=600000"
	jsonStr, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

func UploadDoc(hostname string, port int, core string, path string, done chan bool) {
	// DeleteAll(hostname,port,core)

	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Uploading file to solr...")
	url := "http://" + hostname + ":" + strconv.Itoa(port) + "/solr/" + core + "/update?commit=true&separator=%09&escape=%5c&trim=true&commitWithin=120000"
	fmt.Println("URL:>", url)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/csv")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	done <- true
}
func GetTitle(Url string) {
	resp, err := http.Get(Url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// fmt.Println("response Status:", resp.Status)
	// fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	var data model.Titles
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}

	fmt.Println("\nSearch Results: (Found ", data.Response.NumFound, ")\n")
	// fmt.Println("=============> Basic Info <=============")
	for i, item := range data.Response.Docs {

		fmt.Print("Sl: ", i+1, "\t||\t imdbID: ", item.ID, "\t||\t Title: ", item.PrimaryTitle[0], "\t||\t Type: ", item.TitleType, "\t||\t Year: ", item.StartYear, "\t||\t Genres: ", item.Genres, "\n\nRuntime: ", item.RuntimeMinutes, " minutes \t||\t")
		fmt.Println("Rating: ", item.AverageRating, "\t||\t Votes: ", item.NumVotes)
		getCast := make(chan bool)
		fmt.Println("\n=============> Casts <=============\n")
		go GetCast(item.ID, getCast)
		<-getCast
		fmt.Println("\n===============================================================================================================================================================================================\n")
	}
}

func GetCast(tconst string, done chan bool) {
	q := "tconst:" + `%22` + tconst + `%22` + "%20AND%20category:*"
	resp, err := http.Get("http://" + SolrConfig.Hostname + ":" + strconv.Itoa(SolrConfig.Port) + "/solr/" + SolrConfig.CastCore + "/select?q=" + q + "&sort=ordering%20asc")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var data model.Cast
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}
	for _, item := range data.Response.Docs {
		// fmt.Println("--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------")
		fmt.Println(GetCastName(item.Nconst), "\t||\t ", item.Category, "\t||\t ", item.Characters)
		// fmt.Println("")
	}
	done <- true
}

func GetCastName(id string) string {
	q := "id:" + `%22` + id + `%22` //+ "%20AND%20primaryName:*"

	resp, err := http.Get("http://" + SolrConfig.Hostname + ":" + strconv.Itoa(SolrConfig.Port) + "/solr/" + "person" + "/select?q=" + q) //+ "&fl=primaryName")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var data model.Person
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}
	return data.Response.Docs[0].PrimaryName[0]
}
