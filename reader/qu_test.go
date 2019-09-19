package reader

import (
	"testing"
)

func Test_QuGetBooks(t *testing.T) {
	// urlStr := "http://feeds.twit.tv/twit.xml"
	// urlStr := "http://feed.williamlong.info/"
	urlStr := "https://m.uxiaoshuo.com/sort/3/1.html"
	urlStr = "https://m.qu.la/wapsort/0_1.html"
	urlStr = "https://m.qu.la/waptop/week3.html"
	reader := QuReader{}
	list, err := reader.GetList(urlStr)
	// list, err = reader.GetCategories()
	if err != nil {

	}
	t.Fatal(list)
}
func Test_QuGetChapters(t *testing.T) {
	// urlStr := "http://feeds.twit.tv/twit.xml"
	// urlStr := "http://feed.williamlong.info/"
	// http://book.zongheng.com/chapter/777234/43415281.html
	urlStr := "https://m.uxiaoshuo.com/282/282134/all.html"
	urlStr = "https://m.qu.la/booklist/142095.html"
	reader := QuReader{}
	list, err := reader.GetCatalog(urlStr)
	// list, err = reader.GetCategories()
	if err != nil {

	}
	t.Fatal(list)
}

func Test_QuGetChapter(t *testing.T) {
	urlStr := "https://m.uxiaoshuo.com/282/282134/1460954.html"
	urlStr = "https://m.uxiaoshuo.com/140/140420/7432883.html"
	urlStr = "https://m.uxiaoshuo.com/278/278598/1741780.html"
	urlStr = "https://m.uxiaoshuo.com/281/281973/1795867.html"
	urlStr = "https://m.uxiaoshuo.com/281/281973/1796673.html"
	urlStr = "https://m.qu.la/book/142095/7545899.html"
	reader := QuReader{}
	list, err := reader.GetInfo(urlStr)
	// list, err = reader.GetCategories()
	if err != nil {

	}
	t.Fatal(list)
}
