# goimdb (under-development)
IMDB movie database using Golang (only for personal and non-commercial use)

goimdb downloads data from datasets.imdbws.com, currently provided by imdb. These are the subsets of IMDb data that are available for personal and non-commercial use. Each dataset is gzipped, TSV (tab-separated-values) formatted file. The first line of each file contains headers.

goimdb downloads, extracts and converts these file to Solr compatible JSON format to be able to insert them into a Solr database. Solr is capable of conducting full-text search and provides apis for searching. The current latest version of Solr, 7.2.1 is tested with goimdb.

# Requirement

    Go 1.x
    Solr7.2.x
    JDK/JRE 8.x
    curl 7.x
# Installation
goimdb is go-gettable. The follwing command will download the full projet:

    go get github.com/ashik112/goimb

# Current Status
At this stage of development, goimdb only downloads and extracts the necessary files and certain folders need to be created on the root directory of the project.

    files
    files/archive
    files/decompressed

DownloadFiles() and GetFiles() have to be uncommented in main.go. In ReadFile() more lines should be added with the all the TSV files. Only title.ratings.tsv and title.basics.tsv are currently supported. goimdb will make title.ratings.json and title.ratings.json that are compatible to use with Solr database. 

In Solr, a core needs be created and field names should be manually created for goimdb to work. The following command will insert/update Solr database at core "imdb".

    curl 'http://localhost:8983/solr/imdb/update?commit=true' -H 'Content-type:application/json' --d
    ata-binary @E:\workspace\go\src\github.com\ashik112\goimdb\files\decompressed\title.ratings.json

Most of the process will be automated in the future.