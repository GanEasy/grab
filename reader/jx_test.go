package reader

import (
	"testing"
)

func Test_JxReaderGetBooks(t *testing.T) {
	// urlStr := "http://feeds.twit.tv/twit.xml"
	// urlStr := "http://feed.williamlong.info/"
	urlStr := "https://m.uxiaoshuo.com/sort/3/1.html"
	urlStr = "https://m.jx.la/wapsort/0_1.html"
	urlStr = "https://m.jx.la/waptop/week3.html"
	urlStr = "https://m.jx.la/wapsort/0_1.html"
	urlStr = "https://m.jx.la/waptop/month.html"
	urlStr = "https://m.jx.la/xuanhuanxiaoshuo/"
	// urlStr = "https://m.jx.la/paihangbang/"
	reader := JxReader{}
	list, err := reader.GetList(urlStr)
	// list, err = reader.GetCategories()
	if err != nil {

	}
	t.Fatal(list)
}
func Test_JxReaderSearch(t *testing.T) {
	reader := JxReader{}
	list, err := reader.Search(`点道`)
	// list, err = reader.GetCategories()
	if err != nil {

	}
	t.Fatal(list)
}
func Test_JxReaderGetChapters(t *testing.T) {
	// urlStr := "http://feeds.twit.tv/twit.xml"
	// urlStr := "http://feed.williamlong.info/"
	// http://book.zongheng.com/chapter/777234/43415281.html
	urlStr := "https://m.uxiaoshuo.com/282/282134/all.html"
	urlStr = "https://m.jx.la/booklist/142095.html"
	// urlStr = "https://m.jx.la/book/142095/"
	reader := JxReader{}
	list, err := reader.GetCatalog(urlStr)
	// list, err = reader.GetCategories()
	if err != nil {

	}
	t.Fatal(list)
}

func Test_JxReaderGetChapter(t *testing.T) {
	urlStr := "https://m.uxiaoshuo.com/282/282134/1460954.html"
	urlStr = "https://m.jx.la/book/39775/2494931.html"
	// urlStr = "https://m.jx.la/book/142095/7545899.html"
	reader := JxReader{}
	list, err := reader.GetInfo(urlStr)
	// list, err = reader.GetCategories()
	if err != nil {

	}
	t.Fatal(list)
}
