package main

import (
	"fmt"
    "time"
    "github.com/ashik112/goimdb/downloader"
    "github.com/ashik112/goimdb/decompresser"
    _"github.com/ashik112/goimdb/reader"
    "github.com/ashik112/goimdb/httpreq"
)

/*DownloadFiles does..*/
func DownloadFiles(){
    downloader.Download("./files/archive","https://datasets.imdbws.com/title.ratings.tsv.gz")
    downloader.Download("./files/archive","https://datasets.imdbws.com/title.principals.tsv.gz")
    downloader.Download("./files/archive","https://datasets.imdbws.com/title.episode.tsv.gz")
    downloader.Download("./files/archive","https://datasets.imdbws.com/title.crew.tsv.gz")
    downloader.Download("./files/archive","https://datasets.imdbws.com/title.basics.tsv.gz")
    downloader.Download("./files/archive","https://datasets.imdbws.com/title.akas.tsv.gz")
    downloader.Download("./files/archive","https://datasets.imdbws.com/name.basics.tsv.gz")
}

/*GetFiles does ...*/
func GetFiles(){
    startDecompress:=time.Now()

    doneRatings := make(chan int)
    donePrincipals := make(chan int)
    doneEpisode := make(chan int)
    doneCrew:=make(chan int)
    doneTitleBasics:=make(chan int)
    doneAkas:=make(chan int)
    doneNameBasics:=make(chan int)
    go decompresser.UnGzip("./files/archive/title.ratings.tsv.gz","./files/decompressed",doneRatings)
    go decompresser.UnGzip("./files/archive/title.principals.tsv.gz","./files/decompressed",donePrincipals)
    go decompresser.UnGzip("./files/archive/title.episode.tsv.gz","./files/decompressed",doneEpisode)
    go decompresser.UnGzip("./files/archive/title.crew.tsv.gz","./files/decompressed",doneCrew)
    go decompresser.UnGzip("./files/archive/title.basics.tsv.gz","./files/decompressed",doneTitleBasics)
    go decompresser.UnGzip("./files/archive/title.akas.tsv.gz","./files/decompressed",doneAkas)
    go decompresser.UnGzip("./files/archive/name.basics.tsv.gz","./files/decompressed",doneNameBasics)
    <- doneRatings
    <-donePrincipals
    <-doneEpisode
    <-doneCrew
    <-doneTitleBasics
    <-doneAkas
    <-doneNameBasics
    elsapsedDecompress:=time.Since(startDecompress)
    fmt.Println("Decompression Process took ",elsapsedDecompress);
}

/*ReadFile does..*/
func ReadFile(){
    // directory:="./files/decompressed/"
    // reader.ReadTSV(directory,"title.crew.tsv")
    // reader.ReadTSV(directory,"title.basics.tsv")
    // reader.ReadTSV(directory,"title.crew.tsv")
    // reader.ReadTSV(directory,"name.basics.tsv")
    // reader.ReadTSV(directory,"title.principals.tsv")

}
func CreateSolrFields(){
    directory:="./files/json/"
    doneTitles:=make(chan bool)
    go httpreq.CreateSolrFields("localhost",8983,"test",directory+"field_titles.json",doneTitles)
    // donePersons:=make(chan bool)
    // go httpreq.CreateSolrFields("localhost",8983,"persons",directory+"field_persons.json",donePersons)
    <-doneTitles
    // <-donePersons
}
func UploadSolrData(){
    directory:="./files/decompressed/"
    doneRatings:=make(chan bool)
    go httpreq.UploadDoc("localhost",8983,"test",directory+"title.ratings.tsv",doneRatings)
    // donePersons:=make(chan bool)   
    // go httpreq.UploadDoc("localhost",8983,"persons",directory+"name.basics.tsv",donePersons)
    // doneTitles:=make(chan bool)
    // go httpreq.UploadDoc("localhost",8983,"test",directory+"title.basics.tsv",doneTitles)
    <- doneRatings
    // <- donePersons
    // <- doneTitles
}
func main() {
    // DownloadFiles()
    // GetFiles()
    // ReadFile()    
    // CreateSolrFields()
    UploadSolrData()
   
}
