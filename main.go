package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/jmoiron/jsonq"
)

func GetJSON() (urls [][]string) {
	client := &http.Client{}
	URL := "https://api.datamarket.azure.com/Bing/Search/Image"
	apikey := os.Getenv("BING_API_KEY")
	if apikey == "" {
		fmt.Println("ENV is not set")
		panic(apikey)
	}

	values := url.Values{}
	values.Add("Query", "'三森すずこ'")
	values.Add("$format", "json")
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		println("err: ", err)
	}

	req.URL.RawQuery = values.Encode()
	req.SetBasicAuth(apikey, apikey)
	response, _ := client.Do(req)
	body, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	jsonStr := string(body)
	data := map[string]interface{}{}
	dec := json.NewDecoder(strings.NewReader(jsonStr))
	dec.Decode(&data)
	jq := jsonq.NewQuery(data)

	for i := 0; i < 50; i++ {
		url, _ := jq.String("d", "results", strconv.Itoa(i), "MediaUrl")
		contentType, _ := jq.String("d", "results", strconv.Itoa(i), "ContentType")
		var imageType string
		switch contentType {
		case "image/jpeg":
			imageType = "jpeg"
		case "image/png":
			imageType = "png"
		case "image/gif":
			imageType = "gif"
		default:
			imageType = "jpeg"
		}

		urls = append(urls, []string{url, imageType})
	}

	return
}

func saveImageFile(url string, filePath string) (err error) {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	io.Copy(file, response.Body)

	return err
}

func main() {

	urls := GetJSON()
	timeStamp := time.Now().Format("20060102150405")

	dirName := "mimorin-" + timeStamp
	if err := os.Mkdir(dirName, 0777); err != nil {
		panic(err)
	}
	cpus := runtime.NumCPU()
	runtime.GOMAXPROCS(cpus)
	start := time.Now()
	statusChan := make(chan string)
	for idx, url := range urls {
		filePath := dirName + "/" + "mimorin" + strconv.Itoa(idx) + "." + url[1]
		go func(url, filePath string) {
			saveImageFile(url, filePath)
			statusChan <- ("Downloading... " + filePath)
		}(url[0], filePath)
	}
	for i := 0; i < len(urls); i++ {
		fmt.Println(<-statusChan)
	}
	end := time.Now()

	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())
}
