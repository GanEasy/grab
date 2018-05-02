package grab

import (
	"testing"
)

func Test_FindContentForHTML(t *testing.T) {
	// urlStr := "http://feeds.twit.tv/twit.xml"
	urlStr := "http://www.williamlong.info/"
	htmlStr, err := GetHTML(urlStr, ``)
	if err != nil {

	}
	h, err := FindContentForHTML(htmlStr, `#divMain`)
	if err != nil {

	}
	t.Fatal(h)
}
