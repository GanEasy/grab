package reader

import (
	"testing"
)

func Test_Xin18Search(t *testing.T) {
	reader := Xin18Reader{}
	list, err := reader.Search(`大明`)
	// list, err = reader.GetCategories()
	if err != nil {

	}
	t.Fatal(list)
}

func Test_Xin18GetBooks(t *testing.T) {
	urlStr := "https://www.mcmssc.com/xuanhuanxiaoshuo/"
	urlStr = "https://www.mcmssc.com/chuanyuexiaoshuo/"
	urlStr = "https://www.0335jjlm.com/03/18.html"
	reader := Xin18Reader{}
	list, err := reader.GetList(urlStr)
	// list, err = reader.GetCategories()
	if err != nil {

	}
	t.Fatal(list)
}

func Test_Xin18GetChapters(t *testing.T) {
	urlStr := "http://www.xbiquge.la/39/39720/"
	urlStr = "https://www.0335jjlm.com/0335/45234.html"
	urlStr = "https://www.0335jjlm.com/other/chapters/id/19207.html"
	// urlStr = "http://www.xbiquge.la/15/15484/"
	reader := Xin18Reader{}
	list, err := reader.GetCatalog(urlStr)
	// list, err = reader.GetCategories()
	if err != nil {

	}
	t.Fatal(list)
}

func Test_Xin18GetChapter(t *testing.T) {
	urlStr := "http://www.xbiquge.la/29/29211/18370058.html"
	urlStr = "https://www.0335jjlm.com/03351/45828/O1tSgYNpJbmOn.html"
	reader := Xin18Reader{}
	list, err := reader.GetInfo(urlStr)
	// list, err = reader.GetCategories()
	if err != nil {

	}
	t.Fatal(list)
}
