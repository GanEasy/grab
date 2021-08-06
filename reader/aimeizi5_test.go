package reader

import (
	"testing"
)

func Test_Aimeizi5ReaderGetInfo(t *testing.T) {
	urlStr := `https://5aimeizi.com/chapter/23768`
	// urlStr = `http://www.xinshubao.net/18/18685/2515188_2.html`
	// urlStr = `https://m.35xs.com/book/237551/51896850.html`
	reader := Aimeizi5Reader{}
	list, err := reader.GetInfo(urlStr)
	if err != nil {

	}
	t.Fatal(list)
}

func Test_Aimeizi5ReaderGetCatalog(t *testing.T) {
	// https://www.feixuemh.com/chapter/54380
	urlStr := `https://m.booktxt.net/wapbook/4891.html`
	urlStr = `https://5aimeizi.com/Comic/3216`
	reader := Aimeizi5Reader{}
	list, err := reader.GetCatalog(urlStr)
	if err != nil {

	}
	t.Fatal(list)
}

func Test_Aimeizi5ReaderGetList(t *testing.T) {
	urlStr := `https://m.booktxt.net/wapbook/4891.html`
	urlStr = `https://5aimeizi.com/booklist?page=2`
	reader := Aimeizi5Reader{}
	list, err := reader.GetList(urlStr)
	if err != nil {

	}
	t.Fatal(list)
}
