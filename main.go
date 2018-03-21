package main

import (
	"fmt"
	"os"
	"time"
    "strconv"
    "net/url"
    "strings"
    "bufio"
	"github.com/ashik112/goimdb/decompresser"
	"github.com/ashik112/goimdb/downloader"
	"github.com/ashik112/goimdb/gosolr"
)

/*DownloadFiles does..*/
func DownloadFiles() {
	downloader.Download("./files/archive", "https://datasets.imdbws.com/title.ratings.tsv.gz")
	downloader.Download("./files/archive", "https://datasets.imdbws.com/title.principals.tsv.gz")
	downloader.Download("./files/archive", "https://datasets.imdbws.com/title.episode.tsv.gz")
	downloader.Download("./files/archive", "https://datasets.imdbws.com/title.crew.tsv.gz")
	downloader.Download("./files/archive", "https://datasets.imdbws.com/title.basics.tsv.gz")
	downloader.Download("./files/archive", "https://datasets.imdbws.com/title.akas.tsv.gz")
	downloader.Download("./files/archive", "https://datasets.imdbws.com/name.basics.tsv.gz")
}

/*GetFiles does ...*/
func GetFiles() {
	startDecompress := time.Now()

	doneRatings := make(chan int)
	donePrincipals := make(chan int)
	doneEpisode := make(chan int)
	doneCrew := make(chan int)
	doneTitleBasics := make(chan int)
	doneAkas := make(chan int)
	doneNameBasics := make(chan int)
	go decompresser.UnGzip("./files/archive/title.ratings.tsv.gz", "./files/decompressed", doneRatings)
	go decompresser.UnGzip("./files/archive/title.principals.tsv.gz", "./files/decompressed", donePrincipals)
	go decompresser.UnGzip("./files/archive/title.episode.tsv.gz", "./files/decompressed", doneEpisode)
	go decompresser.UnGzip("./files/archive/title.crew.tsv.gz", "./files/decompressed", doneCrew)
	go decompresser.UnGzip("./files/archive/title.basics.tsv.gz", "./files/decompressed", doneTitleBasics)
	go decompresser.UnGzip("./files/archive/title.akas.tsv.gz", "./files/decompressed", doneAkas)
	go decompresser.UnGzip("./files/archive/name.basics.tsv.gz", "./files/decompressed", doneNameBasics)
	<-doneRatings
	<-donePrincipals
	<-doneEpisode
	<-doneCrew
	<-doneTitleBasics
	<-doneAkas
	<-doneNameBasics
	elsapsedDecompress := time.Since(startDecompress)
	fmt.Println("Decompression Process took ", elsapsedDecompress)
}

/*ReadFile does..*/
func ReadFile() {
	// directory:="./files/decompressed/"
	// reader.ReadTSV(directory,"title.crew.tsv")
	// reader.ReadTSV(directory,"title.basics.tsv")
	// reader.ReadTSV(directory,"title.crew.tsv")
	// reader.ReadTSV(directory,"name.basics.tsv")
	// reader.ReadTSV(directory,"title.principals.tsv")

}
func CreateSolrFields() {
	directory := "./files/json/"
	doneTitles := make(chan bool)
	go gosolr.CreateSolrFields("localhost", 8983, "imdb", directory+"all_fields.json", doneTitles)
	// donePersons:=make(chan bool)
	// go gosolr.CreateSolrFields("localhost",8983,"persons",directory+"field_persons.json",donePersons)
	<-doneTitles
	// <-donePersons
}
func UploadSolrData() {
	directory := "./files/decompressed/"

	start := time.Now()
	donePrincipals := make(chan bool)
	go gosolr.UploadDoc("localhost", 8983, "imdb", directory+"title.principals.tsv", donePrincipals)
	<-donePrincipals
	fmt.Println("Uploading Principals took ", time.Since(start))

	donePersons := make(chan bool)
	go gosolr.UploadDoc("localhost", 8983, "imdb", directory+"name.basics.tsv", donePersons)
	<-donePersons

	doneRatings := make(chan bool)
	go gosolr.UploadDoc("localhost", 8983, "imdb", directory+"title.ratings.tsv", doneRatings)
	doneCrew := make(chan bool)
	go gosolr.UploadDoc("localhost", 8983, "imdb", directory+"title.crew.tsv", doneCrew)
	<-doneRatings
	<-doneCrew

	doneTitles := make(chan bool)
	go gosolr.UploadDoc("localhost", 8983, "imdb", directory+"title.basics.tsv", doneTitles)
	doneEpisodes := make(chan bool)
	go gosolr.UploadDoc("localhost", 8983, "imdb", directory+"title.episode.tsv", doneEpisodes)
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
		gosolr.DeleteAll("localhost", 8983, "imdb")
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

func main() {
    start := time.Now()
	// DownloadFiles()
	// GetFiles()
    // ReadFile()
    // var title string
    fmt.Print("Enter Movie title: ")

    // fmt.Scan(&title)
    reader := bufio.NewReader(os.Stdin)
    title, _ := reader.ReadString('\n')
    fmt.Println(title)



    title=`"`+title+`"`
    titleType:=`"`+"movie"+`"`
    q:="primaryTitle:"+title+"AND titleType:"+titleType
    t := &url.URL{Fragment: q}
	q = strings.Trim(t.String(),"#")
	// gosolr.DeleteAll("localhost", 8983, "imdb")
	// CreateSolrFields()
    // UploadSolrData()
    url:="http://"+"localhost"+":"+strconv.Itoa(8983)+"/solr/"+"imdb"+"/select?q="+q
    gosolr.Get(url)
	fmt.Println("... took ", time.Since(start))

}
