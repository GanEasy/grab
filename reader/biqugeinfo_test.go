package reader

import (
	"testing"
)

func Test_BiqugeinfoGetInfo(t *testing.T) {
	urlStr := `https://m.biquge.info/10_10218/5001515.html`
	// urlStr = `http://www.xinshubao.net/18/18685/2515188_2.html`
	// urlStr = `https://m.35xs.com/book/237551/51896850.html`
	reader := BiqugeinfoReader{}
	list, err := reader.GetInfo(urlStr)
	if err != nil {

	}
	t.Fatal(list)
}
func Test_BiqugeinfoGetCatalog(t *testing.T) {
	urlStr := `https://m.booktxt.net/wapbook/4891.html`
	urlStr = `https://m.biquge.info/51_51498/`
	reader := BiqugeinfoReader{}
	list, err := reader.GetCatalog(urlStr)
	if err != nil {

	}
	t.Fatal(list)
}

func Test_BiqugeinfoGetList(t *testing.T) {
	urlStr := `https://m.booktxt.net/wapsort/1_1.html`
	urlStr = `https://m.biquge.info/list/1_1.html`
	reader := BiqugeinfoReader{}
	list, err := reader.GetList(urlStr)
	if err != nil {

	}
	t.Fatal(list)
}
