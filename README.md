# goimdb
IMDB movie database using Golang (only for personal and non-commercial use)

goimdb downloads data from datasets.imdbws.com, currently provided by imdb. These are the subsets of IMDb data that are available for personal and non-commercial use. Each dataset is gzipped, TSV (tab-separated-values) formatted file. The first line of each file contains headers.

goimdb downloads, extracts these file and inserts them into  Solr. Solr is capable of conducting full-text search. The latest version of Solr, 7.2.1 is tested with goimdb.
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
In Solr, a core needs be created for goimdb to work.

Create core:
```sh
$ solr create -c imdb
```

### Current Status :

<Partially Complete: Download, Decompression, Data insertion in Solr> 

<Under-development & Testing: API, User Input>

At this stage of development, goimdb only downloads and extracts the necessary files and certain folders need to be created on the root directory of the project.

```sh
    files
    files/archive
    files/decompressed
    files/json
```
File download, decompression, data insertion are managed by goimdb. Some parts of the project are hard-coded which will be refactored and fixed soon.

License
----
Apache-2.0

[//]: # 
   [Golang]: <https://golang.org/>
   [Solr]: <http://lucene.apache.org/solr/>