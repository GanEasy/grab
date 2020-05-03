package reader

import (
	"testing"
)

func Test_ShugeGetBooks(t *testing.T) {
	// urlStr := "http://feeds.twit.tv/twit.xml"
	// urlStr := "http://feed.williamlong.info/"
	urlStr := "https://m.shuge.la/sort/1_1/"
	urlStr = "https://m.shuge.la/sort/2_1/"
	reader := ShugeReader{}
	list, err := reader.GetList(urlStr)
	// list, err = reader.GetCategories()
	if err != nil {

	}
	t.Fatal(list)
}

func Test_ShugeGetChapters(t *testing.T) {
	urlStr := "https://m.shuge.la/read/5/5804/"
	urlStr = "https://m.shuge.la/read/6/6352/"
	// urlStr = "http://www.xbiquge.la/15/15484/"
	reader := ShugeReader{}
	list, err := reader.GetCatalog(urlStr)
	// list, err = reader.GetCategories()
	if err != nil {

	}
	t.Fatal(list)
}

func Test_ShugeGetChapter(t *testing.T) {
	urlStr := "http://www.xbiquge.la/29/29211/18370058.html"
	urlStr = "http://www.xbiquge.la/29/29211/18725612.html"
	urlStr = "https://m.shuge.la/read/6/6352/2937826.html"
	reader := ShugeReader{}
	list, err := reader.GetInfo(urlStr)
	// list, err = reader.GetCategories()
	if err != nil {

	}
	t.Fatal(list)
}

func Test_ShugeSearch(t *testing.T) {
	reader := ShugeReader{}
	list, err := reader.Search(`个个`)
	// list, err = reader.GetCategories()
	if err != nil {

	}
	t.Fatal(list)
}
