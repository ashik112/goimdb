package gosolr

import(
	"net/http"
	"fmt"
	"strconv"
	"bytes"
	"io/ioutil"
)

func DeleteAll(hostname string, port int, core string){
	url:="http://"+hostname+":"+strconv.Itoa(port)+"/solr/"+core+"/update?commit=true"
	fmt.Println("URL:>",url) 
	var xmlStr = []byte(`<delete><query>*:*</query></delete>`)
	req,err:=http.NewRequest("POST",url,bytes.NewBuffer(xmlStr))
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

func CreateSolrFields(hostname string, port int, core string,path string,done chan bool){
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Creating fields from...")
	url:="http://"+hostname+":"+strconv.Itoa(port)+"/solr/"+core+"/schema"
	fmt.Println("URL:>",url) 
	req,err:=http.NewRequest("POST",url,bytes.NewBuffer(data))
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
func UploadDoc(hostname string, port int, core string,path string,done chan bool){
	// DeleteAll(hostname,port,core)
	
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Uploading file to solr...")
	url:="http://"+hostname+":"+strconv.Itoa(port)+"/solr/"+core+"/update?commit=true&separator=%09&escape=%5c&trim=true&commitWithin=120000"
	fmt.Println("URL:>",url) 
	req,err:=http.NewRequest("POST",url,bytes.NewBuffer(data))
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