package main

import (
	"fmt"
	"net/http"
)

func main() {
	url := "https://pastebin.com/"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("unable to get %s: %v", url, err)
	}
	fmt.Println(resp)
}
