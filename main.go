package main

import "fmt"

// hit URL
func main() {
	urls := []string{
		"www.naver.com",
		"www.google.com",
	}
	for _, url := range urls {
		fmt.Println(url)
	}
}
