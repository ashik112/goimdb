package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ashik112/goimdb/decompresser"
	"github.com/ashik112/goimdb/downloader"
	"github.com/ashik112/goimdb/gosolr"
	"github.com/ashik112/goimdb/model"
	_"github.com/ashik112/goimdb/reader"
)

var FilePath = "./files/"
var ArchivePath = FilePath + "archive/"
var DecompressedPath = FilePath + "decompressed/"
var JsonPath = FilePath + "json/"
var GzipFile = model.Files{"title.basics.tsv.gz", "title.ratings.tsv.gz", "title.principals.tsv.gz", "name.basics.tsv.gz", "title.crew.tsv.gz", "title.episode.tsv.gz"}
var TsvFile = model.Files{"title.basics.tsv", "title.ratings.tsv", "title.principals.tsv", "name.basics.tsv", "title.crew.tsv", "title.episode.tsv"}
var SolrConfig = model.Solr{"localhost", 8983, "imdb"}
var Imdb = model.Imdb{"https://datasets.imdbws.com/"}

/*DownloadFiles does..*/
func DownloadFiles() {
	downloader.Download(ArchivePath, Imdb.URL+GzipFile.Title)
	downloader.Download(ArchivePath, Imdb.URL+GzipFile.Ratings)
	downloader.Download(ArchivePath, Imdb.URL+GzipFile.Persons)
	downloader.Download(ArchivePath, Imdb.URL+GzipFile.Crew)
	downloader.Download(ArchivePath, Imdb.URL+GzipFile.People)
	downloader.Download(ArchivePath, Imdb.URL+GzipFile.Episode)
}

/*GetFiles does ...*/
func GetFiles() {
	startDecompress := time.Now()
	doneRatings := make(chan int)
	donePrincipals := make(chan int)
	doneEpisode := make(chan int)
	doneCrew := make(chan int)
	doneTitleBasics := make(chan int)
	doneNameBasics := make(chan int)
	go decompresser.UnGzip(ArchivePath+GzipFile.Title, DecompressedPath+GzipFile.Title, doneTitleBasics)
	go decompresser.UnGzip(ArchivePath+GzipFile.Ratings, DecompressedPath+GzipFile.Ratings, doneRatings)
	go decompresser.UnGzip(ArchivePath+GzipFile.People, DecompressedPath+GzipFile.People, donePrincipals)
	go decompresser.UnGzip(ArchivePath+GzipFile.Persons, DecompressedPath+GzipFile.Persons, doneNameBasics)
	go decompresser.UnGzip(ArchivePath+GzipFile.Crew, DecompressedPath+GzipFile.Crew, doneCrew)
	go decompresser.UnGzip(ArchivePath+GzipFile.Episode, DecompressedPath+GzipFile.Episode, doneEpisode)
	<-doneRatings
	<-donePrincipals
	<-doneEpisode
	<-doneCrew
	<-doneTitleBasics
	<-doneNameBasics
	elsapsedDecompress := time.Since(startDecompress)
	fmt.Println("Decompression Process took ", elsapsedDecompress)
}

func CreateSolrFields() {
	doneTitles := make(chan bool)
	go gosolr.CreateSolrFields(SolrConfig.Hostname, SolrConfig.Port, SolrConfig.Core, JsonPath+"all_fields.json", doneTitles)
	<-doneTitles
}
func UploadSolrData() {
	start := time.Now()
	donePrincipals := make(chan bool)
	go gosolr.UploadDoc(SolrConfig.Hostname, SolrConfig.Port, SolrConfig.Core, DecompressedPath+TsvFile.People, donePrincipals)
	<-donePrincipals
	fmt.Println("Uploading Principals took ", time.Since(start))

	donePersons := make(chan bool)
	go gosolr.UploadDoc(SolrConfig.Hostname, SolrConfig.Port, SolrConfig.Core, DecompressedPath+TsvFile.Persons, donePersons)
	<-donePersons

	doneRatings := make(chan bool)
	go gosolr.UploadDoc(SolrConfig.Hostname, SolrConfig.Port, SolrConfig.Core, DecompressedPath+TsvFile.Ratings, doneRatings)
	doneCrew := make(chan bool)
	go gosolr.UploadDoc(SolrConfig.Hostname, SolrConfig.Port, SolrConfig.Core, DecompressedPath+TsvFile.Crew, doneCrew)
	<-doneRatings
	<-doneCrew

	doneTitles := make(chan bool)
	go gosolr.UploadDoc(SolrConfig.Hostname, SolrConfig.Port, SolrConfig.Core, DecompressedPath+TsvFile.Title, doneTitles)
	doneEpisodes := make(chan bool)
	go gosolr.UploadDoc(SolrConfig.Hostname, SolrConfig.Port, SolrConfig.Core, DecompressedPath+TsvFile.Episode, doneEpisodes)
	<-doneTitles
	<-doneEpisodes
}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func CMD(args []string) {
	switch {
	case contains(os.Args[1:], "update"):
		gosolr.DeleteAll(SolrConfig.Hostname, SolrConfig.Port, SolrConfig.Core)
		UploadSolrData()
	case contains(os.Args[1:], "insert"):
		UploadSolrData()
	case contains(os.Args[1:], "init"):
		DownloadFiles()
		GetFiles()
		CreateSolrFields()
		UploadSolrData()
	default:
		fmt.Println("No valid param found")
	}
}

func SearchMovie() {
	start := time.Now()
	fmt.Print("Enter Movie title: ")
	reader := bufio.NewReader(os.Stdin)
	title, _ := reader.ReadString('\n')
	title = strings.Trim(title, "\n")
	title = `"` + title + `"`
	titleType := `"` + "movie" + `"`
	q := "primaryTitle:" + title + "AND titleType:" + titleType
	t := &url.URL{Fragment: q}
	q = strings.Trim(t.String(), "#")
	url := "http://" + SolrConfig.Hostname + ":" + strconv.Itoa(SolrConfig.Port) + "/solr/" + SolrConfig.Core + "/select?q=" + q + "&rows=1000&sort=numVotes%20desc"
	fmt.Println(url)
	gosolr.GetTitle(url)
	fmt.Println("... took ", time.Since(start))
}

func main() {
	// gosolr.DeleteAll("localhost", 8983, "imdb")
	// CreateSolrFields()
	// DownloadFiles()
	// GetFiles()
	// reader.ReadTSV("./files/decompressed/", "title.basics.tsv")
	// reader.ReadTSV("./files/decompressed/", "title.ratings.tsv")
	SearchMovie()
}
