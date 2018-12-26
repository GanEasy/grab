package reader

import (
	"testing"
)

func Test_LuoqiuGetBooks(t *testing.T) {
	// urlStr := "http://feeds.twit.tv/twit.xml"
	// urlStr := "http://feed.williamlong.info/"
	urlStr := "http://www.luoqiu.com/class/1_1.html"
	reader := LuoqiuReader{}
	list, err := reader.GetBooks(urlStr)
	// list, err = reader.GetCategories()
	if err != nil {

	}
	t.Fatal(list)
}

func Test_LuoqiuGetChapters(t *testing.T) {
	urlStr := "http://www.luoqiu.com/read/329070/"
	reader := LuoqiuReader{}
	list, err := reader.GetChapters(urlStr)
	// list, err = reader.GetCategories()
	if err != nil {

	}
	t.Fatal(list)
}
