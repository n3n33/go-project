package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request Failed")

// hit URL
func main() {
	urls := []string{
		"www.naver.com",
		"www.google.com",
	}
	for _, url := range urls {
		//fmt.Println(url)
		hitURL(url)
	}
}

//웹사이트에 접속
func hitURL(url string) error {
	fmt.Println("Checking: ", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		return errRequestFailed
	}
	return nil
}
