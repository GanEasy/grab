package grab

import (
	"testing"
)

func Test_GetHTMLList(t *testing.T) {
	// urlStr := "http://feeds.twit.tv/twit.xml"
	urlStr := "http://www.williamlong.info/"
	list, err := GetHTMLList(urlStr, `#divMain`)
	if err != nil {

	}
	l := Cleaning(list.Links)
	t.Fatal(l)
}
