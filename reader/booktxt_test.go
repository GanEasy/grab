package reader

import (
	"testing"
)

func Test_BooktxtGetInfo(t *testing.T) {
	urlStr := `https://m.booktxt.net/wapbook/4891_4943641.html`
	// urlStr = `http://www.xinshubao.net/18/18685/2515188_2.html`
	// urlStr = `https://m.35xs.com/book/237551/51896850.html`
	reader := BooktxtReader{}
	list, err := reader.GetInfo(urlStr)
	if err != nil {

	}
	t.Fatal(list)
}
func Test_BooktxtGetCatalog(t *testing.T) {
	urlStr := `https://m.booktxt.net/wapbook/4891.html`
	urlStr = `https://m.booktxt.net/wapbook/6454.html`
	urlStr = `https://www.booktxt.net/book/goto/id/4243`
	urlStr = `https://www.booktxt.net/10_10269/`
	reader := BooktxtReader{}
	list, err := reader.GetCatalog(urlStr)
	if err != nil {

	}
	t.Fatal(list)
}

func Test_BooktxtSearch(t *testing.T) {
	reader := BooktxtReader{}
	list, err := reader.Search(`圣王`)
	if err != nil {

	}
	t.Fatal(list)
}

func Test_BooktxtGetList(t *testing.T) {
	urlStr := `https://m.booktxt.net/wapsort/1_1.html`
	urlStr = `https://m.booktxt.net/wapsort/5_1.html`
	reader := BooktxtReader{}
	list, err := reader.GetList(urlStr)
	if err != nil {

	}
	t.Fatal(list)
}
