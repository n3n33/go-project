package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

var baseUrl string = "https://kr.indeen.com/jobs?q=python&limit=50"

func main() {
	pages := getPages()
}

func getPages() int {
	res, err := http.Get(baseUrl)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()
	checkErr(err)

	doc, err := goquery.NewDocumentFromReader(res.Body)

	fmt.Println(doc)
	return 0
}

func checkErr(err error) error {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with status: ", res.Status)
	}
}
