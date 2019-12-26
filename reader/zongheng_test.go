package reader

import (
	"testing"
)

func Test_ZonghengGetList(t *testing.T) {
	// urlStr := "http://feeds.twit.tv/twit.xml"
	// urlStr := "http://feed.williamlong.info/"
	urlStr := "http://book.zongheng.com/store/c1/c0/b0/u6/p1/v9/s9/t0/u0/i1/ALL.html"
	reader := ZonghengReader{}
	list, err := reader.GetList(urlStr)
	// list, err = reader.GetCategories()
	if err != nil {

	}
	t.Fatal(list)
}

func Test_ZonghengGetCatalog(t *testing.T) {
	// urlStr := "http://feeds.twit.tv/twit.xml"
	// urlStr := "http://feed.williamlong.info/"
	// http://book.zongheng.com/chapter/777234/43415281.html
	urlStr := "http://book.zongheng.com/showchapter/777234.html"
	reader := ZonghengReader{}
	list, err := reader.GetCatalog(urlStr)
	// list, err = reader.GetCategories()
	if err != nil {

	}
	t.Fatal(list)
}

func Test_ZonghengGetInfo(t *testing.T) {
	// urlStr := "http://feeds.twit.tv/twit.xml"
	// urlStr := "http://feed.williamlong.info/"
	urlStr := "https://www.xxsy.net/chapter/79588269.html"
	urlStr = "http://book.zongheng.com/chapter/777234/43415281.html"
	reader := ZonghengReader{}
	list, err := reader.GetInfo(urlStr)
	// list, err = reader.GetCategories()
	if err != nil {

	}
	t.Fatal(list)
}
