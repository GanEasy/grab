package reader

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

func Test_GetURLStringParam(t *testing.T) {
	// urlStr := "http://feeds.twit.tv/twit.xml"
	urlStr := "/pages/catalog?drive=xbiquge&url=aHR0cDovL3d3dy54YmlxdWdlLmxhLzE1LzE1MDIxLw%3D%3D"

	t.Fatal(GetURLStringParam(urlStr, `url`))
}
