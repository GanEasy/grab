package reader

import (
	"testing"
)

func Test_HaimaobaReaderGetInfo(t *testing.T) {
	urlStr := `http://www.haimaoba.com/catalog/4030/68270.html`
	// urlStr = `http://www.xinshubao.net/18/18685/2515188_2.html`
	// urlStr = `https://m.35xs.com/book/237551/51896850.html`
	reader := HaimaobaReader{}
	list, err := reader.GetInfo(urlStr)
	if err != nil {

	}
	t.Fatal(list)
}

func Test_HaimaobaReaderGetCatalog(t *testing.T) {
	urlStr := `https://m.booktxt.net/wapbook/4891.html`
	urlStr = `http://www.haimaoba.com/catalog/4032/`
	reader := HaimaobaReader{}
	list, err := reader.GetCatalog(urlStr)
	if err != nil {

	}
	// sort.Reverse(list.Cards)
	// t.Fatal(sort.Sort(sort.Reverse(list.Cards)))
	t.Fatal(list)
}

func Test_HaimaobaReaderGetList(t *testing.T) {
	urlStr := `https://m.booktxt.net/wapbook/4891.html`
	urlStr = `http://www.haimaoba.com/list/0/`
	reader := HaimaobaReader{}
	list, err := reader.GetList(urlStr)
	if err != nil {

	}
	t.Fatal(list)
}
