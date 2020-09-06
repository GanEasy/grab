package reader

import (
	"testing"
)

func Test_BxksSearch(t *testing.T) {
	reader := BxksReader{}
	list, err := reader.Search(`大明`)
	// list, err = reader.GetCategories()
	if err != nil {

	}
	t.Fatal(list)
}

func Test_BxksGetBooks(t *testing.T) {
	urlStr := "https://www.mcmssc.com/xuanhuanxiaoshuo/"
	urlStr = "https://www.mcmssc.com/chuanyuexiaoshuo/"
	urlStr = "https://www.jininggeyin.com/ji/20.html"
	reader := BxksReader{}
	list, err := reader.GetList(urlStr)
	// list, err = reader.GetCategories()
	if err != nil {

	}
	t.Fatal(list)
}

func Test_BxksGetChapters(t *testing.T) {
	urlStr := "http://www.xbiquge.la/39/39720/"
	urlStr = "https://www.jininggeyin.com/jin/52006.html"
	// urlStr = "http://www.xbiquge.la/15/15484/"
	reader := BxksReader{}
	list, err := reader.GetCatalog(urlStr)
	// list, err = reader.GetCategories()
	if err != nil {

	}
	t.Fatal(list)
}

func Test_BxksGetChapter(t *testing.T) {
	urlStr := "http://www.xbiquge.la/29/29211/18370058.html"
	urlStr = "https://www.jininggeyin.com/jin1/52022/tEsuRzftbEJWm.html"
	reader := BxksReader{}
	list, err := reader.GetInfo(urlStr)
	// list, err = reader.GetCategories()
	if err != nil {

	}
	t.Fatal(list)
}
