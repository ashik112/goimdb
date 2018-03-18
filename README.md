# goimdb
IMDB movie database using Golang (only for personal and non-commercial use)

goimdb downloads data from datasets.imdbws.com, currently provided by imdb. These are the subsets of IMDb data that are available for personal and non-commercial use. Each dataset is gzipped, TSV (tab-separated-values) formatted file. The first line of each file contains headers.

goimdb downloads, extracts and converts these file to Solr compatible JSON format to be able to insert them into a Solr database. Solr is capable of conducting full-text search and provides apis for searching. The current latest version of Solr, 7.2.1 is tested with goimdb.
### Tech
goimdb uses a number of open source projects to work properly:

* [Golang] 
* [Solr] 

### Installation

goimdb is go-gettable. The follwing command will download the full project:

```sh
$ go get github.com/ashik112/goimb
```

### Solr
Create core:
```sh
$ solr create -c imdb
```
Command for adding required fields:
```sh
$ curl -X POST -H 'Content-type:application/json' --data-binary '{
  "add-field": {
    "name": "averageRating",
    "type": "pfloat",
    "stored": true,
    "indexed": true,
    "docValues": true
  },
  "add-field": {
    "name": "numVotes",
    "type": "pint",
    "stored": true,
    "indexed": true,
    "docValues": true
  },
  "add-field": {
    "name": "primaryName",
    "type": "text_general",
    "stored": true,
    "indexed": true
  },
  "add-field": {
    "name": "birthYear",
    "type": "pint",
    "stored": true,
    "indexed": true,
    "docValues": true
  },
  "add-field": {
    "name": "deathYear",
    "type": "pint",
    "stored": true,
    "indexed": true,
    "docValues": true
  },
  "add-field": {
    "name": "primaryProfession",
    "type": "string",
    "stored": true,
    "indexed": true,
    "docValues": true,
    "multiValued": true
  },
  "add-field": {
    "name": "knownForTitles",
    "type": "string",
    "stored": true,
    "indexed": true,
    "docValues": true,
    "multiValued": true
  },
  "add-field": {
    "name": "titleType",
    "type": "string",
    "stored": true,
    "indexed": true,
    "docValues": true
  },
  "add-field": {
    "name": "primaryTitle",
    "type": "text_general",
    "stored": true,
    "indexed": true
  },
  "add-field": {
    "name": "originalTitle",
    "type": "text_general",
    "stored": true,
    "indexed": true
  },
  "add-field": {
    "name": "isAdult",
    "type": "pint",
    "stored": true,
    "indexed": true,
    "docValues": true
  },
  "add-field": {
    "name": "startYear",
    "type": "pint",
    "stored": true,
    "indexed": true,
    "docValues": true
  },
  "add-field": {
    "name": "endYear",
    "type": "pint",
    "stored": true,
    "indexed": true,
    "docValues": true
  },
  "add-field": {
    "name": "runtimeMinutes",
    "type": "pint",
    "stored": true,
    "indexed": true,
    "docValues": true
  },
  "add-field": {
    "name": "genres",
    "type": "string",
    "stored": true,
    "indexed": true,
    "docValues": true,
    "multiValued": true
  },
  "add-field": {
    "name": "directors",
    "type": "string",
    "stored": true,
    "indexed": true,
    "docValues": true,
    "multiValued": true
  },
  "add-field": {
    "name": "writers",
    "type": "string",
    "stored": true,
    "indexed": true,
    "docValues": true,
    "multiValued": true
  }    ,
  "add-field": {
    "name": "principalCast",
    "type": "string",
    "stored": true,
    "indexed": true,
    "docValues": true,
    "multiValued": true
  }  
}' http://localhost:8983/solr/imdb/schema
```

### Current Status : under-development
At this stage of development, goimdb only downloads and extracts the necessary files and certain folders need to be created on the root directory of the project.

```sh
    files
    files/archive
    files/decompressed
    files/json
```
DownloadFiles() and GetFiles() have to be uncommented in main.go. In ReadFile() more lines should be added with the all the TSV files. Only title.ratings.tsv and title.basics.tsv are currently supported. goimdb will make title.ratings.json and title.ratings.json that are compatible to use with Solr database. 

In Solr, a core needs be created and field names should be manually created for goimdb to work. The following command will insert/update Solr database at core "imdb".

```sh
$ curl 'http://localhost:8983/solr/imdb/update?commit=true' -H 'Content-type:application/json' --data-binary @title.ratings.json
```

Sample search query:

```sh
http://localhost:8983/solr/imdb/select?q=averageRating:[8 TO *] AND numVotes:[20000 TO *] AND titleType:"movie"&rows=20&sort=numVotes desc
```
Most of the process will be automated in the future.

License
----
Apache-2.0

[//]: # 
   [Golang]: <https://golang.org/>
   [Solr]: <http://lucene.apache.org/solr/>