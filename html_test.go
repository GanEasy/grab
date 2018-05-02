package grab

import (
	"testing"
)

func Test_GetHTML(t *testing.T) {
	// urlStr := "http://feeds.twit.tv/twit.xml"
	urlStr := "http://feed.williamlong.info/"
	list, err := GetHTML(urlStr, ``)
	if err != nil {

	}
	t.Fatal(list)
}
