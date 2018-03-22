package decompresser

import (
	"compress/gzip"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

/*UnGzip does...*/
func UnGzip(source, target string, done chan int) {
	start := time.Now()
	log.Println("Decompressing ", source)

	file := path.Base(source)

	reader, err := os.Open(source)
	if err != nil {
		log.Println(err)
	}
	defer reader.Close()

	archive, err := gzip.NewReader(reader)
	if err != nil {
		log.Println(err)
	}
	defer archive.Close()

	// target = filepath.Join(target, archive.Name)
	archiveName := strings.Replace(file, ".gz", "", -1)
	target = filepath.Join(target, archiveName)
	// log.Println(target)
	writer, err := os.Create(target)
	if err != nil {
		log.Println(err)
	}
	defer writer.Close()

	_, err = io.Copy(writer, archive)
	if err != nil {
		log.Println(err)
	}

	elapsed := time.Since(start)
	log.Println("Decompressed ", archiveName, " in ", elapsed)
	done <- 1
}
