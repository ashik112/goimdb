package httpreq

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