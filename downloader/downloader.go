package downloader

import (
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"
)

/*Download does ...*/
func Download(destination string, url string) {
	start := time.Now()
	//Create the file
	file := path.Base(url)
	log.Printf("Downloading %s", file)
	path := destination + "/" + file
	out, err := os.Create(path)
	if err != nil {
		log.Println(err)
	}
	defer out.Close()

	headResp, err := http.Head(url)

	if err != nil {
		panic(err)
	}

	defer headResp.Body.Close()

	size, err := strconv.Atoi(headResp.Header.Get("Content-Length"))
	if err != nil {
		log.Println(err)
	}
	//Get data
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		panic(err)
	}
	elapsed := time.Since(start)
	sizeMb := float64(size) * 0.000001
	log.Printf("%0.3f MB downloaded completed in %s", sizeMb, elapsed)
}
