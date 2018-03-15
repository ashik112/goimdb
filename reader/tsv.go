package reader

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

/*ReadTSV does ...*/
func ReadTSV(target string) {
	start:=time.Now()
	csvFile, err := os.Open(target)

	if err != nil {
		fmt.Println(err)
	}

	defer csvFile.Close()

	reader := csv.NewReader(csvFile)

	reader.Comma = '\t' // Use tab-delimited instead of comma <---- here!
	reader.LazyQuotes = true
	 reader.FieldsPerRecord = -1

	fmt.Println("Reading Data from file...")
	csvData, err := reader.ReadAll()	
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Reading data took ",time.Since(start))
	count := 0
	 //fmt.Println(csvData)
	for _, each := range csvData {
		fmt.Println(each[0], " ", each[1], " ", each[2])
		count = count + 1
	}
	fmt.Println("total: ", count)
	
}
