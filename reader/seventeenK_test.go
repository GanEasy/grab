package reader

import (
	"testing"
)

func Test_SeventeenKGetBooks(t *testing.T) {
	// urlStr := "http://feeds.twit.tv/twit.xml"
	// urlStr := "http://feed.williamlong.info/"
	urlStr := "http://www.17k.com/all/book/2_21_0_0_0_0_0_0_1.html"
	reader := SeventeenKReader{}
	list, err := reader.GetList(urlStr)
	// list, err = reader.GetCategories()
	if err != nil {

	}
	t.Fatal(list)
}

func Test_SeventeenKGetChapters(t *testing.T) {
	// urlStr := "http://feeds.twit.tv/twit.xml"
	// urlStr := "http://feed.williamlong.info/"
	// http://book.zongheng.com/chapter/777234/43415281.html
	urlStr := "https://www.17k.com/list/2842794.html"
	reader := SeventeenKReader{}
	list, err := reader.GetCatalog(urlStr)
	// list, err = reader.GetCategories()
	if err != nil {

	}
	t.Fatal(list)
}

func Test_SeventeenKGetChapter(t *testing.T) {
	urlStr := "http://www.17k.com/chapter/493239/10060592.html"
	reader := SeventeenKReader{}
	list, err := reader.GetInfo(urlStr)
	// list, err = reader.GetCategories()
	if err != nil {

	}
	t.Fatal(list)
}
