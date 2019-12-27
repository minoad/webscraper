package collect

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// #content_left > div:nth-child(7) > table > tbody > tr:nth-child(2) > td:nth-child(1) > a
func getPage(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("unable to get %s: %v", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode > 400 {
		return fmt.Errorf("error getting %s: %v", url, resp.Status)
	}

	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("unable to download webpage at %s: %v", url, err)
	}

	//htmlString :=
	fmt.Println("calling")
	collectPasteIDs(fmt.Sprintf("%s\n", html))
	return nil
}

func collectPasteIDs(htmlBody string) {
	//fmt.Println(htmlBody)
	//re := regexp.MustCompile(`<a.*href\s*=\s*["'](http[s]{0,1}:\/\/.[^\s]*)["'].*>`)
	// `<a.*href="/([A-Za-z-0-9]) style`
	fmt.Println("a ...interface{}")
	regexString := `<a href=(.+)`
	re := regexp.MustCompile(regexString)
	linkMatches := re.FindAllStringSubmatch(htmlBody, -1)
	fmt.Println(linkMatches)
	//fmt.Println(len(linkMatches))
}
