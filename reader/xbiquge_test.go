package reader

import (
	"testing"
)

func Test_XbiqugeGetBooks(t *testing.T) {
	// urlStr := "http://feeds.twit.tv/twit.xml"
	// urlStr := "http://feed.williamlong.info/"
	urlStr := "http://www.xbiquge.la/chuanyuexiaoshuo/"
	urlStr = "http://www.xbiquge.la/xuanhuanxiaoshuo/"
	reader := XbiqugeReader{}
	list, err := reader.GetList(urlStr)
	// list, err = reader.GetCategories()
	if err != nil {

	}
	t.Fatal(list)
}

func Test_XbiqugeGetChapters(t *testing.T) {
	urlStr := "http://www.xbiquge.la/39/39720/"
	urlStr = "http://www.xbiquge.la/29/29211/"
	// urlStr = "http://www.xbiquge.la/15/15484/"
	reader := XbiqugeReader{}
	list, err := reader.GetCatalog(urlStr)
	// list, err = reader.GetCategories()
	if err != nil {

	}
	t.Fatal(list)
}

func Test_XbiqugeGetChapter(t *testing.T) {
	urlStr := "http://www.xbiquge.la/29/29211/18370058.html"
	urlStr = "http://www.xbiquge.la/29/29211/18725612.html"
	urlStr = "http://www.xbiquge.la/15/15409/8163967.html"
	reader := XbiqugeReader{}
	list, err := reader.GetInfo(urlStr)
	// list, err = reader.GetCategories()
	if err != nil {

	}
	t.Fatal(list)
}
