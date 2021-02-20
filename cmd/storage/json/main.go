package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

//Rsp is
type Rsp struct {
	Cards []Card
}

//Card is
type Card struct {
	//Id is
	ID    int
	Image string
}

var (
	fileName    string
	fullURLFile string
)

func main() {
	url := "https://api.blizzard.com/hearthstone/cards?gameMode=battlegrounds&pageSize=200&locale=en_US"
	bearer := "Bearer EUcxr3binCHCQEF0kcgO6oKCY4u85oK2zY"

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}

	data := Rsp{}
	err = json.Unmarshal(body, &data)

	//fmt.Printf("Results: %v\n", data.Cards)
	// В data.Cards теперь лежит то, что можно перебрать циклом

	for _, v := range data.Cards {
		fmt.Println(v.Image)
		fullURLFile = v.Image

		// Build fileName from fullPath
		buildFileName()

		// Create blank file
		file := createFile()

		// Put content on file
		putFile(file, httpClient())
		if err != nil {
			panic(err)
		}
		fmt.Println("Downloaded: " + v.Image)
	}
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func putFile(file *os.File, client *http.Client) {
	resp, err := client.Get(fullURLFile)

	checkError(err)

	defer resp.Body.Close()

	size, err := io.Copy(file, resp.Body)

	defer file.Close()

	checkError(err)

	fmt.Printf("%v%v%v", "Just Downloaded a file %s with size %d", fileName, size)
}
func buildFileName() {
	fileURL, err := url.Parse(fullURLFile)
	checkError(err)

	path := fileURL.Path
	segments := strings.Split(path, "/")

	fileName = segments[len(segments)-1]
}

func httpClient() *http.Client {
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	return &client
}

func createFile() *os.File {
	file, err := os.Create(fileName)

	checkError(err)
	return file
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
