package reader

import (
	"testing"
)

func Test_GetQkshu6InfoReader(t *testing.T) {
	urlStr := `https://m.qkshu6.com/book/szdfmmgj/10010.html`
	// urlStr = `http://www.xinshubao.net/18/18685/2515188_2.html`
	// urlStr = `https://m.35xs.com/book/237551/51896850.html`
	reader := Qkshu6Reader{}
	list, err := reader.GetInfo(urlStr)
	if err != nil {

	}
	t.Fatal(list)
}
func Test_Qkshu6GetCatalog(t *testing.T) {
	urlStr := `https://m.qkshu6.com/book/ylshzwlxs/`
	urlStr = `https://m.qkshu6.com/book/szdfmmgj/`
	// urlStr = `https://m.uxiaoshuo.com/155/155018/all.html`
	reader := Qkshu6Reader{}
	list, err := reader.GetCatalog(urlStr)
	if err != nil {

	}
	t.Fatal(list)
}

func Test_Qkshu6GetBooks(t *testing.T) {
	// urlStr := "http://feeds.twit.tv/twit.xml"
	// urlStr := "http://feed.williamlong.info/"
	urlStr := "https://m.qkshu6.com/top.php"
	urlStr = "https://m.qkshu6.com/dushi/"
	reader := Qkshu6Reader{}
	list, err := reader.GetList(urlStr)
	// list, err = reader.GetCategories()
	if err != nil {

	}
	t.Fatal(list)
}
