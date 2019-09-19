package reader

import (
	"testing"
)

func Test_HanmanwoReaderGetInfo(t *testing.T) {
	urlStr := `https://r2hm.com/chapter/25486`
	urlStr = `http://www.hanmanwo.com/chapter/25757`
	// urlStr = `https://m.35xs.com/book/237551/51896850.html`
	reader := HanmanwoReader{}
	list, err := reader.GetInfo(urlStr)
	if err != nil {

	}
	t.Fatal(list)
}

func Test_HanmanwoReaderGetCatalog(t *testing.T) {
	urlStr := `https://m.booktxt.net/wapbook/4891.html`
	urlStr = `http://www.hanmanwo.com/book/468`
	reader := HanmanwoReader{}
	list, err := reader.GetCatalog(urlStr)
	if err != nil {

	}
	t.Fatal(list)
}

func Test_HanmanwoReaderGetList(t *testing.T) {
	urlStr := `https://m.booktxt.net/wapbook/4891.html`
	urlStr = `http://www.hanmanwo.com/booklist`
	reader := HanmanwoReader{}
	list, err := reader.GetList(urlStr)
	if err != nil {

	}
	t.Fatal(list)
}
