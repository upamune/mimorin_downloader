package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"time"
)

func GetJSON() (urls []string) {
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

	// println(string(body))

	// re := regexp.MustCompile(`"MediaUrl":"(.+?)",`)
	re := regexp.MustCompile(`"Title":.+?"MediaUrl":"(.+?)",`)
	matches := re.FindAllStringSubmatch(string(body), -1)

	for _, url := range matches {
		urls = append(urls, url[1])
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

	dirName := "mimorin"
	if err := os.Mkdir(dirName, 0777); err != nil {
		panic(err)
	}
	start := time.Now()
	statusChan := make(chan string)
	for idx, url := range urls {
		filePath := dirName + "/" + "mimorin" + strconv.Itoa(idx) + ".jpg"
		go func(url, filePath string){
			saveImageFile(url, filePath)
			statusChan <- ("Downloading... " + filePath)
		}(url,filePath)
	}
	for i := 0; i < len(urls); i++{
		fmt.Println(<-statusChan)
	}
	end := time.Now()

	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())
}
