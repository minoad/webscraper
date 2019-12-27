package collect

import (
	"fmt"
	"testing"
)

func Test_getPage(t *testing.T) {
	url := "https://pastebin.com/archive"
	n := getPage(url)
	fmt.Println(n)
	//getPage("http://www.example.com/index.html")
}
