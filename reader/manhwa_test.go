package reader

import (
	"testing"
)

func Test_ManhwaReaderGetInfo(t *testing.T) {
	urlStr := `https://r2hm.com/chapter/25486`
	urlStr = `https://www.manhwa.cc/chapter/49266`
	// urlStr = `https://m.35xs.com/book/237551/51896850.html`
	reader := ManhwaReader{}
	list, err := reader.GetInfo(urlStr)
	if err != nil {

	}
	t.Fatal(list)
}

func Test_ManhwaReaderGetCatalog(t *testing.T) {
	urlStr := `https://m.booktxt.net/wapbook/4891.html`
	urlStr = `https://www.manhwa.cc/book/997`
	reader := ManhwaReader{}
	list, err := reader.GetCatalog(urlStr)
	if err != nil {

	}
	t.Fatal(list)
}

func Test_ManhwaReaderGetList(t *testing.T) {
	urlStr := `https://m.booktxt.net/wapbook/4891.html`
	urlStr = `https://www.manhwa.cc/booklist`
	urlStr = `https://www.manhwa.cc/booklist/%E9%9F%A9%E5%9B%BD`
	reader := ManhwaReader{}
	list, err := reader.GetList(urlStr)
	if err != nil {

	}
	t.Fatal(list)
}
