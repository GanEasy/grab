package grab

import (
	"testing"
)

func Test_GetRssList(t *testing.T) {
	// urlStr := "http://feeds.twit.tv/twit.xml"
	urlStr := "http://feed.williamlong.info/"
	list, err := GetRssList(urlStr)
	if err != nil {

	}
	t.Fatal(list)
}
