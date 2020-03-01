package reader

import (
	"regexp"
	"testing"
)

func Test_Xs280GetBooks(t *testing.T) {
	urlStr := "https://www.Xs280.com/xuanhuanxiaoshuo/"
	urlStr = "https://www.Xs280.com/chuanyuexiaoshuo/"
	urlStr = "https://www.280xs.com/book_1_1/"
	reader := Xs280Reader{}
	list, err := reader.GetList(urlStr)
	// list, err = reader.GetCategories()
	if err != nil {

	}
	t.Fatal(list)
}

func Test_Xs280Search(t *testing.T) {
	reader := Xs280Reader{}
	list, err := reader.Search(`修仙`)
	// list, err = reader.GetCategories()
	if err != nil {

	}
	t.Fatal(list)
}
func Test_Xs280GetChapters(t *testing.T) {
	urlStr := "http://www.xbiquge.la/39/39720/"
	urlStr = "https://www.280xs.com/dingdian/34_34693/"
	// urlStr = "http://www.xbiquge.la/15/15484/"
	reader := Xs280Reader{}
	list, err := reader.GetCatalog(urlStr)
	// list, err = reader.GetCategories()
	if err != nil {

	}
	t.Fatal(list)
}

func Test_Xs280GetChapter(t *testing.T) {
	urlStr := "http://www.xbiquge.la/29/29211/18370058.html"
	urlStr = "https://www.280xs.com/dingdian/26_26019/n7169097.html"
	reader := Xs280Reader{}
	list, err := reader.GetInfo(urlStr)
	// list, err = reader.GetCategories()
	if err != nil {

	}
	t.Fatal(list)
}

func Test_Xs280MustCompile(t *testing.T) {
	str := `window.location.href = "/chuanyuexiaoshuo/4_" + obj.curr + ".html";`

	reg := regexp.MustCompile(`window.location.href = "(?P<pagemod>[^"]+)"`)

	str = reg.ReplaceAllString(str, "")

	t.Fatal(str)
}
